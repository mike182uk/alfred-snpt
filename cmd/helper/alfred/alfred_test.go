package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateSnippetsItemList(t *testing.T) {
	input := []string{"foo - bar [baz]", "baz - bar [foo]", "bar - baz [foo]"}
	expected := `<items>
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

	assert.Equal(t, expected, GenerateSnippetsItemList(input))

	badInput := []string{"foo bar baz", "foo-", "\n", "bar - baz [foo]"}
	badInputExpected := `<items>
	<item arg="bar - baz [foo]">
		<title>bar</title>
		<subtitle>baz</subtitle>
	</item>
</items>`

	assert.Equal(t, badInputExpected, GenerateSnippetsItemList(badInput))
}

func TestGenerateErrorItemList(t *testing.T) {
	expected := `<items>
	<item>
		<title>foo bar baz</title>
		<icon>icon-error.png</icon>
	</item>
</items>`

	assert.Equal(t, expected, GenerateErrorItemList("foo bar baz", ""))

	expectedWithSubtitle := `<items>
	<item>
		<title>foo bar baz</title>
		<subtitle>baz bar foo</subtitle>
		<icon>icon-error.png</icon>
	</item>
</items>`

	assert.Equal(t, expectedWithSubtitle, GenerateErrorItemList("foo bar baz", "baz bar foo"))
}
