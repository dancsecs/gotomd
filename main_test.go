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
	"path/filepath"
	"strings"
	"testing"

	"github.com/dancsecs/gotomd/internal/update"
	"github.com/dancsecs/sztestlog"
)

const (
	example1     = "example1"
	example1Path = "." + string(os.PathSeparator) +
		example1 + string(os.PathSeparator)

	example2     = "example2"
	example2Path = "." + string(os.PathSeparator) +
		example2 + string(os.PathSeparator)
)

//nolint:goCheckNoGlobals // Ok for test.
var usage = []string{
	"# Usage: programName",
	"",
	"Synchronize GitHub README.md files with Go source code,",
	"documentation, tests, and command output. gotomd processes",
	"Markdown templates or existing README files, replacing special",
	"directives with content generated directly from your Go",
	"codebase. This ensures your documentation is always accurate",
	"and in sync with the source.",
	"",
	"    programName [-v | --verbose ...] [-l | --license] [-h | --help]",
	"                [-f | --force] [-z | --colorize] [-o | --output <dir>]",
	"                [-p | --permission <perm>] [path ...]",
	"",
	"    [-v | --verbose ...]",
	"        Increase the verbose level for each v provided.",
	"",
	"",
	"    [-l | --license]",
	"        Display license before program exits.",
	"",
	"",
	"    [-h | --help]",
	"        Display program usage information.",
	"",
	"",
	"    [-f | --force]",
	"        Do not confirm overwrite of destination.",
	"",
	"",
	"    [-z | --colorize]",
	"        Colorize go test output.",
	"",
	"",
	"    [-o | --output <dir>]",
	"        Direct all output to the specified directory.",
	"",
	"",
	"    [-p | --permission <perm>]",
	"        Permissions to use when creating new file.",
	"",
	"        (can only set RW bits)",
	"",
	"",
	"    [path ...]",
	"        A specific gotomd file template with the extension '*.gtm.md'" +
		" or a",
	"        directory which will be searched for all matching template",
	"        '*.gtm.md' files.  It defaults to the current directory: '.'",
	"",
}

func Test_Example1ExpandTargetOverwriteDirVerbose(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	dir := chk.CreateTmpDir()
	chk.NoErr(
		setup(
			dir,
			"README.md",
			".README.gtm.md",
			"example1_test.go",
			"example1.go",
		),
	)

	rName := filepath.Join(dir, ".README.gtm.md")
	tName := filepath.Join(dir, "README.md")

	chk.SetArgs(
		"programName",
		"-v",
		"-z",
		rName,
	)

	chk.SetStdinData("Y\n")

	chk.NoErr(os.Truncate(tName, 2))

	// Run command expecting the overwrite to succeed.
	main()

	got, wnt, err := getTestFiles(dir, "README.md")
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Stdout(
		"File to process: '"+rName+"'",
		"Expanding "+rName+" to: "+tName,
		"Loading package info for: .",
		"getInfo(\"package\")",
		"getInfo(\"TimesTwo\")",
		"getInfo(\"TimesThree\")",
		"getInfo(\"InterfaceType\")",
		"getInfo(\"StructureType\")",
		"getInfo(\"TimesTwo\")",
		"getInfo(\"TimesTwo\")",
		"getInfo(\"TimesThree\")",
		"getInfo(\"TimesTwo\")",
		"getInfo(\"TimesTwo\")",
		"getInfo(\"TimesThree\")",
		"getInfo(\"TimesTwo\")",
		"getInfo(\"TimesTwo\")",
		"getInfo(\"TimesThree\")",
		"getInfo(\"ConstantGroup1\")",
		"getInfo(\"ConstantGroup1\")",
		"getInfo(\"ConstantGroupA\")",
		fmt.Sprintf(update.ConfirmMsg, tName),
	)
}

////////////

func setup(dir string, files ...string) error {
	const ext = ".example"

	var (
		err       error
		fileBytes []byte
	)

	files = append(files, "go.mod"+ext, "go.sum"+ext)
	for i, mi := 0, len(files); i < mi && err == nil; i++ {
		fileBytes, err = os.ReadFile(filepath.Join("example1", files[i]))
		if err == nil {
			err = os.WriteFile(
				filepath.Join(dir, strings.TrimSuffix(files[i], ext)),
				fileBytes,
				os.FileMode(defaultPerm), //nolint:gosec // Ok.
			)
		}
	}

	return err
}

//nolint:gosec // Ok.
func getTestFiles(dir, fName string) ([]string, []string, error) {
	gotBytes, err := os.ReadFile(filepath.Join(dir, fName))
	if err != nil {
		return nil, nil, err
	}

	wntBytes, err := os.ReadFile(filepath.Join("example1", fName))
	if err != nil {
		return nil, nil, err
	}

	return strings.Split(string(gotBytes), "\n"),
		strings.Split(string(wntBytes), "\n"),
		nil
}

func Test_JustHelp(t *testing.T) {
	chk := sztestlog.CaptureLogAndStdout(t)
	defer chk.Release()

	chk.SetArgs(
		"programName",
		"-v",
		"-l",
		"-h",
	)

	// Nor Run the main function with no -f arg requiring confirmation
	main()

	chk.Stdout(
		license + strings.Join(usage, "\n"),
	)

	chk.Log()
}
