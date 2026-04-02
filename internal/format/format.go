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

package format

import "strings"

//nolint:goCheckNoGlobals // Ok.
var formatForGo bool

// ForMarkdown sets the format style.
func ForMarkdown() {
	formatForGo = false
}

// ForGoDoc code for markdown templates.
func ForGoDoc() {
	formatForGo = true
}

func markForGoPackageInline(content string) string {
	contentLines := strings.Split(content, "\n")
	newContent := make([]string, len(contentLines))

	for i, l := range contentLines {
		newContent[i] = "\t" + l
	}

	return "\n" + strings.Join(newContent, "\n") + "\n"
}

// Inline frames the content in a ```language ... ``` multiline block for an
// .md output and prefixes each body line with a tab "\t" character for a
// go package document.
func Inline(language, body string) string {
	body = strings.Trim(body, "\n \t")
	if body == "" {
		return ""
	}

	if formatForGo {
		return markForGoPackageInline(body)
	}

	return "```" + language + "\n" + body + "\n```"
}

// Comment creates a stand alone comment.
func Comment(line string) string {
	if formatForGo {
		return "// " + line + ".\n"
	}

	return "<!--- " + line + " -->\n"
}

// BalancedComment returns the string centered in a comment line.
func BalancedComment(line string) string {
	extra := 79

	if formatForGo {
		extra -= 4
	} else {
		extra -= 10
	}

	extra = (extra - len(line)) / 2 //nolint:mnd // Ok Half.
	if extra > 0 {
		line = strings.Repeat(" ", extra) + line
	}

	return Comment(line)
}

// HLine returns a horizontal line.
func HLine() string {
	const lineLength = 78
	if formatForGo {
		return strings.Repeat("-", lineLength) + "\n"
	}

	return "---\n"
}
