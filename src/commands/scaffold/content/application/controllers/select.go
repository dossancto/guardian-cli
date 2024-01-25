package controllers

import (
	"bytes"
	"html/template"

	. "github.com/lu-css/guardian-cli/src/commands/scaffold/structs"
)

func SelectController(config FeatureConfiguration) string {

	var selectTemplate = `using Microsoft.AspNetCore.Mvc;
using {{.SlnName}}.Application.Features.{{.FeatureName}}.Entities;

namespace {{.SlnName}}.UI.Controllers.{{.FeatureName}};

public partial class {{.FeatureName}}Controller
{
    /// <summary>
    /// List All {{.EntityName}}s
    /// </summary>
    /// <remarks>List all {{.EntityName}}s</remarks>
    /// <response code="200">The list of {{.EntityName}}s</response>
    /// <response code="500">Fail while searching for {{.EntityName}}</response>
    [ProducesResponseType(typeof(List<{{.EntityName}}>), 200)]
    [ProducesResponseType(500)]
    [HttpGet]
    public async Task<IActionResult> List{{.EntityName}}()
    {
        var result = await _select{{.EntityName}}.All();

        return Ok(result);
    }

    /// <summary>
    /// Get {{.EntityName}} by Id
    /// </summary>
    /// <remarks>Get a {{.EntityName}} by id</remarks>
    /// <response code="200">Find a {{.EntityName}} by id</response>
    /// <response code="500">Fail while searching for {{.EntityName}}</response>
    [ProducesResponseType(typeof(List<{{.EntityName}}>), 200)]
    [ProducesResponseType(500)]
    [HttpGet("{id}")]
    public async Task<IActionResult> ById(string id)
    {
        var result = await _select{{.EntityName}}.ById(id);

        return Ok(result);
    }
}`
	t, err := template.New("controllerSelect").Parse(selectTemplate)

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
