package infra

import (
	"bytes"
	"github.com/lu-css/guardian-cli/src/commands/scaffold/structs"
	"text/template"
)

func DefaultRepositoryImplContent(config structs.FeatureConfiguration) string {
	tmpl := `
using {{.SlnName}}.Application.Features.{{.FeatureName}}.Data;
using {{.SlnName}}.Application.Features.{{.FeatureName}}.Entities;
using {{.SlnName}}.Infra.Database.EF.Contexts;

using NanoidDotNet;
using Microsoft.EntityFrameworkCore;

namespace {{.SlnName}}.Infra.Database.EF.Repositories;

public class EF{{.EntityName}}Repository : I{{.EntityName}}Repository
{
    private readonly ApplicationDbContext _context;

    public EF{{.EntityName}}Repository(ApplicationDbContext context)
    => _context = context;

    public async Task<List<{{.EntityName}}>> All()
      => await _context.{{.EntityName}}s
                       .AsNoTrackingWithIdentityResolution()
                       .ToListAsync();

    public async Task Delete(string id)
    {
        var entity = await _context.{{.EntityName}}s.Where(x => x.Id == id).FirstOrDefaultAsync();

        if (entity is null)
            return;

        _context.{{.EntityName}}s.Remove(entity);

        await _context.SaveChangesAsync();
    }

    public async Task<{{.EntityName}}?> FindById(string id)
      => await _context.{{.EntityName}}s
                       .AsNoTrackingWithIdentityResolution()
                       .Where(x => x.Id == id)
                       .SingleOrDefaultAsync();

    public async Task<{{.EntityName}}> Save({{.EntityName}} entity)
    {
        entity.Id = Nanoid.Generate();

        var createdUser = _context.{{.EntityName}}s.Add(entity);

        await _context.SaveChangesAsync();

        return createdUser.Entity;
    }

    public async Task<{{.EntityName}}> Update({{.EntityName}} entity)
    {
        var updatedPlayer = _context.{{.EntityName}}s.Update(entity);

        await _context.SaveChangesAsync();

        return updatedPlayer.Entity;
    }
}
`
	t, err := template.New("Entity Configuration Content").Parse(tmpl)

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
