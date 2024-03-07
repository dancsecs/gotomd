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
	"errors"
	"os"
	"sort"
	"strings"
)

const sztestPrefix = "<!--- gotomd::"
const szAutoPrefix = sztestPrefix + "Auto::"
const sztestBgnPrefix = sztestPrefix + "Bgn::"
const sztestEndPrefix = sztestPrefix + "End::"
const szDocPrefix = "doc::"
const szTstPrefix = "tst::"

type commandAction struct {
	cmdPrefix []string
	cmdAction []func(string) (string, error)
}

func (c *commandAction) add(p string, a func(string) (string, error)) {
	c.cmdPrefix = append(c.cmdPrefix, p)
	c.cmdAction = append(c.cmdAction, a)
}

func (c *commandAction) sort() {
	sort.Sort(c)
}

func (c *commandAction) search(cmd string) int {
	cmdIdx := sort.SearchStrings(c.cmdPrefix, cmd)
	if cmdIdx == len(c.cmdPrefix) || c.cmdPrefix[cmdIdx] != cmd {
		cmdIdx = -1
	}

	return cmdIdx
}

func (c *commandAction) run(idx int, cmd string) (string, error) {
	return c.cmdAction[idx](cmd)
}

func (c *commandAction) Len() int {
	return len(c.cmdPrefix)
}

func (c *commandAction) Less(i, j int) bool {
	return c.cmdPrefix[i] < c.cmdPrefix[j]
}

func (c *commandAction) Swap(i, j int) {
	c.cmdPrefix[i], c.cmdPrefix[j] = c.cmdPrefix[j], c.cmdPrefix[i]
	c.cmdAction[i], c.cmdAction[j] = c.cmdAction[j], c.cmdAction[i]
}

//nolint:goCheckNoGlobals // Ok.
var action = new(commandAction)

//nolint:goCheckNoInits // Ok.
func init() {
	action.add("doc::", getDoc)
	action.add("dcl::", getDocDecl)
	action.add("dcln::", getDocDeclNatural)
	action.add("dcls::", getDocDeclSingle)
	action.add("file::", getGoFile)
	action.add("tst::", getGoTst)
	action.sort()
}

//nolint:cyclop // Ok.
func cleanMarkDownDocument(fData string) (string, error) {
	var (
		err         error
		skipBlank   = false
		updatedFile string
		skipTo      string
	)

	lines := strings.Split(fData+"\n", "\n")
	for i, mi := 0, len(lines)-1; i < mi && err == nil; i++ {
		l := strings.TrimRight(lines[i], " ")

		switch {
		case skipBlank:
			if l != "" {
				err = errors.New("missing blank line in auto generated output")
			}

			skipBlank = false
		case skipTo != "":
			switch {
			case strings.HasPrefix(l, skipTo):
				skipTo = ""
			case strings.HasPrefix(l, sztestEndPrefix):
				err = errors.New("out of sequence: End before begin: " + l)
			}
		case strings.HasPrefix(l, sztestBgnPrefix):
			skipTo = sztestEndPrefix + l[len(sztestBgnPrefix):]
			// Add unexpanded line.
			updatedFile += sztestPrefix + l[len(sztestEndPrefix):] + "\n"
		case strings.HasPrefix(l, szAutoPrefix):
			// Do not add auto generated line or next blank line to output.
			skipBlank = true
		default:
			updatedFile += l + "\n"
		}
	}

	if err != nil {
		return "", err
	}

	return strings.TrimRight(updatedFile, "\n"), nil
}

func expand(prefix, cmd, content string) string {
	return "" +
		sztestBgnPrefix + prefix + cmd + " -->\n" +
		content + "\n" +
		sztestEndPrefix + prefix + cmd + " -->\n"
}

func isCmd(l string) (int, int, error) {
	const sep = "::"

	var (
		cmdIdx = -1
		end    = 0
	)

	if strings.HasPrefix(l, sztestPrefix) {
		cmd := l[len(sztestPrefix):]
		end = strings.Index(cmd, sep)

		if end >= 0 {
			cmd = cmd[:end+2]
			cmdIdx = action.search(cmd)
		}

		if end < 0 || cmdIdx == -1 {
			return 0, 0, errors.New("unknown cmd: " + l)
		}
	}

	return cmdIdx, len(sztestPrefix) + end + len(sep), nil
}

//nolint:cyclop // Ok.
func updateMarkDownDocument(dir, fData string) (string, error) {
	const (
		skipDirBlank   = ""
		skipDirThis    = "."
		skipDirThisDir = skipDirThis + string(os.PathSeparator)
	)

	var (
		res              string
		cmd              string
		err              error
		cmdIdx, cmdStart int
	)

	if !(dir == skipDirBlank || dir == skipDirThis || dir == skipDirThisDir) {
		var cwd string
		cwd, err = os.Getwd()

		if err == nil {
			defer func() {
				_ = os.Chdir(cwd)
			}()

			err = os.Chdir(dir)
		}
	}

	updatedFile := szAutoPrefix + " See github.com/dancsecs/gotomd "
	if !replace {
		updatedFile += "**DO NOT MODIFY** "
	}

	updatedFile += "-->\n\n"
	lines := strings.Split(fData+"\n", "\n")

	for i, mi := 0, len(lines)-1; i < mi && err == nil; i++ {
		l := strings.TrimRight(lines[i], " ")
		cmdIdx, cmdStart, err = isCmd(l)

		if err == nil {
			if cmdIdx >= 0 {
				cmd = l[cmdStart : len(l)-len(" -->")]
				res, err = action.run(cmdIdx, cmd)

				if err == nil {
					updatedFile += expand(action.cmdPrefix[cmdIdx], cmd, res)
				}
			} else {
				updatedFile += l + "\n"
			}
		}
	}

	if err != nil {
		return "", err
	}

	return strings.TrimRight(updatedFile, "\n"), nil
}
