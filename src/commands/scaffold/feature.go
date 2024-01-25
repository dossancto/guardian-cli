package scaffold

import (
	"errors"
	"log"
	"strings"

	"github.com/lu-css/guardian-cli/src/commands/filters"
	"github.com/lu-css/guardian-cli/src/commands/scaffold/mount"
	"github.com/lu-css/guardian-cli/src/commands/scaffold/structs"
	"github.com/lu-css/guardian-cli/src/utils"
	"github.com/manifoldco/promptui"
)

func HandleNewFeature(slnName string) {
	config := NewFeatureConfiguration(slnName)
	mount.CreateFeature(config)
}

func NewFeatureConfiguration(slnName string) structs.FeatureConfiguration {
	featureName := getFeatureName()
	entityName := getEntityName()
	fields := getFields()
	useControllers := getUseControllers()

	return structs.FeatureConfiguration{
		FeatureName: featureName,
		EntityName:  entityName,
		SlnName:     slnName,
		Fields:      fields,
		ScaffoldConfiguration: structs.ScaffoldConfiguration{
			GenerateController: useControllers,
		},
	}
}

func getFeatureName() string {
	validate := func(input string) error {
		myFilters := []filters.TextFilter{
			filters.BlankTextFilter,
			filters.SpecialCharFilter,
		}

		return filters.CascateValidator(myFilters, input)
	}

	prompt := promptui.Prompt{
		Label:    "Feature name",
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}

	return result
}

func paramValidate(text string) error {
	tokens := strings.Split(text, " ")

	for _, token := range tokens {
		parts := strings.Split(token, ":")

		if len(parts) != 2 {
			return errors.New("Invalid field")
		}
	}

	return nil
}
func getParamsList(text string) []structs.FeatureClass {
	fields := []structs.FeatureClass{}

	tokens := strings.Split(text, " ")

	for _, token := range tokens {
		parts := strings.Split(token, ":")
		fieldName := parts[0]
		fieldType := parts[1]

		fields = append(fields, structs.FeatureClass{FieldName: fieldName, FieldType: fieldType})
	}

	return fields
}

func getFields() []structs.FeatureClass {
	validate := func(input string) error {
		myFilters := []filters.TextFilter{
			filters.BlankTextFilter,
			paramValidate,
		}

		return filters.CascateValidator(myFilters, input)
	}

	prompt := promptui.Prompt{
		Label:    "Fields list",
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}

	return getParamsList(result)
}

func getEntityName() string {
	validate := func(input string) error {
		myFilters := []filters.TextFilter{
			filters.BlankTextFilter,
		}

		return filters.CascateValidator(myFilters, input)
	}

	prompt := promptui.Prompt{
		Label:    "Entity Name",
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}

	return result
}

func getUseControllers() bool {
	return utils.TrueOrFalse("Use controllers", false)
}
