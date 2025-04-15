// utils/api.go
package utils

import (
    "encoding/json"
    "fmt"
    "net/http"
    "os"
)

const API_URL = "https://api.themoviedb.org/3"
var API_KEY = os.Getenv("TMDB_API_KEY")

type Movie struct {
    ID         int     `json:"id"`
    Title      string  `json:"title"`
    Rating     float64 `json:"vote_average"` // Changed to float64
    PosterPath string  `json:"poster_path"`
}

type MovieResponse struct {
    Results    []Movie `json:"results"`
    TotalPages int     `json:"total_pages"`
    TotalItems int     `json:"total_results"`
}

// FetchPopularMovies fetches movies based on the page, query, genre, and actors.
func FetchPopularMovies(page int, query string, genre string, actors string) (MovieResponse, error) {
    var url string
    if query != "" {
        url = fmt.Sprintf("%s/search/movie?api_key=%s&query=%s&page=%d", API_URL, API_KEY, query, page)
    } else {
        url = fmt.Sprintf("%s/discover/movie?api_key=%s&page=%d", API_URL, API_KEY, page)

        if genre != "" {
            url = fmt.Sprintf("%s&with_genres=%s", url, genre)
        }
        if actors != "" {
            url = fmt.Sprintf("%s&with_cast=%s", url, actors)
        }
    }
    
    res, err := http.Get(url)
    if err != nil {
        return MovieResponse{}, err
    }
    defer res.Body.Close()

    var result MovieResponse
    if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
        return MovieResponse{}, err
    }

    return result, nil
}
