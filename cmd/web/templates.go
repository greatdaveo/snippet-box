package main

import (
	"html/template"
	"path/filepath"
	"snippet-box/pkg/forms"
	"snippet-box/pkg/models"
	"time"
)

// To set the holding structure for any dynamic data to be passed to HTML templates
type templateData struct {
	CSRFToken         string          // For CSRF
	AuthenticatedUser *models.User    // To pass the user details value from the request context
	CurrentYear       int             // Field for Current Year
	Flash             string          // Flash field for the flash confirmation message
	Form              *forms.Form     // Pointer to single form field
	Snippet           *models.Snippet // A pointer to a single Snippet from models package
	// To include a Snippets field in the templateData struct
	Snippets []*models.Snippet // A slice of Snippet pointers, holding multiple snippets

}

// Human Date Function
func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

// To create an in memory map to cache the templates
func newTemplateCache(dir string) (map[string]*template.Template, error) {
	// To initialize a new map to act as the cache || empty map to store parsed templates
	cache := map[string]*template.Template{}
	// To get a slice of all filepaths with the extension ".page.tmpl"
	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}

	// Loop through the pages one by one
	for _, page := range pages {
		// extract each file name, from the full file path & assign it to the name variable
		name := filepath.Base(page)
		// Parse the page template in to a template set
		ts, err := template.ParseFiles(page)
		if err != nil {
			return nil, err
		}
		// To add any layout template to the template set (e.g "base" layout)
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}

		// To add any partial templates to the template set (e.g footer)
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}

		// for k := range cache {
		// 	fmt.Println("Template in cache:", k)
		// }

		// To add the template set to the cache, using the name of the page e.g home.page.tmpl as the key
		cache[name] = ts

		// To know the cached templates
		// fmt.Println("Template in cache:", name)

	}

	return cache, nil

}
