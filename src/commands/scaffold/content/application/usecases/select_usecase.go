package usecases

import (
	"bytes"
	"text/template"

	"github.com/lu-css/guardian-cli/src/commands/scaffold/structs"
)

func GenerateSelectUseCase(config structs.FeatureConfiguration) string {
	tmpl := `
using {{.SlnName}}.Application.Features.{{.FeatureName}}.Data;
using {{.SlnName}}.Application.Features.{{.FeatureName}}.Entities;

namespace {{.SlnName}}.Application.Features.{{.FeatureName}}.UseCases;

/// <summary>
/// This class is responsible getting data from {{.EntityName}}s
/// </summary>
public class Select{{.EntityName}}UseCase
{
    private readonly I{{.EntityName}}Repository _{{.EntityName}}Repository;

    /// <summary>
    /// Initializes a new instance of the <see cref="Select{{.EntityName}}UseCase"/> class.
    /// </summary>
    /// <param name="{{.EntityName}}Repository">The {{.EntityName}} repository.</param>
    public Select{{.EntityName}}UseCase(I{{.EntityName}}Repository {{.EntityName}}Repository)
    {
        _{{.EntityName}}Repository = {{.EntityName}}Repository;
    }

    /// <summary>
    /// Retrieves a note by its unique identifier.
    /// </summary>
    /// <param name="id">The unique identifier of the note.</param>
    /// <returns>The task result contains the note if found, null otherwise.</returns>
    public Task<{{.EntityName}}?> ById(string id)
    => _{{.EntityName}}Repository.FindById(id);

    /// <summary>
    /// Retrieves all notes.
    /// </summary>
    /// <returns>The task result contains a list of all notes.</returns>
    public Task<List<{{.EntityName}}>> All()
    => _{{.EntityName}}Repository.All();
}

  `
	t, err := template.New("selectEntityUseCase").Parse(tmpl)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	err = t.Execute(&buf, config)
	if err != nil {
		panic(err)
	}

	usecase := buf.String()

	return usecase
}

