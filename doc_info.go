/*
   Golang To Github Markdown Utility: gotomd
   Copyright (C) 2023, 2024 Leslie Dancsecs

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

package main

import (
	"strings"
)

type docInfo struct {
	header []string
	body   []string
	doc    []string
}

func (di *docInfo) oneLine() string {
	res := ""

	switch len(di.header) {
	case 0:
		switch len(di.body) {
		case 0:
			res = "UNKNOWN DECLARATION"
		case 1:
			res = di.body[0]
		default:
			res = di.body[0] + " ..."
		}
	case 1:
		res = di.header[0]
	default:
		for _, l := range di.header {
			if strings.HasSuffix(res, ",") {
				res += " "
			}

			res += strings.TrimSpace(l)
		}

		res = strings.ReplaceAll(res, ", )", ")")
	}

	return res
}

func (di *docInfo) naturalComments() string {
	res := ""
	for _, l := range di.doc {
		if res != "" {
			res += "\n"
		}

		res += "// " + l
	}

	return res
}

func (di *docInfo) declGoLang() string {
	return markGoCode(strings.Join(di.header, "\n"))
}

func (di *docInfo) docMarkdown() string {
	return strings.Join(di.doc, "\n")
}
