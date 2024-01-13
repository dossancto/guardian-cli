package utils

import (
	"errors"
	"log"
	"os"
)

func FolderExists() error {
	_, err := os.Stat("/path/to/directory")

	if err == nil {
		return nil
	}

	if !os.IsNotExist(err) {
		log.Fatal(err)
	}

	return errors.New("File does not exists")
}

func CreateFolderIfNotExists(path string) {
	err := os.MkdirAll(path, os.ModePerm)

	if err != nil {
		log.Fatal(err)
	}

}

func CreateFile(path string, content string) {
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}
}
