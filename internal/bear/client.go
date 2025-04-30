package bear

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

// Client provides methods to interact with Bear app
type Client struct {
	DryRun bool
}

// NewClient creates a new Bear client
func NewClient(dryRun bool) *Client {
	return &Client{
		DryRun: dryRun,
	}
}

// CreateNote creates a new note in Bear with the specified title, content, and tags
func (c *Client) CreateNote(title, content string, tags []string) (*Note, error) {
	params := map[string]string{
		"title": title,
		"text":  content,
	}

	if len(tags) > 0 {
		params["tags"] = strings.Join(tags, ",")
	}

	return c.callbackCreate(params)
}

// OpenNote opens a note by its ID in Bear
func (c *Client) OpenNote(id string) error {
	params := map[string]string{
		"id": id,
	}

	_, err := c.executeCallback("open-note", params, false)
	return err
}

// UpdateNote replaces the content of a note identified by ID
func (c *Client) UpdateNote(id, content string) error {
	params := map[string]string{
		"id":   id,
		"text": content,
		"mode": "replace",
	}

	_, err := c.executeCallback("add-text", params, false)
	return err
}

// AddText adds text to a note with the specified mode (append, prepend, replace)
func (c *Client) AddText(id, text string, mode string) error {
	if mode != "append" && mode != "prepend" && mode != "replace" {
		mode = "append" // Default to append
	}

	params := map[string]string{
		"id":   id,
		"text": text,
		"mode": mode,
	}

	_, err := c.executeCallback("add-text", params, false)
	return err
}

// GetTags retrieves all tags used in Bear
func (c *Client) GetTags() ([]string, error) {
	if c.DryRun {
		return []string{}, nil
	}

	response, err := c.executeCallback("tags", map[string]string{}, true)
	if err != nil {
		return nil, fmt.Errorf("failed to get tags: %w", err)
	}

	// Extract tags from response
	tagsStr, ok := response["tags"]
	if !ok {
		return []string{}, nil
	}

	return parseTagsString(tagsStr), nil
}

// VerifyBearInstalled checks if Bear app is installed
func (c *Client) VerifyBearInstalled() bool {
	if runtime.GOOS != "darwin" {
		return false
	}

	cmd := exec.Command("osascript", "-e", `tell application "System Events" to return exists application process "Bear"`)
	output, err := cmd.Output()
	if err != nil {
		return false
	}

	return strings.TrimSpace(string(output)) == "true"
}