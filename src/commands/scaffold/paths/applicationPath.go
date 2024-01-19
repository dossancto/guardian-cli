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

func GetCreateUseCaseFilePath(config structs.FeatureConfiguration) string {
	entitiesPath := GetCreateUseCaseLayerPath(config)

	fileName := fmt.Sprintf("Create%sUseCase.cs", config.EntityName)

	return filepath.Join(entitiesPath, fileName)
}

func GetCreateUseCaseLayerPath(config structs.FeatureConfiguration) string {
	return GetApplicationFeatureLayer(config, "UseCases")
}

func GetUpdateUseCaseFilePath(config structs.FeatureConfiguration) string {
	entitiesPath := GetUpdateUseCaseLayerPath(config)

	fileName := fmt.Sprintf("Update%sUseCase.cs", config.EntityName)

	return filepath.Join(entitiesPath, fileName)
}

func GetUpdateUseCaseLayerPath(config structs.FeatureConfiguration) string {
	return GetApplicationFeatureLayer(config, "UseCases")
}
