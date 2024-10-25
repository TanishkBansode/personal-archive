package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "strings"
)


func main() {	 

    // Step 1: Get the filenames and their content from the blogs directory
    dir, err := os.Getwd()
    parentDir := filepath.Join(dir, "..", "blogs")
    CheckNilError(err)

    files, err := ioutil.ReadDir(parentDir)
    CheckNilError(err)

    fmt.Println(files)
    var fileNames []string
    for _, file := range files {
        if !file.IsDir() && strings.HasSuffix(file.Name(), ".txt") { // Check if it's not a directory
	    fileNames = append(fileNames, file.Name())
        }
    }
	
    readit(fileNames)
}
	
	
// Reads a file, and currently prints in format "{Filename}: {content}"
func readit(filenames []string) {
    dir, err := os.Getwd()
	
    CheckNilError(err)

    for _, filename := range filenames {
        path := filepath.Join(dir, "..", "blogs", filename)
	data, err := ioutil.ReadFile(path)
	CheckNilError(err)
	fmt.Printf("%s: %s\n", filename, string(data))
		
    }
	
	
}


// Checks err, it pretty handy
func CheckNilError(err error){
	if err != nil {
		panic(err)
	}
}
