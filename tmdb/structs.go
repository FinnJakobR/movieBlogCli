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


type MovieDetailGenres struct {
	ID int `json:"id"`
	Name string `json:"name"`
}


type MovieDetailCompanie struct {
	ID int `json:"id"`
	LogoPath string `json:"logo_path"`
	Name string `json:"name"`
	OriginCountry string `json:"origin_country"`
}

type MovieDetailProductionCountry struct {
	Iso31661 string `json:"iso_3166_1"`
	Name string `json:"name"`
}

type MovieDetailMovieLanguage struct {
	EnglishName string `json:"english_name"`
	Iso6391 string `json:"iso_639_1"`
	Name string `json:"name"`
}

type MovieDetailCollection struct {
	ID int `json:"id"`
	Name string `json:"name"`
	PosterPath string `json:"poster_path"`
	BackdropPath string `json:"backdrop_path"`
}

type MovieDetailResponse struct {
	Adult bool `json:"adult"`
	BackdropPath string `json:"backdrop_path"`
	BelongsToCollection MovieDetailCollection `json:"belongs_to_collection"`
	Budget int `json:"budget"`
	Genres  []MovieDetailGenres `json:"genres"`
	Homepage string `json:"homepage"`
	ID int `json:"id"`
	ImdbID string `json:"imdb_id"`
	OriginalLanguage string `json:"original_language"`
	OriginalTitle string `json:"original_title"`
	Overview string `json:"overview"`
	Popularity float64 `json:"popularity"`
	PosterPath string `json:"poster_path"`
	ProductionCompanies []MovieDetailCompanie `json:"production_companies"`
	ProductionCountries []MovieDetailProductionCountry `json:"production_countries"`
	ReleaseDate string `json:"release_date"`
	Revenue int `json:"revenue"`
	Runtime int `json:"runtime"`
	SpokenLanguages []MovieDetailMovieLanguage `json:"spoken_languages"`
	Status string `json:"status"`
	Tagline string `json:"tagline"`
	Title string `json:"title"`
	Video bool `json:"video"`
	VoteAverage float64 `json:"vote_average"`
	VoteCount float64 `json:"vote_count"`
}