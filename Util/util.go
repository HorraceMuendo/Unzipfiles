package util

import (
	"archive/zip"
	"fmt"
	"os"
	"path/filepath"
	"strings"
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

func unzipFile(f *zip.File, destination string) {
	filepath := filepath.Join(destination, f.Name)

	if !strings.HasPrefix(filepath, filepath.Clean(destination)+string(os.PathListSeparator)) {
		return fmt.Errorf("invalid file path: %s", filepath)
	}

	if f.FileInfo().IsDir() {
		if err := os.MkdirAll(filepath, os.ModePerm); err != nil {
			return err
		}
		return nil
	}
}
