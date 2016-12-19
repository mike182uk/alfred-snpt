package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func searchCommand(cliArgs []string, snptBin string, fzfBin string) {
	if len(cliArgs) < 2 {
		fmt.Print("Usage: snpt-alfred-workflow search <query>")

		os.Exit(1)
	}

	// get snippets
	snptLsCmd := exec.Command(snptBin, "ls") // #nosec

	var snptLsCmdOut bytes.Buffer
	snptLsCmd.Stdout = &snptLsCmdOut

	err := snptLsCmd.Run()

	if err != nil {
		// @TODO do something useful with error

		os.Exit(1)
	}

	// search snippets
	fzfCmd := exec.Command(fzfBin, "-f", cliArgs[1]) // #nosec
	fzfCmd.Stdin = strings.NewReader(snptLsCmdOut.String())

	var fzfCmdOut bytes.Buffer
	fzfCmd.Stdout = &fzfCmdOut

	err = fzfCmd.Run()

	if err != nil {
		// @TODO do something useful with error

		os.Exit(1)
	}

	// generate output and write to stdout
	fmt.Print(xml.Header)
	fmt.Print(generateSnippetsItemList(strings.Split(fzfCmdOut.String(), "\n")))

	os.Exit(0)
}
