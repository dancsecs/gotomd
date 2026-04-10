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
	"os"
	"sort"
	"strings"

	"github.com/dancsecs/gotomd/internal/errs"
	"github.com/dancsecs/gotomd/internal/file"
	"github.com/dancsecs/gotomd/internal/format"
	"github.com/dancsecs/gotomd/internal/godoc"
	"github.com/dancsecs/gotomd/internal/gorun"
	"github.com/dancsecs/gotomd/internal/gotest"
)

const (
	cmdSep     = "::"
	szCmdLabel = "gotomd"
)

const (
	sztestPrefix = "<!--- " + szCmdLabel + cmdSep
	szDocPrefix  = "doc::"
)

const (
	szAutoHeader1 = "*****  AUTO GENERATED:  DO NOT MODIFY  *****"
	szAutoHeader2 = "MODIFY TEMPLATE: "
	szAutoHeader3 = "See: 'https://github.com/dancsecs/gotomd'"
)

type commandAction struct {
	cmdPrefix []string
	cmdAction []func(string) (string, error)
}

// func buildCommand(cmd string, args ...string) string {
// 	rootCmd := szCmdLabel + cmdSep + cmd

// 	if len(args) > 0 {
// 		rootCmd += cmdSep + strings.Join(args, cmdSep)
// 	}

// 	return format.Comment(rootCmd)
// }

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
	action.add("doc::", godoc.GetDoc)
	action.add("docConstGrp::", godoc.GetDocDeclConstantBlock)
	action.add("dcl::", godoc.GetDocDecl)
	action.add("dcln::", godoc.GetDocDeclNatural)
	action.add("dcls::", godoc.GetDocDeclSingle)
	action.add("file::", file.GetGoFile)
	action.add("run::", gorun.GetGoRun)
	action.add("inline-run::", gorun.RawGoRun)
	action.add("tst::", gotest.GetGoTst)
	action.sort()
}

func isCmd(line string) (int, int, error) {
	const sep = "::"

	var (
		cmdIdx = -1
		end    = 0
	)

	if strings.HasPrefix(line, sztestPrefix) {
		cmd := line[len(sztestPrefix):]
		end = strings.Index(cmd, sep)

		if end >= 0 {
			cmd = cmd[:end+2]
			cmdIdx = action.search(cmd)
		}

		if end < 0 || cmdIdx == -1 {
			return 0, 0, fmt.Errorf("%w: %q", errs.ErrUnknownCommand, line)
		}
	}

	return cmdIdx, len(sztestPrefix) + end + len(sep), nil
}

func getBlock(
	str *strings.Builder,
	i, cmdStart int,
	lines []string,
	terminator, cutSet string,
) (int, error) {
	var (
		addSpace  bool
		cleanLine string
		err       error
	)

	line := lines[i]

	for err == nil && !strings.HasSuffix(line, terminator) {
		if addSpace {
			str.WriteRune(' ')
		} else {
			addSpace = true
		}

		str.WriteString(line[cmdStart:])
		cmdStart = 0
		i++

		if i < len(lines) {
			line = strings.Trim(lines[i], " ")
		} else {
			err = errs.ErrBlockNotTerminated
		}
	}

	cleanLine = strings.TrimRight(line[cmdStart:], cutSet)
	if err == nil {
		if addSpace && len(cleanLine) > 0 {
			_, err = str.WriteRune(' ')
		}
	}

	if err == nil {
		str.WriteString(cleanLine)
	}

	return i, err
}

//nolint:cyclop,funlen // Ok.
func parse(dir, fPath, fData string) (string, error) {
	const (
		skipDirBlank   = ""
		skipDirThis    = "."
		skipDirThisDir = skipDirThis + string(os.PathSeparator)
	)

	var (
		res              string
		cmd              strings.Builder
		err              error
		cmdIdx, cmdStart int
		updatedFile      strings.Builder
	)

	lines := strings.Split(fData+"\n", "\n")

	if !(dir == skipDirBlank || dir == skipDirThis || dir == skipDirThisDir) {
		// Need to change he current working directory and change it back
		// at the end of the parse.
		var cwd string
		cwd, err = os.Getwd()

		if err == nil {
			defer func() {
				_ = os.Chdir(cwd)
			}()

			err = os.Chdir(dir)
		}
	}

	if err == nil {
		_, err = updatedFile.WriteString("" +
			format.BalancedComment(szAutoHeader1) +
			format.BalancedComment(szAutoHeader2+"'"+fPath+"'") +
			format.BalancedComment(szAutoHeader3) +
			"\n",
		)
	}

	for i, mi := 0, len(lines)-1; i < mi && err == nil; i++ {
		line := strings.TrimRight(lines[i], " ")
		cmdIdx, cmdStart, err = isCmd(line)

		if err == nil && cmdIdx >= 0 {
			i, err = getBlock(&cmd, i, cmdStart, lines, "-->", " ->")

			if err == nil {
				res, err = action.run(cmdIdx, cmd.String())
			}

			if err == nil {
				_, err = updatedFile.WriteString(res + "\n")
			}

			continue
		}

		if err == nil {
			// Not a command. Just append the line.
			_, err = updatedFile.WriteString(line + "\n")

			continue
		}
	}

	if err == nil {
		return strings.TrimRight(updatedFile.String(), "\n"), nil
	}

	return "", fmt.Errorf("%w: %w", errs.ErrParseError, err)
}
