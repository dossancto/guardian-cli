package scaffold

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/samber/lo"
)

func buildProperty(field FeatureClass) string {
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
    `,field.FieldName, field.FieldName, field.FieldType, field.FieldName, useDefaultReference)

	return template
}

func generateEntityClass(config FeatureConfiguration) string {
	fildStrs := lo.Map(config.Fields, func(field FeatureClass, _ int) string { return buildProperty(field) })
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

func generateFeatureIfNotExists(path string) {
	err := os.MkdirAll(path, os.ModePerm)

	if err != nil {
		log.Fatal(err)
	}

}
