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

package args

import (
	"path/filepath"
	"testing"

	"github.com/dancsecs/sztestlog"
)

func TestFiles_Expand_Empty(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	Reset()

	dir := chk.CreateTmpDir()
	file1 := chk.CreateTmpFileAs("", ".1"+GoTemplate, nil)
	file2 := chk.CreateTmpFileAs("", ".2"+MdTemplate, nil)

	t.Chdir(dir)

	chk.NoErr(expand([]string{""}))
	chk.NoErr(expand([]string{"."}))

	chk.StrSlice(
		GoFiles(),
		[]string{
			filepath.Base(file1),
		},
	)

	chk.StrSlice(
		MdFiles(),
		[]string{
			filepath.Base(file2),
		},
	)

	chk.Stdout(
		"File to process: '"+".1"+GoTemplate+"'",
		"File to process: '"+".2"+MdTemplate+"'",
		"Excluding redundant template: '"+".1"+GoTemplate+"'",
		"Excluding redundant template: '"+".2"+MdTemplate+"'",
	)
}

//nolint:funlen // Ok.
func TestFiles_Expand(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	Reset()

	dir := chk.CreateTmpDir()
	file1 := chk.CreateTmpFileAs("", ".1"+GoTemplate, nil)
	file2 := chk.CreateTmpFileAs("", ".2"+MdTemplate, nil)

	sub1 := chk.CreateTmpSubDir("sub1")
	file1_1 := chk.CreateTmpFileAs(sub1, ".1_1"+GoTemplate, nil)
	file1_2 := chk.CreateTmpFileAs(sub1, ".1_2"+MdTemplate, nil)
	_ = chk.CreateTmpFileAs(sub1, "ignore.txt", nil)

	sub2 := chk.CreateTmpSubDir("sub2")
	file2_1 := chk.CreateTmpFileAs(sub2, ".1_1"+GoTemplate, nil)
	file2_2 := chk.CreateTmpFileAs(sub2, ".1_2"+MdTemplate, nil)

	chk.Err(
		expand([]string{"DOES_NOT_EXIST"}),
		chk.ErrChain(
			ErrInvalidTemplate,
			ErrUnknownObject,
			"stat DOES_NOT_EXIST",
			"no such file or directory",
		),
	)

	chk.NoErr(expand([]string{dir + "/..."}))

	chk.StrSlice(
		GoFiles(),
		[]string{
			file1,
			file1_1,
			file2_1,
		},
	)

	chk.StrSlice(
		MdFiles(),
		[]string{
			file2,
			file1_2,
			file2_2,
		},
	)

	// Will be ignored.
	_ = chk.CreateTmpFileAs(sub2, "2."+MdTemplate, nil)

	chk.NoErr(expand([]string{dir + "/..."}))

	chk.StrSlice(
		GoFiles(),
		[]string{
			file1,
			file1_1,
			file2_1,
		},
	)

	chk.StrSlice(
		MdFiles(),
		[]string{
			file2,
			file1_2,
			file2_2,
		},
	)

	fileBad := chk.CreateTmpFileAs("", "."+MdTemplate, nil)

	chk.Err(
		expand([]string{fileBad}),
		chk.ErrChain(
			ErrInvalidTemplate,
			ErrUnknownObject,
			ErrInvalidArgument,
			"'"+fileBad+"'",
			"expected - ("+GoTemplate+" or "+MdTemplate+")",
		),
	)

	fileGoodGo := chk.CreateTmpFileAs("", ".fileGood"+GoTemplate, nil)
	fileGoodMD := chk.CreateTmpFileAs("", ".fileGood"+MdTemplate, nil)

	chk.NoErr(expand([]string{fileGoodGo}))

	chk.StrSlice(
		GoFiles(),
		[]string{
			file1,
			file1_1,
			file2_1,
			fileGoodGo,
		},
	)

	chk.NoErr(expand([]string{fileGoodMD}))

	chk.StrSlice(
		MdFiles(),
		[]string{
			file2,
			file1_2,
			file2_2,
			fileGoodMD,
		},
	)

	chk.Stdout(
		"File to process: '"+file1+"'",
		"File to process: '"+file2+"'",
		"File to process: '"+file1_1+"'",
		"File to process: '"+file1_2+"'",
		"File to process: '"+file2_1+"'",
		"File to process: '"+file2_2+"'",
		"Excluding redundant template: '"+file1+"'",
		"Excluding redundant template: '"+file2+"'",
		"Excluding redundant template: '"+file1_1+"'",
		"Excluding redundant template: '"+file1_2+"'",
		"Excluding redundant template: '"+file2_1+"'",
		"Excluding redundant template: '"+file2_2+"'",
		"File to process: '"+fileGoodGo+"'",
		"File to process: '"+fileGoodMD+"'",
	)
}
