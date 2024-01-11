package scaffold

func CreateFeature(config FeatureConfiguration) {
	generateEntities(config)
}

func generateEntities(config FeatureConfiguration) {
	entity := generateEntityClass(config)
	entityPath := getEntityPath(config)
	generateFeatureIfNotExists(getEntityFolderPath(config))

	CreateFile(entityPath, entity)
}

func generateUseCase(config FeatureConfiguration) {
	entity := generateEntityClass(config)
	entityPath := getEntityPath(config)

	CreateFile(entityPath, entity)
}
