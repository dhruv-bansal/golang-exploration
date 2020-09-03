package main

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func main() {

	var libRegEx, e = regexp.Compile("(target|.git)")
	if e != nil {
		log.Fatal(e)
	}

	currentPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	currentPath = "/Users/dhruvbansal/Documents/code/CAFU/release1.1/idea/issue/cobra-auth-server"
	log.Print("Scanning path ", currentPath)

	directoriesToBeDeleted := make([]string, 0, 10)
	var err1 = filepath.Walk(currentPath, func(path string, fileInfo os.FileInfo, err error) error {
		if err == nil && fileInfo.IsDir() && libRegEx.MatchString(fileInfo.Name()) {
			log.Print("Found: ", path)
			directoriesToBeDeleted = append(directoriesToBeDeleted, path)
		}
		return nil
	})

	if err1 != nil {
		log.Fatal(err1)
	}

	log.Print("Total number of directories to be deleted ", len(directoriesToBeDeleted))
	log.Print("Directories to be deleted ", directoriesToBeDeleted)

	for index, directory := range directoriesToBeDeleted {
		log.Print("Deleting directory: ", index, "  ", directory)
		//err := os.RemoveAll(directory)
		if err != nil {
			log.Fatal(err)
		}
	}

}
