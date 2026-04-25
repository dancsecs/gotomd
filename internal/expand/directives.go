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
	"path/filepath"
	"sort"
	"strings"

	"github.com/dancsecs/gotomd/internal/cmds"
	"github.com/dancsecs/gotomd/internal/file"
	"github.com/dancsecs/gotomd/internal/godoc"
	"github.com/dancsecs/gotomd/internal/gorun"
	"github.com/dancsecs/gotomd/internal/gotest"
)

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
	action.add("doc::", godoc.GetDoc)
	action.add("dcl::", godoc.GetDocDecl)
	action.add("dclg::", godoc.GetDocDeclConstantBlock)
	action.add("dcln::", godoc.GetDocDeclNatural)
	action.add("dcls::", godoc.GetDocDeclSingle)
	action.add("src::", file.GetGoFile)
	action.add("run::", gorun.GetGoRun)
	action.add("irun::", gorun.RawGoRun)
	action.add("tst::", gotest.GetGoTst)
	action.add("snip::", includeSnip)
	action.sort()
}

// IncludeSnip retrieves a gotomd template snippet file expanding all
// directives..
func includeSnip(cmd string) (string, error) {
	var (
		action          string
		dir             string
		cmdArgs         []string
		fName           string
		startAfter      string
		expandedSnippet string
		err             error
	)

	dir, action, err = cmds.ParseCmd(cmd)

	if err == nil {
		const expectedArgCount = 2

		cmdArgs = strings.SplitN(action, " ", expectedArgCount)
		fName = cmdArgs[0]

		if len(cmdArgs) > 1 {
			startAfter = cmdArgs[1]
		}
	}

	if err == nil {
		expandedSnippet, err = parse(
			filepath.Join(dir, fName),
			startAfter,
		)
	}

	if err == nil {
		return expandedSnippet, nil
	}

	return "", err
}
