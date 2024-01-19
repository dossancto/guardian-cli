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

func mountCreateUseCase(config structs.FeatureConfiguration){
	content := application.GenerateUseCase(config)

	entityPath := entity_path.GetCreateUseCaseFilePath(config)
	utils.CreateFolderIfNotExists(entity_path.GetCreateUseCaseLayerPath(config))

	utils.CreateFile(entityPath, content)
}

