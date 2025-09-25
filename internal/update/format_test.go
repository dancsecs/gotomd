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

package update_test

import (
	"testing"

	"github.com/dancsecs/gotomd/internal/update"
	"github.com/dancsecs/sztestlog"
)

func TestUpdate_Format(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	update.FormatForMarkdown()

	chk.Str(
		update.MarkGoCode("ABC\n"),
		"```go\nABC\n```",
	)

	chk.Str(
		update.MarkBashCode("ABC\n"),
		"```bash\nABC\n```",
	)

	update.FormatForGoDoc()

	chk.Str(
		update.MarkGoCode("ABC\n"),
		"\tABC",
	)

	chk.Str(
		update.MarkBashCode("ABC\n"),
		"\tABC",
	)
}
