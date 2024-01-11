package commands

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/lu-css/guardian-cli/src/general"
	"github.com/lu-css/guardian-cli/src/utils"
	"github.com/manifoldco/promptui"
)

func MigrateCommand(slnName string) {
	actions := []string{
		"new migration",
		"update",
	}

	prompt := promptui.Select{
		Label:     "What you want to do",
		Items:     actions,
		Templates: utils.GetDefaultTemplate(),
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

func NewMigration(migrationName string, slnName string) {
	os.Setenv("TERM", "xterm-256color")
	cmd := exec.Command("dotnet", "ef", "migrations", "add", migrationName, "-s", slnName+".UI", "-p", slnName+".Infra")

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		log.Fatalf("cmd.CombinedOutput() failed with %s\n", err)
	}

	cmd.Start()

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, "info: ") {
			color.Set(color.FgGreen)
			color.Set(color.BgBlack)
		} else {
			color.Set(color.FgWhite)
		}
		fmt.Println(text)
		color.Unset()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
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

func UpdateDatabase(slnName string) {

	os.Setenv("TERM", "xterm-256color")
	cmd := exec.Command("dotnet", "ef", "database", "update", "-s", slnName+".UI")

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		log.Fatalf("cmd.CombinedOutput() failed with %s\n", err)
	}

	cmd.Start()

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, "info: ") {
			color.Set(color.FgGreen)
			color.Set(color.BgBlack)
		} else {
			color.Set(color.FgWhite)
		}
		fmt.Println(text)
		color.Unset()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
