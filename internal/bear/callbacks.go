package bear

import (
	"fmt"
	"net/http"
	"net/url"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

// callbackServer represents a local HTTP server to handle Bear callbacks
type callbackServer struct {
	server   *http.Server
	port     int
	response chan callbackResponse
}

// callbackResponse contains the response data from a Bear callback
type callbackResponse struct {
	Data map[string]string
	Err  error
}

// startCallbackServer starts a local HTTP server to receive callback responses
func startCallbackServer() (*callbackServer, error) {
	// Use a fixed port for now - could be made configurable
	server := &callbackServer{
		port:     8789,
		response: make(chan callbackResponse, 1),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", server.handleCallback)

	server.server = &http.Server{
		Addr:    fmt.Sprintf("127.0.0.1:%d", server.port),
		Handler: mux,
	}

	go func() {
		if err := server.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			server.response <- callbackResponse{Err: err}
		}
	}()

	// Wait for server to start
	time.Sleep(100 * time.Millisecond)

	return server, nil
}

// handleCallback processes the callback request from Bear
func (s *callbackServer) handleCallback(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]string)

	// Parse query parameters
	query := r.URL.Query()
	for key, values := range query {
		if len(values) > 0 {
			data[key] = values[0]
		}
	}

	// Debug log
	fmt.Printf("Received callback with data: %v\n", data)

	s.response <- callbackResponse{Data: data}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Callback received"))

	// Close the server after handling the callback
	go func() {
		time.Sleep(100 * time.Millisecond)
		s.server.Close()
	}()
}

// executeCallback executes an x-callback-url request to Bear
func (c *Client) executeCallback(action string, params map[string]string, expectResponse bool) (map[string]string, error) {
	if c.DryRun {
		fmt.Printf("Would execute Bear callback: %s with params: %v\n", action, params)
		return make(map[string]string), nil
	}

	if runtime.GOOS != "darwin" {
		return nil, fmt.Errorf("bear callbacks are only supported on macOS")
	}

	var callbackData map[string]string = make(map[string]string)
	var server *callbackServer

	if expectResponse {
		// Start callback server
		var err error
		server, err = startCallbackServer()
		if err != nil {
			return nil, fmt.Errorf("failed to start callback server: %w", err)
		}

		// Add callback parameters
		callbackURLBase := fmt.Sprintf("http://127.0.0.1:%d", server.port)
		params["x-success"] = callbackURLBase
		params["x-error"] = callbackURLBase + "?error=1&message=" + url.QueryEscape("Error processing request")
	}

	// Build the callback URL
	baseURL := "bear://x-callback-url/" + action
	callbackURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("invalid URL: %w", err)
	}

	query := callbackURL.Query()
	for key, value := range params {
		query.Add(key, value)
	}
	callbackURL.RawQuery = query.Encode()

	// Replace + with %20 for proper space encoding
	urlStr := strings.Replace(callbackURL.String(), "+", "%20", -1)
	
	// Execute the callback URL using 'open' command
	fmt.Printf("Executing Bear URL: %s\n", urlStr)
	cmd := exec.Command("open", urlStr)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed to execute Bear callback: %w, output: %s", err, output)
	}

	if expectResponse && server != nil {
		// Wait for response from callback server
		select {
		case response := <-server.response:
			if response.Err != nil {
				return nil, fmt.Errorf("callback server error: %w", response.Err)
			}
			callbackData = response.Data
			fmt.Printf("Received callback data: %v\n", callbackData)
		case <-time.After(10 * time.Second):
			return nil, fmt.Errorf("timeout waiting for callback response")
		}

		// Check for error response
		if _, ok := callbackData["error"]; ok {
			return nil, fmt.Errorf("bear error: %s", callbackData["message"])
		}
	}

	return callbackData, nil
}

// callbackCreate handles the /create x-callback and returns the created note
func (c *Client) callbackCreate(params map[string]string) (*Note, error) {
	// Validate title parameter
	title, ok := params["title"]
	if !ok || title == "" {
		return nil, fmt.Errorf("title is required for creating a note")
	}

	if c.DryRun {
		fmt.Printf("Would create Bear note with title: %s\n", title)

		// For dry run, return a fake note
		tags := []string{}
		if tagsStr, ok := params["tags"]; ok && tagsStr != "" {
			tags = parseTagsString(tagsStr)
		}

		return &Note{
			ID:      "dry-run-id",
			Title:   title,
			Content: params["text"],
			Tags:    tags,
		}, nil
	}

	// Add the show_window parameter to ensure Bear gets focus
	params["show_window"] = "yes"
	
	// Execute the callback with response handling
	response, err := c.executeCallback("create", params, true)
	if err != nil {
		return nil, err
	}

	// Extract note information from response
	noteID, ok := response["note_id"]
	if !ok {
		// If we didn't get a note_id back, create a note with the title as ID
		// This is a fallback in case the callback fails but the note was created
		return &Note{
			ID:      "unknown",
			Title:   title,
			Content: params["text"],
		}, nil
	}

	// Extract tags from parameters
	tags := []string{}
	if tagsStr, ok := params["tags"]; ok && tagsStr != "" {
		tags = parseTagsString(tagsStr)
	}

	return &Note{
		ID:      noteID,
		Title:   title,
		Content: params["text"],
		Tags:    tags,
	}, nil
}