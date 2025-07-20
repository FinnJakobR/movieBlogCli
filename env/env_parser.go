package env

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

func Read_env_file(path string) error {

 	file, err := os.Open(path)
    
 	if err != nil {
        return  errors.New(err.Error());
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)

	line := 1;
	
	for scanner.Scan() {

		option := scanner.Text();

		if(!strings.Contains(option, "=")) {
			return errors.New("could not parse env value at line: " + strconv.Itoa(line))
		}

		splitted_option := strings.Split(option, "=");
		os.Setenv(splitted_option[0], splitted_option[1]);
		line++;
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	return nil;
}