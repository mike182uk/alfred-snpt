package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractSnippetTitle(t *testing.T) {
	assert.Equal(t, "foo.bar", extractSnippetTitle("foo.bar - baz qux [abc123]"))
	assert.Equal(t, "", extractSnippetTitle("foo bar baz"))
}

func TestExtractSnippetSubTitle(t *testing.T) {
	assert.Equal(t, "baz qux", extractSnippetSubTitle("foo.bar - baz qux [abc123]"))
	assert.Equal(t, "", extractSnippetSubTitle("foo bar baz"))
}

func TestGenerateSnippetsItemList(t *testing.T) {
	var input = []string{"foo - bar [baz]", "baz - bar [foo]", "bar - baz [foo]"}
	var expected = `<items>
	<item arg="foo - bar [baz]">
		<title>foo</title>
		<subtitle>bar</subtitle>
	</item>
	<item arg="baz - bar [foo]">
		<title>baz</title>
		<subtitle>bar</subtitle>
	</item>
	<item arg="bar - baz [foo]">
		<title>bar</title>
		<subtitle>baz</subtitle>
	</item>
</items>`

	assert.Equal(t, expected, generateSnippetsItemList(input))

	var badInput = []string{"foo bar baz", "foo-", "\n", "bar - baz [foo]"}
	var badInputExpected = `<items>
	<item arg="bar - baz [foo]">
		<title>bar</title>
		<subtitle>baz</subtitle>
	</item>
</items>`

	assert.Equal(t, badInputExpected, generateSnippetsItemList(badInput))
}

func TestGenerateErrorItemList(t *testing.T) {
	var expected = `<items>
	<item>
		<title>foo bar baz</title>
		<icon>icon-error.png</icon>
	</item>
</items>`

	assert.Equal(t, expected, generateErrorItemList("foo bar baz", ""))

	var expectedWithSubtitle = `<items>
	<item>
		<title>foo bar baz</title>
		<subtitle>baz bar foo</subtitle>
		<icon>icon-error.png</icon>
	</item>
</items>`

	assert.Equal(t, expectedWithSubtitle, generateErrorItemList("foo bar baz", "baz bar foo"))
}
