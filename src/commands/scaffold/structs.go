package scaffold

type FeatureConfiguration struct {
	FeatureName string
	EntityName  string
	SlnName     string
	Fields      []FeatureClass
}

type FeatureClass struct {
	FieldName string
	FieldType string
}
