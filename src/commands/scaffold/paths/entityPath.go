package paths

import (
	"fmt"
	"path/filepath"

	"github.com/lu-css/guardian-cli/src/commands/scaffold/structs"
)

func GetEntitiesFilePath(config structs.FeatureConfiguration) string {
	entitiesPath := GetEntitiesLayerPath(config)

	fileName := fmt.Sprintf("%s.cs", config.EntityName)

	return filepath.Join(entitiesPath, fileName)
}

func GetEntitiesLayerPath(config structs.FeatureConfiguration) string {
	return GetApplicationFeatureLayer(config, "Entities")
}
