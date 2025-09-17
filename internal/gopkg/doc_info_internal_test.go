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
	"testing"

	"github.com/dancsecs/sztestlog"
)

func Test_DocInfo_OneLine(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	dInfo := new(DocInfo)

	chk.Str(dInfo.OneLine(), "UNKNOWN DECLARATION")
}

func Test_DocInfo_NaturalComments(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	dInfo := new(DocInfo)
	dInfo.doc = append(dInfo.doc, "a", "b")

	chk.Str(
		dInfo.NaturalComment(),
		"// a\n// b",
	)
}
