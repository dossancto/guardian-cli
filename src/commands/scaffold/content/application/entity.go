package application

import (
	"fmt"
	"strings"

	"github.com/lu-css/guardian-cli/src/commands/scaffold/structs"
	"github.com/samber/lo"
)

func GenerateUpdateDTO(config structs.FeatureConfiguration) string {

	docs := fmt.Sprintf("Represents a data object for updating a %s", config.EntityName)
	classname := fmt.Sprintf("Update%s", config.EntityName)
	return GenerateDTO(docs, classname, config)
}

func GenerateDTO(docs string, classname string, config structs.FeatureConfiguration) string {
	fildStrs := lo.Map(config.Fields, func(field structs.FeatureClass, _ int) string { return GenerateProperty(field) })
	fields := strings.Join(fildStrs, "\n\n\t")

	convertMethod := toModelMethod("Create", config)

	template := fmt.Sprintf(`
/// <summary>
/// %s
/// </summary>
public class %s
{
    %s

    %s

}
    `, docs, classname, fields, convertMethod)

	return template
}

func toModelMethod(action string, config structs.FeatureConfiguration) string {
	cnf := lo.Map(config.Fields, func(field structs.FeatureClass, _ int) string {
		return field.FieldName + " = " + field.FieldName
	})

	convertFields := strings.Join(cnf, ",\n\t")
	template := fmt.Sprintf(`
    /// <summary>
    /// Converts the DTO to a model.
    /// </summary>
    /// <returns>The note model.</returns>
    public %s ToModel()
    => new()
    {
      %s
    };
    `, config.EntityName, convertFields)
	return template
}

func GenerateProperty(field structs.FeatureClass) string {
	isString := field.FieldType == "string"

	useDefaultReference := ""

	if isString {
		useDefaultReference = "= default!;"
	}

	template := fmt.Sprintf(`
    /// <summary>
    /// %s
    /// </summary>
    /// <example>%s</example>
    public %s %s {get; set;} %s
    `, field.FieldName, field.FieldName, field.FieldType, field.FieldName, useDefaultReference)

	return template
}

func GenerateEntityClass(config structs.FeatureConfiguration) string {
	fildStrs := lo.Map(config.Fields, func(field structs.FeatureClass, _ int) string { return GenerateProperty(field) })
	fields := strings.Join(fildStrs, "\n\n\t")

	template := fmt.Sprintf(`
namespace %s.Application.Features.%s.Entities;

/// <summary>
/// %s
/// </summary>
public class %s
{
        %s
}
    `, config.SlnName, config.FeatureName, config.EntityName, config.EntityName, fields)

	return template
}
