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

package format_test

import (
	"strings"
	"testing"

	"github.com/dancsecs/gotomd/internal/format"
	"github.com/dancsecs/sztestlog"
)

func TestFormat_Inline(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	format.ForMarkdown()

	chk.Str(format.Inline("go", "\n"), "")
	chk.Str(format.Inline("bash", "\n"), "")

	chk.Str(
		format.Inline("go", "ABC\n"),
		"```go\nABC\n```",
	)

	chk.Str(
		format.Inline("bash", "ABC\n"),
		"```bash\nABC\n```",
	)

	format.ForGoDoc()

	chk.Str(format.Inline("go", "\n"), "")
	chk.Str(format.Inline("bash", "\n"), "")

	chk.Str(
		format.Inline("go", "ABC\n"),
		"\tABC",
	)

	chk.Str(
		format.Inline("bash", "ABC\n\nDEF"),
		"\tABC\n\n\tDEF",
	)
}

func TestFormat_Comment(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	format.ForMarkdown()

	chk.Str(
		format.Comment("|---|"),
		"<!--- |---| -->\n",
	)

	format.ForGoDoc()

	chk.Str(
		format.Comment("|---|"),
		"// |---|.\n",
	)
}

func TestFormat_BalancedComment(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	format.ForMarkdown()

	chk.Str(
		format.BalancedComment("|---|"),
		"<!---                                 |---| -->\n",
	)

	format.ForGoDoc()

	chk.Str(
		format.BalancedComment("|---|"),
		"//                                    |---|.\n",
	)
}

func TestFormat_HLine(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	format.ForMarkdown()
	chk.True(format.IsForMarkdown())

	chk.Str(
		format.HLine(),
		"---",
	)

	format.ForGoDoc()

	chk.Str(
		format.HLine(),
		strings.Repeat("-", 78),
	)
}
