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

	"github.com/dancsecs/sztestlog"
)

func TestDiff_Same(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	diffFile(
		"diff",
		"abc\n",
		"abc\n",
	)

	chk.Stdout("")
}

func TestDiff_Diff(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	diffFile(
		"diff",
		"abc\na1c\ndef\n",
		"abc\na2c\ndef\n",
	)

	chk.Stdout(
		"--- Old_diff",
		"+++ New_diff",
		"@@ -1,3 +1,3 @@",
		" abc",
		"-a1c",
		"+a2c",
		" def",
		"",
	)
}
