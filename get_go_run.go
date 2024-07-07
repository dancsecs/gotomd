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
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func runGo(dir, cmd string) (string, string, error) {
	var (
		rawRes []byte
		args   []string
		res    string
	)

	stat, err := os.Stat(dir)
	if err == nil && !stat.IsDir() {
		err = ErrInvalidDirectory
	}

	if err == nil {
		cmdArgs := strings.Split(cmd, " ")
		args = append(
			[]string{"run", filepath.Join(dir, cmdArgs[0])},
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
			err = ErrNoPackageToRun
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

func getGoRun(cmd string) (string, error) {
	var (
		res    string
		runRes string
		runCmd string
	)

	dir, action, err := parseCmd(cmd)
	if err == nil {
		runCmd, runRes, err = runGo(dir, action)
	}

	if err == nil {
		res += markBashCode(runCmd) + "\n\n" + runRes
	}

	if err == nil {
		return res, nil
	}

	return "", err
}
