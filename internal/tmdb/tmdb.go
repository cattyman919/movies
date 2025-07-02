package tmdb

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	apiBaseURL = "https://api.themoviedb.org/3"
)

// Movie defines the structure for a movie from TMDB
type Movie struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	PosterPath  string `json:"poster_path"`
	Overview    string `json:"overview"`
}

// PopularMoviesResponse defines the structure of the popular movies API response
type PopularMoviesResponse struct {
	Page    int     `json:"page"`
	Results []Movie `json:"results"`
}

// Client is a client for the TMDB API
type Client struct {
	apiKey     string
	httpClient *http.Client
}

// NewClient creates a new TMDB API client
func NewClient(apiKey string) *Client {
	return &Client{
		apiKey:     apiKey,
		httpClient: &http.Client{},
	}
}

// GetPopularMovies fetches popular movies from TMDB
func (c *Client) GetPopularMovies() (*PopularMoviesResponse, error) {
	url := fmt.Sprintf("%s/movie/popular?api_key=%s", apiBaseURL, c.apiKey)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var popularMovies PopularMoviesResponse
	if err := json.NewDecoder(resp.Body).Decode(&popularMovies); err != nil {
		return nil, err
	}

	return &popularMovies, nil
}
