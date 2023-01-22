package util

import (
	"archive/zip"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

//"fmt",

func zipFile(source, target string) error {
	// create a zip fle and zip.Writer
	f, err := os.Create(target)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := zip.NewWriter(f)
	defer writer.Close()

	//go through all the files in the source
	return filepath.Walk(source, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		//local file header
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		//set compression
		header.Method = zip.Deflate

		//relative file path as header name
		header.Name, err = filepath.Rel(filepath.Dir(source), path)
		if err != nil {
			return err
		}
		if info.IsDir() {
			header.Name += "/"
		}

		headerWriter, err := writer.CreateHeader(header)
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if err != nil {
			return err
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = io.Copy(headerWriter, f)
		return err

	})

}
