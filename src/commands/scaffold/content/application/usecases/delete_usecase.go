package usecases

import (
	"bytes"
	"text/template"

	"github.com/lu-css/guardian-cli/src/commands/scaffold/structs"
)

func GenerateDeleteUseCase(config structs.FeatureConfiguration) string {
	t, err := template.New("deleteEntityUseCase").Parse(deleteUseCaseTemplate)

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

var deleteUseCaseTemplate =
`using {{.SlnName}}.Application.Features.{{.FeatureName}}.Data;

namespace {{.SlnName}}.Application.Features.{{.FeatureName}}.UseCases;

/// <summary>
/// This class encapsulates the use case of deleting a {{.EntityName}}.
/// </summary>
public class Delete{{.EntityName}}UseCase
{
    private readonly I{{.EntityName}}Repository _{{.EntityName}}Repository;

    /// <summary>
    /// Initializes a new instance of the <see cref="Delete{{.EntityName}}UseCase"/> class.
    /// </summary>
    /// <param name="{{.EntityName}}Repository">The {{.EntityName}} repository.</param>
    public Delete{{.EntityName}}UseCase(I{{.EntityName}}Repository {{.EntityName}}Repository)
    {
        _{{.EntityName}}Repository = {{.EntityName}}Repository;
    }

    /// <summary>
    /// Deletes a {{.EntityName}} by its unique identifier.
    /// </summary>
    /// <param name="id">The unique identifier of the {{.EntityName}}.</param>
    public Task Execute(string id)
    => _{{.EntityName}}Repository.Delete(id);
}`
