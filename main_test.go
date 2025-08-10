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
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/dancsecs/sztest"
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
	"Golang to 'github' markdown.",
	"",
	"\\s   programName [-v | --verbose ...] [-c | --clean] [-r | --replace]",
	"\\s               [-l | --license] [-h | --help] [-f | --force]",
	"\\s               [-z | --colorize] [-o | --output <dir>]",
	"\\s               [-u | --usage <filename>] [-p | --permission <perm>]",
	"\\s               [path ...]",
	"",
	"\\s   [-v | --verbose ...]",
	"\\s       Provide more information when processing.",
	"",
	"",
	"\\s   [-c | --clean]",
	"\\s       Reverse operation and remove generated markdown",
	"",
	"\\s       (Cannot be used with the [-r | --replace] option).",
	"",
	"",
	"\\s   [-r | --replace]",
	"\\s       Replace the *.MD in place\n",
	"\\s       (Cannot be used with the [-c | --clean] option).",
	"",
	"",
	"\\s   [-l | --license]",
	"\\s       Display license before program exits.",
	"",
	"",
	"\\s   [-h | --help]",
	"\\s       Display program usage information.",
	"",
	"",
	"\\s   [-f | --force]",
	"\\s       Do not confirm overwrite of destination.",
	"",
	"",
	"\\s   [-z | --colorize]",
	"\\s       Colorize go test output.",
	"",
	"",
	"\\s   [-o | --output <dir>]",
	"\\s       Direct all output to the specified directory.",
	"",
	"",
	"\\s   [-u | --usage <filename>]",
	"\\s       Replace the usage section in the given Go source file using",
	"\\s       content from standard input.  The section is identified as the",
	"\\s       text between the first occurrence of '^\\n/*\\n# Usage$'" +
		" and the",
	"\\s       following package declaration.  This allows keeping " +
		"command-line",
	"\\s       usage output (e.g., from --help) synchronized with the package",
	"\\s       documentation.",
	"",
	"",
	"\\s   [-p | --permission <perm>]",
	"\\s       Permissions to use when creating new file.",
	"",
	"\\s       (can only set RW bits)",
	"",
	"",
	"\\s   [path ...]",
	"\\s       A specific gotomd file template with the extension '*.gtm.md'" +
		" or a",
	"\\s       directory which will be searched for all matching template",
	"\\s       '*.gtm.md' files.  It defaults to the current directory: '.'",
}

func Test_Example1ExpandTargetOverwriteDirVerbose(t *testing.T) {
	chk := sztest.CaptureLogAndStdout(t)
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
		dir,
	)

	chk.SetStdinData("Y\n")

	chk.NoErr(os.Truncate(tName, 2))

	// Run command expecting the overwrite to be cancelled.
	main()

	got, wnt, err := getTestFiles(dir, "README.md")
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Stdout(
		"filesToProcess:  "+rName,
		"Confirm overwrite of "+tName+" (Y to overwrite)?\\s",
	)

	chk.Log(
		"Expanding "+rName+" to: "+tName,
		"Loading Package info for: .",
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
	)
}

////////////

func Test_Example1ReplaceNoTarget(t *testing.T) {
	chk := sztest.CaptureNothing(t)
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
	chk := sztest.CaptureLogAndStdout(t)
	defer chk.Release()

	dir := chk.CreateTmpDir()
	chk.NoErr(
		setup(dir, "README.md", "example1_test.go", "example1.go"),
	)

	fName := filepath.Join(dir, "README.md")
	chk.SetArgs(
		"programName",
		"-v",
		"-r",
		fName,
	)

	chk.SetStdinData("N\n")

	// Run command expecting the overwrite to be cancelled.
	main()

	chk.Stdout(
		"filesToProcess:  "+fName,
		"Confirm overwrite of "+fName+" (Y to overwrite)? "+
			"overwrite cancelled")

	chk.Log(
		"",
		"Expanding "+fName+" <inPlace> to: "+fName,
		"Loading Package info for: .",
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
	)
}

func Test_Example1ReplaceTargetOverwrite(t *testing.T) {
	chk := sztest.CaptureStdout(t)
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

	chk.Stdout("Confirm overwrite of " + fName + " (Y to overwrite)?\\s")
}

func Test_Example1ReplaceTargetOverwriteDir(t *testing.T) {
	chk := sztest.CaptureStdout(t)
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
		"Confirm overwrite of " + dir + "/README.md (Y to overwrite)?\\s",
	)
}

func Test_Example1ReplaceTargetOverwriteDirFromClean(t *testing.T) {
	chk := sztest.CaptureStdout(t)
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
		"Confirm overwrite of " + dir + "/README.md (Y to overwrite)?\\s",
	)
}

func Test_Example1ReplaceTargetOverwriteDirVerbose(t *testing.T) {
	chk := sztest.CaptureLogAndStdout(t)
	defer chk.Release()

	dir := chk.CreateTmpDir()
	chk.NoErr(
		setup(dir, "README.md", "example1_test.go", "example1.go"),
	)

	chk.SetArgs(
		"programName",
		"-v",
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
		"filesToProcess:  "+pName,
		"Confirm overwrite of "+pName+" (Y to overwrite)?\\s",
	)

	chk.Log(
		"Expanding "+pName+" <inPlace> to: "+pName,
		"Loading Package info for: .",
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
	chk := sztest.CaptureLogAndStdout(t)
	defer chk.Release()

	dir := chk.CreateTmpDir()
	altDir := chk.CreateTmpSubDir("altDir")

	chk.NoErr(setup(dir, "README.md", "example1_test.go", "example1.go"))

	chk.SetArgs(
		"programName",
		"-v",
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
		license+strings.Join(usage, "\n"),
		"",
		"filesToProcess:  "+pName+"\n",
	)

	rFile := filepath.Join(dir, "README.md")
	wFile := filepath.Join(altDir, ".README.gtm.md")
	chk.Log("Cleaning " + rFile + " to: " + wFile)
}

func Test_JustHelp(t *testing.T) {
	chk := sztest.CaptureLogAndStdout(t)
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

func Test_Usage_DoesNotExist(t *testing.T) {
	chk := sztest.CaptureLogAndStdout(t)
	defer chk.Release()

	dir := chk.CreateTmpDir()

	chk.SetArgs(
		"programName",
		"-v",
		"-u", filepath.Join(dir, "DOES_NOT_EXIST.go"),
	)

	chk.Panic(
		main,
		"could not read file: "+
			"open /tmp/Test_Usage_DoesNotExist/DOES_NOT_EXIST.go: "+
			"no such file or directory",
	)

	chk.Log()
	chk.Stdout()
}

func Test_Usage_WarningEmptyFile(t *testing.T) {
	chk := sztest.CaptureLogAndStdout(t)
	defer chk.Release()

	const rwAccess = 0o0600

	dir := chk.CreateTmpDir()

	chk.NoErr(os.WriteFile(
		filepath.Join(dir, "file.go"),
		[]byte(""+
			"",
		),
		rwAccess,
	),
	)

	chk.SetArgs(
		"programName",
		"-v",
		"-u", filepath.Join(dir, "file.go"),
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
			"# Usage: ",
			"This line will be the first replaced.",
			"And this line will be second and last replaced.",
			"*/",
		},
	)
	chk.Log()
	chk.Stdout()
}

func Test_Usage_WarningBlankFile(t *testing.T) {
	chk := sztest.CaptureLogAndStdout(t)
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
		"-v",
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
			"# Usage: ",
			"This line will be the first replaced.",
			"And this line will be second and last replaced.",
			"*/",
		},
	)
	chk.Log()
	chk.Stdout()
}

func Test_Usage_WarningJustPackage(t *testing.T) {
	chk := sztestlog.CaptureAll(t)
	defer chk.Release()

	dir := chk.CreateTmpDir()

	goFile := chk.CreateTmpFileAs(dir, "file.go",
		[]byte(""+
			"package name\n"+
			"",
		),
	)

	chk.SetArgs(
		"programName",
		"-v",
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
			"# Usage: ",
			"This line will be the first replaced.",
			"And this line will be second and last replaced.",
			"*/",
			"package name",
			"",
		},
	)
	chk.Stderr()
	chk.Log(
		"W:no previous usage found in: '" + goFile + "'",
	)
	chk.Stdout()
}

func Test_Usage_WarningJustPackageDuplicated(t *testing.T) {
	chk := sztestlog.CaptureAll(t)
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
		"-v",
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
			"# Usage: ",
			"This line will be the first replaced.",
			"And this line will be second and last replaced.",
			"*/",
			"package name",
			"",
		},
	)
	chk.Stderr()
	chk.Log(
		"W:multiple package delimiters.",
		"W:no previous usage found in: '"+goFile+"'",
	)
	chk.Stdout()
}

func Test_Usage_WarningJustUsageWithBlankPrefix(t *testing.T) {
	chk := sztestlog.CaptureAll(t)
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
		"-v",
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
			"# Usage: abc",
			"This line will be the first replaced.",
			"And this line will be second and last replaced.",
			"*/",
		},
	)
	chk.Stderr()
	chk.Log(
		"W:package header not found in: '" + goFile + "'",
	)
	chk.Stdout()
}

func Test_Usage_WarningJustUsageWithNoBlankPrefix(t *testing.T) {
	chk := sztestlog.CaptureAll(t)
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
		"-v",
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
			"# Usage: abc",
			"This line will be the first replaced.",
			"And this line will be second and last replaced.",
			"*/",
		},
	)
	chk.Stderr()
	chk.Log(
		"W:package header not found in: '" + goFile + "'",
	)
	chk.Stdout()
}

func Test_Usage_WarningPreUsageWithBlankPrefix(t *testing.T) {
	chk := sztestlog.CaptureAll(t)
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
		"-v",
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
			"This line is pre usage",
			"",
			"/*",
			"# Usage: abc",
			"This line will be the first replaced.",
			"And this line will be second and last replaced.",
			"*/",
		},
	)
	chk.Stderr()
	chk.Log(
		"W:package header not found in: '" + goFile + "'",
	)
	chk.Stdout()
}

func Test_Usage_WarningPreUsageWithNoBlankPrefix(t *testing.T) {
	chk := sztestlog.CaptureAll(t)
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
		"-v",
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
			"This line is pre usage",
			"",
			"/*",
			"# Usage: abc",
			"This line will be the first replaced.",
			"And this line will be second and last replaced.",
			"*/",
		},
	)
	chk.Stderr()
	chk.Log(
		"W:package header not found in: '" + goFile + "'",
	)
	chk.Stdout()
}

func Test_Usage_WarningMultipleSeparators(t *testing.T) {
	chk := sztestlog.CaptureAll(t)
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
		"-v",
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
	chk.Stderr()
	chk.Log(
		"W:multiple Usage delimiters.",
		"W:package header not found in: '"+goFile+"'",
	)
	chk.Stdout()
}
