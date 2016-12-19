package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func copyCommand(cliArgs []string, snptBin string) {
	if len(cliArgs) < 2 {
		fmt.Print("Usage: snpt-alfred-workflow copy <snippet>")

		os.Exit(1)
	}

	snptCpCmd := exec.Command(snptBin, "cp") // #nosec
	snptCpCmd.Stdin = strings.NewReader(cliArgs[1])

	err := snptCpCmd.Run()

	if err != nil {
		// @TODO do something useful with error

		os.Exit(1)
	}

	os.Exit(0)
}
