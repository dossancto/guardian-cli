package usecases

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/lu-css/guardian-cli/src/commands/scaffold/content/application"
	"github.com/lu-css/guardian-cli/src/commands/scaffold/structs"
)

func GenerateUpdateUseCase(config structs.FeatureConfiguration) string {
	tmpl := `
using {{.SlnName}}.Application.Features.{{.FeatureName}}.Data;
using {{.SlnName}}.Application.Features.{{.FeatureName}}.Entities;

namespace {{.SlnName}}.Application.Features.{{.FeatureName}}.UseCases;

/// <summary>
/// This class is responsible for updating a {{.EntityName}} using a given repository.
/// </summary>
public class Update{{.EntityName}}UseCase
{
    private readonly I{{.EntityName}}Repository _{{.EntityName}}Repository;

    /// <summary>
    /// Initializes a new instance of the <see cref="Update{{.EntityName}}UseCase"/> class.
    /// </summary>
    /// <param name="{{.EntityName}}Repository">The {{.EntityName}} repository.</param>
    public Update{{.EntityName}}UseCase(I{{.EntityName}}Repository {{.EntityName}}Repository)
    {
        _{{.EntityName}}Repository = {{.EntityName}}Repository;
    }

    /// <summary>
    /// Executes the update of a {{.EntityName}}.
    /// </summary>
    /// <param name="dto">The data transfer object containing the details of the {{.EntityName}} to be updated.</param>
    /// <returns>A task representing the asynchronous operation.</returns>
    public Task<{{.EntityName}}> Execute(Update{{.EntityName}} dto)
    => _{{.EntityName}}Repository.Save(dto.ToModel());
}

  `
	t, err := template.New("updateEntityUseCase").Parse(tmpl)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	err = t.Execute(&buf, config)
	if err != nil {
		panic(err)
	}

	usecase := buf.String()

	idField := structs.FeatureClass{
		FieldName: "Id",
		FieldType: "string",
	}

	config.Fields = append([]structs.FeatureClass{idField}, config.Fields...)

	dto := application.GenerateUpdateDTO(config)

	return fmt.Sprintf("%s\n%s", usecase, dto)
}
