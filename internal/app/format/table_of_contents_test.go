package format

import (
	"testing"
)

func TestTableOfContentsFromFilePaths(t *testing.T) {
	paths := []string{
		"json",
		"json/omitempty",
		"json/omitempty/marshal_string.md",
		"json/omitempty/unmarshal.md",
	}

	out := TableOfContentsFromFilePaths(paths)

	expected := `* [json]
  * [omitempty]
    * [marshal_string.md](json/omitempty/marshal_string.md)
    * [unmarshal.md](json/omitempty/unmarshal.md)
`
	if expected != out {
		t.Fail()
	}
}
