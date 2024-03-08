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
)

const (
	example1     = "example1"
	example1Path = "." + string(os.PathSeparator) +
		example1 + string(os.PathSeparator)

	example2     = "example2"
	example2Path = "." + string(os.PathSeparator) +
		example2 + string(os.PathSeparator)
)

func Test_Example1ExpandTargetOverwriteDirVerbose(t *testing.T) {
	chk := sztest.CaptureLogAndStdout(t)
	defer chk.Release()

	dir := chk.CreateTmpDir()
	chk.NoErr(
		setup(dir, "README.md", "README.md.gtm", "example1_test.go", "example1.go"),
	)

	rName := filepath.Join(dir, "README.md.gtm")
	tName := filepath.Join(dir, "README.md")

	chk.SetupArgsAndFlags([]string{
		"programName",
		"-v",
		"-z",
		dir,
	})

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

	chk.SetupArgsAndFlags([]string{
		"programName",
		"-r",
		fName,
	})

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
	chk.SetupArgsAndFlags([]string{
		"programName",
		"-v",
		"-r",
		fName,
	})

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
	chk.SetupArgsAndFlags([]string{
		"programName",
		"-r",
		"-z",
		fName,
	})

	chk.SetStdinData("Y\n")

	// Run command expecting the overwrite to be cancelled.
	main()

	got, wnt, err := getTestFiles(dir, "README.md")
	chk.NoErr(err)

	wnt[0] = strings.ReplaceAll(wnt[0], "**DO NOT MODIFY** ", "")
	chk.StrSlice(got, wnt)

	_, _, err = getTestFiles(dir, "README.md.gtm")
	chk.Err(
		err,
		"open "+fName+".gtm: no such file or directory",
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

	chk.SetupArgsAndFlags([]string{
		"programName",
		"-r",
		"-z",
		dir,
	})

	chk.SetStdinData("Y\n")

	// Run command expecting the overwrite to be cancelled.
	main()

	got, wnt, err := getTestFiles(dir, "README.md")
	chk.NoErr(err)

	wnt[0] = strings.ReplaceAll(wnt[0], "**DO NOT MODIFY** ", "")
	chk.StrSlice(got, wnt)

	_, _, err = getTestFiles(dir, "README.md.gtm")
	chk.Err(
		err,
		"open "+dir+"/README.md.gtm: no such file or directory",
	)

	chk.Stdout("Confirm overwrite of " + dir + "/README.md (Y to overwrite)?\\s")
}

func Test_Example1ReplaceTargetOverwriteDirFromClean(t *testing.T) {
	chk := sztest.CaptureStdout(t)
	defer chk.Release()

	dir := chk.CreateTmpDir()
	chk.NoErr(
		setup(dir, "README.md.gtm", "example1_test.go", "example1.go"),
	)

	chk.NoErr(
		os.Rename(
			filepath.Join(dir, "README.md.gtm"),
			filepath.Join(dir, "README.md"),
		),
	)

	chk.SetupArgsAndFlags([]string{
		"programName",
		"-r",
		"-z",
		dir,
	})

	chk.SetStdinData("Y\n")

	// Run command expecting the overwrite to be cancelled.
	main()

	got, wnt, err := getTestFiles(dir, "README.md")
	chk.NoErr(err)

	wnt[0] = strings.ReplaceAll(wnt[0], "**DO NOT MODIFY** ", "")
	chk.StrSlice(got, wnt)

	_, _, err = getTestFiles(dir, "README.md.gtm")
	chk.Err(
		err,
		"open "+dir+"/README.md.gtm: no such file or directory",
	)

	chk.Stdout("Confirm overwrite of " + dir + "/README.md (Y to overwrite)?\\s")
}

func Test_Example1ReplaceTargetOverwriteDirVerbose(t *testing.T) {
	chk := sztest.CaptureLogAndStdout(t)
	defer chk.Release()

	dir := chk.CreateTmpDir()
	chk.NoErr(
		setup(dir, "README.md", "example1_test.go", "example1.go"),
	)

	chk.SetupArgsAndFlags([]string{
		"programName",
		"-v",
		"-r",
		"-z",
		dir,
	})

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
				os.FileMode(defaultPerm),
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

	chk.SetupArgsAndFlags([]string{
		"programName",
		"-v",
		"-l",
		"-c",
		"-o", altDir,
		filepath.Join(dir, "README.md"),
	})

	// Nor Run the main function with no -f arg requiring confirmation
	main()

	got, wnt, err := getTestFiles(altDir, "README.md.gtm")
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	pName := filepath.Join(dir, "README.md")
	chk.Stdout(
		license +
			"filesToProcess:  " + pName + "\n",
	)

	rFile := filepath.Join(dir, "README.md")
	wFile := filepath.Join(altDir, "README.md.gtm")
	chk.Log("Cleaning " + rFile + " to: " + wFile)
}
