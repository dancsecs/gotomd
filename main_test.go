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
	"path/filepath"
	"strings"
	"testing"

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
	"    programName [-v | --verbose ...] [--quiet] " +
		"[--log <level | (levels)>]",
	"                [--language <lang>] [--long-labels] [-c | --clean]",
	"                [-r | --replace] [-l | --license] [-h | --help]",
	"                [-f | --force] [-z | --colorize] [-o | --output <dir>]",
	"                [-u | --usage <filename>] [-p | --permission <perm>]",
	"                [path ...]",
	"",
	"    [-v | --verbose ...]",
	"        Increase the verbose level for each v provided.",
	"",
	"",
	"    [--quiet]",
	"        Sets the verbose level to -1 squashing all (non-logged) output.",
	"",
	"",
	"    [--log <level | (levels)>]",
	"        Set the level to log (or a custom combination of levels).  Valid",
	"        levels are: None, FATAL, ERROR, WARN, INFO, DEBUG,TRACE, ALL.",
	"",
	"",
	"    [--language <lang>]",
	"        Sets the local language used for formatting.",
	"",
	"",
	"    [--long-labels]",
	"        Use long labels in log output.",
	"",
	"",
	"    [-c | --clean]",
	"        Reverse operation and remove generated markdown",
	"",
	"        (Cannot be used with the [-r | --replace] option).",
	"",
	"",
	"    [-r | --replace]",
	"        Replace the *.MD in place\n",
	"        (Cannot be used with the [-c | --clean] option).",
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
	"    [-u | --usage <filename>]",
	"        Replace the usage section in the given Go source file using",
	"        content from standard input.  The section is identified as the",
	"        text between the first occurrence of '^\\n/*\\n# Usage .*$'" +
		" and the",
	"        following package declaration.  This allows keeping " +
		"command-line",
	"        usage output (e.g., from --help) synchronized with the package",
	"        documentation.",
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
	chk := sztestlog.CaptureLogAndStdout(t)
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
		"-vvvvvv",
		"--log",
		"all",
		"-z",
		dir,
	)

	chk.SetStdinData("Y\n")

	chk.NoErr(os.Truncate(tName, 2))

	// Run command expecting the overwrite to be cancelled.
	main()

	got, wnt, err := getTestFiles(dir, "README.md")
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Stdout(fmt.Sprintf(confirmMsg, tName))

	chk.Log(
		"I:filesToProcess: "+rName,
		"I:Expanding "+rName+" to: "+tName,
		"I:Loading Package info for: .",
		"I:getInfo(\"package\")",
		"I:getInfo(\"TimesTwo\")",
		"I:getInfo(\"TimesThree\")",
		"I:getInfo(\"InterfaceType\")",
		"I:getInfo(\"StructureType\")",
		"I:getInfo(\"TimesTwo\")",
		"I:getInfo(\"TimesTwo\")",
		"I:getInfo(\"TimesThree\")",
		"I:getInfo(\"TimesTwo\")",
		"I:getInfo(\"TimesTwo\")",
		"I:getInfo(\"TimesThree\")",
		"I:getInfo(\"TimesTwo\")",
		"I:getInfo(\"TimesTwo\")",
		"I:getInfo(\"TimesThree\")",
		"I:getInfo(\"ConstantGroup1\")",
		"I:getInfo(\"ConstantGroup1\")",
		"I:getInfo(\"ConstantGroupA\")",
	)
}

////////////

func Test_Example1ReplaceNoTarget(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	dir := chk.CreateTmpDir()
	chk.NoErr(setup(dir, "example1_test.go", "example1.go"))

	fName := filepath.Join(dir, "README.md")

	chk.SetArgs(
		"programName",
		"-r",
		fName,
	)

	chk.Panic(
		main,
		"stat "+filepath.Join(dir, "README.md")+": no such file or directory",
	)

	_, _, err := getTestFiles(dir, "README.md")
	chk.Err(
		err,
		"open "+fName+": no such file or directory",
	)
}

func Test_Example1ReplaceTargetCancel(t *testing.T) {
	chk := sztestlog.CaptureLogAndStdout(t)
	defer chk.Release()

	dir := chk.CreateTmpDir()
	chk.NoErr(
		setup(dir, "README.md", "example1_test.go", "example1.go"),
	)

	fName := filepath.Join(dir, "README.md")
	chk.SetArgs(
		"programName",
		"-vvvvvv",
		"--log",
		"all",
		"-r",
		fName,
	)

	chk.SetStdinData("N\n")

	// Run command expecting the overwrite to be cancelled.
	main()

	chk.Stdout(
		fmt.Sprintf(confirmMsg, fName) +
			confirmCancelled[:len(confirmCancelled)-1],
	)

	chk.Log(
		"I:filesToProcess: "+fName,
		"I:Expanding "+fName+" <inPlace> to: "+fName,
		"I:Loading Package info for: .",
		"I:getInfo(\"package\")",
		"I:getInfo(\"TimesTwo\")",
		"I:getInfo(\"TimesThree\")",
		"I:getInfo(\"InterfaceType\")",
		"I:getInfo(\"StructureType\")",
		"I:getInfo(\"TimesTwo\")",
		"I:getInfo(\"TimesTwo\")",
		"I:getInfo(\"TimesThree\")",
		"I:getInfo(\"TimesTwo\")",
		"I:getInfo(\"TimesTwo\")",
		"I:getInfo(\"TimesThree\")",
		"I:getInfo(\"TimesTwo\")",
		"I:getInfo(\"TimesTwo\")",
		"I:getInfo(\"TimesThree\")",
		"I:getInfo(\"ConstantGroup1\")",
		"I:getInfo(\"ConstantGroup1\")",
		"I:getInfo(\"ConstantGroupA\")",
	)
}

func Test_Example1ReplaceTargetOverwrite(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	dir := chk.CreateTmpDir()
	chk.NoErr(
		setup(dir, "README.md", "example1_test.go", "example1.go"),
	)

	fName := filepath.Join(dir, "README.md")

	chk.SetArgs(
		"programName",
		"-r",
		"-z",
		fName,
	)

	chk.SetStdinData("Y\n")

	// Run command expecting the overwrite to be cancelled.
	main()

	got, wnt, err := getTestFiles(dir, "README.md")
	chk.NoErr(err)

	wnt[0] = strings.ReplaceAll(wnt[0], "**DO NOT MODIFY** ", "")
	chk.StrSlice(got, wnt)

	_, _, err = getTestFiles(dir, ".README.gtm.md")
	chk.Err(
		err,
		"open "+filepath.Join(dir, ".README.gtm.md")+
			": no such file or directory",
	)

	chk.Stdout(
		fmt.Sprintf(confirmMsg, fName),
	)
}

func Test_Example1ReplaceTargetOverwriteDir(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	dir := chk.CreateTmpDir()
	chk.NoErr(
		setup(dir, "README.md", "example1_test.go", "example1.go"),
	)

	chk.SetArgs(
		"programName",
		"-r",
		"-z",
		dir,
	)

	chk.SetStdinData("Y\n")

	// Run command expecting the overwrite to be cancelled.
	main()

	got, wnt, err := getTestFiles(dir, "README.md")
	chk.NoErr(err)

	wnt[0] = strings.ReplaceAll(wnt[0], "**DO NOT MODIFY** ", "")
	chk.StrSlice(got, wnt)

	_, _, err = getTestFiles(dir, ".README.gtm.md")
	chk.Err(
		err,
		"open "+dir+"/.README.gtm.md: no such file or directory",
	)

	chk.Stdout(
		fmt.Sprintf(confirmMsg, filepath.Join(dir, "README.md")),
	)
}

func Test_Example1ReplaceTargetOverwriteDirFromClean(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	dir := chk.CreateTmpDir()
	chk.NoErr(
		setup(dir, ".README.gtm.md", "example1_test.go", "example1.go"),
	)

	chk.NoErr(
		os.Rename(
			filepath.Join(dir, ".README.gtm.md"),
			filepath.Join(dir, "README.md"),
		),
	)

	chk.SetArgs(
		"programName",
		"-r",
		"-z",
		dir,
	)

	chk.SetStdinData("Y\n")

	// Run command expecting the overwrite to be cancelled.
	main()

	got, wnt, err := getTestFiles(dir, "README.md")
	chk.NoErr(err)

	wnt[0] = strings.ReplaceAll(wnt[0], "**DO NOT MODIFY** ", "")
	chk.StrSlice(got, wnt)

	_, _, err = getTestFiles(dir, ".README.gtm.md")
	chk.Err(
		err,
		"open "+dir+"/.README.gtm.md: no such file or directory",
	)

	chk.Stdout(
		fmt.Sprintf(confirmMsg, filepath.Join(dir, "README.md")),
	)
}

func Test_Example1ReplaceTargetOverwriteDirVerbose(t *testing.T) {
	chk := sztestlog.CaptureLogAndStdout(t)
	defer chk.Release()

	dir := chk.CreateTmpDir()
	chk.NoErr(
		setup(dir, "README.md", "example1_test.go", "example1.go"),
	)

	chk.SetArgs(
		"programName",
		"-vvvvvv",
		"--log",
		"all",
		"-r",
		"-z",
		dir,
	)

	chk.SetStdinData("Y\n")

	// Run command expecting the overwrite to be cancelled.
	main()

	got, wnt, err := getTestFiles(dir, "README.md")
	chk.NoErr(err)

	wnt[0] = strings.ReplaceAll(wnt[0], "**DO NOT MODIFY** ", "")
	chk.StrSlice(got, wnt)

	pName := filepath.Join(dir, "README.md")
	chk.Stdout(
		fmt.Sprintf(confirmMsg, pName),
	)

	chk.Log(
		"I:filesToProcess: "+pName,
		"I:Expanding "+pName+" <inPlace> to: "+pName,
		"I:Loading Package info for: .",
		"I:getInfo(\"package\")",
		"I:getInfo(\"TimesTwo\")",
		"I:getInfo(\"TimesThree\")",
		"I:getInfo(\"InterfaceType\")",
		"I:getInfo(\"StructureType\")",
		"I:getInfo(\"TimesTwo\")",
		"I:getInfo(\"TimesTwo\")",
		"I:getInfo(\"TimesThree\")",
		"I:getInfo(\"TimesTwo\")",
		"I:getInfo(\"TimesTwo\")",
		"I:getInfo(\"TimesThree\")",
		"I:getInfo(\"TimesTwo\")",
		"I:getInfo(\"TimesTwo\")",
		"I:getInfo(\"TimesThree\")",
		"I:getInfo(\"ConstantGroup1\")",
		"I:getInfo(\"ConstantGroup1\")",
		"I:getInfo(\"ConstantGroupA\")",
	)
}

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

func Test_Example1CleanNoTargetAlternateOut(t *testing.T) {
	chk := sztestlog.CaptureLogAndStderrAndStdout(t)
	defer chk.Release()

	dir := chk.CreateTmpDir()
	altDir := chk.CreateTmpSubDir("altDir")

	chk.NoErr(setup(dir, "README.md", "example1_test.go", "example1.go"))

	chk.SetArgs(
		"programName",
		"--log",
		"all",
		"-l",
		"-h",
		"-c",
		"-o", altDir,
		filepath.Join(dir, "README.md"),
	)

	// Nor Run the main function with no -f arg requiring confirmation
	main()

	got, wnt, err := getTestFiles(altDir, ".README.gtm.md")
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	pName := filepath.Join(dir, "README.md")
	chk.Stdout(
		license + strings.Join(usage, "\n"),
	)

	rFile := filepath.Join(dir, "README.md")
	wFile := filepath.Join(altDir, ".README.gtm.md")
	chk.Log(
		"I:filesToProcess: "+pName,
		"I:Cleaning "+rFile+" to: "+wFile,
	)
	chk.Stderr()
}

func Test_JustHelp(t *testing.T) {
	chk := sztestlog.CaptureLogAndStdout(t)
	defer chk.Release()

	chk.SetArgs(
		"programName",
		"-vvvvvv",
		"--log",
		"all",
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

func Test_Usage_DoesNotExist(t *testing.T) {
	chk := sztestlog.CaptureLog(t)
	defer chk.Release()

	dir := chk.CreateTmpDir()

	chk.SetArgs(
		"programName",
		"-u", filepath.Join(dir, "DOES_NOT_EXIST.go"),
	)

	chk.Panic(
		main,
		"could not read file: "+
			"open /tmp/Test_Usage_DoesNotExist/DOES_NOT_EXIST.go: "+
			"no such file or directory",
	)

	chk.Log()
}

func Test_Usage_WarningEmptyFile(t *testing.T) {
	chk := sztestlog.CaptureLog(t)
	defer chk.Release()

	dir := chk.CreateTmpDir()

	goFile := chk.CreateTmpFileAs(dir, "file.go",
		[]byte(""+
			"",
		),
	)

	chk.SetArgs(
		"programName",
		"-vvvvvv",
		"--log",
		"all",
		"-u", goFile,
	)

	msg := strings.Join([]string{
		"This line will be the first replaced.",
		"And this line will be second and last replaced.",
	},
		"\n",
	)
	chk.SetStdinData(msg)

	main()

	//nolint:gosec // Ok.
	updatedBytes, err := os.ReadFile(filepath.Join(dir, "file.go"))
	chk.NoErr(err)
	chk.StrSlice(
		strings.Split(string(updatedBytes), "\n"),
		[]string{
			"/*",
			"This line will be the first replaced.",
			"And this line will be second and last replaced.",
			"*/",
		},
	)
	chk.Log(
		"W:blank usage file: '"+goFile+"'",
		"W:no previous usage found in: '"+goFile+"'",
		"W:package header not found in: '"+goFile+"'",
	)
}

func Test_Usage_WarningBlankFile(t *testing.T) {
	chk := sztestlog.CaptureLog(t)
	defer chk.Release()

	dir := chk.CreateTmpDir()

	goFile := chk.CreateTmpFileAs(dir, "file.go",
		[]byte(""+
			"\n"+
			"",
		),
	)

	chk.SetArgs(
		"programName",
		"-vvvvvv",
		"--log",
		"all",
		"-u", goFile,
	)

	msg := strings.Join([]string{
		"This line will be the first replaced.",
		"And this line will be second and last replaced.",
	},
		"\n",
	)
	chk.SetStdinData(msg)

	main()

	//nolint:gosec // Ok.
	updatedBytes, err := os.ReadFile(goFile)
	chk.NoErr(err)
	chk.StrSlice(
		strings.Split(string(updatedBytes), "\n"),
		[]string{
			"/*",
			"This line will be the first replaced.",
			"And this line will be second and last replaced.",
			"*/",
		},
	)
	chk.Log(
		"W:no previous usage found in: '"+goFile+"'",
		"W:package header not found in: '"+goFile+"'",
	)
}

func Test_Usage_WarningJustPackage(t *testing.T) {
	chk := sztestlog.CaptureLog(t)
	defer chk.Release()

	dir := chk.CreateTmpDir()

	goFile := chk.CreateTmpFileAs(dir, "file.go",
		[]byte(""+
			"package name"+
			"",
		),
	)

	chk.SetArgs(
		"programName",
		"-vvvvvv",
		"--log",
		"all",
		"-u", goFile,
	)

	msg := strings.Join([]string{
		"This line will be the first replaced.",
		"And this line will be second and last replaced.",
	},
		"\n",
	)
	chk.SetStdinData(msg)

	main()

	//nolint:gosec // Ok.
	updatedBytes, err := os.ReadFile(goFile)
	chk.NoErr(err)
	chk.StrSlice(
		strings.Split(string(updatedBytes), "\n"),
		[]string{
			"/*",
			"This line will be the first replaced.",
			"And this line will be second and last replaced.",
			"*/",
			"package name",
		},
	)

	chk.Log(
		"W:no previous usage found in: '" + goFile + "'",
	)
}

func Test_Usage_WarningJustPackageDuplicated(t *testing.T) {
	chk := sztestlog.CaptureLog(t)
	defer chk.Release()

	dir := chk.CreateTmpDir()

	goFile := chk.CreateTmpFileAs(dir, "file.go",
		[]byte(""+
			"/*\n"+
			"package notReal\n"+
			"*/\n"+
			"package name\n"+
			"",
		),
	)

	chk.SetArgs(
		"programName",
		"-vvvvvv", "--log", "all",
		"-u", goFile,
	)

	msg := strings.Join([]string{
		"This line will be the first replaced.",
		"And this line will be second and last replaced.",
	},
		"\n",
	)
	chk.SetStdinData(msg)

	main()

	//nolint:gosec // Ok.
	updatedBytes, err := os.ReadFile(goFile)
	chk.NoErr(err)
	chk.StrSlice(
		strings.Split(string(updatedBytes), "\n"),
		[]string{
			"/*",
			"package notReal",
			"*/",
			"",
			"/*",
			"This line will be the first replaced.",
			"And this line will be second and last replaced.",
			"*/",
			"package name",
			"",
		},
	)
	chk.Log(
		"W:multiple package delimiters.",
		"W:no previous usage found in: '"+goFile+"'",
	)
}

func Test_Usage_WarningJustUsageWithBlankPrefix(t *testing.T) {
	chk := sztestlog.CaptureLog(t)
	defer chk.Release()

	dir := chk.CreateTmpDir()

	goFile := chk.CreateTmpFileAs(dir, "file.go",
		[]byte(""+
			"\n"+
			"/*\n"+
			"# Usage: abc\n"+
			"/*\n"+
			"",
		),
	)

	chk.SetArgs(
		"programName",
		"-vvvvvv", "--log", "all", // Setup szlog for maximum output.
		"-u", goFile,
	)

	msg := strings.Join([]string{
		"# Usage: abc",
		"This line will be the first replaced.",
		"And this line will be second and last replaced.",
	},
		"\n",
	)
	chk.SetStdinData(msg)

	main()

	//nolint:gosec // Ok.
	updatedBytes, err := os.ReadFile(goFile)
	chk.NoErr(err)
	chk.StrSlice(
		strings.Split(string(updatedBytes), "\n"),
		[]string{
			"/*",
			"# Usage: abc",
			"This line will be the first replaced.",
			"And this line will be second and last replaced.",
			"*/",
		},
	)

	chk.Log(
		"W:package header not found in: '" + goFile + "'",
	)
}

func Test_Usage_WarningJustUsageWithNoBlankPrefix(t *testing.T) {
	chk := sztestlog.CaptureLog(t)
	defer chk.Release()

	dir := chk.CreateTmpDir()

	goFile := chk.CreateTmpFileAs(dir, "file.go",
		[]byte(""+
			"/*\n"+
			"# Usage: abc\n"+
			"/*\n"+
			"",
		),
	)

	chk.SetArgs(
		"programName",
		"-vvvvvv", "--log", "all",
		"-u", goFile,
	)

	msg := strings.Join([]string{
		"# Usage: abc",
		"This line will be the first replaced.",
		"And this line will be second and last replaced.",
	},
		"\n",
	)
	chk.SetStdinData(msg)

	main()

	//nolint:gosec // Ok.
	updatedBytes, err := os.ReadFile(goFile)
	chk.NoErr(err)
	chk.StrSlice(
		strings.Split(string(updatedBytes), "\n"),
		[]string{
			"/*",
			"# Usage: abc",
			"This line will be the first replaced.",
			"And this line will be second and last replaced.",
			"*/",
		},
	)

	chk.Log(
		"W:package header not found in: '" + goFile + "'",
	)
}

func Test_Usage_WarningPreUsageWithBlankPrefix(t *testing.T) {
	chk := sztestlog.CaptureLog(t)
	defer chk.Release()

	dir := chk.CreateTmpDir()

	goFile := chk.CreateTmpFileAs(dir, "file.go",
		[]byte(""+
			"This line is pre usage\n"+
			"\n"+
			"/*\n"+
			"# Usage: abc\n"+
			"/*\n"+
			"",
		),
	)

	chk.SetArgs(
		"programName",
		"-vvvvvv", "--log", "all",
		"-u", goFile,
	)

	msg := strings.Join([]string{
		"# Usage: abc",
		"This line will be the first replaced.",
		"And this line will be second and last replaced.",
	},
		"\n",
	)
	chk.SetStdinData(msg)

	main()

	//nolint:gosec // Ok.
	updatedBytes, err := os.ReadFile(goFile)
	chk.NoErr(err)
	chk.StrSlice(
		strings.Split(string(updatedBytes), "\n"),
		[]string{
			"This line is pre usage",
			"",
			"/*",
			"# Usage: abc",
			"This line will be the first replaced.",
			"And this line will be second and last replaced.",
			"*/",
		},
	)

	chk.Log(
		"W:package header not found in: '" + goFile + "'",
	)
}

func Test_Usage_WarningPreUsageWithNoBlankPrefix(t *testing.T) {
	chk := sztestlog.CaptureLog(t)
	defer chk.Release()

	dir := chk.CreateTmpDir()

	goFile := chk.CreateTmpFileAs(dir, "file.go",
		[]byte(""+
			"This line is pre usage\n"+
			"/*\n"+
			"# Usage: abc\n"+
			"/*\n"+
			"",
		),
	)

	chk.SetArgs(
		"programName",
		"-vvvvvv", "--log", "all",
		"-u", goFile,
	)

	msg := strings.Join([]string{
		"# Usage: abc",
		"This line will be the first replaced.",
		"And this line will be second and last replaced.",
	},
		"\n",
	)
	chk.SetStdinData(msg)

	main()

	//nolint:gosec // Ok.
	updatedBytes, err := os.ReadFile(goFile)
	chk.NoErr(err)
	chk.StrSlice(
		strings.Split(string(updatedBytes), "\n"),
		[]string{
			"This line is pre usage",
			"",
			"/*",
			"# Usage: abc",
			"This line will be the first replaced.",
			"And this line will be second and last replaced.",
			"*/",
		},
	)
	chk.Log(
		"W:package header not found in: '" + goFile + "'",
	)
}

func Test_Usage_WarningMultipleSeparators(t *testing.T) {
	chk := sztestlog.CaptureLog(t)
	defer chk.Release()

	dir := chk.CreateTmpDir()

	goFile := chk.CreateTmpFileAs(dir, "file.go",
		[]byte(""+
			"This line is pre usage\n"+
			"\n"+
			"/*\n"+
			"# Usage: abc\n"+
			"/*\n"+
			"# Usage: def\n"+
			"/*\n"+
			"",
		),
	)

	chk.SetArgs(
		"programName",
		"-vvvvvv", "--log", "all",
		"-u", goFile,
	)

	msg := strings.Join([]string{
		"# Usage: def",
		"This line will be the first replaced.",
		"And this line will be second and last replaced.",
	},
		"\n",
	)
	chk.SetStdinData(msg)

	main()

	//nolint:gosec // Ok.
	updatedBytes, err := os.ReadFile(goFile)
	chk.NoErr(err)
	chk.StrSlice(
		strings.Split(string(updatedBytes), "\n"),
		[]string{
			"This line is pre usage",
			"",
			"/*",
			"# Usage: abc",
			"",
			"/*",
			"# Usage: def",
			"This line will be the first replaced.",
			"And this line will be second and last replaced.",
			"*/",
		},
	)

	chk.Log(
		"W:multiple Usage delimiters.",
		"W:package header not found in: '"+goFile+"'",
	)
}

func Test_Usage_AllGood(t *testing.T) {
	chk := sztestlog.CaptureLog(t)
	defer chk.Release()

	dir := chk.CreateTmpDir()

	bytes, err := os.ReadFile("main_usage.go")
	chk.NoErr(err)

	goFile := chk.CreateTmpFileAs(dir, "file.go", bytes)

	chk.SetArgs(
		"programName",
		"-u", goFile,
	)

	msg := strings.Join([]string{
		"# Usage: abc",
		"    This line will be the first replaced.",
		"        And this line will be second and last replaced.",
	},
		"\n",
	)
	chk.SetStdinData(msg)

	main()

	//nolint:gosec // Ok.
	updatedBytes, err := os.ReadFile(goFile)
	chk.NoErr(err)
	//nolint:lll // Ok.
	chk.StrSlice(
		strings.Split(string(updatedBytes), "\n"),
		[]string{
			"/*",
			"   Golang To Github Markdown Utility: gotomd",
			"   Copyright (C) 2023, 2024 Leslie Dancsecs",
			"",
			"   This program is free software: you can redistribute it and/or modify",
			"   it under the terms of the GNU General Public License as published by",
			"   the Free Software Foundation, either version 3 of the License, or",
			"   (at your option) any later version.",
			"",
			"   This program is distributed in the hope that it will be useful,",
			"   but WITHOUT ANY WARRANTY; without even the implied warranty of",
			"   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the",
			"   GNU General Public License for more details.",
			"",
			"   You should have received a copy of the GNU General Public License",
			"   along with this program.  If not, see <https://www.gnu.org/licenses/>.",
			"*/",
			"",
			"/*",
			"# Usage: abc",
			"\tThis line will be the first replaced.",
			"\t    And this line will be second and last replaced.",
			"*/",
			"package main",
			"",
		},
	)
	chk.Log()
}
