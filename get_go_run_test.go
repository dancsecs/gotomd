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
	"os"
	"testing"

	"github.com/dancsecs/sztest"
)

func Test_GetRun_GetGoRun(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	_, err := getGoRun("")
	chk.Err(
		err,
		ErrInvalidRelativeDir.Error()+": \"\"",
	)

	cmd := "TEST_DIRECTORY_DOES_NOT_EXIST" + string(os.PathSeparator)
	_, err = getGoRun(cmd)
	chk.Err(
		err,
		ErrInvalidRelativeDir.Error()+": \""+cmd+"\"",
	)

	_, err = getGoRun("./TEST_DOES_NOT_EXIST")
	chk.Err(err, ErrNoPackageToRun.Error())
}

func Test_GetRun_RunTestNotDirectory(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	f := chk.CreateTmpFile(nil)

	_, _, err := runGo(f, "")
	chk.Err(
		err,
		ErrInvalidDirectory.Error(),
	)
}

func Test_GetRun_RunTestNoPackage(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	_, _, err := runGo(".", "")
	chk.NoErr(err)
}

func Test_GetRun_RunExampleNoPackage(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	out, err := getGoRun("./example3/main.go -v")
	chk.NoErr(err)
	chk.Str(
		out,
		"```bash\n"+
			"go run example3/main.go -v\n"+
			"```\n"+
			"\n"+
			"<pre>\n"+
			"Running with 1 arguments\n"+
			"-v\n"+
			"</pre>",
	)
}
