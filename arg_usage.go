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
	buildUsage     = ""
	defaultPerm    = defaultPermissions
	showLicense    = false
	showHelp       = false
	setDefault     = false
)

//nolint:cyclop,funlen // Ok for now.
func processArgs() ([]string, string, error) {
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
		"Reverse operation and remove generated markdown\n"+
			"(Cannot be used with the [-r | --replace] option).",
	)

	replace = args.Is(
		"[-r | --replace]",
		"Replace the *.MD in place\n"+
			"(Cannot be used with the [-c | --clean] option).",
	)

	showLicense = args.Is(
		"[-l | --license]",
		"Display license before program exits.",
	)

	showHelp = args.Is(
		"[-h | --help]",
		"Display program usage information.",
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
		"[-o | --output <dir>]",
		"Direct all output to the specified directory.",
	)
	if !found {
		outputDir = "."
	}

	buildUsage, _ = args.ValueString(
		"[-u | --usage <filename>]",
		"Replace the usage section in the given Go source file using "+
			"content from standard input.  The section is identified as "+
			"the text between the first occurrence of '^\\n/*\\n# Usage$' "+
			"and the following package declaration.  This allows keeping "+
			"command-line usage output (e.g., from --help) synchronized "+
			"with the package documentation.",
	)

	defaultPerm, found = args.ValueInt(
		"[-p | --permission <perm>]",
		"Permissions to use when creating new file.\n"+
			"(can only set RW bits)",
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
		setDefault = showLicense || showHelp || buildUsage != ""

		args.PushArg(".") // Default to current directory if no args given.
	}

	var filesToProcess []string
	for args.HasNext() && !args.HasErr() {
		filesToProcess = append(filesToProcess, args.NextString(
			"[path ...]",
			"A specific gotomd file template with the extension '*.gtm.md' "+
				"or a directory which will be searched for all matching "+
				"template '*.gtm.md' files.  It defaults to the current "+
				"directory: '.'",
		))
	}

	if setDefault {
		filesToProcess = nil
	}

	if !args.HasErr() {
		args.Done()
	}

	if args.HasErr() {
		return nil, args.Usage(), args.Err()
	}

	return filesToProcess, args.Usage(), nil
}
