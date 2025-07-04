package tmdb

import (
	"errors"
	"github.com/gin-gonic/gin"
	"movies/internal/tmdb/response"
	"net/http"
	"strconv"
)

const (
	BASE_URL_MOVIE = BASE_URL + "/movie"

	// Movies List
	NOW_PLAYING_MOVIES_URL string = BASE_URL_MOVIE + "/now_playing"
	POPULAR_MOVIES_URL     string = BASE_URL_MOVIE + "/popular"
	TOP_RATED_MOVIES_URL   string = BASE_URL_MOVIE + "/top_rated"
	UPCOMING_MOVIES_URL    string = BASE_URL_MOVIE + "/upcoming"
)

func (c *Client) NowPlaying(context *gin.Context) {

	var nowPlaying response.NowPlaying_Movies_Response

	if err := c.do("GET", NOW_PLAYING_MOVIES_URL, &nowPlaying); err != nil {
		context.Error(err)
		return
	}

	context.JSON(http.StatusOK, nowPlaying)
}

func (c *Client) Popular_Movies(context *gin.Context) {

	var popular_movies response.Popular_Movies_Response

	if err := c.do("GET", POPULAR_MOVIES_URL, &popular_movies); err != nil {
		context.Error(err)
		return
	}

	context.JSON(http.StatusOK, popular_movies)
}

func (c *Client) Top_Rated_Movies(context *gin.Context) {

	var top_rated_movies response.TopRated_Movies_Response

	if err := c.do("GET", TOP_RATED_MOVIES_URL, &top_rated_movies); err != nil {
		context.Error(err)
		return
	}

	context.JSON(http.StatusOK, top_rated_movies)
}

func (c *Client) Upcoming_Movies(context *gin.Context) {

	var upcoming_movies response.TopRated_Movies_Response

	if err := c.do("GET", UPCOMING_MOVIES_URL, &upcoming_movies); err != nil {
		context.Error(err)
		return
	}

	context.JSON(http.StatusOK, upcoming_movies)
}

func (c *Client) Movie_Detail(context *gin.Context) {
	movie_id := context.Param("movie_id")
	if movie_id == "" {
		context.Error(errors.New("Must provide movie_id"))
		return
	}
	_, err := strconv.Atoi(movie_id)
	if err != nil {
		context.Error(err)
		return
	}

	language := context.DefaultQuery("language", "en_US")

	URL := BASE_URL_MOVIE + "/" + movie_id + "?language=" + language

	var movie_detail response.Movie_Detail_Response
	if err := c.do("GET", URL, &movie_detail); err != nil {
		context.Error(err)
		return
	}

	context.JSON(http.StatusOK, movie_detail)
}
