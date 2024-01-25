package mount

import (
	content "github.com/lu-css/guardian-cli/src/commands/scaffold/content/application/controllers"
	entity_path "github.com/lu-css/guardian-cli/src/commands/scaffold/paths"
	"github.com/lu-css/guardian-cli/src/commands/scaffold/structs"
	"github.com/lu-css/guardian-cli/src/utils"
)

func mountControllers(config structs.FeatureConfiguration) {
	mountMannageController(config)
	// mountSelectController(config)
	// mountDtoController(config)
	// mountDependencyInversionController(config)
}

func mountControllerType(config structs.FeatureConfiguration, entityPath string, content string) {
	utils.CreateFolderIfNotExists(entity_path.ControllersLayer(config))
	utils.CreateFile(entityPath, content)
}

func mountMannageController(config structs.FeatureConfiguration) {
	content := content.DependencyInjectionController(config)
	entityPath := entity_path.ControllerManageFile(config)
	mountControllerType(config, entityPath, content)
}

func mountSelectController(config structs.FeatureConfiguration) {
	content := content.DependencyInjectionController(config)
	entityPath := entity_path.ControllerSelectFile(config)
	mountControllerType(config, entityPath, content)
}

func mountDtoController(config structs.FeatureConfiguration) {
	content := content.DependencyInjectionController(config)
	entityPath := entity_path.ControllerDtoFile(config)
	mountControllerType(config, entityPath, content)
}

func mountDependencyInversionController(config structs.FeatureConfiguration) {
	content := content.DependencyInjectionController(config)
	entityPath := entity_path.ControllerDependencyInjectFile(config)
	mountControllerType(config, entityPath, content)
}
