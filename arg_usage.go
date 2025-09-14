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

	"github.com/dancsecs/szargs"
	"github.com/dancsecs/szlog"
)

const defaultPermissions = 0o0644

//nolint:goCheckNoGlobals // Ok.
var (
	forceOverwrite = false
	szColorize     = false
	outputDir      = "."
	defaultPerm    = defaultPermissions
	showLicense    = false
	showHelp       = false
)

//nolint:cyclop,funlen // Ok for now.
func processArgs() ([]string, string, error) {
	var (
		args       *szargs.Args
		found      bool
		setDefault bool
	)

	cleanedArgs, err := szlog.AbsorbArgs(
		easterEgg(os.Args),
		szlog.EnableVerbose,
	)

	args = szargs.New(
		"Synchronize GitHub README.md files with Go source code,\n"+
			"documentation, tests, and command output. gotomd processes\n"+
			"Markdown templates or existing README files, replacing special\n"+
			"directives with content generated directly from your Go\n"+
			"codebase. This ensures your documentation is always accurate\n"+
			"and in sync with the source."+
			"",
		cleanedArgs,
	)

	szlog.ArgUsageInfo(args.RegisterUsage)
	args.PushErr(err)

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

	if outputDir != "." {
		s, err := os.Stat(outputDir)
		if err != nil || !s.IsDir() {
			args.PushErr(
				fmt.Errorf("%w: '%s'", ErrInvalidOutputDir, outputDir),
			)
		}
	}

	if !args.HasNext() {
		setDefault = showLicense || showHelp

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

func easterEgg(args []string) []string {
	cleanedArgs := make([]string, 0, len(args))

	for _, arg := range args {
		if arg == "--Reem" {
			fmt.Print(dedication) //nolint:forbidigo // Ok.

			continue
		}

		cleanedArgs = append(cleanedArgs, arg)
	}

	return cleanedArgs
}
