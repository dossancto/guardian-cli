package generate

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

func GenGuardianProject(projectName string) {
	cmd := exec.Command("dotnet", "new", "guardian", "-o", projectName)

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.CombinedOutput() failed with %s\n", err)
	}

	log.Println(string(output))
}

func RunGuardianProject(slnName string) {
	os.Setenv("TERM", "xterm-256color")
	cmd := exec.Command("dotnet", "run", "--project", slnName+".UI")

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
