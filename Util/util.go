package util

import (
	"archive/zip"
	"fmt"
	"path/filepath"
)

func unzipSrc(source, destination string) error {
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
		err := unzipFile(file, destination)
		if err != nil {
			return err
		}
	}
	return nil

}
