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

/*
Package internal contains the main program entry point (a classic unix Main
returning a error status code. and all sub packages used by the program.
*/
package internal

import (
	"os"

	"github.com/dancsecs/gotomd/internal/args"
	"github.com/dancsecs/gotomd/internal/expand"
	"github.com/dancsecs/gotomd/internal/gopkg"
	"github.com/dancsecs/gotomd/internal/update"
	"github.com/dancsecs/szlog"
)

func processGoFiles(filesToProcess []string) error {
	var err error

	for i, mi := len(filesToProcess)-1, 0; i >= mi && err == nil; i-- {
		gopkg.Reset()

		err = expand.Process(filesToProcess[i])
	}

	return err //nolint:wrapcheck // Ok.
}

func processMDFiles(filesToProcess []string) error {
	var err error

	for i, mi := len(filesToProcess)-1, 0; i >= mi && err == nil; i-- {
		gopkg.Reset()

		err = expand.Process(filesToProcess[i])
	}

	return err //nolint:wrapcheck // Ok.
}

// Main ids the classic unix entry point into the program.
func Main() int {
	const (
		returnGood        = 0
		returnFailed      = 1
		returnNotUpToDate = 2
	)

	var (
		origWd string
		err    error
	)

	// Restore original working directory on exit.
	origWd, err = os.Getwd()
	if err == nil {
		defer func() {
			_ = os.Chdir(origWd)
		}()
	}

	err = args.Process()

	if args.ShowLicense() {
		szlog.Say0(LicenseCopyright, "\n")
	}

	if args.ShowHelp() {
		szlog.Say0(args.Usage(), "\n")
	}

	if args.ShowDirective() {
		szlog.Say0(DirectiveHowTo, "\n")
	}

	update.ResetUpToDate()

	if err == nil {
		err = processGoFiles(args.GoFiles())
	}

	if err == nil {
		err = processMDFiles(args.MdFiles())
	}

	if err == nil {
		if args.CheckUpToDate() {
			if !update.IsUpToDate() {
				szlog.Say1("Documentation is NOT up to date.\n")

				return returnNotUpToDate
			}

			szlog.Say1("Documentation is up to date.\n")
		}

		return returnGood
	}

	szlog.Say0("Failed: ", err, "\n")

	return returnFailed
}
