// Package json provides functionality for working with movie data in JSON format.
package json

import (
	"encoding/json"
	"fmt"
	"log"
)

// Movie represents a movie with its title, release year, color status and actors.
// The JSON tags control how the struct fields are marshaled to JSON.
type Movie struct {
	Title  string
	Year   int      `json:"released"`        // Marshaled as "released" in JSON
	Color  bool     `json:"color,omitempty"` // Only included in JSON if true
	Actors []string `json:"actors"`          // Marshaled as "actors" array in JSON
}

// movies is a slice containing sample movie data
var movies = []Movie{
	{
		Title:  "Casablanca",
		Year:   1942,
		Color:  false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"},
	},
	{
		Title:  "Cool Hand Luke",
		Year:   1967,
		Color:  true,
		Actors: []string{"Paul Newman"},
	},
	{
		Title:  "Bullitt",
		Year:   1968,
		Color:  true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"},
	},
}

// PrintMovies marshals the movies slice to JSON format with indentation
// and prints the output to standard output. If marshaling fails, it logs
// a fatal error and exits the program.
func PrintMovies() {
	data, err := json.MarshalIndent(movies, "", "  ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
}
