package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome! Zip Service is started working.")
	filenames := []string{"a.txt", "b.txt", "c.txt"}
	zipFileName := "output.zip"
	if err := ZipFiles(filenames, zipFileName); err != nil {
		fmt.Println("Error Zipping files.. ", err)
	}
}

// ZipFiles - this functions zip files
func ZipFiles(filenames []string, zipFileName string) error {
	return Zip(filenames, zipFileName)
}
