package paths

import (
	"fmt"
	"path/filepath"

	"github.com/lu-css/guardian-cli/src/commands/scaffold/structs"
)

func UseCaseLayer(config structs.FeatureConfiguration) string {
	return GetApplicationFeatureLayer(config, "UseCases")
}

func useCaseFile(config structs.FeatureConfiguration, t string) string {
	entitiesPath := UseCaseLayer(config)

	fileName := fmt.Sprintf("%s%sUseCase.cs", t, config.EntityName)

	return filepath.Join(entitiesPath, fileName)
}

func CreateFile(config structs.FeatureConfiguration) string {
	return useCaseFile(config, "Create")
}

func UpdateFile(config structs.FeatureConfiguration) string {
	return useCaseFile(config, "Update")
}

func SelectFile(config structs.FeatureConfiguration) string {
	return useCaseFile(config, "Select")
}

func DeleteFile(config structs.FeatureConfiguration) string {
	return useCaseFile(config, "Delete")
}
