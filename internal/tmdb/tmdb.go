package tmdb

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	BASE_URL              = "https://api.themoviedb.org/3"
	NEW_GUEST_SESSION_URL = BASE_URL + "/authentication/guest_session/new"
	NOW_PLAYING_URL       = BASE_URL + "/movie/now_playing"
)

// Movie defines the structure for a movie from TMDB
type Movie struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	PosterPath string `json:"poster_path"`
	Overview   string `json:"overview"`
}

// Client is a client for the TMDB API
type Client struct {
	apiKey         string
	guessSessionId string
	httpClient     *http.Client
}

// NewClient creates a new TMDB API client
func NewClient(apiKey string) *Client {
	return &Client{
		apiKey:         apiKey,
		guessSessionId: "",
		httpClient:     &http.Client{},
	}
}

func (c *Client) do(req *http.Request, v interface{}) error {
	resp, err := c.httpClient.Do(req)
	log.Printf("Request URL to: %s\n", req.URL)
	log.Printf("Response Status: %d\n", resp.StatusCode)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errResp ErrorTMDBResponse
		if err := json.NewDecoder(resp.Body).Decode(&errResp); err != nil {
			return fmt.Errorf("api error: status %d", resp.StatusCode)
		}
		errResp.HTTP_Status = resp.StatusCode
		return &errResp
	}

	return json.NewDecoder(resp.Body).Decode(v)
}

func (c *Client) ReqAuth(method, url string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+c.apiKey)
	return req, nil
}

func (c *Client) NowPlaying(context *gin.Context) {
	req, err := c.ReqAuth("GET", NOW_PLAYING_URL)
	if err != nil {
		context.Error(err)
		return
	}

	var nowPlaying NowPlayingResponse
	if err := c.do(req, &nowPlaying); err != nil {
		context.Error(err)
		return
	}

	context.JSON(http.StatusOK, nowPlaying)
}

func (c *Client) GetGuestSessionId(context *gin.Context) {
	req, err := c.ReqAuth("GET", NEW_GUEST_SESSION_URL)
	if err != nil {
		context.Error(err)
		return
	}

	var guestSession NewGuestSessionResponse
	if err := c.do(req, &guestSession); err != nil {
		context.Error(err)
		return
	}

	context.JSON(http.StatusOK, guestSession)
}
