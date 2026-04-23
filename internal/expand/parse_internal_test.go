/*
   Golang To Github Markdown Utility: gotomd
   Copyright (C) 2023-2025 Leslie Dancsecs

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

package expand

import (
	"testing"

	"github.com/dancsecs/gotomd/internal/errs"
	"github.com/dancsecs/sztestlog"
)

func TestInternalExpand_Parse_InvalidCommandDirectory(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	fName := chk.CreateTmpFileAs(chk.CreateTmpDir(), "file.md",
		[]byte(""+
			szCmdPrefix+szDocPrefix+
			"./INVALID_ROOT_DIRECTORY/action1 -->\n"+
			"",
		),
	)
	updatedDoc, err := parse(fName, "")

	chk.Err(
		err,
		chk.ErrChain(
			errs.ErrParseError,
			errs.ErrInvalidDirectory,
			"\"./INVALID_ROOT_DIRECTORY\"",
		),
	)
	chk.Str(updatedDoc, "")
}

func TestInternalExpand_Parse_UnknownCommand(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	fName := chk.CreateTmpFile(
		[]byte("" +
			szCmdPrefix + "unknownCommand -->\n",
		),
	)
	updatedDoc, err := parse(fName, "")

	chk.Err(
		err,
		chk.ErrChain(
			errs.ErrParseError,
			errs.ErrUnknownCommand,
			"\"<!--- gotomd::unknownCommand -->\"",
		),
	)
	chk.Str(updatedDoc, "")
}

func TestInternalExpand_Parse_SquashMultipleBlankLines(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	fName := chk.CreateTmpFileAs("", "test.gtm.md",
		[]byte("\n\nFirst\n\n\nSecond\n\n\n\nthird\n\n\n\n\nlast\n\n"),
	)
	updatedDoc, err := parse(fName, "")

	chk.NoErr(err)
	chk.Str(updatedDoc, "\nFirst\n\nSecond\n\nthird\n\nlast")
}

func TestInternalExpand_Parse_SplitDir(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	type testRecord struct {
		rawPath string
		expDir  string
		expName string
		expPath string
	}

	for i, tst := range []testRecord{
		/* Idx:  0 */ {"", ".", "", "."},
		/* Idx:  1 */ {".", ".", ".", "."},
		/* Idx:  2 */ {"./", ".", "", "."},
	} {
		dir, name, path, err := splitDir(tst.rawPath)
		chk.NoErr(err, "Idx: ", i, " err")
		chk.Str(dir, tst.expDir, "Idx: ", i, " dir")
		chk.Str(name, tst.expName, "Idx: ", i, " name")
		chk.Str(path, tst.expPath, "Idx: ", i, " path")
	}

	type testErrRecord struct {
		rawPath string
		expErr  string
	}

	for i, tst := range []testErrRecord{
		/* Idx:  0 */ {"/./", chk.ErrChain(errs.ErrNotLocalDir)},
		/* Idx:  1 */ {"/./abc", chk.ErrChain(errs.ErrNotLocalDir)},
		/* Idx:  1 */ {"/./abc", chk.ErrChain(errs.ErrNotLocalDir)},
	} {
		dir, name, path, err := splitDir(tst.rawPath)
		chk.Err(err, tst.expErr, "Idx: ", i, "err")
		chk.Str(dir, "", "Idx: ", i, " dir")
		chk.Str(name, "", "Idx: ", i, " name")
		chk.Str(path, "", "Idx: ", i, " path")
	}
}
