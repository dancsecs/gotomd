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

package gorun_test

import (
	"os"
	"testing"

	"github.com/dancsecs/gotomd/internal/errs"
	"github.com/dancsecs/gotomd/internal/gorun"
	"github.com/dancsecs/sztestlog"
)

func Test_GetRun_GetGoRun(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	_, err := gorun.GetGoRun("")
	chk.Err(
		err,
		errs.ErrInvalidRelativeDir.Error()+": \"\"",
	)

	cmd := "TEST_DIRECTORY_DOES_NOT_EXIST" + string(os.PathSeparator)
	_, err = gorun.GetGoRun(cmd)
	chk.Err(
		err,
		errs.ErrInvalidRelativeDir.Error()+": \""+cmd+"\"",
	)

	_, err = gorun.GetGoRun("./TEST_DOES_NOT_EXIST")
	chk.Err(err, errs.ErrNoPackageToRun.Error())
}

func Test_GetRun_RunTestNotDirectory(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	f := chk.CreateTmpFile(nil)

	_, _, err := gorun.RunGo(f, "")
	chk.Err(
		err,
		errs.ErrInvalidDirectory.Error(),
	)
}

func Test_GetRun_RunTestNoPackage(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	_, _, err := gorun.RunGo(".", "")
	chk.NoErr(err)
}

func Test_GetRun_RunExampleNoPackage(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	out, err := gorun.GetGoRun("./testdata/tstpkg/main.go -v")
	chk.NoErr(err)
	chk.Str(
		out,
		"---\n```bash\n"+
			"go run ./testdata/tstpkg/main.go -v\n"+
			"```\n"+
			"\n"+
			"<pre>\n"+
			"Running with 1 arguments\n"+
			"-v\n"+
			"</pre>\n---",
	)
}
