package search

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
	"os/exec"
	"strings"

	alfredHelper "github.com/mike182uk/snpt-alfred-workflow/cmd/helper/alfred"
)

// Run executes the search command
func Run(args []string, snptBin string, fzfBin string) {
	if len(args) < 2 {
		fmt.Print("Usage: snpt-alfred-workflow search <query>")

		os.Exit(1)
	}

	// get snippets
	snptLsCmd := exec.Command(snptBin, "ls") // #nosec

	var snptLsCmdOut bytes.Buffer
	snptLsCmd.Stdout = &snptLsCmdOut

	err := snptLsCmd.Run()

	if err != nil {
		fmt.Print("Failed to run: snpt ls")

		os.Exit(1)
	}

	// search snippets
	fzfCmd := exec.Command(fzfBin, "-f", args[1]) // #nosec
	fzfCmd.Stdin = strings.NewReader(snptLsCmdOut.String())

	var fzfCmdOut bytes.Buffer
	fzfCmd.Stdout = &fzfCmdOut

	err = fzfCmd.Run()

	if err != nil {
		fmt.Printf("Failed to run: fzf -f %s", args[1])

		os.Exit(1)
	}

	// generate output and write to stdout
	fmt.Print(xml.Header)
	fmt.Print(alfredHelper.GenerateSnippetsItemList(strings.Split(fzfCmdOut.String(), "\n")))

	os.Exit(0)
}
