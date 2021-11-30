package main

import (
	"archive/zip"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {

	create, err := os.Create("zip/aaa.zip")
	if err != nil {
		return
	}
	defer create.Close()
	zipWriter := zip.NewWriter(create)
	defer zipWriter.Close()
	filepath.Walk("zip", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		w, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}
		if !info.IsDir() {
			open, err := os.Open(path)
			if err != nil {
				return err
			}
			defer open.Close()
			_, err = io.Copy(w, open)

		}
		return err
	})
}
