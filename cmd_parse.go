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
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func parseCmd(cmd string) (string, string, error) {
	if !strings.HasPrefix(cmd, "./") {
		return "", "", fmt.Errorf("%w: %q", ErrInvalidRelativeDir, cmd)
	}

	lastSeparatorPos := strings.LastIndex(cmd, string(os.PathSeparator))
	dir := strings.TrimSpace(cmd[:lastSeparatorPos])
	action := strings.TrimSpace(cmd[lastSeparatorPos+1:])
	s, err := os.Stat(dir)

	if err != nil || !s.IsDir() {
		return "",
			"",
			fmt.Errorf("%w: %q", ErrInvalidDirectory, dir)
	}

	if action == "" {
		return "", "", ErrMissingAction
	}

	return dir, action, nil
}

// ParseCmds parses cmd strings into arrays of directories and actions.
// Directories are validated while the actions are context sensitive.
// The first entry must contain a relative directory component however
// subsequent entries that do not specify a directory will default to
// the last directory defined.
func parseCmds(cmdStr string) ([]string, []string, error) {
	var (
		lastDir       string
		dir, action   string
		err           error
		dirs, actions []string
	)

	cmds := regexp.MustCompile(`[\s\t]+`).Split(cmdStr, -1)

	for i, mi := 0, len(cmds); i < mi && err == nil; i++ {
		cmd := cmds[i]
		if lastDir != "" &&
			strings.LastIndex(cmd, string(os.PathSeparator)) < 0 {
			cmd = "." + string(os.PathSeparator) + filepath.Join(lastDir, cmd)
		}

		dir, action, err = parseCmd(cmd)
		if err == nil {
			dirs = append(dirs, dir)
			actions = append(actions, action)
			lastDir = dir
		}
	}

	if err != nil {
		return nil, nil, err
	}

	return dirs, actions, nil
}
