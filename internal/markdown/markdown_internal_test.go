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

package markdown

import (
	"os"
	"testing"

	"github.com/dancsecs/gotomd/internal/errs"
	"github.com/dancsecs/gotomd/internal/gopkg"
	"github.com/dancsecs/sztestlog"
)

const (
	sep        = string(os.PathSeparator)
	tstpkg     = "tstpkg"
	tstpkgPath = "." + sep + "testdata" + sep + tstpkg
)

func Test_Markdown_UpdateMarkDownDocument(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	updatedDoc, err := updateMD("",
		sztestPrefix+szDocPrefix+"./INVALID_ROOT_DIRECTORY/action1 -->\n",
	)

	chk.Err(
		err,
		errs.ErrInvalidDirectory.Error()+": \"./INVALID_ROOT_DIRECTORY\"",
	)
	chk.Str(updatedDoc, "")
}

func Test_Markdown_UpdateMarkDown_InvalidCommand(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	updatedDoc, err := updateMD("",
		sztestPrefix+"unknownCommand -->\n",
	)

	chk.Err(
		err,
		errs.ErrUnknownCommand.Error()+
			": \"<!--- gotomd::unknownCommand -->\"",
	)
	chk.Str(updatedDoc, "")
}

func Test_Markdown_Expand(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	docInfo, err := gopkg.Info(tstpkgPath, "TimesTwo")
	chk.NoErr(err)

	chk.Str(
		expand(szDocPrefix,
			"TimesTwo",
			markGoCode(docInfo.Declaration())+"\n\n"+
				docInfo.Comment(),
		),
		"<!--- gotomd::Bgn::doc::TimesTwo -->\n"+
			"```go\nfunc TimesTwo(i int) int\n```\n"+
			"\n"+
			"TimesTwo returns the value times two.\n"+
			"<!--- gotomd::End::doc::TimesTwo -->\n",
	)

	chk.Stdout(
		"Loading package info for: "+tstpkgPath,
		`getInfo("TimesTwo")`,
	)
}

func Test_Markdown_Search(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	chk.Int(action.search("a"), -1)
	chk.Int(action.search("z"), -1)
}
