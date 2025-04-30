package templates

import (
	"embed"
	"io/fs"
	"text/template"
)

//go:embed files/*
var templateFS embed.FS

// Templates holds all parsed templates
type Templates struct {
	Year         *template.Template
	Quarter      *template.Template
	Month        *template.Template
	Day          *template.Template
	EventSection *template.Template
}

// Load parses all template files and returns a Templates struct
func Load() (*Templates, error) {
	templates := &Templates{}

	// Parse templates
	var err error

	templates.Year, err = template.ParseFS(templateFS, "files/year.tmpl")
	if err != nil {
		return nil, err
	}

	templates.Quarter, err = template.ParseFS(templateFS, "files/quarter.tmpl")
	if err != nil {
		return nil, err
	}

	templates.Month, err = template.ParseFS(templateFS, "files/month.tmpl")
	if err != nil {
		return nil, err
	}

	templates.Day, err = template.ParseFS(templateFS, "files/day.tmpl")
	if err != nil {
		return nil, err
	}

	templates.EventSection, err = template.ParseFS(templateFS, "files/event_section.tmpl")
	if err != nil {
		return nil, err
	}

	return templates, nil
}

// GetTemplateFiles returns a list of all template filenames
func GetTemplateFiles() ([]string, error) {
	files, err := fs.Glob(templateFS, "files/*.tmpl")
	if err != nil {
		return nil, err
	}
	return files, nil
}

