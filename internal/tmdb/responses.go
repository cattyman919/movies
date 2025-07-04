package tmdb

type generalTMDBResponse struct {
	Success bool `json:"success"`
}

type ErrorTMDBResponse struct {
	HTTP_Status    int    `json:"-"`
	Status_Code    int    `json:"status_code"`
	Status_Message string `json:"status_message"`
}

func (e *ErrorTMDBResponse) Error() string {
	return e.Status_Message
}

// PopularMoviesResponse defines the structure of the popular movies API response
type PopularMoviesResponse struct {
	Page    int     `json:"page"`
	Results []Movie `json:"results"`
}

// NowPlayingResponse Response
type NowPlayingResponse struct {
	Dates struct {
		Maximum string `json:"maximum"`
		Minimum string `json:"minimum"`
	}
	Page    int `json:"page"`
	Results []struct {
		Adult            bool    `json:"adult"`
		Backdrop_path    string  `json:"backdrop_path"`
		Genre_ids        []int   `json:"genre_ids"`
		Id               int     `json:"id"`
		OriginalLanguage string  `json:"original_language"`
		OriginalTitle    string  `json:"original_title"`
		Overview         string  `json:"overview"`
		Popularity       float32 `json:"popularity"`
		PosterPath       string  `json:"poster_path"`
		ReleaseDate      string  `json:"release_date"`
		Title            string  `json:"title"`
		Video            bool    `json:"video"`
		VoteAverage      float32 `json:"vote_average"`
		VoteCount        int     `json:"vote_count"`
	} `json:"results"`
	Total_pages   int `json:"total_pages"`
	Total_results int `json:"total_results"`
}

// NewGuessSessionResponse creates a new guest session id
type NewGuestSessionResponse struct {
	generalTMDBResponse
	Guest_Session_Id string `json:"guest_session_id"`
	Expires_At       string `json:"expires_at"`
}
