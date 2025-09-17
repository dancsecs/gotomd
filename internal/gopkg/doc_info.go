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

package gopkg

import (
	"strings"
)

// DocInfo provides functions to return formatted go documentation.
type DocInfo struct {
	header []string
	body   []string
	doc    []string
}

// Header returns the documentation header.
func (di *DocInfo) Header() []string {
	return di.header
}

// Body returns the objects body.
func (di *DocInfo) Body() []string {
	return di.body
}

// Doc returns the objects documentation.
func (di *DocInfo) Doc() []string {
	return di.doc
}

// OneLine returns a string representing the go object's declaration on a
// single line.
func (di *DocInfo) OneLine() string {
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

// NaturalComment returns a go object's comments exactly as they appear in
// its go source file.
func (di *DocInfo) NaturalComment() string {
	res := ""
	for _, l := range di.doc {
		if res != "" {
			res += "\n"
		}

		res += "// " + l
	}

	return res
}

// Declaration returns a go object's declaration exactly as it appears in
// its go source file.
func (di *DocInfo) Declaration() string {
	return strings.Join(di.header, "\n")
}

// Comment returns a clean view of the object's comments.
func (di *DocInfo) Comment() string {
	return strings.Join(di.doc, "\n")
}

// ConstantBlock returns a constant block formatted as it would in a go
// source file.
func (di *DocInfo) ConstantBlock() string {
	return di.NaturalComment() + "\n" +
		strings.Join(di.body, "\n")
}
