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

package gorun

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/dancsecs/gotomd/internal/cmds"
	"github.com/dancsecs/gotomd/internal/errs"
)

func markBashCode(content string) string {
	return "```bash\n" + strings.TrimRight(content, "\n") + "\n```"
}

func joinKeepPrefix(dir, file string) string {
	const relativePrefix = "." + string(os.PathSeparator)

	keep := strings.HasPrefix(dir, relativePrefix)

	joined := filepath.Join(dir, file)

	if keep && !strings.HasPrefix(joined, relativePrefix) {
		joined = relativePrefix + joined
	}

	return joined
}

// RunGo executes the named go package in the provided directory.
func RunGo(dir, cmd string) (string, string, error) {
	var (
		rawRes []byte
		args   []string
		res    string
	)

	stat, err := os.Stat(dir)
	if err == nil && !stat.IsDir() {
		err = errs.ErrInvalidDirectory
	}

	if err == nil {
		cmdArgs := strings.Split(cmd, " ")
		args = append(
			[]string{"run", joinKeepPrefix(dir, cmdArgs[0])},
			cmdArgs[1:]...,
		)
	}

	if err == nil {
		c := exec.Command("go", args...) //nolint:gosec // Ok.
		//	c.Dir = dir
		//  c.Env = setupEnv(os.Environ())
		rawRes, _ = c.CombinedOutput() // We expect a general task error.

		if bytes.HasPrefix(
			rawRes,
			[]byte("package TEST_DOES_NOT_EXIST is not in"),
		) {
			err = errs.ErrNoPackageToRun
		}
	}

	if err == nil {
		res = "<pre>\n" + strings.TrimRight(string(rawRes), "\n") + "\n</pre>"
		//  res = strings.ReplaceAll(res, "\t", tabSpaces)
		//  res = strings.ReplaceAll(res, "%", hardPercent)

		return "go " + strings.Join(args, " "),
			strings.TrimRight(res, "\n"),
			nil
	}

	return "", "", err
}

// GetGoRun runs "go run" the provided package collecting and returning the
// output.
func GetGoRun(cmd string) (string, error) {
	var (
		res    string
		runRes string
		runCmd string
	)

	dir, action, err := cmds.ParseCmd(cmd)
	if err == nil {
		runCmd, runRes, err = RunGo(dir, action)
	}

	if err == nil {
		res += "---\n" + markBashCode(runCmd) + "\n\n" + runRes + "\n---"
	}

	if err == nil {
		return res, nil
	}

	return "", err
}
