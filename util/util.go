package util

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"movieBlog/cli/tmdb"
	"net/http"
	"os"
	"sort"
)




func POST(path string, jsonBody string, authToken string) (string, error) {
	// JSON-Body als Reader
	reqBody := bytes.NewBuffer([]byte(jsonBody))

	// HTTP-POST Request erstellen
	req, err := http.NewRequest("POST", path, reqBody)
	if err != nil {
		return "", err
	}

	// Header setzen
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer " + authToken)

	// HTTP-Client ausführen
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Antwort lesen
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Optional: HTTP-Status prüfen
	if resp.StatusCode >= 300 {
		return "", fmt.Errorf("request failed with status %d: %s", resp.StatusCode, string(respBody))
	}

	return string(respBody), nil
}



func DownloadImage(path string, filePath string) (string, error) {
	resp, err := http.Get(path)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("could not download image")
	}

	// Lies den Body einmal
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Datei speichern
	err = os.WriteFile(filePath, bodyBytes, 0644)
	if err != nil {
		return "", err
	}

	// Base64 erzeugen
	return base64.StdEncoding.EncodeToString(bodyBytes), nil
}

func RemoveDuplicateCrewPerson(crew *[]tmdb.Crew) {

	data := *crew
		seen := make(map[int]bool)
		result := make([]tmdb.Crew, 0, len(data))

		for _, c := range data {
			if !seen[c.ID] {
				seen[c.ID] = true
				result = append(result, c)
			}
		}

		*crew = result
}


func RemoveDuplicateCastPerson(cast *[]tmdb.Casts) {

		data := *cast
		seen := make(map[int]bool)
		result := make([]tmdb.Casts, 0, len(data))

		for _, c := range data {
			if !seen[c.ID] {
				seen[c.ID] = true
				result = append(result, c)
			}
		}

		*cast = result
}




func SortCrewByPopularity(crew *[]tmdb.Crew ){
		
	data := *crew 

	sort.SliceStable(data, func(i, j int) bool {
		return data[i].Popularity > data[j].Popularity
	})

	*crew = data
}


func SortCastByPopularity(cast *[]tmdb.Casts ) {
	
	data := *cast 

	sort.SliceStable(data, func(i, j int) bool {
		return data[i].Popularity > data[j].Popularity
	})

	*cast = data
}