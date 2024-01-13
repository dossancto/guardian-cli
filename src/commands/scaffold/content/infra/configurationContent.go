package infra

import (
	"bytes"
	"github.com/lu-css/guardian-cli/src/commands/scaffold/structs"
	"text/template"
)

func EntityConfigurationContent(config structs.FeatureConfiguration) string {
	tmpl := `
using {{.SlnName}}.Application.Features.{{.FeatureName}}.Entities;
using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Metadata.Builders;

namespace {{.SlnName}}.Infra.Database.EF.Configurations;

public class {{.EntityName}}Configuration : IEntityTypeConfiguration<{{.EntityName}}>
{
    public void Configure(EntityTypeBuilder<{{.EntityName}}> builder)
    {
        builder.HasKey(x => x.Id);
        builder.Property(x => x.Id).ValueGeneratedOnAdd();

        builder.Property(x => x.Id).HasMaxLength(40);
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
