package paths

import (
	"fmt"
	"github.com/lu-css/guardian-cli/src/commands/scaffold/structs"
	"path/filepath"
)

func GetApplicationFeaturePath(config structs.FeatureConfiguration) string {
	pathTemplate := "%s.Application/src/Features/%s/"

	return fmt.Sprintf(pathTemplate, config.SlnName, config.FeatureName)
}

func GetApplicationFeatureLayer(config structs.FeatureConfiguration, layer string) string {
	featurePath := GetApplicationFeaturePath(config)
	return filepath.Join(featurePath, layer)
}

func GetInfraFeaturePath(config structs.FeatureConfiguration) string {
	pathTemplate := "%s.Infra/src/Database/EF/"

	return fmt.Sprintf(pathTemplate, config.SlnName)
}

func GetInfraFeatureLayer(config structs.FeatureConfiguration, layer string) string {
	featurePath := GetApplicationFeaturePath(config)
	return filepath.Join(featurePath, layer)
}
