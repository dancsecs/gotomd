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

package args

import (
	"fmt"
	"os"

	"github.com/dancsecs/gotomd/internal/errs"
	"github.com/dancsecs/szargs"
	"github.com/dancsecs/szlog"
)

// Process parses the command line arguments
//
//nolint:cyclop,funlen // Ok for now.
func Process() error {
	var (
		args        *szargs.Args
		cleanedArgs []string
		found       bool
		permInt     uint32
		stat        os.FileInfo
		foundEgg    bool
		err         error
	)

	Reset()

	cleanedArgs, foundEgg = easterEgg(os.Args)
	cleanedArgs, err = szlog.AbsorbArgs(
		cleanedArgs,
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

	permInt, found = args.ValueUint32(
		"[-p | --permission <perm>]",
		"Permissions to use when creating new file.\n"+
			"(can only set RW bits)",
	)

	if !found {
		perm = defaultPerm
	} else {
		perm = os.FileMode(permInt)
	}

	if int(permInt)&(^0o0666) != 0 {
		args.PushErr(
			fmt.Errorf("%w: '0o%#o'", errs.ErrInvalidDefPerm, permInt),
		)
	}

	if outputDir != "." {
		stat, err = os.Stat(outputDir)
		if err != nil || !stat.IsDir() {
			args.PushErr(
				fmt.Errorf("%w: '%s'", errs.ErrInvalidOutputDir, outputDir),
			)
		}
	}

	if !args.HasNext() && !(showLicense || showHelp || foundEgg) {
		args.PushArg(".") // Default to current directory if no args given.
	}

	var filesToProcess []string
	for args.HasNext() {
		filesToProcess = append(filesToProcess, args.NextString(
			"[path ...]",
			"A specific gotomd file template with the extension '*.gtm.md' "+
				"or a directory which will be searched for all matching "+
				"template '*.gtm.md' files.  It defaults to the current "+
				"directory: '.'",
		))
	}

	args.Done()

	usage = args.Usage()

	err = args.Err()

	if err == nil {
		err = expand(filesToProcess)
	}

	if err == nil {
		return nil
	}

	return err
}

func easterEgg(args []string) ([]string, bool) {
	found := false
	cleanedArgs := make([]string, 0, len(args))

	for _, arg := range args {
		if arg == "--Reem" {
			found = true

			fmt.Print(dedication) //nolint:forbidigo // Ok.

			continue
		}

		cleanedArgs = append(cleanedArgs, arg)
	}

	return cleanedArgs, found
}
