package mount

import (
	"github.com/lu-css/guardian-cli/src/commands/scaffold/content/infra"
	entity_path "github.com/lu-css/guardian-cli/src/commands/scaffold/paths"
	"github.com/lu-css/guardian-cli/src/commands/scaffold/structs"
	"github.com/lu-css/guardian-cli/src/utils"
)

func createRepositoryImpl(config structs.FeatureConfiguration) {
	content := infra.DefaultRepositoryImplContent(config)

	entityPath := entity_path.GetRepositoryImplFilePath(config)
	utils.CreateFolderIfNotExists(entity_path.GetRepositoryImplLayerPath(config))

	utils.CreateFile(entityPath, content)
}

func createEntityConfiguration(config structs.FeatureConfiguration) {
	content := infra.EntityConfigurationContent(config)

	entityPath := entity_path.GetEntityConfigurationFilePath(config)
	utils.CreateFolderIfNotExists(entity_path.GetEntityConfigurationLayerPath(config))

	utils.CreateFile(entityPath, content)
}
