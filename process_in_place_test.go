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

type inPlaceGlobals struct {
	forceOverwrite bool
	logLevel       szlog.LogLevel
}

func setupInPlaceGlobals(
	chk *sztest.Chk, override inPlaceGlobals,
) {
	chk.T().Helper()
	setupTest(chk, true, false, override.forceOverwrite, override.logLevel)
}

func setupInPlaceDirs(makeTarget bool) error {
	const fName = "README_SHORT.md"

	var (
		err   error
		tFile string
		fData []byte
	)

	if makeTarget {
		fData, err = os.ReadFile(filepath.Join(example1Path, fName))
		if err == nil {
			tFile = filepath.Join(outputDir, fName)
			//nolint:gosec // Ok.
			err = os.WriteFile(tFile, fData, fs.FileMode(defaultPerm))
		}
	}

	return err
}

func getInPlaceFiles() (string, []string, []string, error) {
	const fName = "README_SHORT.md"

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

func Test_ProcessInPlace_NoTargetNoForceNoVerbose(t *testing.T) {
	chk := sztestlog.CaptureLogAndStdout(t)
	defer chk.Release()

	setupInPlaceGlobals(
		chk, inPlaceGlobals{forceOverwrite: false, logLevel: szlog.LevelNone},
	)
	chk.NoErr(setupInPlaceDirs(false))

	chk.NoErr(replaceMDInPlace(example1Path + "README_SHORT.md"))

	_, got, wnt, err := getInPlaceFiles()
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Log()

	chk.Stdout()
}

func Test_ProcessInPlace_NoTargetForceNoVerbose(t *testing.T) {
	chk := sztestlog.CaptureLogAndStdout(t)
	defer chk.Release()

	setupInPlaceGlobals(
		chk, inPlaceGlobals{forceOverwrite: true, logLevel: szlog.LevelNone},
	)
	chk.NoErr(setupInPlaceDirs(false))

	chk.NoErr(replaceMDInPlace(example1Path + "README_SHORT.md"))

	_, got, wnt, err := getInPlaceFiles()
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Log()

	chk.Stdout()
}

func Test_ProcessInPlace_NoTargetNoForceVerbose(t *testing.T) {
	chk := sztestlog.CaptureLog(t)
	defer chk.Release()

	setupInPlaceGlobals(
		chk, inPlaceGlobals{forceOverwrite: false, logLevel: szlog.LevelAll},
	)
	chk.NoErr(setupInPlaceDirs(false))

	chk.NoErr(replaceMDInPlace(example1Path + "README_SHORT.md"))

	tFile, got, wnt, err := getInPlaceFiles()
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Log(
		"I:Expanding "+example1Path+"README_SHORT.md <inPlace> to: "+tFile,
		"I:getInfo(\"package\")",
	)
}

func Test_ProcessInPlace_NoTargetForceVerbose(t *testing.T) {
	chk := sztestlog.CaptureLog(t)
	defer chk.Release()

	setupInPlaceGlobals(
		chk, inPlaceGlobals{forceOverwrite: true, logLevel: szlog.LevelAll},
	)
	chk.NoErr(setupInPlaceDirs(false))

	chk.NoErr(replaceMDInPlace(example1Path + "README_SHORT.md"))

	tFile, got, wnt, err := getInPlaceFiles()
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Log(
		"I:Expanding "+example1Path+"README_SHORT.md <inPlace> to: "+tFile,
		"I:getInfo(\"package\")",
	)
}

func Test_ProcessInPlace_CancelOverwriteForceNoVerbose(t *testing.T) {
	chk := sztestlog.CaptureLog(t)
	defer chk.Release()

	setupInPlaceGlobals(
		chk, inPlaceGlobals{forceOverwrite: true, logLevel: szlog.LevelNone},
	)
	chk.NoErr(setupInPlaceDirs(true))

	chk.SetStdinData("N\n")

	chk.NoErr(replaceMDInPlace(example1Path + "README_SHORT.md"))

	_, got, wnt, err := getInPlaceFiles()
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Log()
}

func Test_ProcessInPlace_CancelOverwriteForceVerbose(t *testing.T) {
	chk := sztestlog.CaptureLog(t)
	defer chk.Release()

	setupInPlaceGlobals(
		chk, inPlaceGlobals{forceOverwrite: true, logLevel: szlog.LevelAll},
	)
	chk.NoErr(setupInPlaceDirs(true))

	chk.SetStdinData("N\n")

	chk.NoErr(replaceMDInPlace(example1Path + "README_SHORT.md"))

	tFile, got, wnt, err := getInPlaceFiles()
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Log(
		"I:Expanding "+example1Path+"README_SHORT.md <inPlace> to: "+tFile,
		"I:getInfo(\"package\")",
	)
}

func Test_ProcessInPlace_OverwriteForceNoVerbose(t *testing.T) {
	chk := sztestlog.CaptureLog(t)
	defer chk.Release()

	setupInPlaceGlobals(
		chk, inPlaceGlobals{forceOverwrite: true, logLevel: szlog.LevelNone},
	)
	chk.NoErr(setupInPlaceDirs(true))

	chk.SetStdinData("Y\n")

	chk.NoErr(replaceMDInPlace(example1Path + "README_SHORT.md"))

	_, got, wnt, err := getInPlaceFiles()
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Log()
}

func Test_ProcessInPlace_OverwriteForceVerbose(t *testing.T) {
	chk := sztestlog.CaptureLog(t)
	defer chk.Release()

	setupInPlaceGlobals(
		chk, inPlaceGlobals{forceOverwrite: true, logLevel: szlog.LevelAll},
	)
	chk.NoErr(setupInPlaceDirs(true))

	chk.SetStdinData("Y\n")

	chk.NoErr(replaceMDInPlace(example1Path + "README_SHORT.md"))

	tFile, got, wnt, err := getInPlaceFiles()
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Log(
		"I:Expanding "+example1Path+"README_SHORT.md <inPlace> to: "+tFile,
		"I:getInfo(\"package\")",
	)
}
