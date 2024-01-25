package mount

import (
	"github.com/lu-css/guardian-cli/src/commands/scaffold/content/application"
	entity_path "github.com/lu-css/guardian-cli/src/commands/scaffold/paths"
	"github.com/lu-css/guardian-cli/src/commands/scaffold/structs"
	"github.com/lu-css/guardian-cli/src/utils"
)

func createEntities(config structs.FeatureConfiguration) {
	idField := structs.FeatureClass{
		FieldName: "Id",
		FieldType: "string",
	}

	config.Fields = append([]structs.FeatureClass{idField}, config.Fields...)

	content := application.GenerateEntityClass(config)

	path := entity_path.GetEntitiesFilePath(config)
	utils.CreateFolderIfNotExists(entity_path.GetEntitiesLayerPath(config))

	utils.CreateFile(path, content)
}

func createInterfaceRepository(config structs.FeatureConfiguration) {
	content := application.GenerateRepositoryInterface(config)

	entityPath := entity_path.GetDataFilePath(config)
	utils.CreateFolderIfNotExists(entity_path.GetDataLayerPath(config))

	utils.CreateFile(entityPath, content)
}

func mountCreateUseCase(config structs.FeatureConfiguration) {
	content := application.GenerateCreateUseCase(config)

	entityPath := entity_path.CreateFile(config)
	utils.CreateFolderIfNotExists(entity_path.UseCaseLayer(config))

	utils.CreateFile(entityPath, content)
}

func mountUpdateUseCase(config structs.FeatureConfiguration) {
	content := application.GenerateUpdateUseCase(config)

	entityPath := entity_path.UpdateFile(config)
	utils.CreateFolderIfNotExists(entity_path.UseCaseLayer(config))

	utils.CreateFile(entityPath, content)
}

func mountSelectUseCase(config structs.FeatureConfiguration) {
	content := application.GenerateSelectUseCase(config)

	entityPath := entity_path.SelectFile(config)
	utils.CreateFolderIfNotExists(entity_path.UseCaseLayer(config))

	utils.CreateFile(entityPath, content)
}

func mountDeleteUseCase(config structs.FeatureConfiguration) {
	content := application.GenerateDeleteUseCase(config)

	entityPath := entity_path.DeleteFile(config)
	utils.CreateFolderIfNotExists(entity_path.UseCaseLayer(config))

	utils.CreateFile(entityPath, content)
}
