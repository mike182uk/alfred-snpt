package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

type itemList struct {
	XMLName xml.Name       `xml:"items"`
	Items   []itemListItem `xml:"item"`
}

type itemListItem struct {
	Arg      string `xml:"arg,attr,omitempty"`
	Title    string `xml:"title"`
	SubTitle string `xml:"subtitle,omitempty"`
	Icon     string `xml:"icon,omitempty"`
}

const snptBin = "snpt"
const fzfBin = "fzf"

func main() {
	cliArgs := os.Args[1:]

	if len(cliArgs) == 0 {
		fmt.Print("Usage: snpt-alfred-workflow <search|cp|title>")

		os.Exit(1)
	}

	checkDeps()

	switch cliArgs[0] {
	case "search":
		searchCmd(cliArgs)
	case "cp":
		copyCmd(cliArgs)
	case "title":
		titleCmd(cliArgs)
	}
}

func checkDeps() {
	whichSnptCmd := exec.Command("which", "snpt") // #nosec
	whichFzfCmd := exec.Command("which", "fzf")   // #nosec

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

func searchCmd(cliArgs []string) {
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

func copyCmd(cliArgs []string) {
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

func titleCmd(cliArgs []string) {
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

func generateSnippetsItemList(items []string) string {
	list := itemList{}

	for _, i := range items {
		t := extractSnippetTitle(i)
		st := extractSnippetSubTitle(i)

		// if item is a new line or does not have a title or subtitle, skip it
		if i == "\n" || t == "" || st == "" {
			continue
		}

		list.Items = append(list.Items, itemListItem{
			Arg:      i,
			Title:    t,
			SubTitle: st,
		})
	}

	output, err := xml.MarshalIndent(list, "", "	")

	if err != nil {
		return ""
	}

	return string(output)
}

func generateErrorItemList(title string, subtitle string) string {
	list := itemList{}
	item := itemListItem{}

	item.Title = title
	item.Icon = "icon-error.png"

	if subtitle != "" {
		item.SubTitle = subtitle
	}

	list.Items = append(list.Items, item)

	output, err := xml.MarshalIndent(list, "", "	")

	if err != nil {
		return ""
	}

	return string(output)
}

func extractSnippetTitle(s string) string {
	i := strings.Index(s, " - ")

	if i == -1 {
		return ""
	}

	return s[0:i]
}

func extractSnippetSubTitle(s string) string {
	r := regexp.MustCompile(`[[A-Za-z0-9]+]`)
	s = strings.TrimRight(r.ReplaceAllString(s, ""), " ")
	i := strings.LastIndex(s, " - ")

	if i == -1 {
		return ""
	}

	return s[(i + 3):]
}
