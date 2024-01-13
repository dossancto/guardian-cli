package scaffold

import (
	// "github.com/lu-css/guardian-cli/src/commands/scaffold/content/infra"
	entity_path "github.com/lu-css/guardian-cli/src/commands/scaffold/paths"
	"github.com/lu-css/guardian-cli/src/commands/scaffold/structs"
)

func CreateFeature(config structs.FeatureConfiguration) {
	generateEntities(config)
	generateInterfaceRepository(config)
}

func generateInterfaceRepository(config structs.FeatureConfiguration) {
	content := createRepositoryInterface(config)

	entityPath := entity_path.GetDataFilePath(config)
	generateFeatureIfNotExists(entity_path.GetDataLayerPath(config))

	CreateFile(entityPath, content)
}

func generateEntities(config structs.FeatureConfiguration) {
	config.Fields = append(config.Fields, structs.FeatureClass{
		FieldName: "Id",
		FieldType: "string",
	})

	content := generateEntityClass(config)

	path := entity_path.GetEntitiesFilePath(config)
	generateFeatureIfNotExists(entity_path.GetEntitiesLayerPath(config))

	CreateFile(path, content)
}

// func generateUseCase(config structs.FeatureConfiguration) {
// 	entity := generateEntityClass(config)
// 	entityPath := getEntityPath(config)

// 	CreateFile(entityPath, entity)
// }
