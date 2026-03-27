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
var formatForPackage bool

// ForMarkdown sets the format style.
func ForMarkdown() {
	formatForPackage = false
}

// ForGoDoc code for markdown templates.
func ForGoDoc() {
	formatForPackage = true
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

	if formatForPackage {
		return markForGoPackageInline(body)
	}

	return "```" + language + "\n" + body + "\n```"
}
