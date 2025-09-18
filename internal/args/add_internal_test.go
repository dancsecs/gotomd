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

func TestFiles_Valid_TooShort(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	chk.Err(
		isValidTemplate(""),
		chk.ErrChain(
			ErrInvalidArgument,
			"''",
			"expected - ("+GoTemplate+" or "+MdTemplate+")",
		),
	)
	chk.Err(
		isValidTemplate(GoTemplate),
		chk.ErrChain(
			ErrInvalidArgument,
			"'"+GoTemplate+"'",
			"expected - ("+GoTemplate+" or "+MdTemplate+")",
		),
	)
	chk.Err(
		isValidTemplate(MdTemplate),
		chk.ErrChain(
			ErrInvalidArgument,
			"'"+MdTemplate+"'",
			"expected - ("+GoTemplate+" or "+MdTemplate+")",
		),
	)

	chk.Err(
		isValidTemplate("."+GoTemplate),
		chk.ErrChain(
			ErrInvalidArgument,
			"'."+GoTemplate+"'",
			"expected - ("+GoTemplate+" or "+MdTemplate+")",
		),
	)
	chk.Err(
		isValidTemplate("."+MdTemplate),
		chk.ErrChain(
			ErrInvalidArgument,
			"'."+MdTemplate+"'",
			"expected - ("+GoTemplate+" or "+MdTemplate+")",
		),
	)
}

func TestFiles_Valid_Invalid(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	chk.Err(
		isValidTemplate("a."+GoTemplate),
		chk.ErrChain(
			ErrInvalidArgument,
			"'a."+GoTemplate+"'",
			"expected - ("+GoTemplate+" or "+MdTemplate+")",
		),
	)
	chk.Err(
		isValidTemplate("a."+MdTemplate),
		chk.ErrChain(
			ErrInvalidArgument,
			"'a."+MdTemplate+"'",
			"expected - ("+GoTemplate+" or "+MdTemplate+")",
		),
	)

	chk.Err(
		isValidTemplate(filepath.Join("dir", "a."+GoTemplate)),
		chk.ErrChain(
			ErrInvalidArgument,
			"'"+filepath.Join("dir", "a."+GoTemplate)+"'",
			"expected - ("+GoTemplate+" or "+MdTemplate+")",
		),
	)
	chk.Err(
		isValidTemplate(filepath.Join("dir", "a."+MdTemplate)),
		chk.ErrChain(
			ErrInvalidArgument,
			"'"+filepath.Join("dir", "a."+MdTemplate)+"'",
			"expected - ("+GoTemplate+" or "+MdTemplate+")",
		),
	)
}

func TestFiles_Valid_Good(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	chk.NoErr(isValidTemplate(".a" + GoTemplate))
	chk.NoErr(isValidTemplate(".a" + MdTemplate))

	chk.NoErr(isValidTemplate(filepath.Join("dir", ".a"+GoTemplate)))
	chk.NoErr(isValidTemplate(filepath.Join("dir", ".a"+MdTemplate)))
}

func TestFiles_AddNewFile(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	Reset()

	appendFile(".abc" + GoTemplate)
	appendFile(".def" + MdTemplate)

	appendFile(".ghi" + GoTemplate)
	appendFile(".jkl" + MdTemplate)

	appendFile(".abc" + GoTemplate)
	appendFile(".def" + MdTemplate)

	chk.StrSlice(
		goFiles,
		[]string{
			".abc" + GoTemplate,
			".ghi" + GoTemplate,
		},
	)

	chk.StrSlice(
		mdFiles,
		[]string{
			".def" + MdTemplate,
			".jkl" + MdTemplate,
		},
	)

	chk.Stdout(
		"File to process: '"+".abc"+GoTemplate+"'",
		"File to process: '"+".def"+MdTemplate+"'",
		"File to process: '"+".ghi"+GoTemplate+"'",
		"File to process: '"+".jkl"+MdTemplate+"'",
		"Excluding redundant template: '"+".abc"+GoTemplate+"'",
		"Excluding redundant template: '"+".def"+MdTemplate+"'",
	)
}

//nolint:funlen // Ok.
func TestFiles_AddDir(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	Reset()

	dir := chk.CreateTmpDir()
	file1 := chk.CreateTmpFileAs("", ".1"+GoTemplate, nil)
	file2 := chk.CreateTmpFileAs("", ".2"+MdTemplate, nil)

	sub1 := chk.CreateTmpSubDir("sub1")
	file1_1 := chk.CreateTmpFileAs(sub1, ".1_1"+GoTemplate, nil)
	file1_2 := chk.CreateTmpFileAs(sub1, ".1_2"+MdTemplate, nil)

	sub2 := chk.CreateTmpSubDir("sub2")
	file2_1 := chk.CreateTmpFileAs(sub2, ".1_1"+GoTemplate, nil)
	file2_2 := chk.CreateTmpFileAs(sub2, ".1_2"+MdTemplate, nil)

	chk.Err(
		add("DOES_NOT_EXIST", false),
		chk.ErrChain(
			ErrUnknownObject,
			"stat DOES_NOT_EXIST",
			"no such file or directory",
		),
	)

	chk.NoErr(add(dir, true))

	chk.StrSlice(
		goFiles,
		[]string{
			file1,
			file1_1,
			file2_1,
		},
	)

	chk.StrSlice(
		mdFiles,
		[]string{
			file2,
			file1_2,
			file2_2,
		},
	)

	// Will be ignored.
	_ = chk.CreateTmpFileAs(sub2, "2."+MdTemplate, nil)

	chk.NoErr(add(dir, true))

	chk.StrSlice(
		goFiles,
		[]string{
			file1,
			file1_1,
			file2_1,
		},
	)

	chk.StrSlice(
		mdFiles,
		[]string{
			file2,
			file1_2,
			file2_2,
		},
	)

	fileBad := chk.CreateTmpFileAs("", "."+MdTemplate, nil)

	chk.Err(
		add(fileBad, false),
		chk.ErrChain(
			ErrUnknownObject,
			ErrInvalidArgument,
			"'"+fileBad+"'",
			"expected - ("+GoTemplate+" or "+MdTemplate+")",
		),
	)

	fileGoodGo := chk.CreateTmpFileAs("", ".fileGood"+GoTemplate, nil)
	fileGoodMD := chk.CreateTmpFileAs("", ".fileGood"+MdTemplate, nil)

	chk.NoErr(add(fileGoodGo, false))
	chk.NoErr(add(fileGoodMD, false))

	chk.StrSlice(
		goFiles,
		[]string{
			file1,
			file1_1,
			file2_1,
			fileGoodGo,
		},
	)

	chk.StrSlice(
		mdFiles,
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
