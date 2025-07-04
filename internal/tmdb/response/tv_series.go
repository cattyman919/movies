package response

type TVSeriesResultResponse struct {
	paginationResponse
	Results []struct {
		Backdrop_path    string   `json:"backdrop_path"`
		FirstAirDate     string   `json:"first_air_date"`
		Genre_ids        []int    `json:"genre_ids"`
		Id               int      `json:"id"`
		Name             string   `json:"name"`
		OriginCountry    []string `json:"origin_country"`
		OriginalLanguage string   `json:"original_language"`
		OriginalTitle    string   `json:"original_title"`
		Overview         string   `json:"overview"`
		Popularity       float32  `json:"popularity"`
		PosterPath       string   `json:"poster_path"`
		VoteAverage      float32  `json:"vote_average"`
		VoteCount        int      `json:"vote_count"`
	} `json:"results"`
}

// TMDB Response For "/tv/airing_today"
type AiringToday_TVSeries_Response = TVSeriesResultResponse

// TMDB Response For "/tv/on_the_air"
type OnTheAir_TVSeries_Response = TVSeriesResultResponse

// TMDB Response For "/tv/popular"
type Popular_TVSeries_Response = TVSeriesResultResponse

// TMDB Response For "/tv/top_rated"
type TopRated_TVSeries_Response = TVSeriesResultResponse

// TMDB Response For "/trending/tv/{time_window}"
type Trending_TVSeries_Response = TVSeriesResultResponse
