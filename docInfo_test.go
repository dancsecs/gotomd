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
	"testing"

	"github.com/dancsecs/sztest"
)

func Test_DocInfo_OneLine(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	dInfo := &docInfo{}

	chk.Str(dInfo.oneLine(), "UNKNOWN DECLARATION")
}

func Test_DocInfo_NaturalComments(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	dInfo := &docInfo{}
	dInfo.doc = append(dInfo.doc, "a", "b")

	chk.Str(
		dInfo.naturalComments(),
		"// a\n// b",
	)
}
