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
	szCmdPrefix = "<!--- " + szCmdLabel + cmdSep
	szDocPrefix = "doc::"
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

func getBlock(
	i, cmdStart int,
	lines []string,
	terminator, cutSet, sep string,
) (int, string, error) {
	var (
		str          strings.Builder
		addSeparator bool
		line         string
		err          error
	)

	if i < len(lines) {
		line = strings.Trim(lines[i][cmdStart:], " ")
	} else {
		err = errs.ErrBlockNotTerminated
	}

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
		if i < len(lines) {
			line = strings.Trim(lines[i], " ")
		} else {
			err = errs.ErrBlockNotTerminated

			continue
		}
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

func processCmd(
	file *strings.Builder,
	i,
	cmdIdx, cmdStart int,
	lines []string,
) (int, error) {
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
		file.WriteString(res + "\n")
	}

	return i, err
}

func processCodeBlock(
	file *strings.Builder,
	i int,
	lines []string,
	codeSyntaxName string,
) (int, error) {
	var (
		code string
		err  error
	)

	i, code, err = getBlock(i+1, 0, lines, "```", "`", "\n")
	if err == nil {
		file.WriteString(
			format.Inline(codeSyntaxName, code),
		)
	}

	return i, err
}

func processLines(updatedFile *strings.Builder, lines []string) error {
	var (
		cmdIdx   int
		cmdStart int
		err      error
	)

	for i, mi := 0, len(lines)-1; i < mi; i++ {
		line := strings.TrimRight(lines[i], " ")

		cmdIdx, cmdStart, err = isCmd(line)
		if err == nil && cmdIdx >= 0 {
			i, err = processCmd(updatedFile, i, cmdIdx, cmdStart, lines)
			if err == nil {
				continue
			}
		}

		if err == nil && strings.HasPrefix(line, "```") {
			i, err = processCodeBlock(updatedFile, i, lines, line[3:])

			if err == nil {
				continue
			}
		}

		if err == nil {
			// Not a block. Just append the line.
			updatedFile.WriteString(line + "\n")
		} else {
			break
		}
	}

	return err
}

func parse(dir, fPath, fData string) (string, error) {
	const (
		skipDirBlank   = ""
		skipDirThis    = "."
		skipDirThisDir = skipDirThis + string(os.PathSeparator)
	)

	var (
		err         error
		updatedFile strings.Builder
	)

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
		updatedFile.WriteString("" +
			format.BalancedComment(szAutoHeader1) +
			format.BalancedComment(szAutoHeader2+"'"+fPath+"'") +
			format.BalancedComment(szAutoHeader3) +
			"\n",
		)
	}

	err = processLines(&updatedFile, strings.Split(fData+"\n", "\n"))

	if err == nil {
		return strings.TrimRight(updatedFile.String(), "\n"), nil
	}

	return "", fmt.Errorf("%w: %w", errs.ErrParseError, err)
}
