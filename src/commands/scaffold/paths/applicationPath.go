package paths

import (
	"fmt"
	"path/filepath"

	"github.com/lu-css/guardian-cli/src/commands/scaffold/structs"
)

func GetDataFilePath(config structs.FeatureConfiguration) string {
	entitiesPath := GetDataLayerPath(config)

	fileName := fmt.Sprintf("I%sRepository.cs", config.EntityName)

	return filepath.Join(entitiesPath, fileName)
}

func GetDataLayerPath(config structs.FeatureConfiguration) string {
	return GetApplicationFeatureLayer(config, "Data")
}

func CreateFile(config structs.FeatureConfiguration) string {
	entitiesPath := UseCaseLayer(config)

	fileName := fmt.Sprintf("Create%sUseCase.cs", config.EntityName)

	return filepath.Join(entitiesPath, fileName)
}

func UseCaseLayer(config structs.FeatureConfiguration) string {
	return GetApplicationFeatureLayer(config, "UseCases")
}

func UpdateFile(config structs.FeatureConfiguration) string {
	entitiesPath := UseCaseLayer(config)

	fileName := fmt.Sprintf("Update%sUseCase.cs", config.EntityName)

	return filepath.Join(entitiesPath, fileName)
}

func SelectFile(config structs.FeatureConfiguration) string {
	entitiesPath := UseCaseLayer(config)

	fileName := fmt.Sprintf("Select%sUseCase.cs", config.EntityName)

	return filepath.Join(entitiesPath, fileName)
}

func DeleteFile(config structs.FeatureConfiguration) string {
	entitiesPath := UseCaseLayer(config)

	fileName := fmt.Sprintf("Delete%sUseCase.cs", config.EntityName)

	return filepath.Join(entitiesPath, fileName)
}

