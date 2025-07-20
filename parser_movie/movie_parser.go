package parsermovie

import (
	"errors"
	"os"
	"strings"

	"github.com/goccy/go-yaml"
)


type MovieStickers struct {
	Type        int    `yaml:"type"`
	Description string `yaml:"description"`
	Name string `yaml:"name"`
}

type ParsedMovieFile struct {
	Name     string          `yaml:"name"`
	Title    string          `yaml:"title"`
	Rating   int             `yaml:"rating"`
	Stickers []MovieStickers `yaml:"stickers"`
	Article string
}


func ExtractHeaderAndContent(content string) (string, string, error) {
	
	lines := strings.Split(content, "\n");

	if(lines[0] != "---") {
		return "", "", errors.New("first have to be a ---");
	}

	currentLine := 1;

	for i := 1; i < len(lines); i++ {
	
		if(lines[i] == "---") {
			return  strings.Join(lines[1:currentLine - 1], "\n"), strings.Join(lines[currentLine + 1:], "\n"), nil;
		}
		currentLine++;
		
	};

	return "", "", errors.New("Unclosed Header");

}


func ParseMovieFile(path string) (*ParsedMovieFile,error) {


	if(!strings.Contains(path, ".movieBlog")) {
		return  nil, errors.New("file has to be a .movieBlog File"); 
	}


	file, err := os.ReadFile(path)

	if(err != nil) {
		return  nil, err;
	};

	content := string(file);

	header, article, extract_err := ExtractHeaderAndContent(content);

	if(extract_err != nil) {
		return  nil, extract_err;
	}

	parsedData := &ParsedMovieFile{Article: article};


	parsed_err := yaml.Unmarshal([]byte(header), parsedData);

	

	if(parsed_err != nil) {
		return  nil, parsed_err;
	}


	return parsedData, nil;

}
