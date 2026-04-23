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
	"fmt"
	"strings"

	"github.com/dancsecs/gotomd/internal/errs"
)

func isCmd(line string) (int, int, error) {
	const sep = "::"

	var (
		cmdIdx = -1
		end    = 0
	)

	if strings.HasPrefix(line, szCmdPrefix) {
		cmd := line[len(szCmdPrefix):]
		end = strings.Index(cmd, sep)

		if end >= 0 {
			cmd = cmd[:end+2]
			cmdIdx = action.search(cmd)
		}

		if end < 0 || cmdIdx == -1 {
			return 0, 0, fmt.Errorf("%w: %q", errs.ErrUnknownCommand, line)
		}
	}

	return cmdIdx, len(szCmdPrefix) + end + len(sep), nil
}

func expandCmd(
	i,
	cmdIdx, cmdStart int,
	lines []string,
) (string, int, error) {
	var (
		cmd string
		res string
		err error
	)

	i, cmd, err = getBlock(i, cmdStart, lines, "-->", " ->", " ")

	if err == nil {
		res, err = action.run(cmdIdx, cmd)
	}

	if err == nil {
		return res, i, nil
	}

	return "", i, err
}
