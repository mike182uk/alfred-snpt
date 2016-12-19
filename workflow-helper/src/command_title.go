package main

import (
	"fmt"
	"os"
)

func titleCommand(cliArgs []string) {
	if len(cliArgs) < 2 {
		fmt.Print("Usage: snpt-alfred-workflow title <snippet>")

		os.Exit(1)
	}

	title := extractSnippetTitle(cliArgs[1])

	// default to "Snippet" if a title can not be extracted
	if title == "" {
		title = "Snippet"
	}

	fmt.Print(title)
	os.Exit(0)
}
