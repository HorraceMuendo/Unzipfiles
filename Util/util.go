package util

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func UnzipSrc(source, destination string) error {
	// open the zip file
	reader, err := zip.OpenReader(source)
	if err != nil {
		fmt.Println("cannot open the zip file....")
		return err
	}
	defer reader.Close()
	// get the absolute destination
	destination, err = filepath.Abs(destination)
	if err != nil {
		return err
	}
	// loop through all the zipfiles and unzip each one of them

	for _, file := range reader.File {
		err := UnzipFile(file, destination)
		if err != nil {
			return err
		}
	}
	return nil

}

func UnzipFile(f *zip.File, destination string) {
	filepath := filepath.Join(destination, f.Name)
	// checking vulnerability to zip slip
	if !strings.HasPrefix(filepath, filepath.Clean(destination)+string(os.PathListSeparator)) {

		return fmt.Errorf("invalid file path: %s", filepath)

	}

	//creating a directory tree so that all unziped files match the tree iside zip
	if f.FileInfo().IsDir() {
		if err := os.MkdirAll(filepath, os.ModePerm); err != nil {
			return err
		}
		return nil
	}
	//creating a destination for unzipped files
	destinationFile, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
	if err != nil {
		return err
	}
	defer destinationFile.Close()
	//unzipping file contents
	zippedFile, err := f.Open()
	if err != nil {
		return err
	}
	defer zippedFile.Close()
	if _, err := io.Copy(destinationFile, zippedFile); err != nil {
		return err
	}
	return nil
}
