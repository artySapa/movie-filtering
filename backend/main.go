// main.go
package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "movie-info-service/handlers"
    "movie-info-service/middlewares"
)

func main() {
    r := mux.NewRouter()

    // Define the routes
    r.HandleFunc("/movies", handlers.ListMovies).Methods("GET")       // List all movies
    r.HandleFunc("/find/{external_id}", handlers.GetMovieByExternalID).Methods("GET")  // Get details of a specific movie by external ID
    r.HandleFunc("/search", handlers.SearchMovies).Methods("GET")     // Search movies
    r.HandleFunc("/filter", handlers.FilterMovies).Methods("GET")     // Filter movies
    r.HandleFunc("/genres", handlers.GetGenres).Methods("GET")        // Get all genres
	r.HandleFunc("/actors", handlers.SearchActors).Methods("GET")     // Search for actors

    // Wrap the handlers with the CORS middleware
    corsRouter := middlewares.CORS(r)

    // Start the server on port 8080
    log.Println("Server starting on :8080...")
    log.Fatal(http.ListenAndServe(":8080", corsRouter))
}
