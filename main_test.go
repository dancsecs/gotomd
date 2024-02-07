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

const sampleGoProjectOne = "sample_go_project_one"
const sampleGoProjectOnePath = "." + string(os.PathSeparator) +
	sampleGoProjectOne + string(os.PathSeparator)

const sampleGoProjectTwo = "sample_go_project_two"
const sampleGoProjectTwoPath = "." + string(os.PathSeparator) +
	sampleGoProjectTwo + string(os.PathSeparator)

func Test_SampleGoProjectOneExpandTargetOverwriteDirVerbose(t *testing.T) {
	chk := sztest.CaptureLogAndStdout(t)
	defer chk.Release()

	dir := chk.CreateTmpDir()
	chk.NoErr(
		setup(dir, "README.md", "README.md.gtm", "sample_test.go", "sample.go"),
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

func Test_SampleGoProjectOneReplaceNoTarget(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	dir := chk.CreateTmpDir()
	chk.NoErr(setup(dir, "sample_test.go", "sample.go"))

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

func Test_SampleGoProjectOneReplaceTargetCancel(t *testing.T) {
	chk := sztest.CaptureLogAndStdout(t)
	defer chk.Release()

	dir := chk.CreateTmpDir()
	chk.NoErr(
		setup(dir, "README.md", "sample_test.go", "sample.go"),
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

func Test_SampleGoProjectOneReplaceTargetOverwrite(t *testing.T) {
	chk := sztest.CaptureStdout(t)
	defer chk.Release()

	dir := chk.CreateTmpDir()
	chk.NoErr(
		setup(dir, "README.md", "sample_test.go", "sample.go"),
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

func Test_SampleGoProjectOneReplaceTargetOverwriteDir(t *testing.T) {
	chk := sztest.CaptureStdout(t)
	defer chk.Release()

	dir := chk.CreateTmpDir()
	chk.NoErr(
		setup(dir, "README.md", "sample_test.go", "sample.go"),
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

func Test_SampleGoProjectOneReplaceTargetOverwriteDirFromClean(t *testing.T) {
	chk := sztest.CaptureStdout(t)
	defer chk.Release()

	dir := chk.CreateTmpDir()
	chk.NoErr(
		setup(dir, "README.md.gtm", "sample_test.go", "sample.go"),
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

func Test_SampleGoProjectOneReplaceTargetOverwriteDirVerbose(t *testing.T) {
	chk := sztest.CaptureLogAndStdout(t)
	defer chk.Release()

	dir := chk.CreateTmpDir()
	chk.NoErr(
		setup(dir, "README.md", "sample_test.go", "sample.go"),
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
	var err error
	var b []byte

	const ext = ".sample"

	files = append(files, "go.mod"+ext, "go.sum"+ext)
	for i, mi := 0, len(files); i < mi && err == nil; i++ {
		b, err = os.ReadFile(filepath.Join("sample_go_project_one", files[i]))
		if err == nil {
			err = os.WriteFile(
				filepath.Join(dir, strings.TrimSuffix(files[i], ext)),
				b,
				os.FileMode(defaultPerm),
			)
		}
	}

	return err
}

func getTestFiles(dir, fName string) ([]string, []string, error) {
	gotBytes, err := os.ReadFile(filepath.Join(dir, fName))
	if err != nil {
		return nil, nil, err
	}
	wntBytes, err := os.ReadFile(filepath.Join("sample_go_project_one", fName))
	if err != nil {
		return nil, nil, err
	}
	return strings.Split(string(gotBytes), "\n"),
		strings.Split(string(wntBytes), "\n"),
		nil
}

func Test_SampleGoProjectOneCleanNoTargetAlternateOut(t *testing.T) {
	chk := sztest.CaptureLogAndStdout(t)
	defer chk.Release()

	dir := chk.CreateTmpDir()
	altDir := chk.CreateTmpSubDir("altDir")

	chk.NoErr(setup(dir, "README.md", "sample_test.go", "sample.go"))

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

func Test_SampleGoProjectOne_CpuProfile(t *testing.T) {
	chk := sztest.CaptureLogAndStdout(t)
	defer chk.Release()

	dir := chk.CreateTmpDir()
	altDir := chk.CreateTmpSubDir("altDir")

	pprofFilePath := filepath.Join(dir, "gotomd.pprof")

	chk.NoErr(setup(dir, "README.md.gtm", "sample_test.go", "sample.go"))

	chk.SetupArgsAndFlags([]string{
		"programName",
		"-f",
		"-o", altDir,
		"-U", pprofFilePath,
		"-u", "1",
		"-z",
		filepath.Join(dir, "README.md.gtm"),
	})

	// Nor Run the main function with no -f arg requiring confirmation
	main()

	got, wnt, err := getTestFiles(altDir, "README.md")
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	ppStat, err := os.Stat(pprofFilePath)
	chk.NoErr(err)
	chk.False(ppStat.IsDir())

	chk.Stdout()
	chk.Log()
}
