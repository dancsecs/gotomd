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

	"github.com/dancsecs/gotomd/internal/errs"
)

func getBlock(
	i, cmdStart int,
	lines []string,
	trimLeft bool,
	terminator, cutSet, sep string,
) (int, string, error) {
	var (
		str          strings.Builder
		addSeparator bool
		line         string
		err          error
	)

	nextLine := func() error {
		if i < len(lines) {
			line = strings.TrimRight(lines[i][cmdStart:], " ")
			if trimLeft {
				line = strings.TrimLeft(line, " ")
			}

			cmdStart = 0

			return nil
		}

		return errs.ErrBlockNotTerminated
	}

	err = nextLine()

	for err == nil {
		if strings.HasSuffix(line, terminator) {
			break
		}

		if addSeparator {
			str.WriteString(sep)
		} else {
			addSeparator = true
		}

		str.WriteString(line)

		i++
		err = nextLine()
	}

	if err == nil {
		line = strings.TrimRight(line, cutSet)
		if addSeparator && len(line) > 0 {
			str.WriteString(sep)
		}
	}

	if err == nil {
		str.WriteString(line)
	}

	return i,
		strings.TrimRight(str.String(), "\n"), err
}
