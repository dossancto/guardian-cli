package generate

import (
	"log"

	"github.com/lu-css/guardian-cli/src/commands"
	"github.com/lu-css/guardian-cli/src/commands/scaffold"
	"github.com/lu-css/guardian-cli/src/utils"
	"github.com/manifoldco/promptui"
)

func inSlnActions(inSlnProject bool) []string {
	if inSlnProject {
		return []string{
			"run",
			"migrate",
			"scaffold",
		}
	}

	return []string{
		"init",
	}
}

func ListActions(inSlnProject bool, slnName string) {
	var actions = inSlnActions(inSlnProject)

	prompt := promptui.Select{
		Label:     "What you want to do",
		Items:     actions,
		Templates: utils.GetDefaultTemplate(),
		Size:      8,
	}

	_, cmd, err := prompt.Run()

	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}

	RunCommand(cmd, slnName)
}

func RunCommand(command string, slnName string) {
	switch command {
	case "init":
		commands.GenGuardianProject()
	case "run":
		commands.RunGuardianProject(slnName)
	case "migrate":
		commands.MigrateCommand(slnName)
	case "scaffold":
		scaffold.HandleScaffold(slnName)
	default:
		log.Fatalf("Command '%s' not found", command)
	}
}
