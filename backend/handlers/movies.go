// handlers/movies.go
package handlers

import (
    "encoding/json"
    "fmt"
    "net/http"
    "os"
	"strconv"
    "github.com/gorilla/mux"
    "movie-info-service/utils"
)

// GetMovieByExternalID handles the request to get a specific movie by external ID (e.g., IMDb ID)
func GetMovieByExternalID(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    externalID := params["external_id"]

    if externalID == "" {
        http.Error(w, "External ID is required", http.StatusBadRequest)
        return
    }

    // Build the URL to query the TMDb API with the external ID
    url := fmt.Sprintf("https://api.themoviedb.org/3/movie/%s?api_key=%s", externalID, os.Getenv("TMDB_API_KEY"))

    // Fetch the movie details from the TMDb API
    res, err := http.Get(url)
    if err != nil {
        http.Error(w, "Unable to fetch movie details: "+err.Error(), http.StatusInternalServerError)
        return
    }
    defer res.Body.Close()

    if res.StatusCode != http.StatusOK {
        http.Error(w, "Movie not found", res.StatusCode)
        return
    }

    var movieDetails map[string]interface{}
    if err := json.NewDecoder(res.Body).Decode(&movieDetails); err != nil {
        http.Error(w, "Unable to parse response: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(movieDetails)
}

// ListMovies handles listing movies with optional search query and genre filtering
func ListMovies(w http.ResponseWriter, r *http.Request) {
    pageParam := r.URL.Query().Get("page")
    page, err := strconv.Atoi(pageParam)
    if err != nil || page < 1 {
        page = 1
    }

    query := r.URL.Query().Get("query")  // Get the search query from the URL parameters
    genre := r.URL.Query().Get("genre")  // Get the genre filter from the URL parameters
    actors := r.URL.Query().Get("actors") // Get the actor filter from the URL parameters

    movieResponse, err := utils.FetchPopularMovies(page, query, genre, actors)
    if err != nil {
        http.Error(w, "Unable to fetch movies: "+err.Error(), http.StatusInternalServerError)
        return
    }

    response := map[string]interface{}{
        "movies":      movieResponse.Results,
        "total_pages": movieResponse.TotalPages,
        "total_items": movieResponse.TotalItems,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func GetGenres(w http.ResponseWriter, r *http.Request) {
    url := fmt.Sprintf("https://api.themoviedb.org/3/genre/movie/list?api_key=%s&language=en-US", os.Getenv("TMDB_API_KEY"))
    res, err := http.Get(url)
    if err != nil {
        http.Error(w, "Unable to fetch genres: "+err.Error(), http.StatusInternalServerError)
        return
    }
    defer res.Body.Close()

    if res.StatusCode != http.StatusOK {
        http.Error(w, "Failed to fetch genres from TMDb API", res.StatusCode)
        return
    }

    var genresResponse map[string]interface{}
    if err := json.NewDecoder(res.Body).Decode(&genresResponse); err != nil {
        http.Error(w, "Unable to parse response: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(genresResponse)
}

// GetMovie handles the request to get a specific movie by ID
func GetMovie(w http.ResponseWriter, r *http.Request) {
    // Placeholder response
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"message": "Details of a specific movie"})
}

// SearchMovies handles the request to search for movies
func SearchMovies(w http.ResponseWriter, r *http.Request) {
    // Placeholder response
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"message": "Search results"})
}

// FilterMovies handles the request to filter movies
func FilterMovies(w http.ResponseWriter, r *http.Request) {
    // Placeholder response
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"message": "Filtered movies"})
}

// SearchActors handles the request to search for actors
func SearchActors(w http.ResponseWriter, r *http.Request) {
    query := r.URL.Query().Get("query")
    if query == "" {
        http.Error(w, "Query parameter is required", http.StatusBadRequest)
        return
    }

    url := fmt.Sprintf("https://api.themoviedb.org/3/search/person?api_key=%s&query=%s", os.Getenv("TMDB_API_KEY"), query)
    res, err := http.Get(url)
    if err != nil {
        http.Error(w, "Unable to fetch actors: "+err.Error(), http.StatusInternalServerError)
        return
    }
    defer res.Body.Close()

    if res.StatusCode != http.StatusOK {
        http.Error(w, "Failed to fetch actors from TMDb API", res.StatusCode)
        return
    }

    var actorsResponse map[string]interface{}
    if err := json.NewDecoder(res.Body).Decode(&actorsResponse); err != nil {
        http.Error(w, "Unable to parse response: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(actorsResponse)
}