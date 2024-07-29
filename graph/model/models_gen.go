// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Movie struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	URL         string `json:"url"`
	ReleaseDate string `json:"releaseDate"`
}

type NewMovie struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

type UpdateMovie struct {
	ID    string  `json:"id"`
	Title *string `json:"title,omitempty"`
	URL   *string `json:"url,omitempty"`
}
