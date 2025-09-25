/*
   Golang To Github Markdown Utility: gotomd
   Copyright (C) 2025 Leslie Dancsecs

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package update

import "strings"

//nolint:goCheckNoGlobals // Ok.
var formatForPackage bool

// FormatForMarkdown sets the format style.
func FormatForMarkdown() {
	formatForPackage = false
}

// FormatForGoDoc code for markdown templates.
func FormatForGoDoc() {
	formatForPackage = true
}

func markForGo(content string) string {
	contentLines := strings.Split(content, "\n")
	newContent := make([]string, len(contentLines))

	for i, l := range contentLines {
		newContent[i] = "\t" + l
	}

	return strings.Join(newContent, "\n")
}

// MarkGoCode frames the content in go code braces (```go ...```) if
// processing a markdown document or by prefixing each line with a tab (\t)
// character for go package documentation.
func MarkGoCode(content string) string {
	content = strings.TrimRight(content, "\n")

	if formatForPackage {
		return markForGo(content)
	}

	return "```go\n" + content + "\n```"
}

// MarkBashCode frames the content in go code braces (```go ...```) if
// processing a markdown document or by prefixing each line with a tab (\t)
// character for go package documentation.
func MarkBashCode(content string) string {
	content = strings.TrimRight(content, "\n")

	if formatForPackage {
		return markForGo(content)
	}

	return "```bash\n" + content + "\n```"
}
