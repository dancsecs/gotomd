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
	"bytes"
	"os"
	"path/filepath"
	"strings"

	"github.com/dancsecs/gotomd/internal/args"
	"github.com/dancsecs/gotomd/internal/errs"
	"github.com/dancsecs/gotomd/internal/format"
	"github.com/dancsecs/gotomd/internal/update"
	"github.com/dancsecs/szlog"
)

// Process processes the supplied file data (after switching
// to the supplied directory if necessary) returning the updated data with
// all the gotomd commands expanded.
func Process(rPath string) error {
	var (
		err         error
		rDir, rFile string
		wDir, wFile string
		wPath       string
		fileBytes   []byte
		res         string
	)

	rDir, rFile = filepath.Split(rPath)
	wDir = rDir

	if args.OutputDir() != "." {
		wDir = args.OutputDir()
	}

	var found bool

	wFile, found = strings.CutSuffix(rFile, ".gtm.go")
	if found {
		format.ForGoDoc()

		wFile += ".go"
	} else {
		wFile, found = strings.CutSuffix(rFile, ".gtm.md")
		if found {
			format.ForMarkdown()

			wFile += ".md"
		} else {
			err = errs.ErrUnknownTemplate
		}
	}

	if err == nil {
		wPath = filepath.Join(wDir, strings.TrimPrefix(wFile, "."))

		szlog.Say1f("Expanding %s to: %s\n", rPath, wPath)

		fileBytes, err = os.ReadFile(rPath) //nolint:gosec // Ok.
	}

	if err == nil {
		fileData := string(bytes.TrimRight(fileBytes, "\n"))
		res, err = parse(rDir, rPath, fileData)
	}

	if err == nil {
		res = strings.ReplaceAll(res, "\t", "    ")

		_, err = update.File(
			wPath, args.Force(), args.CheckUpToDate(), res, args.Perm(),
		)
	}

	return err //nolint:wrapcheck // Ok update returns clean errors.
}
