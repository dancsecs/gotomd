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

type expandGlobals struct {
	forceOverwrite bool
	logLevel       szlog.LogLevel
}

func setupExpandGlobals(
	chk *sztest.Chk, override expandGlobals,
) {
	chk.T().Helper()
	setupTest(chk, true, false, override.forceOverwrite, override.logLevel)
}

func setupExpandDirs(makeTarget bool) error {
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

func getExpandFiles() (string, []string, []string, error) {
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

func Test_ProcessExpand_NoTargetNoForceNoVerbose(t *testing.T) {
	chk := sztestlog.CaptureLog(t)
	defer chk.Release()

	// Clear packages from other runs.
	packageCache = make(map[string]*packageInfo)

	setupExpandGlobals(
		chk, expandGlobals{forceOverwrite: false, logLevel: szlog.LevelNone},
	)
	chk.NoErr(setupExpandDirs(false))

	chk.NoErr(expandMD(example1Path + ".README_SHORT.gtm.md"))

	_, got, wnt, err := getExpandFiles()
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Log()
}

func Test_ProcessExpand_NoTargetForceNoVerbose(t *testing.T) {
	chk := sztestlog.CaptureLog(t)
	defer chk.Release()

	setupExpandGlobals(
		chk, expandGlobals{forceOverwrite: true, logLevel: szlog.LevelNone},
	)
	chk.NoErr(setupExpandDirs(false))

	chk.NoErr(expandMD(example1Path + ".README_SHORT.gtm.md"))

	_, got, wnt, err := getExpandFiles()
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Log()
}

func Test_ProcessExpand_NoTargetNoForceVerbose(t *testing.T) {
	chk := sztestlog.CaptureLog(t)
	defer chk.Release()

	setupExpandGlobals(
		chk,
		expandGlobals{forceOverwrite: false, logLevel: szlog.LevelAll},
	)
	chk.NoErr(setupExpandDirs(false))

	chk.NoErr(expandMD(example1Path + ".README_SHORT.gtm.md"))

	tFile, got, wnt, err := getExpandFiles()
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Log(
		"I:Expanding "+example1Path+".README_SHORT.gtm.md to: "+tFile,
		"I:getInfo(\"package\")",
	)
}

func Test_ProcessExpand_NoTargetForceVerbose(t *testing.T) {
	chk := sztestlog.CaptureLog(t)
	defer chk.Release()

	setupExpandGlobals(
		chk,
		expandGlobals{forceOverwrite: true, logLevel: szlog.LevelAll},
	)
	chk.NoErr(setupExpandDirs(false))

	chk.NoErr(expandMD(example1Path + ".README_SHORT.gtm.md"))

	tFile, got, wnt, err := getExpandFiles()
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Log(
		"I:Expanding "+example1Path+".README_SHORT.gtm.md to: "+tFile,
		"I:getInfo(\"package\")",
	)
}

func Test_ProcessExpand_CancelOverwriteTargetForceNoVerbose(t *testing.T) {
	chk := sztestlog.CaptureLog(t)
	defer chk.Release()

	setupExpandGlobals(
		chk, expandGlobals{forceOverwrite: true, logLevel: szlog.LevelNone},
	)
	chk.NoErr(setupExpandDirs(true))

	chk.SetStdinData("N\n")

	chk.NoErr(expandMD(example1Path + ".README_SHORT.gtm.md"))

	_, got, wnt, err := getExpandFiles()
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Log()
}

func Test_ProcessExpand_CancelOverwriteForceVerbose(t *testing.T) {
	chk := sztestlog.CaptureLog(t)
	defer chk.Release()

	setupExpandGlobals(
		chk, expandGlobals{forceOverwrite: true, logLevel: szlog.LevelAll},
	)
	chk.NoErr(setupExpandDirs(true))

	chk.SetStdinData("N\n")

	chk.NoErr(expandMD(example1Path + ".README_SHORT.gtm.md"))

	tFile, got, wnt, err := getExpandFiles()
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Log(
		"I:Expanding "+example1Path+".README_SHORT.gtm.md to: "+tFile,
		"I:getInfo(\"package\")",
	)
}

func Test_ProcessExpand_OverwriteTargetForceNoVerbose(t *testing.T) {
	chk := sztestlog.CaptureLog(t)
	defer chk.Release()

	setupExpandGlobals(
		chk, expandGlobals{forceOverwrite: true, logLevel: szlog.LevelNone},
	)
	chk.NoErr(setupExpandDirs(true))

	chk.SetStdinData("Y\n")

	chk.NoErr(expandMD(example1Path + ".README_SHORT.gtm.md"))

	_, got, wnt, err := getExpandFiles()
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Log()
}

func Test_ProcessExpand_OverwriteForceVerbose(t *testing.T) {
	chk := sztestlog.CaptureLog(t)
	defer chk.Release()

	setupExpandGlobals(
		chk, expandGlobals{forceOverwrite: true, logLevel: szlog.LevelAll},
	)
	chk.NoErr(setupExpandDirs(true))

	chk.SetStdinData("Y\n")

	chk.NoErr(expandMD(example1Path + ".README_SHORT.gtm.md"))

	wFile, got, wnt, err := getExpandFiles()
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Log(
		"I:Expanding "+example1Path+".README_SHORT.gtm.md to: "+wFile,
		"I:getInfo(\"package\")",
	)
}
