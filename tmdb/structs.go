package tmdb

type MovieResponse struct {
	Page         int     `json:"page"`
	Results      []Movie `json:"results"`
	TotalPages   int     `json:"total_pages"`
	TotalResults int     `json:"total_results"`
}

type Movie struct {
	Adult            bool    `json:"adult"`
	BackdropPath     string  `json:"backdrop_path"`
	GenreIDs         []int   `json:"genre_ids"`
	ID               int     `json:"id"`
	OriginalLanguage string  `json:"original_language"`
	OriginalTitle    string  `json:"original_title"`
	Overview         string  `json:"overview"`
	Popularity       float64 `json:"popularity"`
	PosterPath       string  `json:"poster_path"`
	ReleaseDate      string  `json:"release_date"`
	Title            string  `json:"title"`
	Video            bool    `json:"video"`
	VoteAverage      float64 `json:"vote_average"`
	VoteCount        int     `json:"vote_count"`
}


type Casts struct {
	Adult bool `json:"adult"`
	Gender int `json:"gender"`
	ID int `json:"id"`
	KnownForDepartment string `json:"known_for_department"`
	Name string `json:"name"`
	OriginalName string `json:"original_name"`
	Popularity float64 `json:"popularity"`
	ProfilePath string `json:"profile_path"`
	Castid int `json:"cast_id"`
	Character string `json:"character"`
	Creditid string `json:"credit_id"`
	Order int `json:"order"`
}

type Crew struct {
	Adult bool `json:"adult"`
	Gender int `json:"gender"`
	ID int `json:"id"`
	KnownForDepartment string `json:"known_for_department"`
	Name string `json:"name"`
	OriginalName string `json:"original_name"`
	Popularity float64 `json:"popularity"`
	ProfilePath string `json:"profile_path"`
	Creditid string `json:"credit_id"`
	Department string `json:"department"`
	Job string `json:"job"`	
}


type CastRepsonse struct {
	ID int `json:"id"`
	Casts []Casts `json:"cast"`
	Crews []Crew `json:"crew"`
}