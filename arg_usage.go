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
	"fmt"
	"os"

	"github.com/dancsecs/szargs"
)

const defaultPermissions = 0o0644

//nolint:goCheckNoGlobals // Ok.
var (
	cleanOnly      = false
	forceOverwrite = false
	replace        = false
	verbose        = false
	szColorize     = false
	outputDir      = "."
	defaultPerm    = defaultPermissions
	showLicense    = false
)

//nolint:cyclop,funlen // Ok for now.
func processArgs() []string {
	var (
		args  *szargs.Args
		found bool
	)

	args = szargs.New("Golang to 'github' markdown.", os.Args)

	verbose = args.Count(
		"[-v | --verbose ...]",
		"Provide more information when processing.",
	) > 0

	cleanOnly = args.Is(
		"[-c | --clean]",
		"Reverse operation and remove generated markdown "+
			"(Cannot be used with the [-r | --replace] option).",
	)

	replace = args.Is(
		"[-r | --replace]",
		"Replace the *.MD in place "+
			"(Cannot be used with the [-c | --clean] option).",
	)

	showLicense = args.Is(
		"[-l | --license]",
		"Display license before program exits.",
	)

	forceOverwrite = args.Is(
		"[-f | --force]",
		"Do not confirm overwrite of destination.",
	)

	szColorize = args.Is(
		"[-z | --colorize]",
		"Colorize go test output.",
	)

	outputDir, found = args.ValueString(
		"[-o | --output dir]",
		"Direct all output to the specified directory.",
	)
	if !found {
		outputDir = "."
	}

	defaultPerm, found = args.ValueInt(
		"[-p | --permission]",
		"Permissions to use when creating new file"+
			" (can only set RW bits).",
	)

	if !found {
		defaultPerm = defaultPermissions
	}

	if defaultPerm&(^0o0666) != 0 {
		args.PushErr(ErrInvalidDefPerm)
	}

	if replace && cleanOnly {
		args.PushErr(ErrInvalidOptionsRC)
	}

	if outputDir != "." {
		s, err := os.Stat(outputDir)
		if err != nil || !s.IsDir() {
			args.PushErr(
				fmt.Errorf("%w: '%s'", ErrInvalidOutputDir, outputDir),
			)
		}
	}

	if !args.HasNext() {
		args.PushArg(".") // Default to current directory if no args given.
	}

	var filesToProcess []string
	for args.HasNext() && !args.HasErr() {
		filesToProcess = append(filesToProcess, args.NextString(
			"[path ...]",
			"A specific gotomd file template with the extension '*.gtm.md'\n"+
				"or a directory which will be searched for all matching\n"+
				"template '*.gtm.md' files.   It defaults to the current\n"+
				"directory: '.'",
		))
	}

	if !args.HasErr() {
		args.Done()
	}

	if args.HasErr() {
		panic(args.Err().Error() + "\n\n" + args.Usage())
	}

	return filesToProcess
}
