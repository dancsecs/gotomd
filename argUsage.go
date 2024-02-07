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
	"flag"
	"fmt"
	"os"
	"strings"
)

const defaultCpuProfileIterations = uint(5)
const defaultPermissions = 0644

//nolint:goCheckNoGlobals // Ok.
var (
	cleanOnly            = false
	forceOverwrite       = false
	replace              = false
	verbose              = false
	szColorize           = false
	outputDir            = "."
	defaultPerm          = defaultPermissions
	showLicense          = false
	cpuProfile           = ""
	cpuProfileIterations = defaultCpuProfileIterations
)

func usage() {
	cmdName := os.Args[0]
	fmt.Fprint(flag.CommandLine.Output(),
		"Usage of ", cmdName,
		" [-c | -r]"+
			" [-f]"+
			" [-v]"+
			" [-z]"+
			" [-p perm]"+
			" [-o outDir]"+
			" [-U file]"+
			" [-u int]"+
			" file|dir"+
			" [file|dir...]"+
			"\n",
	)
	flag.PrintDefaults()
}

func captureFlagDefaults() string {
	buf := strings.Builder{}
	origOut := flag.CommandLine.Output()
	defer func() {
		flag.CommandLine.SetOutput(origOut)
	}()
	flag.CommandLine.SetOutput(&buf)

	flag.CommandLine.Usage()
	return buf.String()
}

//nolint:funlen // Ok.
func processArgs() {
	flag.BoolVar(&verbose, "v", false,
		"Provide more information when processing.",
	)
	flag.BoolVar(&cleanOnly, "c", false,
		"Reverse operation and remove generated markdown "+
			"(Cannot be used with the -r option).",
	)
	flag.BoolVar(&replace, "r", false,
		"Replace the *.MD in place (Cannot be used with the -c flag).",
	)
	flag.BoolVar(&showLicense, "l", false,
		"Display license before program exits.",
	)
	flag.BoolVar(&forceOverwrite, "f", false,
		"Do not confirm overwrite of destination.",
	)
	flag.BoolVar(&szColorize, "z", false,
		"Colorize go test output.",
	)
	flag.StringVar(&outputDir, "o", ".",
		"Direct all output to the specified directory.",
	)
	flag.IntVar(&defaultPerm, "p", defaultPermissions,
		"Permissions to use when creating new file (can only set RW bits).",
	)
	flag.StringVar(&cpuProfile, "U", "",
		"Collect cpu profile data into named file.",
	)

	flag.UintVar(&cpuProfileIterations, "u", defaultCpuProfileIterations,
		"Number of iterations to run when collecting cpu profile information.",
	)

	flag.CommandLine.Usage = usage
	_ = flag.CommandLine.Parse(os.Args[1:])

	if flag.CommandLine.NArg() < 1 {
		panic("at least one file or directory must be specified\n" +
			captureFlagDefaults(),
		)
	}

	if defaultPerm&(^0666) != 0 {
		panic("invalid default permissions specified\n" +
			captureFlagDefaults(),
		)
	}

	if replace && cleanOnly {
		panic("only one of -c and -r may be specified\n" +
			captureFlagDefaults(),
		)
	}

	if outputDir != "." {
		s, err := os.Stat(outputDir)
		if err != nil || !s.IsDir() {
			panic("invalid output directory specified: " + outputDir + "\n" +
				captureFlagDefaults(),
			)
		}
	}
}
