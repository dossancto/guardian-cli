package controllers

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"

	"github.com/lu-css/guardian-cli/src/commands/scaffold/content/application"
	"github.com/lu-css/guardian-cli/src/commands/scaffold/structs"
	. "github.com/lu-css/guardian-cli/src/commands/scaffold/structs"
	"github.com/samber/lo"
)

func UpdateRequest(config structs.FeatureConfiguration) string {
	cnf := lo.Map(config.Fields, func(field structs.FeatureClass, _ int) string {
		return field.FieldName + " = " + field.FieldName
	})

	convertFields := strings.Join(cnf, ",\n\t")

	template := fmt.Sprintf(`
    /// <summary>
    /// Converts the DTO to a model.
    /// </summary>
    /// <returns>The note model.</returns>
    public Update%s ToModel(string id)
    => new()
    {
      Id = id,
      %s
    };
    `, config.EntityName, convertFields)

	return template
}

func DtoController(config FeatureConfiguration) string {

	var tmpl = `using {{.SlnName}}.Application.Features.{{.FeatureName}}.UseCases;

namespace {{.SlnName}}.UI.Controllers.{{.FeatureName}};

/// <summary>
/// Represents a request for updating a {{.EntityName}}
/// </summary>
public class Update{{.EntityName}}Request
{
    %s

    %s
}
    `

	t, err := template.New("controllerDto").Parse(tmpl)

	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	err = t.Execute(&buf, config)

	if err != nil {
		panic(err)
	}

	fields := lo.Map(config.Fields, func(field structs.FeatureClass, _ int) string {
		return application.GenerateProperty(field)
	})

	convertFields := strings.Join(fields, ",\n\t")

	dto := UpdateRequest(config)

	usecase := fmt.Sprintf(buf.String(), convertFields, dto)

	return usecase
}
