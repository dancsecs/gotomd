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
	"io"
	"testing"

	"github.com/dancsecs/sztestlog"
)

func Test_Process_ConfirmOverwrite_Error(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	ok, err := confirm("the_file", "abc", "abc")
	chk.Err(err, io.EOF.Error())
	chk.False(ok)

	chk.Stdout(
		"Confirm overwrite of: the_file",
		"(Y to overwrite; [N/n] to Cancel; [R/r] to review)? ",
	)
}

func Test_Process_ConfirmOverwrite_InvalidCancel(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	chk.SetStdinData("Q\nN\n")

	ok, err := confirm("the_file", "abc", "abc")
	chk.NoErr(err)
	chk.False(ok)

	chk.Stdout(
		"Confirm overwrite of: the_file",
		"(Y to overwrite; [N/n] to Cancel; [R/r] to review)? "+
			"Unknown response: 'Q'",
		"",
		"Confirm overwrite of: the_file",
		"(Y to overwrite; [N/n] to Cancel; [R/r] to review)? "+
			"Overwrite cancelled",
		"",
	)
}

func Test_Process_ConfirmOverwrite_ReviewCancel(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	chk.SetStdinData("r\nn\n")

	ok, err := confirm("the_file", "abc1", "abc2")
	chk.NoErr(err)
	chk.False(ok)

	chk.Stdout(
		"Confirm overwrite of: the_file",
		"(Y to overwrite; [N/n] to Cancel; [R/r] to review)? "+
			"--- Old_the_file",
		"+++ New_the_file",
		"@@ -1 +1 @@",
		"-abc1",
		"+abc2",
		"",
		"Confirm overwrite of: the_file",
		"(Y to overwrite; [N/n] to Cancel; [R/r] to review)? "+
			"Overwrite cancelled",
		"",
	)
}

func Test_Process_ConfirmOverwrite_ReviewConfirm(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	chk.SetStdinData("R\nY\n")

	ok, err := confirm("the_file", "abc1", "abc2")
	chk.NoErr(err)
	chk.True(ok)

	chk.Stdout(
		"Confirm overwrite of: the_file",
		"(Y to overwrite; [N/n] to Cancel; [R/r] to review)? "+
			"--- Old_the_file",
		"+++ New_the_file",
		"@@ -1 +1 @@",
		"-abc1",
		"+abc2",
		"",
		"Confirm overwrite of: the_file",
		"(Y to overwrite; [N/n] to Cancel; [R/r] to review)? "+
			"",
	)
}
