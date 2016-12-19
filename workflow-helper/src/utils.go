package main

import (
	"encoding/xml"
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
