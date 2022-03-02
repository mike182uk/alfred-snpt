package title

import (
	"fmt"
	"os"

	snippetHelper "alfred-snpt/cmd/helper/snippet"
)

// Run executes the title command
func Run(args []string) {
	if len(args) < 2 {
		fmt.Print("Usage: alfred-snpt title <snippet>")

		os.Exit(1)
	}

	title := snippetHelper.ExtractSnippetTitle(args[1])

	// default to "Snippet" if a title can not be extracted
	if title == "" {
		title = "Snippet"
	}

	fmt.Print(title)

	os.Exit(0)
}
