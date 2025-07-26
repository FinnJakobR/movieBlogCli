package main

import (
	"encoding/json"
	"fmt"
	"log"
	"movieBlog/cli/cliParser"
	"movieBlog/cli/env"
	parsermovie "movieBlog/cli/parser_movie"
	"movieBlog/cli/tmdb"
	"movieBlog/cli/util"
	"os"
	p "path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
);

type MovieBlogMovieImage struct {
	Data string `json:"data"`
	Name string `json:"name"`
}

type MovieBlogMovie struct {
	Title string `json:"title"`
	Release string `json:"release"`
	Img MovieBlogMovieImage `json:"img"`
	Runtime int `json:"runtime"`
	Link string `json:"link"`
	Rating int `json:"rating"`
}

type MovieBlogArticle struct {
	Title string `json:"title"`
	Body string `json:"body"`
	Overview string `json:"overview"`
	Author string `json:"author"`
}

type MovieBlogSticker struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Type int `json:"type"`
}

type MovieBlogPersons struct {
	Firstname string `json:"firstname"`
	Surname string `json:"surname"`
	Job string `json:"job"`
}


type MovieBlog struct {
	Movie MovieBlogMovie `json:"movie"`
	Article MovieBlogArticle `json:"article"`
	Stickers []MovieBlogSticker `json:"stickers"`
	Persons  []MovieBlogPersons `json:"persons"`
}



// Funktion, um '~' im Pfad durch echten Home-Ordner zu ersetzen
func expandPath(path string) string {
    if strings.HasPrefix(path, "~/") {
        home, err := os.UserHomeDir()
        if err != nil {
            return path // falls Fehler, Pfad so lassen
        }
        return filepath.Join(home, path[2:])
    }
    return path
}

func main() {
   path, verbose, flag_err := cliParser.ParseArgs();
   expandedPath := expandPath("~/configs/.movie_blog_env");
   env.Read_env_file(expandedPath);

   if(flag_err != nil) {
		log.Fatal(flag_err.Error());
	}
	

	content, parsed_file_err := parsermovie.ParseMovieFile(*path);

	if(parsed_file_err != nil) {
		log.Fatal(parsed_file_err.Error());
	}


	movie := content.Name
	resp, tmdb_err := tmdb.SearchForMovie(movie);

	if(tmdb_err != nil) {
		log.Fatal(tmdb_err.Error());
	}


	options := []string{};

	for _, movie := range resp.Results {
		title := fmt.Sprintf("%s from %s -- %f / 10.0", movie.Title, movie.ReleaseDate, movie.VoteAverage);
		options = append(options, title);
	}
	prompt := promptui.Select{
			Label: "Wähle den Film aus den du hochladen möchtest!",
			Items: options,
	}

	index, _, err := prompt.Run()
	if err != nil {
		log.Fatal(err.Error());
		return
	}


	choosed_movie := resp.Results[index];
	cast_resp, err := tmdb.GetPersonByMovieID(choosed_movie.ID)
	detail_resp, detail_err := tmdb.GetMovieDetails(choosed_movie.ID);

	if(!*verbose) {
		save_err := tmdb.SaveTmdbData(p.Dir(*path), &choosed_movie, cast_resp, detail_resp);
		
		if(save_err != nil) {
			log.Fatal(save_err.Error());
		}
	}

	if(detail_err != nil) {
		log.Fatal(detail_err.Error());
	}


	if(err != nil) {
		log.Fatal(err.Error());
	}


	util.RemoveDuplicateCrewPerson(&cast_resp.Crews);
	util.RemoveDuplicateCastPerson(&cast_resp.Casts);

	util.SortCastByPopularity(&cast_resp.Casts);
	util.SortCrewByPopularity(&cast_resp.Crews);

	cast_resp.Casts = cast_resp.Casts[0:5]
	cast_resp.Crews = cast_resp.Crews[0:5]

	dir_from_file := filepath.Dir(*path);

	image_name := strings.ReplaceAll(choosed_movie.OriginalTitle," ", "_");
	image_file_path := dir_from_file + "/" + image_name + ".jpg"; 
	image_path := os.Getenv("API_IMAGE_BASE_URL") + choosed_movie.PosterPath;


	base64, download_err := util.DownloadImage(image_path, image_file_path)

	if(download_err != nil) {
		log.Fatal(download_err.Error());
	
	}

	json_obj := MovieBlog{};
	json_obj.Movie = MovieBlogMovie{};
	image_json_obj := MovieBlogMovieImage{Data: "data:image/jpg;base64," + base64, Name:  image_name + ".jpg"};
	json_obj.Movie.Img = image_json_obj;

	json_obj.Movie.Rating = content.Rating;
	json_obj.Movie.Release = choosed_movie.ReleaseDate;
	json_obj.Movie.Title = choosed_movie.OriginalTitle;
	json_obj.Movie.Runtime = detail_resp.Runtime;

	json_obj.Movie.Link = "https://www.themoviedb.org/movie/" + strconv.Itoa(detail_resp.ID);
 
	json_obj.Article = MovieBlogArticle{Title: content.Title};
	json_obj.Article.Body = content.Article;
	json_obj.Article.Overview = detail_resp.Overview;
	json_obj.Article.Author = content.Author;
	
	json_obj.Stickers = []MovieBlogSticker{};

	for _, sticker := range content.Stickers {
		json_obj.Stickers = append(json_obj.Stickers, MovieBlogSticker{Name: sticker.Name, Description: sticker.Description, Type: sticker.Type});
	}

	json_obj.Persons = []MovieBlogPersons{};

	for _, cast := range cast_resp.Casts {

		names := strings.Split(cast.Name, " ");
		json_obj.Persons = append(json_obj.Persons, MovieBlogPersons{Firstname: names[0], Surname: names[1], Job: "Cast"})
	}

	for _, crew := range cast_resp.Crews {
		names := strings.Split(crew.Name, " ");
		json_obj.Persons = append(json_obj.Persons, MovieBlogPersons{Firstname: names[0], Surname: names[1], Job: crew.Name})	
	}



	parsed_movie_blog, json_err := json.Marshal(&json_obj);

	if(json_err != nil) {
		log.Fatal(json_err.Error());
	}

	upload_resp, upload_err := util.POST(os.Getenv("UPLOAD_URL"), string(parsed_movie_blog), os.Getenv("AUTH_KEY"));

	if(upload_err != nil) {
		log.Fatal(upload_err.Error());
		return;
	}

	fmt.Println(upload_resp);

	log.Print("succesfull uploaded new Movie Blog!");





}

