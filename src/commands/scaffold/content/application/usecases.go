package application

import (
	"fmt"
	"strings"

	"github.com/lu-css/guardian-cli/src/commands/scaffold/structs"
	"github.com/samber/lo"
)

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

func generateDTO(docs string, classname string, config structs.FeatureConfiguration) string {
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
