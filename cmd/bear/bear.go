package bear

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/varunshoor/noteworthy/internal/bear"
)

var (
	title   string
	content string
	tags    string
	id      string
	mode    string
	dryRun  bool
)

// BearCmd is the command for interacting with Bear
var BearCmd = &cobra.Command{
	Use:   "bear",
	Short: "Interact with Bear app",
	Long:  `Perform operations on Bear notes such as create, open, update, etc.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Skip verification for the test command
		if cmd.Name() == "test" {
			return
		}

		client := bear.NewClient(dryRun)
		if !client.VerifyBearInstalled() {
			fmt.Println("Bear app is not installed or not running.")
			fmt.Println("Please install Bear from https://bear.app/ or launch it if already installed.")
		}
	},
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new note in Bear",
	Run: func(cmd *cobra.Command, args []string) {
		client := bear.NewClient(dryRun)
		
		tagsList := []string{}
		if tags != "" {
			tagsList = strings.Split(tags, ",")
		}
		
		note, err := client.CreateNote(title, content, tagsList)
		if err != nil {
			fmt.Printf("Error creating note: %v\n", err)
			return
		}
		
		fmt.Printf("Created note with ID: %s\n", note.ID)
		fmt.Printf("Title: %s\n", note.Title)
		if len(note.Tags) > 0 {
			fmt.Printf("Tags: %s\n", strings.Join(note.Tags, ", "))
		}
	},
}

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open a note in Bear",
	Run: func(cmd *cobra.Command, args []string) {
		client := bear.NewClient(dryRun)
		
		err := client.OpenNote(id)
		if err != nil {
			fmt.Printf("Error opening note: %v\n", err)
			return
		}
		
		fmt.Println("Note opened successfully")
	},
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a note in Bear",
	Run: func(cmd *cobra.Command, args []string) {
		client := bear.NewClient(dryRun)
		
		err := client.UpdateNote(id, content)
		if err != nil {
			fmt.Printf("Error updating note: %v\n", err)
			return
		}
		
		fmt.Println("Note updated successfully")
	},
}

var addTextCmd = &cobra.Command{
	Use:   "add-text",
	Short: "Add text to a note in Bear",
	Run: func(cmd *cobra.Command, args []string) {
		client := bear.NewClient(dryRun)
		
		err := client.AddText(id, content, mode)
		if err != nil {
			fmt.Printf("Error adding text to note: %v\n", err)
			return
		}
		
		fmt.Println("Text added to note successfully")
	},
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a note from Bear by ID",
	Run: func(cmd *cobra.Command, args []string) {
		client := bear.NewClient(dryRun)
		
		note, err := client.GetNoteByID(id)
		if err != nil {
			fmt.Printf("Error getting note: %v\n", err)
			return
		}
		
		fmt.Printf("Title: %s\n", note.Title)
		fmt.Printf("Content: %s\n", note.Content)
		if len(note.Tags) > 0 {
			fmt.Printf("Tags: %s\n", strings.Join(note.Tags, ", "))
		}
		if note.Created != "" {
			fmt.Printf("Created: %s\n", note.Created)
		}
		if note.Modified != "" {
			fmt.Printf("Modified: %s\n", note.Modified)
		}
	},
}

var getByTitleCmd = &cobra.Command{
	Use:   "get-by-title",
	Short: "Get a note from Bear by title",
	Run: func(cmd *cobra.Command, args []string) {
		client := bear.NewClient(dryRun)
		
		note, err := client.GetNoteByTitle(title)
		if err != nil {
			fmt.Printf("Error getting note by title: %v\n", err)
			return
		}
		
		fmt.Printf("ID: %s\n", note.ID)
		fmt.Printf("Title: %s\n", note.Title)
		fmt.Printf("Content: %s\n", note.Content)
		if len(note.Tags) > 0 {
			fmt.Printf("Tags: %s\n", strings.Join(note.Tags, ", "))
		}
		if note.Created != "" {
			fmt.Printf("Created: %s\n", note.Created)
		}
		if note.Modified != "" {
			fmt.Printf("Modified: %s\n", note.Modified)
		}
	},
}

var existsCmd = &cobra.Command{
	Use:   "exists",
	Short: "Check if a note exists in Bear",
	Run: func(cmd *cobra.Command, args []string) {
		client := bear.NewClient(dryRun)
		
		exists, noteID, err := client.NoteExists(title)
		if err != nil {
			fmt.Printf("Error checking if note exists: %v\n", err)
			return
		}
		
		if exists {
			fmt.Printf("Note exists with ID: %s\n", noteID)
		} else {
			fmt.Println("Note does not exist")
		}
	},
}

var tagsCmd = &cobra.Command{
	Use:   "tags",
	Short: "Get all tags from Bear",
	Run: func(cmd *cobra.Command, args []string) {
		client := bear.NewClient(dryRun)
		
		tagsList, err := client.GetTags()
		if err != nil {
			fmt.Printf("Error getting tags: %v\n", err)
			return
		}
		
		if len(tagsList) == 0 {
			fmt.Println("No tags found")
			return
		}
		
		fmt.Println("Tags:", strings.Join(tagsList, ", "))
	},
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Test Bear client functionality",
	Run: func(cmd *cobra.Command, args []string) {
		client := bear.NewClient(dryRun)
		
		fmt.Println("Testing Bear client...")
		installed := client.VerifyBearInstalled()
		fmt.Println("Bear installed:", installed)
		
		if !installed && !dryRun {
			fmt.Println("Bear app is not installed or not running.")
			fmt.Println("Please install Bear from https://bear.app/ or launch it if already installed.")
		}
	},
}

func init() {
	// Add commands to BearCmd
	BearCmd.AddCommand(createCmd)
	BearCmd.AddCommand(openCmd)
	BearCmd.AddCommand(updateCmd)
	BearCmd.AddCommand(addTextCmd)
	BearCmd.AddCommand(getCmd)
	BearCmd.AddCommand(getByTitleCmd)
	BearCmd.AddCommand(existsCmd)
	BearCmd.AddCommand(tagsCmd)
	BearCmd.AddCommand(testCmd)
	
	// Add global flag
	BearCmd.PersistentFlags().BoolVar(&dryRun, "dry-run", false, "Print operations without executing them")
	
	// Add command-specific flags
	createCmd.Flags().StringVar(&title, "title", "", "Title of the note")
	createCmd.Flags().StringVar(&content, "content", "", "Content of the note")
	createCmd.Flags().StringVar(&tags, "tags", "", "Comma-separated list of tags")
	createCmd.MarkFlagRequired("title")
	
	openCmd.Flags().StringVar(&id, "id", "", "ID of the note")
	openCmd.MarkFlagRequired("id")
	
	updateCmd.Flags().StringVar(&id, "id", "", "ID of the note")
	updateCmd.Flags().StringVar(&content, "content", "", "New content for the note")
	updateCmd.MarkFlagRequired("id")
	updateCmd.MarkFlagRequired("content")
	
	addTextCmd.Flags().StringVar(&id, "id", "", "ID of the note")
	addTextCmd.Flags().StringVar(&content, "content", "", "Text to add to the note")
	addTextCmd.Flags().StringVar(&mode, "mode", "append", "Mode (append, prepend, replace)")
	addTextCmd.MarkFlagRequired("id")
	addTextCmd.MarkFlagRequired("content")
	
	getCmd.Flags().StringVar(&id, "id", "", "ID of the note")
	getCmd.MarkFlagRequired("id")
	
	getByTitleCmd.Flags().StringVar(&title, "title", "", "Title of the note")
	getByTitleCmd.MarkFlagRequired("title")
	
	existsCmd.Flags().StringVar(&title, "title", "", "Title of the note")
	existsCmd.MarkFlagRequired("title")
}