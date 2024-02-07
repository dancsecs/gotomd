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
		"relative directory must be specified in cmd: \""+tstDir+"\"",
	)

	_, err = getGoTst(sampleGoProjectOnePath + "TEST_DOES_NOT_EXIST")
	chk.Err(err, "no tests to run")
}

func Test_GetFile_GetGoFile(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	d, err := getGoFile(sampleGoProjectOnePath + "crumb.go")
	chk.NoErr(err)
	chk.Str(
		d,
		""+
			markBashCode("cat "+sampleGoProjectOnePath+"crumb.go")+
			"\n\n"+
			markGoCode("package "+sampleGoProjectOne),
	)
}

func Test_GetFile_GetGoFile2(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	file1 := sampleGoProjectOnePath + "crumb.go"
	file2 := sampleGoProjectTwoPath + "crumb.go"

	d, err := getGoFile(file1 + " " + file2)
	chk.NoErr(err)
	chk.Str(
		d,
		""+
			markBashCode("cat "+file1)+
			"\n\n"+
			markGoCode("package "+sampleGoProjectOne)+
			"\n\n"+
			markBashCode("cat "+file2)+
			"\n\n"+
			markGoCode("package "+sampleGoProjectTwo)+
			"",
	)
}
