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
	"github.com/dancsecs/gotomd/internal/format"
	"github.com/dancsecs/sztestlog"
)

func TestInternalExpand_ExpandPreFormatted_Invalid(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	lines := []string{
		"line:0",
		"```bash",
		"# ScriptComment",
		"echo \"Hello\"",
		"line:4",
	}

	format.ForMarkdown()

	i := 1
	block, i, err := expandPreFormatted(i, lines)

	chk.Err(
		err,
		chk.ErrChain(
			errs.ErrBlockNotTerminated,
		),
	)
	chk.Int(i, 5)
	chk.Str(block, "")

	format.ForGoDoc()

	i = 1
	block, i, err = expandPreFormatted(i, lines)

	chk.Err(
		err,
		chk.ErrChain(
			errs.ErrBlockNotTerminated,
		),
	)
	chk.Int(i, 5)
	chk.Str(block, "")
}

func TestInternalExpand_ExpandPreFormatted(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	lines := []string{
		"line:0",
		"```bash",
		"# ScriptComment",
		"echo \"Hello\"",
		"```",
		"line:5",
	}

	format.ForMarkdown()

	i := 1
	block, i, err := expandPreFormatted(i, lines)

	chk.NoErr(err)
	chk.Int(i, 4)
	chk.StrSlice(
		strings.Split(block, "\n"),
		[]string{
			"```bash",
			"# ScriptComment",
			"echo \"Hello\"",
			"```",
		},
	)

	format.ForGoDoc()

	i = 1
	block, i, err = expandPreFormatted(i, lines)

	chk.NoErr(err)
	chk.Int(i, 4)
	chk.StrSlice(
		strings.Split(block, "\n"),
		[]string{
			"\t# ScriptComment",
			"\techo \"Hello\"",
		},
	)
}
