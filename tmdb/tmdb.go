package tmdb

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
)

func saveTmdbDataInFile(strc any, path string) error {

	str, str_error := json.Marshal(strc);

	if(str_error != nil) {
		return str_error;
	}

	write_err := os.WriteFile(path, str, 0755);

	if(write_err != nil) {
		return  write_err;
	}

	return nil;

}

func SaveTmdbData(p string, movieSearchResponse *Movie, movieCreditResponse *CastRepsonse, movieDetailResponse *MovieDetailResponse) error{

	save_file_path := path.Join(p, "tmdb_" + strconv.Itoa(movieSearchResponse.ID) + "_data.json");
	

	save_value := &SaveMovieStruct{};

	save_value.Credits = *movieCreditResponse;
	save_value.Movie = *movieSearchResponse;
	save_value.Details = *movieDetailResponse;

	save_err := saveTmdbDataInFile(save_value, save_file_path);

	if(save_err != nil) {
		return  save_err;
	}

	return nil;

}



func addParamsToUrl(url string, params []string) string {

	url += "?";
	
	for _, param := range params {
		param = strings.ReplaceAll(param, " ", "%20")
		url+=param;
		url+="&";
	}

	//remove last &
	url = url[:len(url) - 1];
	return url;
}




func GET(url string, params []string, api_key string) (string, error) {

	client := &http.Client{}
	url_with_params := addParamsToUrl(url, params);
	req, err := http.NewRequest("GET", url_with_params, nil)

	if(err != nil) {		
		return  "", errors.New(err.Error());
	}

	 req.Header.Set("Authorization", "Bearer " + api_key);
	 req.Header.Set("accept", "application/json");

	resp, err := client.Do(req)
    
	
	//err nur dann wenn es einen internen Fehler gibt!
	if err != nil {
       return  "", errors.New(err.Error());
    }
	
	defer resp.Body.Close()


    // Antwort lesen
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", errors.New(err.Error());
    }

	if(resp.StatusCode != 200) {
		return "", errors.New("Request for url: " + url_with_params + " failed with StatusCode: " + strconv.Itoa(resp.StatusCode) + "and message: " + string(body));
	}

	return string(body), nil;

}

func GetMovieDetails(movie_id int) (*MovieDetailResponse, error) {

	api_key := os.Getenv("API_KEY");
	url := os.Getenv("API_MOVIE_DETAIL_URL");
	url = strings.ReplaceAll(url, "movie_id", strconv.Itoa(movie_id));

	params := []string{"language=de"};
	
	response, err := GET(url, params, api_key); 

	if(err  != nil) {
		return nil, err;
	}

	resp := &MovieDetailResponse{};

	json_err := json.Unmarshal([]byte(response), resp);

	
	if(json_err != nil) {
		return nil, json_err;
	};

	return resp, nil;

}


func SearchForMovie(movie string) (*MovieResponse, error) {

	api_key := os.Getenv("API_KEY");
	url := os.Getenv("API_SEARCH_FOR_MOVIE_URL");
	


	if(api_key == "" || url == "") {
		return nil, errors.New("could not extract all Infos please provide one");
	}


	params := [4]string{"query=" + movie, "include_adult=false", "language=en-US", "page=1"};

	body, err := GET(url, params[:], api_key)
	
	
	if(err != nil) {
		return  nil, errors.New(err.Error());
	}


	resp := &MovieResponse{};

	json_err := json.Unmarshal([]byte(body), resp);
	if json_err != nil {
		return  nil, errors.New(json_err.Error());
	}
	
	return resp, nil;
}


func GetPersonByMovieID(id int) (*CastRepsonse, error){


	api_key := os.Getenv("API_KEY");
	url := os.Getenv("API_CASTS_FOR_MOVIE_URL");
	
	url = strings.Replace(url, "movie_id", strconv.Itoa(id), 1);


	if(api_key == "" || url == "") {
		return nil, errors.New("could not extract all Infos please provide one");
	}
	
	params := [1]string{"language=en-US"};

	body, err := GET(url, params[:], api_key)
	
	
	if(err != nil) {
		return  nil, errors.New(err.Error());
	}

	resp := &CastRepsonse{};

	json_err := json.Unmarshal([]byte(body), resp);
	if json_err != nil {
		return  nil, errors.New(json_err.Error());
	}
	
	return resp, nil;

}