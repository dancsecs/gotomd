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

package file_test

import (
	"os"
	"strings"
	"testing"

	"github.com/dancsecs/gotomd/internal/errs"
	"github.com/dancsecs/gotomd/internal/file"
	"github.com/dancsecs/sztestlog"
)

const (
	catCmd   = "cat "
	pkgLabel = "package"
	sep      = string(os.PathSeparator)

	tstpkg1     = "tstpkg1"
	tstpkg1Path = "." + sep + "testdata" + sep + tstpkg1 + sep

	tstpkg2     = "tstpkg2"
	tstpkg2Path = "." + sep + "testdata" + sep + tstpkg2 + sep
)

func markGoCode(content string) string {
	return "```go\n" + strings.TrimRight(content, "\n") + "\n```"
}

func markBashCode(content string) string {
	return "```bash\n" + strings.TrimRight(content, "\n") + "\n```"
}

func Test_GetFile_GetGoFileInvalid(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	tstDir := "TEST_DIRECTORY_DOES_NOT_EXIST" + string(os.PathSeparator)
	_, err := file.GetGoFile(tstDir)
	chk.Err(
		err,
		errs.ErrInvalidRelativeDir.Error()+": \""+tstDir+"\"",
	)
}

func Test_GetFile_GetGoFile(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	d, err := file.GetGoFile(tstpkg1Path + "crumb.go")
	chk.NoErr(err)
	chk.Str(
		d,
		""+
			markBashCode(catCmd+tstpkg1Path+"crumb.go")+
			"\n\n"+
			markGoCode(pkgLabel+" "+tstpkg1),
	)
}

func Test_GetFile_GetGoFile2(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	file1 := tstpkg1Path + "crumb.go"
	file2 := tstpkg2Path + "crumb.go"

	d, err := file.GetGoFile(file1 + " " + file2)
	chk.NoErr(err)
	chk.Str(
		d,
		""+
			markBashCode(catCmd+file1)+
			"\n\n"+
			markGoCode(pkgLabel+" "+tstpkg1)+
			"\n\n"+
			markBashCode(catCmd+file2)+
			"\n\n"+
			markGoCode(pkgLabel+" "+tstpkg2)+
			"",
	)
}
