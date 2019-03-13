package main

import "io/ioutil"

// WalkThroughFolder - This function walk through the specified folder
func WalkThroughFolder(folderName string) ([]string, error) {
	files, err := ioutil.ReadDir(folderName)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if !file.IsDir() {
			fileNames = append(fileNames, folderName+"/"+file.Name())
		} else {
			newFolderName := folderName + "/" + file.Name()
			WalkThroughFolder(newFolderName)
		}
	}
	return fileNames, nil
}
