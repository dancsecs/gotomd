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
		permInt     uint32
		stat        os.FileInfo
		foundEgg    bool
		foundOutput bool
		foundPerm   bool
		err         error
	)

	Reset()

	cleanedArgs, foundEgg = easterEgg(os.Args)
	cleanedArgs, err = szlog.AbsorbArgs(
		cleanedArgs,
		szlog.EnableVerbose,
	)

	args = szargs.New(programFunction, cleanedArgs)

	szlog.ArgUsageInfo(args.RegisterUsage)
	args.PushErr(err)

	showDirective = args.Is(directiveFlag, prepareDesc(directiveDesc))
	showLicense = args.Is(licenseFlag, prepareDesc(licenseDesc))
	showHelp = args.Is(helpFlag, prepareDesc(helpDesc))
	forceOverwrite = args.Is(forceFlag, prepareDesc(forceDesc))
	upToDate = args.Is(upToDateFlag, prepareDesc(upToDateDesc))

	outputDir, foundOutput = args.ValueString(
		outputDirFlag,
		prepareDesc(outputDirDesc),
	)
	permInt, foundPerm = args.ValueUint32(
		permFlag,
		prepareDesc(permDesc),
	)

	args.RegisterUsage(pathArg, prepareDesc(pathDesc))

	if !foundOutput {
		outputDir = "."
	}

	if !foundPerm {
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

	if upToDate && foundOutput {
		args.PushErr(errs.ErrUpToDateWithOutput)
	}

	if upToDate && foundPerm {
		args.PushErr(errs.ErrUpToDateWithPerm)
	}

	if upToDate && forceOverwrite {
		args.PushErr(errs.ErrUpToDateWithForce)
	}

	infoShown := showLicense || showHelp || foundEgg || showDirective
	if !args.HasNext() && !infoShown {
		args.PushArg(".") // Default to current directory if no args given.
	}

	var filesToProcess []string
	for args.HasNext() {
		filesToProcess = append(filesToProcess, args.NextString(
			pathArg,
			"",
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
