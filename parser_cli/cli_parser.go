package cliParser

import (
	"errors"
	"flag"
	"os"
)




func ParseArgs() (*string, *bool, error){	
	pathFlag := flag.String("path", "./", "set a custom Path for uploading the movie data");
	verboseFlag := flag.Bool("verbose", false, "Set verbose to true will not save the tmdb data")
	flag.Parse()

	_, os_err := os.Open(*pathFlag);

	if(os_err != nil) {
		return  nil, nil, errors.New("path is not a valid Path");
	}
	

	return pathFlag, verboseFlag, nil;
}
