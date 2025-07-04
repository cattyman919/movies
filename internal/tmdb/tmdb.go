package tmdb

import (
	"encoding/json"
	"fmt"
	"log"
	"movies/internal/tmdb/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	BASE_URL              string = "https://api.themoviedb.org/3"
	NEW_GUEST_SESSION_URL string = BASE_URL + "/authentication/guest_session/new"
)

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

func (c *Client) do(method, url string, v interface{}) error {
	req, err := c.ReqAuth(method, url)
	if err != nil {
		return err
	}

	resp, err := c.httpClient.Do(req)
	log.Printf("Request URL to: %s\n", req.URL)
	log.Printf("Response Status: %d\n", resp.StatusCode)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errResp response.ErrorTMDBResponse
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

	// Asumes the Api Key is not empty
	req.Header.Add("Authorization", "Bearer "+c.apiKey)
	return req, nil
}

func (c *Client) GetGuestSessionId(context *gin.Context) {
	var guestSession response.NewGuestSessionResponse
	if err := c.do("GET", NEW_GUEST_SESSION_URL, &guestSession); err != nil {
		context.Error(err)
		return
	}

	context.JSON(http.StatusOK, guestSession)
}
