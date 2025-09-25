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
	"fmt"
	"os"

	"github.com/dancsecs/gotomd/internal/args"
	"github.com/dancsecs/gotomd/internal/gopkg"
	"github.com/dancsecs/gotomd/internal/markdown"
	"github.com/dancsecs/szlog"
)

// License contains the internal copy if the program's license.
const License = `
Golang To Github Markdown: gotomd.
Copyright (C) 2023-2025  Leslie Dancsecs

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
`

func processFiles(filesToProcess []string) error {
	var err error

	for i, mi := 0, len(filesToProcess); i < mi && err == nil; i++ {
		gopkg.Reset()

		err = markdown.ExpandMD(filesToProcess[i])
	}

	return err //nolint:wrapcheck // Ok.
}

// Main ids the classic unix entry point into the program.
func Main() int {
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
		fmt.Print(License) //nolint:forbidigo  // Ok.
	}

	if args.ShowHelp() {
		fmt.Println(args.Usage()) //nolint:forbidigo  // Ok.
	}

	if err == nil {
		err = processFiles(args.MdFiles())
	}

	if err == nil {
		return 0
	}

	szlog.Say0("Failed: ", err, "\n")

	return 1
}
