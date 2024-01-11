package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/lu-css/aspgen/src/generate"
	"github.com/lu-css/aspgen/src/validations"
)

func main() {
	inSlnProject, slnNameFile := validations.ExistsSln()
	slnName := strings.TrimSuffix(slnNameFile, filepath.Ext(slnNameFile))
  fmt.Println(slnName)

	generate.ListActions(inSlnProject, slnName)
}
