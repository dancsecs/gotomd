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

//nolint:goCheckNoGlobals,lll // Ok for test.
var usage = []string{
	"# Usage: programName",
	"",
	"",
	"Synchronize Go package and GitHub style README.md documentation by",
	"embedding Go documentation, source code, test and command output directly",
	"from the Go codebase. This ensures that program documentation is kept in",
	"one place—the Go code itself—while keeping the README and package",
	"documentation automatically up to date. It does this by processing template",
	"files containing markdown formatting and replacing embedded directives with",
	"content generated directly from your Go codebase. This ensures your",
	"documentation is always accurate and in sync with the source.",
	"",
	"",
	"    programName [-v | --verbose ...] [-l | --license] [-h | --help]",
	"                [-f | --force] [-z | --colorize] [-u | --uptodate]",
	"                [-o | --output <dir>] [-p | --permission <perm>]",
	"                [path ...]",
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
	"    [-u | --uptodate]",
	"        Returns 0 if no changes would have been made. No writes are",
	"        performed.",
	"",
	"",
	"    [-o | --output <dir>]",
	"        Direct all output to the specified directory.",
	"",
	"",
	"    [-p | --permission <perm>]",
	"        Permissions to use when creating new file.",
	"",
	"        (can only set RW bits).",
	"",
	"",
	"    [path ...]",
	"        Specific template files (named like '.*.gtm.md' or '.*.gtm.go') or",
	"        a directory which will be searched for all matching template",
	"        files. All subdirectories may be searched by using the special",
	"        './...' path. It defaults to search the current directory: '.'",
	"",
}

//nolint:gosec // Ok.
func getTestFiles(gotFName, wntFName string) ([]string, []string, error) {
	gotBytes, err := os.ReadFile(gotFName)
	if err != nil {
		return nil, nil, err
	}

	wntBytes, err := os.ReadFile(wntFName)
	if err != nil {
		return nil, nil, err
	}

	return strings.Split(string(gotBytes), "\n"),
		strings.Split(string(wntBytes), "\n"),
		nil
}

func Test_Example1ExpandMDTargetOverwriteDirVerbose(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	dir := chk.CreateTmpDir()

	var (
		templatePath = filepath.Join(tstpkgPath, ".README.gtm.md")
		gotPath      = filepath.Join(dir, "README.md")
		wntPath      = filepath.Join(tstpkgPath, "README.md")
	)

	chk.SetArgs(
		"programName",
		"-v",
		"-z",
		"-o", dir,
		templatePath,
	)

	chk.Int(internal.Main(), 0)

	got, wnt, err := getTestFiles(wntPath, gotPath)
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Stdout(
		"File to process: '"+templatePath+"'",
		"Expanding "+templatePath+" to: "+gotPath,
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

func Test_Example1ExpandGoTargetOverwriteDirVerbose(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	dir := chk.CreateTmpDir()

	var (
		templatePath = filepath.Join(tstpkgPath, ".doc.gtm.go")
		wntPath      = filepath.Join(tstpkgPath, "doc.go")
		gotPath      = filepath.Join(dir, "doc.go")
	)

	chk.SetArgs(
		"programName",
		"-v",
		"-z",
		"-o", dir,
		templatePath,
	)

	chk.Int(internal.Main(), 0)

	got, wnt, err := getTestFiles(wntPath, gotPath)
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Stdout(
		"File to process: '"+templatePath+"'",
		"Expanding "+templatePath+" to: "+gotPath,
	)
}

func Test_Example1ExpandGoUpToDateVerbose(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	var (
		docName      = "doc.go"
		templatePath = filepath.Join(tstpkgPath, ".doc.gtm.go")
	)

	chk.SetArgs(
		"programName",
		"-v",
		"-z",
		"--uptodate",
		templatePath,
	)

	chk.Int(internal.Main(), 0)

	chk.Stdout(
		"File to process: '"+templatePath+"'",
		"Expanding "+templatePath+" to: "+docName,
		"No change: "+docName,
		"Documentation is up to date.",
	)
}

func Test_Example1ExpandGoUpToDateVerboseWithChange(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	var (
		docName      = "doc_not_there.go"
		templatePath = filepath.Join(tstpkgPath, ".doc_not_there.gtm.go")
	)

	chk.SetArgs(
		"programName",
		"-v",
		"-z",
		"--uptodate",
		templatePath,
	)

	chk.Int(internal.Main(), 2)

	chk.Stdout(
		"File to process: '"+templatePath+"'",
		"Expanding "+templatePath+" to: "+docName,
		"Would have created: "+docName,
		"Documentation is NOT up to date.",
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
