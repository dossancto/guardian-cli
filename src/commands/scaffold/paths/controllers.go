package paths

import (
	"fmt"
	"path"

	"github.com/lu-css/guardian-cli/src/commands/scaffold/structs"
)

func ControllersLayer(config structs.FeatureConfiguration) string {
	return GetUILayer(config, "Controllers")
}

func ControllerItemPath(config structs.FeatureConfiguration, t string) string {
	controllerLayer := ControllersLayer(config)

	fileName := fmt.Sprintf("%sController%s.cs", config.FeatureName, t)

	return path.Join(controllerLayer, fileName)
}

func ControllerManageFile(config structs.FeatureConfiguration) string {
	return ControllerItemPath(config, "Manage")
}

func ControllerDtoFile(config structs.FeatureConfiguration) string {
	return ControllerItemPath(config, "Dto")
}

func ControllerSelectFile(config structs.FeatureConfiguration) string {
	return ControllerItemPath(config, "Select")
}

func ControllerDependencyInjectFile(config structs.FeatureConfiguration) string {
	return ControllerItemPath(config, "")
}
