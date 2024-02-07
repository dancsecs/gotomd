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
	"testing"

	"github.com/dancsecs/sztest"
)

func Test_Markdown_CleanMarkDownDocument(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	md, err := cleanMarkDownDocument(
		sztestBgnPrefix + szDocPrefix + "action1 -->\n" +
			sztestEndPrefix + szTstPrefix + "action2 -->\n",
	)

	chk.Err(
		err,
		"out of sequence: End before begin: <!--- gotomd::End::tst::action2 -->",
	)
	chk.Str(md, "")
}

func Test_Markdown_CleanMarkDownDocumentMissingBlankAfterAuto(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	md, err := cleanMarkDownDocument(
		szAutoPrefix + " -->\nThis is not a blank line.",
	)

	chk.Err(
		err,
		"missing blank line in auto generated output",
	)
	chk.Str(md, "")
}

func Test_Markdown_UpdateMarkDownDocument(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	md, err := updateMarkDownDocument("",
		sztestPrefix+szDocPrefix+"./INVALID_ROOT_DIRECTORY/action1 -->\n",
	)

	chk.Err(
		err,
		"invalid directory specified as: \"./INVALID_ROOT_DIRECTORY\"",
	)
	chk.Str(md, "")
}

func Test_Markdown_UpdateMarkDown_InvalidCommand(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	md, err := updateMarkDownDocument("",
		sztestPrefix+"unknownCommand -->\n",
	)

	chk.Err(
		err,
		"unknown cmd: <!--- gotomd::unknownCommand -->",
	)
	chk.Str(md, "")
}

func Test_Markdown_Expand(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	d, err := getInfo("./sample_go_project_one", "TimesTwo")
	chk.NoErr(err)

	chk.Str(
		expand(szDocPrefix,
			"TimesTwo",
			d.declGoLang()+"\n\n"+
				d.docMarkdown(),
		),
		"<!--- gotomd::Bgn::doc::TimesTwo -->\n"+
			"```go\nfunc TimesTwo(i int) int\n```\n"+
			"\n"+
			"TimesTwo returns the value times two.\n"+
			"<!--- gotomd::End::doc::TimesTwo -->\n",
	)
}

func Test_Markdown_Search(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	chk.Int(action.search("a"), -1)
	chk.Int(action.search("z"), -1)
}
