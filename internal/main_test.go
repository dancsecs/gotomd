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

package internal_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/dancsecs/gotomd/internal"
	"github.com/dancsecs/gotomd/internal/errs"
	"github.com/dancsecs/sztestlog"
)

const (
	tstpkgName = "tstpkg"
	sep        = string(os.PathSeparator)
	tstpkgPath = "." + sep + "testdata" + sep + tstpkgName
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

//nolint:gosec // Ok.
func getTestFiles(fName, tName string) ([]string, []string, error) {
	gotBytes, err := os.ReadFile(fName)
	if err != nil {
		return nil, nil, err
	}

	wntBytes, err := os.ReadFile(tName)
	if err != nil {
		return nil, nil, err
	}

	return strings.Split(string(gotBytes), "\n"),
		strings.Split(string(wntBytes), "\n"),
		nil
}

func Test_Example1ExpandTargetOverwriteDirVerbose(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	dir := chk.CreateTmpDir()

	fName := tstpkgPath + sep + ".README.gtm.md"
	tName := filepath.Join(dir, "README.md")

	chk.SetArgs(
		"programName",
		"-v",
		"-z",
		"-o", dir,
		fName,
	)

	chk.Int(internal.Main(), 0)

	got, wnt, err := getTestFiles(tstpkgPath+sep+"README.md", tName)
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Stdout(
		"File to process: '"+fName+"'",
		"Expanding "+fName[2:]+" to: "+tName,
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
	)
}

////////////

func Test_JustHelp(t *testing.T) {
	chk := sztestlog.CaptureLogAndStdout(t)
	defer chk.Release()

	chk.SetArgs(
		"programName",
		"-v",
		"-l",
		"-h",
		"DOES_NOT_EXIST",
	)

	chk.Int(internal.Main(), 1)

	chk.Stdout(
		internal.License+strings.Join(usage, "\n"),
		chk.ErrChain(
			"Failed",
			errs.ErrInvalidTemplate,
			errs.ErrUnknownObject,
			"stat DOES_NOT_EXIST",
			"no such file or directory",
		),
	)

	chk.Log()
}
