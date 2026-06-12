package models

import "strings"

// Movie describes one movie shown by the website and JSON API.
type Movie struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Genre       string  `json:"genre"`
	Year        int     `json:"year"`
	Rating      float64 `json:"rating"`
	Description string  `json:"description"`
}

// Movies is an in-memory data store. It is recreated whenever the server starts.
var Movies = []Movie{
	{ID: 1, Title: "Interstellar", Genre: "Sci-Fi", Year: 2014, Rating: 8.7, Description: "Explorers travel through a wormhole in space to find a new home for humanity."},
	{ID: 2, Title: "Arrival", Genre: "Sci-Fi", Year: 2016, Rating: 7.9, Description: "A linguist works to communicate with mysterious visitors from another world."},
	{ID: 3, Title: "The Martian", Genre: "Sci-Fi", Year: 2015, Rating: 8.0, Description: "An astronaut uses science and determination to survive alone on Mars."},
	{ID: 4, Title: "The Shawshank Redemption", Genre: "Drama", Year: 1994, Rating: 9.3, Description: "Two imprisoned men form a lasting friendship over many years."},
	{ID: 5, Title: "Mad Max: Fury Road", Genre: "Action", Year: 2015, Rating: 8.1, Description: "Survivors race across a desert wasteland while escaping a tyrant."},
	{ID: 6, Title: "The Dark Knight", Genre: "Action", Year: 2008, Rating: 9.0, Description: "Batman faces a criminal mastermind who brings chaos to Gotham."},
}

// FindMovie returns the movie with the given ID.
func FindMovie(id int) (Movie, bool) {
	for _, movie := range Movies {
		if movie.ID == id {
			return movie, true
		}
	}
	return Movie{}, false
}

// MoviesByGenre performs a case-insensitive genre search.
func MoviesByGenre(genre string) []Movie {
	matches := make([]Movie, 0)
	for _, movie := range Movies {
		if strings.EqualFold(movie.Genre, genre) {
			matches = append(matches, movie)
		}
	}
	return matches
}
