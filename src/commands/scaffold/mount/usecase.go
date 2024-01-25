package mount

import (
	content "github.com/lu-css/guardian-cli/src/commands/scaffold/content/application/usecases"
	entity_path "github.com/lu-css/guardian-cli/src/commands/scaffold/paths"
	"github.com/lu-css/guardian-cli/src/commands/scaffold/structs"
	"github.com/lu-css/guardian-cli/src/utils"
)

func mountUseCases(config structs.FeatureConfiguration) {
	mountCreateUseCase(config)
	mountUpdateUseCase(config)
	mountSelectUseCase(config)
	mountDeleteUseCase(config)
}

func mountUseCaseType(config structs.FeatureConfiguration, entityPath string, content string) {
	utils.CreateFolderIfNotExists(entity_path.UseCaseLayer(config))
	utils.CreateFile(entityPath, content)
}

func mountCreateUseCase(config structs.FeatureConfiguration) {
	content := content.GenerateCreateUseCase(config)
	entityPath := entity_path.CreateFile(config)
	mountUseCaseType(config, entityPath, content)
}

func mountUpdateUseCase(config structs.FeatureConfiguration) {
	content := content.GenerateUpdateUseCase(config)
	entityPath := entity_path.UpdateFile(config)
	mountUseCaseType(config, entityPath, content)
}

func mountSelectUseCase(config structs.FeatureConfiguration) {
	content := content.GenerateSelectUseCase(config)
	entityPath := entity_path.SelectFile(config)
	mountUseCaseType(config, entityPath, content)
}

func mountDeleteUseCase(config structs.FeatureConfiguration) {
	content := content.GenerateDeleteUseCase(config)
	entityPath := entity_path.DeleteFile(config)
	mountUseCaseType(config, entityPath, content)
}
