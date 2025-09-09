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

	"github.com/dancsecs/sztest"
)

const (
	tabSpaces = "    "
	// Using a no break space \u00A0 with a "Combining Grapheme Joiner" \u034F
	// which has no visible display but prevents GitHub's LaTeX from merging
	// consecutive spaces.  Two in a row is the same width as a fixed LaTeX
	// font character.
	hardSpace = "&#xA0;&#x34F;&#xA0;&#x34F;" //
	// Using a no break space \u00A0 with a "Combining Low Line" \u0332 to
	// simulate an underscore which GitHub LaTeX only permits in math mode.
	// Two in a row is the same width as a fixed LaTeX font character.
	hardUnderscore = "&#xA0;&#x332;&#xA0;&#x332;"
	// Using a "SMALL PERCENT SIGN" \uFE6A in place of a regular percent sign
	// GitHub markdown processes normal percent signs.
	hardPercent = "&#xFE6A;"
)

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
			e != "\""+sztest.EnvTmpDir+"="
		if add {
			newEnv = append(newEnv, e)
		}
	}

	return append(newEnv, szEnv...)
}

//nolint:funlen // Ok.
func runTest(dir, tests string) (string, string, error) {
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
		args = []string{"test", "-v", "-cover"}

		if tests != pkgLabel {
			args = append(args, "-run", tests)
		}

		args = append(args, dir)
		c := exec.Command("go", args...) //nolint:gosec // Ok.
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
		if szColorize {
			res = translateToTestSymbols(string(rawRes))
			res = squashTestTime.ReplaceAllString(res, `${1} (0.0s)`)
			res = squashAllTestTime.ReplaceAllString(res, `FAIL ${1} 0.0s`)
			res = squashCached.ReplaceAllString(res, `${1}${2}`)
			// Replacing hyphens with an 'FIGURE DASH' u2012 as a regular
			// hyphen in LaTeX is too short (compared to a = used in the
			// corresponding '=== RUN' test bracket.)
			const dashes = "\u2012\u2012\u2012"
			res = strings.ReplaceAll(res, "--- PASS: ", dashes+" PASS:  ")
			res = strings.ReplaceAll(res, "--- FAIL: ", dashes+" FAIL:  ")
			res = strings.ReplaceAll(res, "\t", tabSpaces)
			res = strings.ReplaceAll(res, "%", hardPercent)
			res = strings.ReplaceAll(res, " ", hardSpace)
			res = strings.ReplaceAll(res, "_", hardUnderscore)

			latexRes := ""
			lines := strings.Split(res, "\n")

			for _, line := range lines[:len(lines)-1] {
				if latexRes != "" {
					latexRes += "\n"
				}

				latexRes += "$\\small{\\texttt{" +
					line +
					"}}$\n<br>"
			}

			res = latexRes
		} else {
			res = translateToBlankSymbols(string(rawRes))
			res = "<pre>\n" + strings.TrimRight(res, "\n") + "\n</pre>"
			res = squashTestTime.ReplaceAllString(res, `${1} (0.0s)`)
			res = squashAllTestTime.ReplaceAllString(res, `FAIL ${1} 0.0s`)
			res = squashCached.ReplaceAllString(res, `${1}${2}`)
			res = strings.ReplaceAll(res, "\t", tabSpaces)
			res = strings.ReplaceAll(res, "%", hardPercent)
		}

		return "go " + strings.Join(args, " "),
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
