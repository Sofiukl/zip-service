package main

import (
	"archive/zip"
	"io"
	"os"
)

var fileNames []string

//ZipFiles - This func zips all the mentioned files
func ZipFiles(filenames []string, zipFileName string) error {
	return Zip(filenames, zipFileName)
}

// Zip - This is the main work horse functions
func Zip(filenames []string, zipFileName string) error {
	newZipFile, err := os.Create(zipFileName)
	if err != nil {
		return err
	}
	defer newZipFile.Close()

	zipWritter := zip.NewWriter(newZipFile)
	defer zipWritter.Close()
	for _, filename := range filenames {
		if err := addFileToZip(zipWritter, filename); err != nil {
			return err
		}
	}
	return nil
}

func addFileToZip(zipWritter *zip.Writer, filename string) error {
	/**
		The below section needs to be extensible
		as file can be in server, S3, or other location
	**/
	file, fileOpenErr := os.Open(filename)
	if fileOpenErr != nil {
		return fileOpenErr
	}
	defer file.Close()
	fileInfo, statErr := file.Stat()
	if statErr != nil {
		return statErr
	}
	fileHeader, err := zip.FileInfoHeader(fileInfo)
	if err != nil {
		return err
	}
	fileHeader.Name = filename
	fileHeader.Method = zip.Deflate

	writter, err := zipWritter.CreateHeader(fileHeader)
	if err != nil {
		return err
	}
	_, err = io.Copy(writter, file)
	return err
}
