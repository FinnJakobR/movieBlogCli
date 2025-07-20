package cliParser

import (
	"errors"
	"flag"
	"os"
)




func ParseArgs() (*string, error){	
	var pathFlag = flag.String("path", "./", "set a custom Path for uploading the movie data");
	flag.Parse()

	_, os_err := os.Open(*pathFlag);

	if(os_err != nil) {
		return  nil, errors.New("path is not a valid Path");
	}
	

	return pathFlag, nil;
}
