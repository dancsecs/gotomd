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

import (
	"strings"
	"testing"

	"github.com/dancsecs/sztestlog"
)

func TestDiff_Same(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	chk.Str(
		diffFile(
			"diff",
			"abc",
			"abc\n",
		),
		"",
	)
}

func TestDiff_Diff(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	diff := diffFile(
		"diff",
		"abc\na1c\ndef",
		"abc\na2c\ndef",
	)

	chk.StrSlice(
		strings.Split(diff, "\n"),
		[]string{
			"--- Old_diff",
			"+++ New_diff",
			"@@ -1,3 +1,3 @@",
			" abc",
			"-a1c",
			"+a2c",
			" def",
			"",
		},
	)
}
