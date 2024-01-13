package paths

import (
	"fmt"
	"path/filepath"

	"github.com/lu-css/guardian-cli/src/commands/scaffold/structs"
)

func GetEntityConfigurationFilePath(config structs.FeatureConfiguration) string {
	entitiesPath := GetEntityConfigurationLayerPath(config)

	fileName := fmt.Sprintf("%sConfiguration.cs", config.EntityName)

	return filepath.Join(entitiesPath, fileName)
}

func GetEntityConfigurationLayerPath(config structs.FeatureConfiguration) string {
	return GetInfraFeatureLayer(config, "Configurations")
}

func GetRepositoryImplFilePath(config structs.FeatureConfiguration) string {
	entitiesPath := GetRepositoryImplLayerPath(config)

	fileName := fmt.Sprintf("EF%sRepository.cs", config.EntityName)

	return filepath.Join(entitiesPath, fileName)
}

func GetRepositoryImplLayerPath(config structs.FeatureConfiguration) string {
	return GetInfraFeatureLayer(config, "Repositories")
}
