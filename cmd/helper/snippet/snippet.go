package snippet

import (
	"regexp"
	"strings"
)

// ExtractSnippetTitle extracts the title of a snippet from a string
func ExtractSnippetTitle(s string) string {
	i := strings.Index(s, " - ")

	if i == -1 {
		return ""
	}

	return s[0:i]
}

// ExtractSnippetSubTitle extracts the subtitle of a snippet from a string
func ExtractSnippetSubTitle(s string) string {
	r := regexp.MustCompile(`[[A-Za-z0-9]+]`)
	s = strings.TrimRight(r.ReplaceAllString(s, ""), " ")
	i := strings.LastIndex(s, " - ")

	if i == -1 {
		return ""
	}

	return s[(i + 3):]
}
