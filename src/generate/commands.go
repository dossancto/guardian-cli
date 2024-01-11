package generate

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
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
		fmt.Println(text)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
