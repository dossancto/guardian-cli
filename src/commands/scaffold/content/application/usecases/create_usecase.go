package usecases

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/lu-css/guardian-cli/src/commands/scaffold/structs"
)

func GenerateCreateDTO(config structs.FeatureConfiguration) string {
	docs := fmt.Sprintf("Represents a data object for creating a %s", config.EntityName)
	classname := fmt.Sprintf("Create%s", config.EntityName)
	return generateDTO(docs, classname, config)
}

func GenerateCreateUseCase(config structs.FeatureConfiguration) string {
	tmpl := `
using {{.SlnName}}.Application.Features.{{.FeatureName}}.Data;
using {{.SlnName}}.Application.Features.{{.FeatureName}}.Entities;

namespace {{.SlnName}}.Application.Features.{{.FeatureName}}.UseCases;

/// <summary>
/// This class is responsible for creating a {{.EntityName}} using a given repository.
/// </summary>
public class Create{{.EntityName}}UseCase
{
    private readonly I{{.EntityName}}Repository _{{.EntityName}}Repository;

    /// <summary>
    /// Initializes a new instance of the <see cref="Create{{.EntityName}}UseCase"/> class.
    /// </summary>
    /// <param name="{{.EntityName}}Repository">The {{.EntityName}} repository.</param>
    public Create{{.EntityName}}UseCase(I{{.EntityName}}Repository {{.EntityName}}Repository)
    {
        _{{.EntityName}}Repository = {{.EntityName}}Repository;
    }

    /// <summary>
    /// Executes the creation of a {{.EntityName}}.
    /// </summary>
    /// <param name="dto">The data transfer object containing the details of the {{.EntityName}} to be created.</param>
    /// <returns>A task representing the asynchronous operation.</returns>
    public Task Execute(Create{{.EntityName}} dto)
    => _{{.EntityName}}Repository.Save(dto.ToModel());
}
  `
	t, err := template.New("createEntityUseCase").Parse(tmpl)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	err = t.Execute(&buf, config)
	if err != nil {
		panic(err)
	}

	usecase := buf.String()
	dto := GenerateCreateDTO(config)

	return fmt.Sprintf("%s\n%s", usecase, dto)
}
