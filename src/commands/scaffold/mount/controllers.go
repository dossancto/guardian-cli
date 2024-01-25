package mount

import (
	content "github.com/lu-css/guardian-cli/src/commands/scaffold/content/application/usecases"
	entity_path "github.com/lu-css/guardian-cli/src/commands/scaffold/paths"
	"github.com/lu-css/guardian-cli/src/commands/scaffold/structs"
	"github.com/lu-css/guardian-cli/src/utils"
)

func mountControllers(config structs.FeatureConfiguration) {
	mountMannageController(config)
	mountSelectController(config)
	mountDtoController(config)
	mountDependencyInversionController(config)
}

func mountControllerType(config structs.FeatureConfiguration, entityPath string, content string) {
	utils.CreateFolderIfNotExists(entity_path.UseCaseLayer(config))
	utils.CreateFile(entityPath, content)
}

func mountMannageController(config structs.FeatureConfiguration) {
	content := content.GenerateCreateUseCase(config)
	entityPath := entity_path.CreateFile(config)
	mountControllerType(config, entityPath, content)
}

func mountSelectController(config structs.FeatureConfiguration) {
	content := content.GenerateUpdateDTO(config)
	entityPath := entity_path.UpdateFile(config)
	mountControllerType(config, entityPath, content)
}

func mountDtoController(config structs.FeatureConfiguration) {
	content := content.GenerateSelectUseCase(config)
	entityPath := entity_path.SelectFile(config)
	mountControllerType(config, entityPath, content)
}

func mountDependencyInversionController(config structs.FeatureConfiguration) {
	content := content.GenerateDeleteUseCase(config)
	entityPath := entity_path.DeleteFile(config)
	mountControllerType(config, entityPath, content)
}
