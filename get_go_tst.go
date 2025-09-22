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
	"regexp"
	"strings"

	"github.com/dancsecs/gotomd/internal/ansi"
	"github.com/dancsecs/gotomd/internal/args"
)

//nolint:goCheckNoGlobals // Ok.
var szEnvSetup = []string{
	"SZTEST_MARK_WNT_ON=" + ansi.Blue,
	"SZTEST_MARK_WNT_OFF=" + ansi.Off,
	"SZTEST_MARK_GOT_ON=" + ansi.Magenta,
	"SZTEST_MARK_GOT_OFF=" + ansi.Off,
	"SZTEST_MARK_MSG_ON=" + ansi.Italic,
	"SZTEST_MARK_MSG_OFF=" + ansi.Off,
	"SZTEST_MARK_INS_ON=" + ansi.Green,
	"SZTEST_MARK_INS_OFF=" + ansi.Off,
	"SZTEST_MARK_DEL_ON=" + ansi.Red,
	"SZTEST_MARK_DEL_OFF=" + ansi.Off,
	"SZTEST_MARK_CHG_ON=" + ansi.Cyan,
	"SZTEST_MARK_CHG_OFF=" + ansi.Off,
	"SZTEST_MARK_SEP_ON=" + ansi.Yellow,
	"SZTEST_MARK_SEP_OFF=" + ansi.Off,
}

// "--- PASS: Test_PASS_Example1 (0.0s)".
// "--- FAIL: Test_FAIL_Example1 (0.0s)".
var squashTestTime = regexp.MustCompile(
	`(?m)^(--- (?:PASS|FAIL): .+?) \(\d+\.\d+s\)$`,
)

// Squash runtimes for all tests.
// "FAIL\tgithub.com/dancsecs/sztestToMarkdown/example1\t0.0s".
var squashAllTestTime = regexp.MustCompile(`(?m)^FAIL\s(.+?)\s\d+\.\d+s$`)

// Squash (cached) indicator on tests.
// "ok  	example	(cached) coverage: 100.0% of statements".
// "ok  	example	0.269s	coverage: 100.0% of statements".

var squashCached = regexp.MustCompile(
	`(?m)^(ok\s+.+?\s+)(?:\(cached\)|\d+\.\d+s)\s+(.*)$`,
)

func setupEnv(env []string) []string {
	szEnv := szEnvSetup
	newEnv := make([]string, 0, len(env)+len(szEnv))

	for _, e := range env {
		add := !strings.HasPrefix(e, "\"SZTEST_") &&
			e != "\"SZTEST_TMP_DIR="
		if add {
			newEnv = append(newEnv, e)
		}
	}

	return append(newEnv, szEnv...)
}

func runTest(dir, tests string) (string, string, error) {
	var (
		rawRes  []byte
		tstArgs []string
		res     string
	)

	stat, err := os.Stat(dir)
	if err == nil && !stat.IsDir() {
		err = ErrInvalidDirectory
	}

	if err == nil {
		tstArgs = []string{"test", "-v", "-cover"}

		if tests != "package" {
			tstArgs = append(tstArgs, "-run", tests)
		}

		tstArgs = append(tstArgs, dir)
		c := exec.Command("go", tstArgs...) //nolint:gosec // Ok.
		//	c.Dir = dir
		c.Env = setupEnv(os.Environ())
		rawRes, _ = c.CombinedOutput() // We expect a general task error.

		if bytes.HasPrefix(
			rawRes,
			[]byte("testing: warning: no tests to run"),
		) {
			err = ErrNoTestToRun
		}
	}

	if err == nil {
		if args.Colorize() {
			res = squashTestTime.ReplaceAllString(
				string(rawRes), `${1} (0.0s)`,
			)
			res = squashAllTestTime.ReplaceAllString(res, `FAIL ${1} 0.0s`)
			res = squashCached.ReplaceAllString(res, `${1}${2}`)
			res = strings.ReplaceAll(res, " PASS: ", " PASS:  ")
			res = strings.ReplaceAll(res, " FAIL: ", " FAIL:  ")
			res = ansi.Colorize(res, true)
		} else {
			res = "<pre>\n" + strings.TrimRight(string(rawRes), "\n") +
				"\n</pre>"
			res = squashTestTime.ReplaceAllString(res, `${1} (0.0s)`)
			res = squashAllTestTime.ReplaceAllString(res, `FAIL ${1} 0.0s`)
			res = squashCached.ReplaceAllString(res, `${1}${2}`)
			res = strings.ReplaceAll(res, "\t", ansi.TabSpaces)
			res = ansi.Colorize(res, false)
		}

		return "go " + strings.Join(tstArgs, " "),
			strings.TrimRight(res, "\n"),
			nil
	}

	return "", "", err
}

func buildTestCmds(dir, action []string) ([]string, []string) {
	if len(dir) < 1 {
		return nil, nil
	}

	var (
		newDir    []string
		newAction []string
		cDir      string
		cAction   string
		i, mi     int
	)

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
	var (
		res    string
		tstRes string
		tstCmd string
	)

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
