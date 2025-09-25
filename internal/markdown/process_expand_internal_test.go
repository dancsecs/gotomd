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

package markdown

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/dancsecs/gotomd/internal/args"
	"github.com/dancsecs/szlog"
	"github.com/dancsecs/sztest"
	"github.com/dancsecs/sztestlog"
)

type expandGlobals struct {
	forceOverwrite bool
	verboseLevel   szlog.VerboseLevel
}

func setupTest(
	chk *sztest.Chk,
	tForceOverwrite bool,
	tVerbose szlog.VerboseLevel,
) {
	chk.T().Helper()

	origCWD, err := os.Getwd()
	chk.NoErr(err)

	origVerboseLevel := szlog.Verbose()

	outputDir := chk.CreateTmpDir()

	if tForceOverwrite {
		chk.SetArgs(
			"noProgName",
			"-f",
			"-o", outputDir,
		)
	} else {
		chk.SetArgs(
			"noProgName",
			"-o", outputDir,
		)
	}

	chk.NoErr(args.Process())

	szlog.SetVerbose(tVerbose)

	if chk.NoErr(err) {
		chk.PushPostReleaseFunc(func() error {
			args.Reset()
			szlog.SetVerbose(origVerboseLevel)

			return os.Chdir(origCWD)
		})
	}
}

func setupExpandGlobals(
	chk *sztest.Chk, override expandGlobals,
) {
	chk.T().Helper()
	setupTest(chk, override.forceOverwrite, override.verboseLevel)
}

func setupExpandDirs(makeTarget bool) error {
	const fName = "README.md"

	var (
		err   error
		tFile string
		fData []byte
	)

	if makeTarget {
		fData, err = os.ReadFile(filepath.Join(tstpkgPath, fName))
		if err == nil {
			tFile = filepath.Join(args.OutputDir(), fName)
			err = os.WriteFile(tFile, fData, args.Perm())
		}
	}

	return err
}

func getExpandFiles() (string, []string, []string, error) {
	const fName = "README.md"

	var (
		targetPath string
		err        error
		gotBytes   []byte
		wntBytes   []byte
	)

	targetPath = filepath.Join(args.OutputDir(), fName)
	gotBytes, err = os.ReadFile(targetPath) //nolint:gosec // Ok.

	if err == nil {
		wntBytes, err = os.ReadFile(tstpkgPath + sep + fName)
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

	setupExpandGlobals(
		chk, expandGlobals{forceOverwrite: false, verboseLevel: 0},
	)
	chk.NoErr(setupExpandDirs(false))

	chk.NoErr(ExpandMD(tstpkgPath + sep + ".README.gtm.md"))

	_, got, wnt, err := getExpandFiles()
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Log()
}

func Test_ProcessExpand_NoTargetForceNoVerbose(t *testing.T) {
	chk := sztestlog.CaptureLog(t)
	defer chk.Release()

	setupExpandGlobals(
		chk, expandGlobals{forceOverwrite: true, verboseLevel: 0},
	)
	chk.NoErr(setupExpandDirs(false))

	chk.NoErr(ExpandMD(tstpkgPath + sep + ".README.gtm.md"))

	_, got, wnt, err := getExpandFiles()
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Log()
}

func Test_ProcessExpand_NoTargetNoForceVerbose(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	setupExpandGlobals(
		chk,
		expandGlobals{forceOverwrite: false, verboseLevel: 6},
	)
	chk.NoErr(setupExpandDirs(false))

	chk.NoErr(ExpandMD(tstpkgPath + sep + ".README.gtm.md"))

	tFile, got, wnt, err := getExpandFiles()
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Stdout(
		"Expanding "+tstpkgPath+sep+".README.gtm.md to: "+tFile,
		"getInfo(\"package\")",
	)
}

func Test_ProcessExpand_NoTargetForceVerbose(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	setupExpandGlobals(
		chk,
		expandGlobals{forceOverwrite: true, verboseLevel: 6},
	)
	chk.NoErr(setupExpandDirs(false))

	chk.NoErr(ExpandMD(tstpkgPath + sep + ".README.gtm.md"))

	tFile, got, wnt, err := getExpandFiles()
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Stdout(
		"Expanding "+tstpkgPath+sep+".README.gtm.md to: "+tFile,
		"getInfo(\"package\")",
	)
}

func Test_ProcessExpand_CancelOverwriteTargetForceNoVerbose(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	setupExpandGlobals(
		chk, expandGlobals{forceOverwrite: true, verboseLevel: 0},
	)
	chk.NoErr(setupExpandDirs(true))

	chk.SetStdinData("N\n")

	chk.NoErr(ExpandMD(tstpkgPath + sep + ".README.gtm.md"))

	_, got, wnt, err := getExpandFiles()
	chk.NoErr(err)
	chk.StrSlice(got, wnt)
}

func Test_ProcessExpand_CancelOverwriteForceVerbose(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	setupExpandGlobals(
		chk, expandGlobals{forceOverwrite: true, verboseLevel: 6},
	)
	chk.NoErr(setupExpandDirs(true))

	chk.SetStdinData("N\n")

	chk.NoErr(ExpandMD(tstpkgPath + sep + ".README.gtm.md"))

	tFile, got, wnt, err := getExpandFiles()
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Stdout(
		"Expanding "+tstpkgPath+sep+".README.gtm.md to: "+tFile,
		"getInfo(\"package\")",
	)
}

func Test_ProcessExpand_OverwriteTargetForceNoVerbose(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	setupExpandGlobals(
		chk, expandGlobals{forceOverwrite: true, verboseLevel: 0},
	)
	chk.NoErr(setupExpandDirs(true))

	chk.SetStdinData("Y\n")

	chk.NoErr(ExpandMD(tstpkgPath + sep + ".README.gtm.md"))

	_, got, wnt, err := getExpandFiles()
	chk.NoErr(err)
	chk.StrSlice(got, wnt)
}

func Test_ProcessExpand_OverwriteForceVerbose(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	setupExpandGlobals(
		chk, expandGlobals{forceOverwrite: true, verboseLevel: 6},
	)
	chk.NoErr(setupExpandDirs(true))

	chk.SetStdinData("Y\n")

	chk.NoErr(ExpandMD(tstpkgPath + sep + ".README.gtm.md"))

	wFile, got, wnt, err := getExpandFiles()
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Stdout(
		"Expanding "+tstpkgPath+sep+".README.gtm.md to: "+wFile,
		"getInfo(\"package\")",
	)
}
