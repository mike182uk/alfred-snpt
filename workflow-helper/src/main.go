package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	snptBin := "snpt"
	fzfBin := "fzf"
	cliArgs := os.Args[1:]

	if len(cliArgs) == 0 {
		fmt.Print("Usage: snpt-alfred-workflow <search|cp|title>")

		os.Exit(1)
	}

	checkDeps(snptBin, fzfBin)

	switch cliArgs[0] {
	case "search":
		searchCommand(cliArgs, snptBin, fzfBin)
	case "cp":
		copyCommand(cliArgs, snptBin)
	case "title":
		titleCommand(cliArgs)
	}
}

func checkDeps(snptBin string, fzfBin string) {
	whichSnptCmd := exec.Command("which", snptBin) // #nosec
	whichFzfCmd := exec.Command("which", fzfBin)   // #nosec

	_, err := whichSnptCmd.Output()

	if err != nil {
		fmt.Print(generateErrorItemList("snpt was not found", "Make sure you have installed snpt to use this workflow."))

		os.Exit(1)
	}

	_, err = whichFzfCmd.Output()

	if err != nil {
		fmt.Print(generateErrorItemList("fzf was not found", "Make sure you have installed fzf to use this workflow."))

		os.Exit(1)
	}
}
