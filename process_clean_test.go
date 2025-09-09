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

	"github.com/dancsecs/szlog"
	"github.com/dancsecs/sztest"
	"github.com/dancsecs/sztestlog"
)

type cleanGlobals struct {
	forceOverwrite bool
	logLevel       szlog.LogLevel
}

func getCleanedFiles() (string, []string, []string, error) {
	const fName = ".README.gtm.md"

	var (
		targetPath string
		err        error
		gotBytes   []byte
		wntBytes   []byte
	)

	targetPath = filepath.Join(outputDir, fName)

	gotBytes, err = os.ReadFile(targetPath) //nolint:gosec // Ok.
	if err == nil {
		wntBytes, err = os.ReadFile(example1Path + fName)
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
	const fName = ".README.gtm.md"

	var err error

	if makeTarget {
		tFile := filepath.Join(outputDir, fName)
		//nolint:gosec // Ok.
		err = os.WriteFile(tFile, nil, fs.FileMode(defaultPerm))
	}

	return err
}

func setupCleanGlobals(
	chk *sztest.Chk, override cleanGlobals,
) {
	chk.T().Helper()
	setupTest(chk, true, false, override.forceOverwrite, override.logLevel)
}

// +-------------------------------------------------------+
// | Option possibilities for type of test.                |
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
	chk := sztestlog.CaptureLog(t)
	defer chk.Release()

	setupCleanGlobals(
		chk,
		cleanGlobals{forceOverwrite: false, logLevel: szlog.LevelNone},
	)
	chk.NoErr(setupCleanDirs(false))

	chk.NoErr(cleanMD(example1Path + "README.md"))

	_, got, wnt, err := getCleanedFiles()
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Log()
}

func Test_ProcessClean_NoTargetForceNoVerbose(t *testing.T) {
	chk := sztestlog.CaptureLog(t)
	defer chk.Release()

	setupCleanGlobals(
		chk,
		cleanGlobals{forceOverwrite: true, logLevel: szlog.LevelNone},
	)
	chk.NoErr(setupCleanDirs(false))

	chk.NoErr(cleanMD(example1Path + "README.md"))

	_, got, wnt, err := getCleanedFiles()
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Log()
}

func Test_ProcessClean_NoTargetNoForceVerbose(t *testing.T) {
	chk := sztestlog.CaptureLog(t)
	defer chk.Release()

	setupCleanGlobals(
		chk,
		cleanGlobals{forceOverwrite: false, logLevel: szlog.LevelAll},
	)
	chk.NoErr(setupCleanDirs(false))

	chk.NoErr(cleanMD(example1Path + "README.md"))

	tPath, got, wnt, err := getCleanedFiles()
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Log("I:Cleaning " + example1Path + "README.md to: " + tPath)
}

func Test_ProcessClean_NoTargetForceVerbose(t *testing.T) {
	chk := sztestlog.CaptureLog(t)
	defer chk.Release()

	setupCleanGlobals(
		chk,
		cleanGlobals{forceOverwrite: true, logLevel: szlog.LevelAll},
	)
	chk.NoErr(setupCleanDirs(false))

	chk.NoErr(cleanMD(example1Path + "README.md"))

	tPath, got, wnt, err := getCleanedFiles()
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Log("I:Cleaning " + example1Path + "README.md to: " + tPath)
}

func Test_ProcessClean_CancelOverwriteNoForceNoVerbose(t *testing.T) {
	chk := sztestlog.CaptureLogAndStdout(t)
	defer chk.Release()

	setupCleanGlobals(
		chk,
		cleanGlobals{forceOverwrite: false, logLevel: szlog.LevelNone},
	)
	chk.NoErr(setupCleanDirs(true))

	chk.SetStdinData("N\n")

	// Run command expecting the overwrite to be cancelled.
	chk.NoErr(cleanMD(example1Path + "README.md"))

	tFile := filepath.Join(outputDir, ".README.gtm.md")
	chk.Stdout("Confirm overwrite of " + tFile + " (Y to overwrite)? " +
		"overwrite cancelled",
	)

	chk.Log()
}

func Test_ProcessClean_CancelOverwriteForceNoVerbose(t *testing.T) {
	chk := sztestlog.CaptureLog(t)
	defer chk.Release()

	setupCleanGlobals(
		chk,
		cleanGlobals{forceOverwrite: true, logLevel: szlog.LevelNone},
	)
	chk.NoErr(setupCleanDirs(true))

	// Run command expecting the overwrite to be cancelled.
	chk.NoErr(cleanMD(example1Path + "README.md"))

	chk.Log()
}

func Test_ProcessClean_CancelOverwriteNoForceVerbose(t *testing.T) {
	chk := sztestlog.CaptureLogAndStdout(t)
	defer chk.Release()

	setupCleanGlobals(
		chk,
		cleanGlobals{forceOverwrite: false, logLevel: szlog.LevelAll},
	)
	chk.NoErr(setupCleanDirs(true))

	chk.SetStdinData("N\n")

	// Run command expecting the overwrite to be cancelled.
	chk.NoErr(cleanMD(example1Path + "README.md"))

	tFile := filepath.Join(outputDir, ".README.gtm.md")
	chk.Stdout("Confirm overwrite of " + tFile + " (Y to overwrite)? " +
		"overwrite cancelled",
	)

	chk.Log("I:Cleaning " + example1Path + "README.md to: " + tFile)
}

func Test_ProcessClean_CancelOverwriteForceVerbose(t *testing.T) {
	chk := sztestlog.CaptureLog(t)
	defer chk.Release()

	setupCleanGlobals(
		chk,
		cleanGlobals{forceOverwrite: true, logLevel: szlog.LevelAll},
	)
	chk.NoErr(setupCleanDirs(true))

	// Run command expecting the overwrite to be cancelled.
	chk.NoErr(cleanMD(example1Path + "README.md"))

	tFile := filepath.Join(outputDir, ".README.gtm.md")
	chk.Log("I:Cleaning " + example1Path + "README.md to: " + tFile)
}

func Test_ProcessClean_OverwriteNoForceNoVerbose(t *testing.T) {
	chk := sztestlog.CaptureLogAndStdout(t)
	defer chk.Release()

	setupCleanGlobals(
		chk,
		cleanGlobals{forceOverwrite: false, logLevel: szlog.LevelNone},
	)
	chk.NoErr(setupCleanDirs(true))

	chk.SetStdinData("Y\n")

	// Run command expecting the overwrite to be cancelled.
	chk.NoErr(cleanMD(example1Path + "README.md"))

	_, got, wnt, err := getCleanedFiles()
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	tFile := filepath.Join(outputDir, ".README.gtm.md")
	chk.Stdout("Confirm overwrite of " + tFile + " (Y to overwrite)? ")

	chk.Log()
}

func Test_ProcessClean_OverwriteForceNoVerbose(t *testing.T) {
	chk := sztestlog.CaptureLog(t)
	defer chk.Release()

	setupCleanGlobals(
		chk,
		cleanGlobals{forceOverwrite: true, logLevel: szlog.LevelNone},
	)
	chk.NoErr(setupCleanDirs(true))

	// Run command expecting the overwrite to be cancelled.
	chk.NoErr(cleanMD(example1Path + "README.md"))

	_, got, wnt, err := getCleanedFiles()
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Log()
}

func Test_ProcessClean_OverwriteNoForceVerbose(t *testing.T) {
	chk := sztestlog.CaptureLogAndStdout(t)
	defer chk.Release()

	setupCleanGlobals(
		chk,
		cleanGlobals{forceOverwrite: false, logLevel: szlog.LevelAll},
	)
	chk.NoErr(setupCleanDirs(true))

	chk.SetStdinData("Y\n")

	// Run command expecting the overwrite to be cancelled.
	chk.NoErr(cleanMD(example1Path + "README.md"))

	_, got, wnt, err := getCleanedFiles()
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	tFile := filepath.Join(outputDir, ".README.gtm.md")
	chk.Stdout("Confirm overwrite of " + tFile + " (Y to overwrite)? ")

	chk.Log("I:Cleaning " + example1Path + "README.md to: " + tFile)
}

func Test_ProcessClean_OverwriteForceVerbose(t *testing.T) {
	chk := sztestlog.CaptureLog(t)
	defer chk.Release()

	setupCleanGlobals(
		chk, cleanGlobals{forceOverwrite: true, logLevel: szlog.LevelAll},
	)
	chk.NoErr(setupCleanDirs(true))

	// Run command expecting the overwrite to be cancelled.
	chk.NoErr(cleanMD(example1Path + "README.md"))

	_, got, wnt, err := getCleanedFiles()
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	tFile := filepath.Join(outputDir, ".README.gtm.md")
	chk.Log("I:Cleaning " + example1Path + "README.md to: " + tFile)
}
