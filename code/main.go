package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"

)


func main() {	 
    // Step 1: Checks for all files in the blogs directory
    dir, err := os.Getwd()
    parentDir := filepath.Join(dir, "..", "blogs")

    CheckNilError(err)

    files, err := ioutil.ReadDir(parentDir)
    CheckNilError(err)

    fmt.Println("Files in directory:", parentDir)
    for _, file := range files {
        if !file.IsDir() { // Check if it's not a directory
            fmt.Println(file.Name())
        }
    }
}
	


func CheckNilError(err error){
	if err != nil {
		panic(err)
	}
}
