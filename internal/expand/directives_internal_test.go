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
	"strings"
	"testing"

	"github.com/dancsecs/gotomd/internal/errs"
	"github.com/dancsecs/sztestlog"
)

func TestInternalExpand_Directive_ActionSearch(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	chk.Int(action.search("a"), -1)
	chk.Int(action.search("z"), -1)
}

func TestInternalExpand_Directive_IncludeSnippet_NOT_EXIST(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	res, err := includeSnip("DOES_NOT_EXIST")
	chk.Err(
		err,
		chk.ErrChain(
			errs.ErrInvalidRelativeDir,
			"\"DOES_NOT_EXIST\"",
		),
	)
	chk.Str(res, "")

	res, err = includeSnip("./DOES_NOT_EXIST  <!--- START SNIPPET -->")
	chk.Err(
		err,
		chk.ErrChain(
			errs.ErrParseError,
			"open DOES_NOT_EXIST",
			"no such file or directory",
		),
	)
	chk.Str(res, "")
}

func TestInternalExpand_Directive_IncludeSnippet(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	res, err := includeSnip(
		"./testdata/tstpkg/.sharedTemplate.sds.md # START SNIPPET",
	)
	chk.StrSlice(
		strings.Split(res, "\n"),
		[]string{
			"# Common Snippet Inclusion",
			"",
			"```bash",
			"#!/bin/bash",
			"echo \"Hello, world.\"",
			"```",
		},
	)
	chk.NoErr(err)
}

func TestInternalExpand_Directive_IncludeSnippetAsString(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	res, err := includeSnip(
		"./testdata/tstpkg/.sharedTemplate.sds.md string # START SNIPPET",
	)
	chk.StrSlice(
		strings.Split(res, "\n"),
		[]string{
			"\t\"# Common Snippet Inclusion\"" + ` + "\n" +`,
			"\t\"\"" + ` + "\n" +`,
			"\t\"```bash\"" + ` + "\n" +`,
			"\t\"#!/bin/bash\"" + ` + "\n" +`,
			"\t\"echo \\\"Hello, world.\\\"\"" + ` + "\n" +`,
			"\t\"```\"" + ` + "\n" +`,
		},
	)
	chk.NoErr(err)
}

func TestInternalExpand_Directive_IncludeTxtSnippetAsString(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	res, err := includeSnip(
		"./testdata/tstpkg/.sharedTemplate.sds.txt string",
	)
	chk.StrSlice(
		strings.Split(res, "\n"),
		[]string{
			"\t\"Just plain text.\"" + ` + "\n" +`,
		},
	)
	chk.NoErr(err)
}
