package application

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/lu-css/guardian-cli/src/commands/scaffold/structs"
	"github.com/samber/lo"
)

func GenerateCreateDTO(config structs.FeatureConfiguration) string {
	docs := fmt.Sprintf("Represents a data object for creating a %s", config.EntityName)
	classname := fmt.Sprintf("Create%s", config.EntityName)
	return generateDTO(docs, classname, config)
}

func GenerateUpdateDTO(config structs.FeatureConfiguration) string {
	idField := structs.FeatureClass{
		FieldName: "Id",
		FieldType: "string",
	}

	config.Fields = append([]structs.FeatureClass{idField}, config.Fields...)

	docs := fmt.Sprintf("Represents a data object for updating a %s", config.EntityName)
	classname := fmt.Sprintf("Update%s", config.EntityName)
	return generateDTO(docs, classname, config)
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
	dto := GenerateUpdateDTO(config)

	return fmt.Sprintf("%s\n%s", usecase, dto)
}
