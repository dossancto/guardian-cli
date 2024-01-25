package commands

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/manifoldco/promptui"
)

func GenGuardianProject() {
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

	result = strings.Replace(result, " ", ".", -1)

	createGuardianProject(result)
  postGuardianCreate(result)
}

func createGuardianProject(projectName string) {
	cmd := exec.Command("dotnet", "new", "guardian", "-o", projectName)

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.CombinedOutput() failed with %s\n", err)
	}

	log.Println(string(output))
}

func postGuardianCreate(projectName string) {
	log.Println("Generating Makefile...")
	log.Println("Generating Dockerfile...")
	log.Println("Generating Readme...")
}
