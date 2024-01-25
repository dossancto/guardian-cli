package mount

import (
	// "github.com/lu-css/guardian-cli/src/commands/scaffold/content/infra"
	"github.com/lu-css/guardian-cli/src/commands/scaffold/structs"
)

func CreateFeature(config structs.FeatureConfiguration) {
	createEntities(config)
	createInterfaceRepository(config)
	createEntityConfiguration(config)
	createRepositoryImpl(config)
	mountCreateUseCase(config)
  mountUpdateUseCase(config)
  mountSelectUseCase(config)
  mountDeleteUseCase(config)
}
