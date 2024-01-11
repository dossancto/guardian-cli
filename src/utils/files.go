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
