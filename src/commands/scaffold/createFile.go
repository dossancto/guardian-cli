package scaffold

import (
	"fmt"
	"log"
	"os"
)

func getEntityFolderPath(config FeatureConfiguration) string {
	var pathTemplate = "%s.Application/src/Features/%s/Entities/"

	return fmt.Sprintf(pathTemplate, config.SlnName, config.FeatureName)
}

func getEntityPath(config FeatureConfiguration) string {
  folder := getEntityFolderPath(config)
	var pathTemplate = "%s/%s.cs"

	return fmt.Sprintf(pathTemplate, folder, config.EntityName)
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
