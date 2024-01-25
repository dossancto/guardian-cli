package paths

import (
	"fmt"
	"path"

	"github.com/lu-css/guardian-cli/src/commands/scaffold/structs"
)

func GetUILayer(config structs.FeatureConfiguration, folder string) string {
	pathTemplate := "%s.UI/src/%s/"

	return fmt.Sprintf(pathTemplate, config.SlnName, folder)

}

func ControllersLayer(config structs.FeatureConfiguration) string {
	return GetUILayer(config, "Controllers")
}

func ControllerItemPath(config structs.FeatureConfiguration, t string) string {
	controllerLayer := ControllersLayer(config)

	fileName := fmt.Sprintf("%sController%s.cs", config.FeatureName, t)

	return path.Join(controllerLayer, fileName)
}
