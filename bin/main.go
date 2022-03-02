package main

import (
	"bytes"
	"os/exec"
	"regexp"
	"strings"

	aw "github.com/deanishe/awgo"
)

var (
	wf      *aw.Workflow
	errIcon *aw.Icon = &aw.Icon{Value: "icon-error.png"}
)

func getSnippets() (snippets []string, err error) {
	var cmdOut bytes.Buffer

	cmd := exec.Command("snpt", "ls")
	cmd.Stdout = &cmdOut

	err = cmd.Run()

	if err != nil {
		return
	}

	snippets = strings.Split(cmdOut.String(), "\n")

	return
}

func getTitle(s string) string {
	i := strings.Index(s, " - ")

	if i == -1 {
		return ""
	}

	return s[0:i]
}

func getSubtitle(s string) string {
	r := regexp.MustCompile(`[[A-Za-z0-9]+]`)
	s = strings.TrimRight(r.ReplaceAllString(s, ""), " ")
	i := strings.LastIndex(s, " - ")

	if i == -1 {
		return ""
	}

	return s[(i + 3):]
}

func run() {
	// Retrieve snippets
	snippets, err := getSnippets()

	if err != nil {
		wf.NewItem("An error occured whilst executing snpt").
			Subtitle(err.Error()).
			Icon(errIcon)

		wf.SendFeedback()

		return
	}

	// Build items
	for _, snippet := range snippets {
		wf.NewItem(snippet).
			Title(getTitle(snippet)).
			Subtitle(getSubtitle(snippet)).
			Arg(snippet).
			Valid(true)
	}

	// Filter items
	args := wf.Args()

	if len(args) > 0 {
		query := args[0]

		if query != "" {
			wf.Filter(query)
		}
	}

	// Send Feedback
	if wf.IsEmpty() {
		wf.Feedback.Clear()

		wf.NewItem("No matching snippets found").
			Subtitle("Try updating your query").
			Icon(errIcon)
	}

	wf.SendFeedback()
}

func main() {
	wf = aw.New()

	wf.Run(run)
}
