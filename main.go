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

package main

import (
	"fmt"
	"os"

	"github.com/dancsecs/gotomd/internal/args"
	"github.com/dancsecs/gotomd/internal/gopkg"
)

const license = `
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

		err = expandMD(filesToProcess[i])
	}

	return err
}

func main() {
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
		fmt.Print(license) //nolint:forbidigo  // Ok.
	}

	if args.ShowHelp() {
		fmt.Println(args.Usage()) //nolint:forbidigo  // Ok.
	}

	if err == nil {
		err = processFiles(args.MdFiles())
	}

	if err != nil {
		panic(err)
	}
}
