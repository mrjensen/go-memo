package format

import (
	"fmt"
	"os"
	"strings"
)

const indentDepth = "  "

// TableOfContentsFromFilePaths returns a markdown list
// Input must be in alphabetical order (as from filepath.Walk)
func TableOfContentsFromFilePaths(paths []string) string {
	// store if a part of the table of contents already is outputted - needed for sections with multiple child elements
	isAlreadyOutputted := make(map[int]map[string]bool)

	output := ""
	for _, path := range paths {
		pathParts := strings.Split(path, string(os.PathSeparator))
		for depth, part := range pathParts {
			isMarkdown := strings.HasSuffix(part, ".md")
			if !isAlreadyOutputted[depth][part] {
				indent := strings.Repeat(indentDepth, depth)
				listItem := ""
				if isMarkdown {
					listItem = fmt.Sprintf("%s* [%s](%s)\n", indent, part, path)
				} else {
					listItem = fmt.Sprintf("%s* [%s]\n", indent, part)
				}
				output = output + listItem
				if len(isAlreadyOutputted[depth]) == 0 {
					isAlreadyOutputted[depth] = make(map[string]bool)
				}
				isAlreadyOutputted[depth][part] = true
			}
		}
	}
	return output
}
