package generate

import (
	"errors"
	"fmt"

	"github.com/lu-css/aspgen/src/general"
	"github.com/manifoldco/promptui"
)

var defaultTempalte = &promptui.SelectTemplates{
	Label:    "{{ . }}?",
	Active:   "\U000027A1 {{ . | cyan }}",
	Inactive: "  {{ . | cyan }} ",
	Selected: "\U0001F336 {{ . | red | cyan }}",
}

func ListActions(inSlnProject bool, slnName string) {
	inSlnActions := []string{
		"new",
		"run",
	}

	notInSlnActions := []string{
		"init",
	}

	var actions = []string{}

	if inSlnProject {
		actions = inSlnActions
	} else {
		actions = notInSlnActions
	}

	prompt := promptui.Select{
		Label:     "What you want to do",
		Items:     actions,
		Templates: defaultTempalte,
		Size:      8,
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		general.Exit()
		return
	}

	switch actions[i] {
	case "init":
		genGuardianProject()
	case "run":
		RunGuardianProject(slnName)
	}
}

func genGuardianProject() {
	validate := func(input string) error {
		if input == "" {
			return errors.New("Blank Text")
		}

		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Project name",
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	GenGuardianProject(result)
}
