// models/movie.go
package models

type Movie struct {
    ID          int    `json:"id"`
    Title       string `json:"title"`
    Rating      string `json:"vote_average"`
    PosterPath  string `json:"poster_path"`
}
