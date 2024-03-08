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

func Test_GetFile_GetGoFileInvalid(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	tstDir := "TEST_DIRECTORY_DOES_NOT_EXIST" + string(os.PathSeparator)
	_, err := getGoFile(tstDir)
	chk.Err(
		err,
		ErrInvalidRelativeDir.Error()+": \""+tstDir+"\"",
	)

	_, err = getGoTst(example1Path + "TEST_DOES_NOT_EXIST")
	chk.Err(err, ErrNoTestToRun.Error())
}

func Test_GetFile_GetGoFile(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	d, err := getGoFile(example1Path + "crumb.go")
	chk.NoErr(err)
	chk.Str(
		d,
		""+
			markBashCode(catCmd+example1Path+"crumb.go")+
			"\n\n"+
			markGoCode(pkgLabel+" "+example1),
	)
}

func Test_GetFile_GetGoFile2(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	file1 := example1Path + "crumb.go"
	file2 := example2Path + "crumb.go"

	d, err := getGoFile(file1 + " " + file2)
	chk.NoErr(err)
	chk.Str(
		d,
		""+
			markBashCode(catCmd+file1)+
			"\n\n"+
			markGoCode(pkgLabel+" "+example1)+
			"\n\n"+
			markBashCode(catCmd+file2)+
			"\n\n"+
			markGoCode(pkgLabel+" "+example2)+
			"",
	)
}
