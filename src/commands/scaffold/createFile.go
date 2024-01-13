package scaffold

import (
	"fmt"

	"github.com/lu-css/guardian-cli/src/commands/scaffold/structs"
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
