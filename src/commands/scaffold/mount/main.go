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

	mountUseCases(config)
	mountControllers(config)
}

