package copy

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Run executes the copy command
func Run(args []string, snptBin string) {
	if len(args) < 2 {
		fmt.Print("Usage: snpt-alfred-workflow copy <snippet>")

		os.Exit(1)
	}

	snptCpCmd := exec.Command(snptBin, "cp") // #nosec
	snptCpCmd.Stdin = strings.NewReader(args[1])

	err := snptCpCmd.Run()

	if err != nil {
		fmt.Printf("Failed to run: snpt cp %s", args[1])

		os.Exit(1)
	}

	os.Exit(0)
}
