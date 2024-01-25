package controllers

import (
	"bytes"
	"html/template"

	. "github.com/lu-css/guardian-cli/src/commands/scaffold/structs"
)

func DependencyInjectionController(config FeatureConfiguration) string {
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

var deleteUseCaseTemplate = `using Microsoft.AspNetCore.Mvc;

using {{.SlnName}}.Application.Features.{{.FeatureName}}.UseCases;

namespace {{.SlnName}}.UI.Controllers.{{.FeatureName}};

/// <summary>
/// Controller responsible for handling operations related to {{.FeatureName}}.
/// </summary>
[ApiController]
[Route("[controller]")]
public partial class {{.FeatureName}}Controller : ControllerBase
{
    private readonly Create{{.EntityName}}UseCase _create{{.EntityName}};
    private readonly Update{{.EntityName}}UseCase _update{{.EntityName}};
    private readonly Delete{{.EntityName}}UseCase _delete{{.EntityName}};
    private readonly Select{{.EntityName}}UseCase _select{{.EntityName}};

    /// <summary>
    /// Initializes a new instance of the <see cref="{{.FeatureName}}Controller"/> class.
    /// </summary>
    public {{.FeatureName}}Controller(Select{{.EntityName}}UseCase select{{.EntityName}}, Delete{{.EntityName}}UseCase delete{{.EntityName}}, Update{{.EntityName}}UseCase update{{.EntityName}}, Create{{.EntityName}}UseCase create{{.EntityName}})
    {
        _select{{.EntityName}} = select{{.EntityName}};
        _delete{{.EntityName}} = delete{{.EntityName}};
        _update{{.EntityName}} = update{{.EntityName}};
        _create{{.EntityName}} = create{{.EntityName}};
    }
}`
