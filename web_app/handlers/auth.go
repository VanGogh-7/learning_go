package handlers

import (
	"net/http"
	"net/url"
	"strings"

	"web_app/models"
)

// Login shows the form for GET requests and handles its values for POST requests.
func Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		render(w, "login", nil)
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Could not read form", http.StatusBadRequest)
			return
		}

		username := strings.TrimSpace(r.FormValue("username"))
		if username == "" {
			username = "Guest"
		}
		http.Redirect(w, r, "/profile?user="+url.QueryEscape(username), http.StatusSeeOther)
	default:
		w.Header().Set("Allow", "GET, POST")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Profile reads a query parameter and turns it into template data.
func Profile(w http.ResponseWriter, r *http.Request) {
	username := strings.TrimSpace(r.URL.Query().Get("user"))
	if username == "" {
		username = "Guest"
	}

	user := models.User{
		Username:        username,
		PreferredGenres: []string{"Sci-Fi", "Drama", "Action"},
	}
	render(w, "profile", user)
}
