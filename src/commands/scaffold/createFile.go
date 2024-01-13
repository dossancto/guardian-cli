package scaffold

import (
	"fmt"
	"github.com/lu-css/guardian-cli/src/commands/scaffold/structs"
	"log"
	"os"
)

func getEntityFolderPath(config structs.FeatureConfiguration, layer string) string {
	var pathTemplate = "%s.Application/src/Features/%s/%s/"

	return fmt.Sprintf(pathTemplate, config.SlnName, config.FeatureName, layer)
}

func getDataFolderPath(config structs.FeatureConfiguration) string {
	folder := getEntityFolderPath(config, "Data")
	var pathTemplate = "%s/%s.cs"

	return fmt.Sprintf(pathTemplate, folder, config.EntityName)
}

func getEntityPath(config structs.FeatureConfiguration) string {
	folder := getEntityFolderPath(config, "Entities")
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
