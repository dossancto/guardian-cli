package generate

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/lu-css/guardian-cli/src/general"
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
		"run",
		"migrate",
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

	RunCommand(actions[i], slnName)
}

func RunCommand(command string, slnName string) {
	switch command {
	case "init":
		genGuardianProject()
	case "run":
		RunGuardianProject(slnName)
	case "migrate":
		migrateProject(slnName)
	default:
		log.Fatal("Command not found")
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

func migrateProject(slnName string) {
	actions := []string{
		"new migration",
		"update",
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

	switch i {
	case (0):
		NewMigration(GetMigrationName(), slnName)
	case (1):
		UpdateDatabase(slnName)
	}
}

func GetMigrationName() string {
	validate := func(input string) error {
		if input == "" {
			return errors.New("Blank Text")
		}

		if strings.Contains(input, " ") {
			return errors.New("Can't have space")
		}

		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Migration name",
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}

	return result
}
