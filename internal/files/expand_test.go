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

package files_test

import (
	"path/filepath"
	"testing"

	"github.com/dancsecs/gotomd/internal/files"
	"github.com/dancsecs/sztestlog"
)

func TestFiles_Expand_Empty(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	dir := chk.CreateTmpDir()
	file1 := chk.CreateTmpFileAs("", ".1"+files.GoTemplate, nil)
	file2 := chk.CreateTmpFileAs("", ".2"+files.MdTemplate, nil)

	t.Chdir(dir)

	chk.NoErr(files.Expand([]string{""}))
	chk.NoErr(files.Expand([]string{"."}))

	chk.StrSlice(
		files.GoFiles(),
		[]string{
			filepath.Base(file1),
		},
	)

	chk.StrSlice(
		files.MdFiles(),
		[]string{
			filepath.Base(file2),
		},
	)

	chk.Stdout(
		"File to process: '"+".1"+files.GoTemplate+"'",
		"File to process: '"+".2"+files.MdTemplate+"'",
		"File to process: '"+".1"+files.GoTemplate+"'",
		"File to process: '"+".2"+files.MdTemplate+"'",
	)
}

//nolint:funlen // Ok.
func TestFiles_Expand(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	dir := chk.CreateTmpDir()
	file1 := chk.CreateTmpFileAs("", ".1"+files.GoTemplate, nil)
	file2 := chk.CreateTmpFileAs("", ".2"+files.MdTemplate, nil)

	sub1 := chk.CreateTmpSubDir("sub1")
	file1_1 := chk.CreateTmpFileAs(sub1, ".1_1"+files.GoTemplate, nil)
	file1_2 := chk.CreateTmpFileAs(sub1, ".1_2"+files.MdTemplate, nil)
	_ = chk.CreateTmpFileAs(sub1, "ignore.txt", nil)

	sub2 := chk.CreateTmpSubDir("sub2")
	file2_1 := chk.CreateTmpFileAs(sub2, ".1_1"+files.GoTemplate, nil)
	file2_2 := chk.CreateTmpFileAs(sub2, ".1_2"+files.MdTemplate, nil)

	chk.Err(
		files.Expand([]string{"DOES_NOT_EXIST"}),
		chk.ErrChain(
			files.ErrInvalidTemplate,
			files.ErrUnknownObject,
			"stat DOES_NOT_EXIST",
			"no such file or directory",
		),
	)

	chk.NoErr(files.Expand([]string{dir + "/..."}))

	chk.StrSlice(
		files.GoFiles(),
		[]string{
			file1,
			file1_1,
			file2_1,
		},
	)

	chk.StrSlice(
		files.MdFiles(),
		[]string{
			file2,
			file1_2,
			file2_2,
		},
	)

	// Will be ignored.
	_ = chk.CreateTmpFileAs(sub2, "2."+files.MdTemplate, nil)

	chk.NoErr(files.Expand([]string{dir + "/..."}))

	chk.StrSlice(
		files.GoFiles(),
		[]string{
			file1,
			file1_1,
			file2_1,
		},
	)

	chk.StrSlice(
		files.MdFiles(),
		[]string{
			file2,
			file1_2,
			file2_2,
		},
	)

	fileBad := chk.CreateTmpFileAs("", "."+files.MdTemplate, nil)

	chk.Err(
		files.Expand([]string{fileBad}),
		chk.ErrChain(
			files.ErrInvalidTemplate,
			files.ErrUnknownObject,
			files.ErrInvalidArgument,
			"'"+fileBad+"'",
			"expected - ("+files.GoTemplate+" or "+files.MdTemplate+")",
		),
	)

	fileGoodGo := chk.CreateTmpFileAs("", ".fileGood"+files.GoTemplate, nil)
	fileGoodMD := chk.CreateTmpFileAs("", ".fileGood"+files.MdTemplate, nil)

	chk.NoErr(files.Expand([]string{fileGoodGo}))

	chk.StrSlice(
		files.GoFiles(),
		[]string{
			fileGoodGo,
		},
	)

	chk.NoErr(files.Expand([]string{fileGoodMD}))

	chk.StrSlice(
		files.MdFiles(),
		[]string{
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
		"File to process: '"+file1+"'",
		"File to process: '"+file2+"'",
		"File to process: '"+file1_1+"'",
		"File to process: '"+file1_2+"'",
		"File to process: '"+file2_1+"'",
		"File to process: '"+file2_2+"'",
		"File to process: '"+fileGoodGo+"'",
		"File to process: '"+fileGoodMD+"'",
	)
}
