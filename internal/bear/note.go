package bear

import (
	"fmt"
	"strings"
)

// Note represents a Bear note
type Note struct {
	ID       string
	Title    string
	Content  string
	Tags     []string
	Created  string // ISO 8601 date format
	Modified string // ISO 8601 date format
}

// GetNoteByID retrieves a note from Bear by its ID
func (c *Client) GetNoteByID(id string) (*Note, error) {
	if c.DryRun {
		return &Note{
			ID:      id,
			Title:   "Dry Run Note",
			Content: "This is a dry run, not a real note",
			Tags:    []string{},
		}, nil
	}

	params := map[string]string{
		"id": id,
	}

	response, err := c.executeCallback("open-note", params, true)
	if err != nil {
		return nil, fmt.Errorf("failed to get note by ID: %w", err)
	}

	// Extract note data from response
	note := &Note{
		ID:       id,
		Title:    response["title"],
		Content:  response["text"],
		Created:  response["created"],
		Modified: response["modified"],
	}

	// Parse tags if present
	if tags, ok := response["tags"]; ok && tags != "" {
		note.Tags = parseTagsString(tags)
	}

	return note, nil
}

// GetNoteByTitle retrieves a note from Bear by its title
func (c *Client) GetNoteByTitle(title string) (*Note, error) {
	if c.DryRun {
		return &Note{
			ID:      "dry-run-id",
			Title:   title,
			Content: "This is a dry run, not a real note",
			Tags:    []string{},
		}, nil
	}

	// Use search to find the note by title
	params := map[string]string{
		"term":        title,
		"show_window": "yes",
	}

	response, err := c.executeCallback("search", params, true)
	if err != nil {
		return nil, fmt.Errorf("failed to search for note with title '%s': %w", title, err)
	}

	// Check if a note was found
	noteID, ok := response["note_id"]
	if !ok || noteID == "" {
		return nil, fmt.Errorf("no note found with title '%s'", title)
	}

	// Now get the full note by ID
	return c.GetNoteByID(noteID)
}

// NoteExists checks if a note with the given title exists
func (c *Client) NoteExists(title string) (bool, string, error) {
	if c.DryRun {
		return false, "", nil
	}

	if title == "" {
		return false, "", fmt.Errorf("title cannot be empty")
	}

	// Search for the note by title
	params := map[string]string{
		"term":        title,
		"show_window": "no",
	}

	response, err := c.executeCallback("search", params, true)
	if err != nil {
		return false, "", fmt.Errorf("failed to search for note: %w", err)
	}

	// Check if a note was found
	noteID, ok := response["note_id"]
	if !ok || noteID == "" {
		return false, "", nil
	}

	return true, noteID, nil
}

// Helper function to parse tags string
func parseTagsString(tagsStr string) []string {
	// Bear returns tags as a comma-separated string
	if tagsStr == "" {
		return []string{}
	}
	return splitAndTrim(tagsStr, ",")
}

// Helper function to split a string and trim whitespace
func splitAndTrim(s, sep string) []string {
	if s == "" {
		return []string{}
	}

	parts := strings.Split(s, sep)
	trimmed := make([]string, 0, len(parts))

	for _, part := range parts {
		trimmed = append(trimmed, strings.TrimSpace(part))
	}

	return trimmed
}