package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/lu-css/aspgen/src/generate"
	"github.com/lu-css/aspgen/src/validations"
)

func main() {
	args := os.Args[1:]

	inSlnProject, slnNameFile := validations.ExistsSln()

	slnName := strings.TrimSuffix(slnNameFile, filepath.Ext(slnNameFile))

	if len(args) > 0 {
		cmd := args[0]
		generate.RunCommand(cmd, slnName)
    return
	}

	generate.ListActions(inSlnProject, slnName)
}
