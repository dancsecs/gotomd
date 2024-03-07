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
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/dancsecs/sztest"
)

type cleanGlobals struct {
	forceOverwrite bool
	verbose        bool
}

func getCleanedFiles() (string, []string, []string, error) {
	const fName = "README.md.gtm"

	var (
		targetPath string
		err        error
		gotBytes   []byte
		wntBytes   []byte
	)

	targetPath = filepath.Join(outputDir, fName)

	gotBytes, err = os.ReadFile(targetPath)
	if err == nil {
		wntBytes, err = os.ReadFile(sampleGoProjectOnePath + fName)
	}

	if err != nil {
		return "", nil, nil, err
	}

	return targetPath,
		strings.Split(string(gotBytes), "\n"),
		strings.Split(string(wntBytes), "\n"),
		nil
}

func setupCleanDirs(makeTarget bool) error {
	const fName = "README.md.gtm"

	var err error

	if makeTarget {
		tFile := filepath.Join(outputDir, fName)
		err = os.WriteFile(tFile, nil, fs.FileMode(defaultPerm))
	}

	return err
}

func setupCleanGlobals(
	chk *sztest.Chk, override cleanGlobals,
) {
	chk.T().Helper()
	setupTest(chk, true, false, override.forceOverwrite, override.verbose)
}

// +-------------------------------------------------------+
// | Flag possibilities for type of test.                  |
// +------------+-----------+------------------+-----------+
// | cleanOnly  |  replace  |  forceOverwrite  |  verbose  |
// +------------+-----------+------------------+-----------+
// |  false     |   false   |     false        |   false   |
// |  false     |   true    |     false        |   false   |
// |  true      |   false   |     false        |   false   |
// +------------+-----------+------------------+-----------+
// |  false     |   false   |     false        |   true    |
// |  false     |   true    |     false        |   true    |
// |  true      |   false   |     false        |   true    |
// +------------+-----------+------------------+-----------+
// |  false     |   false   |     true         |   false   |
// |  false     |   true    |     true         |   false   |
// |  true      |   false   |     true         |   false   |
// +------------+-----------+------------------+-----------+
// |  false     |   false   |     true         |   true    |
// |  false     |   true    |     true         |   true    |
// |  true      |   false   |     true         |   true    |
// +------------+-----------+------------------+-----------+.

func Test_ProcessClean_NoTargetNoForceNoVerbose(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	setupCleanGlobals(chk, cleanGlobals{forceOverwrite: false, verbose: false})
	chk.NoErr(setupCleanDirs(false))

	chk.NoErr(cleanMD(sampleGoProjectOnePath + "README.md"))

	_, got, wnt, err := getCleanedFiles()
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Log()
}

func Test_ProcessClean_NoTargetForceNoVerbose(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	setupCleanGlobals(chk, cleanGlobals{forceOverwrite: true, verbose: false})
	chk.NoErr(setupCleanDirs(false))

	chk.NoErr(cleanMD(sampleGoProjectOnePath + "README.md"))

	_, got, wnt, err := getCleanedFiles()
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Log()
}

func Test_ProcessClean_NoTargetNoForceVerbose(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	setupCleanGlobals(chk, cleanGlobals{forceOverwrite: false, verbose: true})
	chk.NoErr(setupCleanDirs(false))

	chk.NoErr(cleanMD(sampleGoProjectOnePath + "README.md"))

	tPath, got, wnt, err := getCleanedFiles()
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Log("Cleaning " + sampleGoProjectOnePath + "README.md to: " + tPath)
}

func Test_ProcessClean_NoTargetForceVerbose(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	setupCleanGlobals(chk, cleanGlobals{forceOverwrite: true, verbose: true})
	chk.NoErr(setupCleanDirs(false))

	chk.NoErr(cleanMD(sampleGoProjectOnePath + "README.md"))

	tPath, got, wnt, err := getCleanedFiles()
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Log("Cleaning " + sampleGoProjectOnePath + "README.md to: " + tPath)
}

func Test_ProcessClean_CancelOverwriteNoForceNoVerbose(t *testing.T) {
	chk := sztest.CaptureLogAndStdout(t)
	defer chk.Release()

	setupCleanGlobals(chk, cleanGlobals{forceOverwrite: false, verbose: false})
	chk.NoErr(setupCleanDirs(true))

	chk.SetStdinData("N\n")

	// Run command expecting the overwrite to be cancelled.
	chk.NoErr(cleanMD(sampleGoProjectOnePath + "README.md"))

	tFile := filepath.Join(outputDir, "README.md.gtm")
	chk.Stdout("Confirm overwrite of " + tFile + " (Y to overwrite)? " +
		"overwrite cancelled",
	)

	chk.Log()
}

func Test_ProcessClean_CancelOverwriteForceNoVerbose(t *testing.T) {
	chk := sztest.CaptureLogAndStdout(t)
	defer chk.Release()

	setupCleanGlobals(chk, cleanGlobals{forceOverwrite: true, verbose: false})
	chk.NoErr(setupCleanDirs(true))

	// Run command expecting the overwrite to be cancelled.
	chk.NoErr(cleanMD(sampleGoProjectOnePath + "README.md"))

	chk.Stdout()

	chk.Log()
}

func Test_ProcessClean_CancelOverwriteNoForceVerbose(t *testing.T) {
	chk := sztest.CaptureLogAndStdout(t)
	defer chk.Release()

	setupCleanGlobals(chk, cleanGlobals{forceOverwrite: false, verbose: true})
	chk.NoErr(setupCleanDirs(true))

	chk.SetStdinData("N\n")

	// Run command expecting the overwrite to be cancelled.
	chk.NoErr(cleanMD(sampleGoProjectOnePath + "README.md"))

	tFile := filepath.Join(outputDir, "README.md.gtm")
	chk.Stdout("Confirm overwrite of " + tFile + " (Y to overwrite)? " +
		"overwrite cancelled",
	)

	chk.Log("Cleaning " + sampleGoProjectOnePath + "README.md to: " + tFile)
}

func Test_ProcessClean_CancelOverwriteForceVerbose(t *testing.T) {
	chk := sztest.CaptureLogAndStdout(t)
	defer chk.Release()

	setupCleanGlobals(chk, cleanGlobals{forceOverwrite: true, verbose: true})
	chk.NoErr(setupCleanDirs(true))

	// Run command expecting the overwrite to be cancelled.
	chk.NoErr(cleanMD(sampleGoProjectOnePath + "README.md"))

	chk.Stdout()

	tFile := filepath.Join(outputDir, "README.md.gtm")
	chk.Log("Cleaning " + sampleGoProjectOnePath + "README.md to: " + tFile)
}

func Test_ProcessClean_OverwriteNoForceNoVerbose(t *testing.T) {
	chk := sztest.CaptureLogAndStdout(t)
	defer chk.Release()

	setupCleanGlobals(chk, cleanGlobals{forceOverwrite: false, verbose: false})
	chk.NoErr(setupCleanDirs(true))

	chk.SetStdinData("Y\n")

	// Run command expecting the overwrite to be cancelled.
	chk.NoErr(cleanMD(sampleGoProjectOnePath + "README.md"))

	_, got, wnt, err := getCleanedFiles()
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	tFile := filepath.Join(outputDir, "README.md.gtm")
	chk.Stdout("Confirm overwrite of " + tFile + " (Y to overwrite)?\\s")

	chk.Log()
}

func Test_ProcessClean_OverwriteForceNoVerbose(t *testing.T) {
	chk := sztest.CaptureLogAndStdout(t)
	defer chk.Release()

	setupCleanGlobals(chk, cleanGlobals{forceOverwrite: true, verbose: false})
	chk.NoErr(setupCleanDirs(true))

	// Run command expecting the overwrite to be cancelled.
	chk.NoErr(cleanMD(sampleGoProjectOnePath + "README.md"))

	_, got, wnt, err := getCleanedFiles()
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Stdout()

	chk.Log()
}

func Test_ProcessClean_OverwriteNoForceVerbose(t *testing.T) {
	chk := sztest.CaptureLogAndStdout(t)
	defer chk.Release()

	setupCleanGlobals(chk, cleanGlobals{forceOverwrite: false, verbose: true})
	chk.NoErr(setupCleanDirs(true))

	chk.SetStdinData("Y\n")

	// Run command expecting the overwrite to be cancelled.
	chk.NoErr(cleanMD(sampleGoProjectOnePath + "README.md"))

	_, got, wnt, err := getCleanedFiles()
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	tFile := filepath.Join(outputDir, "README.md.gtm")
	chk.Stdout("Confirm overwrite of " + tFile + " (Y to overwrite)?\\s")

	chk.Log("Cleaning " + sampleGoProjectOnePath + "README.md to: " + tFile)
}

func Test_ProcessClean_OverwriteForceVerbose(t *testing.T) {
	chk := sztest.CaptureLogAndStdout(t)
	defer chk.Release()

	setupCleanGlobals(chk, cleanGlobals{forceOverwrite: true, verbose: true})
	chk.NoErr(setupCleanDirs(true))

	// Run command expecting the overwrite to be cancelled.
	chk.NoErr(cleanMD(sampleGoProjectOnePath + "README.md"))

	_, got, wnt, err := getCleanedFiles()
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Stdout()

	tFile := filepath.Join(outputDir, "README.md.gtm")
	chk.Log("Cleaning " + sampleGoProjectOnePath + "README.md to: " + tFile)
}
