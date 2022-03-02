package snippet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractSnippetTitle(t *testing.T) {
	assert.Equal(t, "foo.bar", ExtractSnippetTitle("foo.bar - baz qux [abc123]"))
	assert.Equal(t, "", ExtractSnippetTitle("foo bar baz"))
}

func TestExtractSnippetSubTitle(t *testing.T) {
	assert.Equal(t, "baz qux", ExtractSnippetSubTitle("foo.bar - baz qux [abc123]"))
	assert.Equal(t, "", ExtractSnippetSubTitle("foo bar baz"))
}
