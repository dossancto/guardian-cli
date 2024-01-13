package scaffold

import (
	"bytes"
	"text/template"
	"github.com/lu-css/guardian-cli/src/commands/scaffold/structs"
)

func createRepositoryInterface(config structs.FeatureConfiguration) string {
	tmpl := `
using {{.SlnName}}.Application.Features.{{.FeatureName}}.Entities;

namespace {{.SlnName}}.Application.Features.{{.FeatureName}}.Data;

/// <summary>
/// Defines operations for a {{.EntityName}} repository.
/// </summary>
public interface I{{.EntityName}}Repository
{
   /// <summary>
   /// Save a new {{.EntityName}}
   /// </summary>
   /// <param name="entity">The entity to save</param>
   /// <returns>The created {{.EntityName}}</returns>
   Task<{{.EntityName}}> Save({{.EntityName}} entity);

   /// <summary>
   /// Update a existint {{.EntityName}}
   /// </summary>
   /// <param name="entity">The entity to update</param>
   /// <returns>The updated {{.EntityName}}</returns>
   Task<{{.EntityName}}> Update({{.EntityName}} entity);

   /// <summary>
   /// Search for a {{.EntityName}} using it Id
   /// </summary>
   /// <param name="id">The Entity Id</param>
   /// <returns>The search result or null, if not found</returns>
   Task<{{.EntityName}}?> FindById(string id);

   /// <summary>
   /// List all {{.EntityName}}
   /// </summary>
   /// <returns>A list of all {{.EntityName}}</returns>
   Task<List<{{.EntityName}}>> All();

   /// <summary>
   /// Delete a {{.EntityName}}
   /// </summary>
   Task Delete(string id);
}`
	t, err := template.New("createRepositoryInterface").Parse(tmpl)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	err = t.Execute(&buf, config)
	if err != nil {
		panic(err)
	}

	return buf.String()
}

func createRepositoryImpl(config structs.FeatureConfiguration) string {
	return ""
}
