package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"web_app/models"
)

type movieDetailData struct {
	Movie           models.Movie
	Recommendations []models.Movie
}

// Movies sends all in-memory movies to the movies template.
func Movies(w http.ResponseWriter, r *http.Request) {
	render(w, "movies", models.Movies)
}

// MovieDetail reads an ID from /movie?id=1 and displays one movie.
func MovieDetail(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		return
	}

	movie, found := models.FindMovie(id)
	if !found {
		http.NotFound(w, r)
		return
	}

	recommendations := make([]models.Movie, 0)
	for _, recommendation := range models.MoviesByGenre(movie.Genre) {
		if recommendation.ID != movie.ID {
			recommendations = append(recommendations, recommendation)
		}
	}

	render(w, "movie_detail", movieDetailData{
		Movie:           movie,
		Recommendations: recommendations,
	})
}

// RecommendationsAPI encodes Go values as JSON for frontend JavaScript.
func RecommendationsAPI(w http.ResponseWriter, r *http.Request) {
	genre := strings.TrimSpace(r.URL.Query().Get("genre"))
	if genre == "" {
		http.Error(w, "The genre query parameter is required", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(models.MoviesByGenre(genre)); err != nil {
		http.Error(w, "Could not encode recommendations", http.StatusInternalServerError)
	}
}
