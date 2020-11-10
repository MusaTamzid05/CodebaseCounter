package codebaseLib

import (
	"log"
	"os"
	"path/filepath"
)

func GetFilePath(path string) ([]string, error) {

	paths := []string{}

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			log.Fatal(err)
		}

		paths = append(paths, path)
		return nil
	})

	if err != nil {
		return paths, err
	}
	return paths, nil
}
