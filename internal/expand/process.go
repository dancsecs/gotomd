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

package expand

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/dancsecs/gotomd/internal/args"
	"github.com/dancsecs/gotomd/internal/errs"
	"github.com/dancsecs/gotomd/internal/format"
	"github.com/dancsecs/gotomd/internal/update"
	"github.com/dancsecs/szlog"
)

func setTarget(fPath string) (string, string, error) {
	var (
		wFile string
		found bool
		err   error
	)

	wFile, found = strings.CutSuffix(fPath, ".gtm.go")
	if found {
		format.ForGoDoc()

		wFile += ".go"
	} else {
		wFile, found = strings.CutSuffix(fPath, ".gtm.md")
		if found {
			format.ForMarkdown()

			wFile += ".md"
		} else {
			err = errs.ErrUnknownTemplate
		}
	}

	if err == nil {
		return args.OutputDir(), wFile, nil
	}

	return "", "", err
}

func isCwd(dir string) bool {
	const (
		skipDirBlank   = ""
		skipDirThis    = "."
		skipDirThisDir = skipDirThis + string(os.PathSeparator)
	)

	return dir == skipDirBlank ||
		dir == skipDirThis ||
		dir == skipDirThisDir
}

// Process processes the supplied file data (after switching
// to the supplied directory if necessary) returning the updated data with
// all the gotomd commands expanded.
func Process(rPath string) error {
	var (
		err         error
		rDir, rFile string
		wDir, wFile string
		wPath       string
		res         string
	)

	rDir, rFile = filepath.Split(rPath)
	wDir, wFile, err = setTarget(rFile)

	if !isCwd(rDir) {
		// Need to change he current working directory and change it back
		// at the end of the parse.
		var cwd string
		cwd, err = os.Getwd()

		if err == nil {
			defer func() {
				_ = os.Chdir(cwd)
			}()

			err = os.Chdir(rDir)
		}
	}

	if err == nil {
		wPath = filepath.Join(wDir, strings.TrimPrefix(wFile, "."))

		szlog.Say1f("Expanding %s to: %s\n", rPath, wPath)

		res, err = parse(rFile)
	}

	if err == nil {
		res = "" +
			format.BalancedComment(szAutoHeader1) +
			format.BalancedComment(szAutoHeader2+"'"+rPath+"'") +
			format.BalancedComment(szAutoHeader3) +
			"\n" +
			res

		if format.IsForMarkdown() {
			res = strings.ReplaceAll(res, "\t", "    ")
		}

		_, err = update.File(
			wPath, args.Force(), args.CheckUpToDate(), res, args.Perm(),
		)
	}

	return err //nolint:wrapcheck // Ok update returns clean errors.
}
