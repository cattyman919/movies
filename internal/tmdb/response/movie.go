package response

// === Movies List ===
type MovieResultResponse struct {
	paginationResponse
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
}

type MovieResultDateResponse struct {
	paginationResponse
	MovieResultResponse
	Dates struct {
		Maximum string `json:"maximum"`
		Minimum string `json:"minimum"`
	}
}

// TMDB Response For "/movie/now_playing"
type NowPlaying_Movies_Response = MovieResultDateResponse

// TMDB Response For "/movie/popular"
type Popular_Movies_Response = MovieResultResponse

// TMDB Response For "/movie/top_rated"
type TopRated_Movies_Response = MovieResultResponse

// TMDB Response For "/movie/upcoming"
type Upcoming_Movies_Response = MovieResultDateResponse

// TMDB Response For "trending/movie/{time_window}"
type Trending_Movies_Response = MovieResultResponse

// Movies Detail
// TMDB Response For "movie/{movie_id}"
type Movie_Detail_Response struct {
	Adult         bool   `json:"adult"`
	Backdrop_path string `json:"backdrop_path"`
	Genre_ids     []struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"genre_ids"`
	Id                  int     `json:"id"`
	Imdb_Id             string  `json:"imdb_id"`
	OriginalLanguage    string  `json:"original_language"`
	OriginalTitle       string  `json:"original_title"`
	Overview            string  `json:"overview"`
	Popularity          float32 `json:"popularity"`
	PosterPath          string  `json:"poster_path"`
	ProductionCompanies []struct {
		Id             int    `json:"id"`
		Logo_path      string `json:"logo_path"`
		Name           string `json:"name"`
		Origin_country string `json:"origin_country"`
	} `json:"production_companies"`
	ProductionCountries []struct {
		Iso_3166_1 string `json:"iso_3166_1"`
		Name       string `json:"name"`
	} `json:"production_countries"`
	ReleaseDate      string `json:"release_date"`
	Revenue          int    `json:"revenue"`
	Runtime          int    `json:"runtime"`
	Spoken_Languages []struct {
		English_Name string `json:"english_name"`
		Iso_639_1    string `json:"iso_639_1"`
		Name         string `json:"name"`
	} `json:"spoken_languages"`
	Status      string  `json:"status"`
	Tagline     string  `json:"tagline"`
	Title       string  `json:"title"`
	Video       bool    `json:"video"`
	VoteAverage float32 `json:"vote_average"`
	VoteCount   int     `json:"vote_count"`
}

// Movies Image Object
type Movie_Image_Object struct {
	Aspect_Ratio int     `json:"aspect_ratio"`
	Height       int     `json:"height"`
	Width        int     `json:"width"`
	Iso_639_1    string  `json:"iso_639_1"`
	File_path    string  `json:"file_path"`
	VoteAverage  float32 `json:"vote_average"`
	VoteCount    int     `json:"vote_count"`
}

// Movies Images
// TMDB Response For "movie/{movie_id}/images"
type Movie_Images_Response struct {
	Id        int                  `json:"id"`
	Backdrops []Movie_Image_Object `json:"backdrops"`
	Logos     []Movie_Image_Object `json:"logos"`
	Posters   []Movie_Image_Object `json:"posters"`
}

// Movies Video
// TMDB Response For "movie/{movie_id}/videos"
type Movie_Videos_Response struct {
	Id      int `json:"id"`
	Results []struct {
		Iso_639_1    string `json:"iso_639_1"`
		Iso_3166_1   string `json:"iso_3166_1"`
		Name         string `json:"name"`
		Key          string `json:"key"`
		Site         string `json:"site"`
		Size         int    `json:"size"`
		Type         string `json:"type"`
		Official     bool   `json:"official"`
		Published_at string `json:"published_at"`
		Id           string `json:"id"`
	} `json:"results"`
}
