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
	"errors"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

const tabSPaces = "    "
const hardSpace = "\\unicode{160}"
const hardUnderscore = "&#x332;"
const hardPercent = "&#xFE6A;"

// "--- PASS: Test_PASS_SampleGoProjectOne (0.0s)".
// "--- FAIL: Test_FAIL_SampleGoProjectOne (0.0s)".
var squashTestTime = regexp.MustCompile(
	`(?m)^(--- (?:PASS|FAIL): .+?) \(\d+\.\d+s\)$`,
)

// Squash runtimes for all tests.
// "FAIL\tgithub.com/dancsecs/sztestToMarkdown/sample_go_project_one\t0.0s".
var squashAllTestTime = regexp.MustCompile(`(?m)^FAIL\s(.+?)\s\d+\.\d+s$`)

// Squash (cached) indicator on tests.
// "ok  	example	(cached) coverage: 100.0% of statements".
// "ok  	example	0.269s	coverage: 100.0% of statements".

var squashCached = regexp.MustCompile(
	`(?m)^(ok\s+.+?\s+)(?:\(cached\)|\d+\.\d+s)\s+(.*)$`,
)

//nolint:funlen // Ok.
func runTest(dir, tests string) (string, string, error) {
	var rawRes []byte
	var args []string
	res := ""

	stat, err := os.Stat(dir)
	if err == nil && !stat.IsDir() {
		err = errors.New("not a directory")
	}
	if err == nil {
		args = []string{"test", "-v", "-cover"}

		if tests != "package" {
			args = append(args, "-run", tests)
		}
		args = append(args, dir)
		c := exec.Command("go", args...) //nolint:gosec // Ok.
		//	c.Dir = dir
		rawRes, _ = c.CombinedOutput() // We expect a general task error.
		if bytes.HasPrefix(rawRes, []byte("testing: warning: no tests to run")) {
			err = errors.New("no tests to run")
		}
	}
	if err == nil {
		res, err = marksToMarkdownHTML(string(rawRes))
	}
	if err == nil && szColorize {
		res = squashTestTime.ReplaceAllString(res, `${1} (0.0s)`)
		res = squashAllTestTime.ReplaceAllString(res, `FAIL ${1} 0.0s`)
		res = squashCached.ReplaceAllString(res, `${1}${2}`)
		res = strings.ReplaceAll(res, "\t", tabSPaces)
		res = strings.ReplaceAll(res, "%", hardPercent)
		res = strings.ReplaceAll(res, " ", hardSpace)
		res = strings.ReplaceAll(res, "_", hardUnderscore)

		latexRes := ""
		lines := strings.Split(res, "\n")
		for _, line := range lines[:len(lines)-1] {
			if latexRes != "" {
				latexRes += "\n"
			}
			latexRes += "$\\small{\\texttt{" + line + "}}$\n<br>"
		}
		res = latexRes
		//  res = "<---\n" + string(rawRes) + "\n -->\n\n" + latexRes
	}
	if err == nil && !szColorize {
		res = "<pre>\n" + strings.TrimRight(res, "\n") + "\n</pre>"
		res = squashTestTime.ReplaceAllString(res, `${1} (0.0s)`)
		res = squashAllTestTime.ReplaceAllString(res, `FAIL ${1} 0.0s`)
		res = squashCached.ReplaceAllString(res, `${1}${2}`)
		res = strings.ReplaceAll(res, "\t", tabSPaces)
		res = strings.ReplaceAll(res, "%", hardPercent)
	}

	if err == nil {
		return "go " + strings.Join(args, " "), strings.TrimRight(res, "\n"), nil
	}
	return "", "", err
}

func buildTestCmds(dir, action []string) ([]string, []string) {
	if len(dir) < 1 {
		return nil, nil
	}

	var newDir []string
	var newAction []string
	var cDir string
	var cAction string
	var i, mi int

	cDir = dir[0]
	cAction = action[0]

	for i, mi = 1, len(dir); i < mi; i++ {
		if dir[i] == cDir {
			cAction += " " + action[i]
		} else {
			newDir = append(newDir, cDir)
			newAction = append(newAction, cAction)
			cDir = dir[i]
			cAction = action[i]
		}
	}
	if cDir != "" {
		newDir = append(newDir, cDir)
		newAction = append(newAction, cAction)
	}
	return newDir, newAction
}

func getGoTst(cmd string) (string, error) {
	var res string
	var tstRes string
	var tstCmd string

	dir, action, err := parseCmds(cmd)
	if err == nil {
		dir, action = buildTestCmds(dir, action)
	}

	for i, mi := 0, len(dir); i < mi && err == nil; i++ {
		tstCmd, tstRes, err = runTest(dir[i], action[i])
		if err == nil {
			if res != "" {
				res += "\n\n"
			}
			res += markBashCode(tstCmd) + "\n\n" + tstRes
		}
	}

	if err == nil {
		return res, nil
	}
	return "", err
}
