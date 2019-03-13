package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome! Zip Service is started working.")
	fileExts := []string{".txt", ".pdf", ".docs", ".jpg", ".jpeg", ".png"}
	isFile := false
	args := os.Args[1:]
	folderName := args[0]
	zipFileName := args[1]

	var fileNames []string
	for _, fileExt := range fileExts {
		if strings.Contains(folderName, fileExt) {
			isFile = true
		}
	}

	if isFile {
		fileNames = append(fileNames, folderName)
	} else {
		fileNames, _ = WalkThroughFolder(folderName)
	}

	fmt.Println("Files to be Zipped: ", fileNames)
	if err := ZipFiles(fileNames, zipFileName); err != nil {
		fmt.Println("Error Zipping files.. ", err)
	}
	fmt.Println("Completed zipping. Zip File Name- ", zipFileName)
}
