package handlers

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var pageTemplates = loadTemplates()

// loadTemplates combines the shared layout with each individual page.
func loadTemplates() map[string]*template.Template {
	pages := []string{"home", "login", "movies", "movie_detail", "profile"}
	templates := make(map[string]*template.Template, len(pages))

	for _, page := range pages {
		files := []string{
			filepath.Join("templates", "layout.html"),
			filepath.Join("templates", page+".html"),
		}
		templates[page] = template.Must(template.ParseFiles(files...))
	}
	return templates
}

// render writes one HTML template as the HTTP response.
func render(w http.ResponseWriter, page string, data any) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := pageTemplates[page].ExecuteTemplate(w, "layout", data); err != nil {
		log.Printf("render %s: %v", page, err)
		http.Error(w, "Could not render page", http.StatusInternalServerError)
	}
}

// Home handles only the exact "/" path.
func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	render(w, "home", nil)
}
