package helper

import (
	"encoding/xml"

	snippetHelper "alfred-snpt/cmd/helper/snippet"
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

// GenerateSnippetsItemList generates an XML item list for use by Alfred
func GenerateSnippetsItemList(items []string) string {
	list := itemList{}

	for _, i := range items {
		t := snippetHelper.ExtractSnippetTitle(i)
		st := snippetHelper.ExtractSnippetSubTitle(i)

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

// GenerateErrorItemList generates an XML error item list for use by Alfred
func GenerateErrorItemList(title string, subtitle string) string {
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
