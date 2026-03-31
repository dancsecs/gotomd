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

package expand_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/dancsecs/gotomd/internal/args"
	"github.com/dancsecs/gotomd/internal/expand"
	"github.com/dancsecs/gotomd/internal/gopkg"
	"github.com/dancsecs/szlog"
	"github.com/dancsecs/sztest"
	"github.com/dancsecs/sztestlog"
)

const (
	sep        = string(os.PathSeparator)
	tstpkg     = "tstpkg"
	tstpkgPath = "." + sep + "testdata" + sep + tstpkg
)

type expandGlobals struct {
	forceOverwrite bool
	verboseLevel   szlog.VerboseLevel
}

func templateName(chk *sztest.Chk, fName string) string {
	rootFName, isGo := strings.CutSuffix(fName, ".go")
	if isGo {
		return tstpkgPath + sep + "." + rootFName + ".gtm.go"
	}

	rootFName, isMD := strings.CutSuffix(fName, ".md")
	if isMD {
		return tstpkgPath + sep + "." + rootFName + ".gtm.md"
	}

	chk.T().Helper()
	chk.Error("unknown template: " + fName)

	return fName
}

func setupTest(
	chk *sztest.Chk,
	tForceOverwrite bool,
	tVerbose szlog.VerboseLevel,
) {
	chk.T().Helper()

	gopkg.Reset()

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

func setupExpandDirs(makeTarget bool, fName string) error {
	var (
		err   error
		tFile string
		fData []byte
	)

	if makeTarget {
		//nolint:gosec // Ok.
		fData, err = os.ReadFile(filepath.Join(tstpkgPath, fName))
		if err == nil {
			tFile = filepath.Join(args.OutputDir(), fName)
			err = os.WriteFile(tFile, fData, args.Perm())
		}
	}

	return err
}

func getExpandFiles(fName string) (string, []string, []string, error) {
	var (
		targetPath string
		err        error
		gotBytes   []byte
		wntBytes   []byte
	)

	targetPath = filepath.Join(args.OutputDir(), fName)
	gotBytes, err = os.ReadFile(targetPath) //nolint:gosec // Ok.

	if err == nil {
		//nolint:gosec // Ok.
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

func Test_ProcessExpand_MD_NoTargetNoForceNoVerbose(t *testing.T) {
	chk := sztestlog.CaptureLog(t)
	defer chk.Release()

	const targetFName = "README.md"

	setupExpandGlobals(
		chk, expandGlobals{forceOverwrite: false, verboseLevel: 0},
	)
	chk.NoErr(setupExpandDirs(false, targetFName))

	chk.NoErr(expand.Process(templateName(chk, targetFName)))

	_, got, wnt, err := getExpandFiles(targetFName)
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Log()
}

func Test_ProcessExpand_MD_NoTargetForceNoVerbose(t *testing.T) {
	chk := sztestlog.CaptureLog(t)
	defer chk.Release()

	const targetFName = "README.md"

	setupExpandGlobals(
		chk, expandGlobals{forceOverwrite: true, verboseLevel: 0},
	)
	chk.NoErr(setupExpandDirs(false, targetFName))

	chk.NoErr(expand.Process(templateName(chk, targetFName)))

	_, got, wnt, err := getExpandFiles(targetFName)
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Log()
}

func Test_ProcessExpand_MD_NoTargetNoForceVerbose(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	const (
		targetFName = "README.md"
		sourceFName = ".README.gtm.md"
	)

	setupExpandGlobals(
		chk,
		expandGlobals{forceOverwrite: false, verboseLevel: 6},
	)
	chk.NoErr(setupExpandDirs(false, targetFName))

	chk.NoErr(expand.Process(templateName(chk, targetFName)))

	tFile, got, wnt, err := getExpandFiles(targetFName)
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Stdout(
		"Expanding "+tstpkgPath+sep+sourceFName+" to: "+tFile,
		"Loading package info for: .",
		"getInfo(\"package\")",
	)
}

func Test_ProcessExpand_MD_NoTargetForceVerbose(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	const (
		targetFName = "README.md"
		sourceFName = ".README.gtm.md"
	)

	setupExpandGlobals(
		chk,
		expandGlobals{forceOverwrite: true, verboseLevel: 6},
	)
	chk.NoErr(setupExpandDirs(false, targetFName))

	chk.NoErr(expand.Process(templateName(chk, targetFName)))

	tFile, got, wnt, err := getExpandFiles(targetFName)
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Stdout(
		"Expanding "+tstpkgPath+sep+sourceFName+" to: "+tFile,
		"Loading package info for: .",
		"getInfo(\"package\")",
	)
}

func Test_ProcessExpand_MD_CancelOverwriteTargetForceNoVerbose(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	const targetFName = "README.md"

	setupExpandGlobals(
		chk, expandGlobals{forceOverwrite: true, verboseLevel: 0},
	)
	chk.NoErr(setupExpandDirs(true, targetFName))

	chk.SetStdinData("N\n")

	chk.NoErr(expand.Process(templateName(chk, targetFName)))

	_, got, wnt, err := getExpandFiles(targetFName)
	chk.NoErr(err)
	chk.StrSlice(got, wnt)
}

func Test_ProcessExpand_MD_CancelOverwriteForceVerbose(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	const (
		targetFName = "README.md"
		sourceFName = ".README.gtm.md"
	)

	setupExpandGlobals(
		chk, expandGlobals{forceOverwrite: true, verboseLevel: 6},
	)
	chk.NoErr(setupExpandDirs(true, targetFName))

	chk.SetStdinData("N\n")

	chk.NoErr(expand.Process(templateName(chk, targetFName)))

	tFile, got, wnt, err := getExpandFiles(targetFName)
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Stdout(
		"Expanding "+tstpkgPath+sep+sourceFName+" to: "+tFile,
		"Loading package info for: .",
		"getInfo(\"package\")",
	)
}

func Test_ProcessExpand_MD_OverwriteTargetForceNoVerbose(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	const targetFName = "README.md"

	setupExpandGlobals(
		chk, expandGlobals{forceOverwrite: true, verboseLevel: 0},
	)
	chk.NoErr(setupExpandDirs(true, targetFName))

	chk.SetStdinData("Y\n")

	chk.NoErr(expand.Process(templateName(chk, targetFName)))

	_, got, wnt, err := getExpandFiles(targetFName)
	chk.NoErr(err)
	chk.StrSlice(got, wnt)
}

func Test_ProcessExpand_MD_OverwriteForceVerbose(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	const (
		targetFName = "README.md"
		sourceFName = ".README.gtm.md"
	)

	setupExpandGlobals(
		chk, expandGlobals{forceOverwrite: true, verboseLevel: 6},
	)
	chk.NoErr(setupExpandDirs(true, targetFName))

	chk.SetStdinData("Y\n")

	chk.NoErr(expand.Process(templateName(chk, targetFName)))

	wFile, got, wnt, err := getExpandFiles(targetFName)
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Stdout(
		"Expanding "+tstpkgPath+sep+sourceFName+" to: "+wFile,
		"Loading package info for: .",
		"getInfo(\"package\")",
	)
}

//
//
//
//
//
//
//
//

func Test_ProcessExpand_GO_NoTargetNoForceNoVerbose(t *testing.T) {
	chk := sztestlog.CaptureLog(t)
	defer chk.Release()

	const targetFName = "doc.go"

	setupExpandGlobals(
		chk, expandGlobals{forceOverwrite: false, verboseLevel: 0},
	)
	chk.NoErr(setupExpandDirs(false, targetFName))

	chk.NoErr(expand.Process(templateName(chk, targetFName)))

	_, got, wnt, err := getExpandFiles(targetFName)
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Log()
}

func Test_ProcessExpand_GO_NoTargetForceNoVerbose(t *testing.T) {
	chk := sztestlog.CaptureLog(t)
	defer chk.Release()

	const targetFName = "doc.go"

	setupExpandGlobals(
		chk, expandGlobals{forceOverwrite: true, verboseLevel: 0},
	)
	chk.NoErr(setupExpandDirs(false, targetFName))

	chk.NoErr(expand.Process(templateName(chk, targetFName)))

	_, got, wnt, err := getExpandFiles(targetFName)
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Log()
}

func Test_ProcessExpand_GO_NoTargetNoForceVerbose(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	const (
		targetFName = "doc.go"
		sourceFName = ".doc.gtm.go"
	)

	setupExpandGlobals(
		chk,
		expandGlobals{forceOverwrite: false, verboseLevel: 8},
	)
	chk.NoErr(setupExpandDirs(false, targetFName))

	chk.NoErr(expand.Process(templateName(chk, targetFName)))

	tFile, got, wnt, err := getExpandFiles(targetFName)
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Stdout(
		"Expanding " + tstpkgPath + sep + sourceFName + " to: " + tFile,
	)
}

func Test_ProcessExpand_GO_NoTargetForceVerbose(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	const (
		targetFName = "doc.go"
		sourceFName = ".doc.gtm.go"
	)

	setupExpandGlobals(
		chk,
		expandGlobals{forceOverwrite: true, verboseLevel: 6},
	)
	chk.NoErr(setupExpandDirs(false, targetFName))

	chk.NoErr(expand.Process(templateName(chk, targetFName)))

	tFile, got, wnt, err := getExpandFiles(targetFName)
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Stdout(
		"Expanding " + tstpkgPath + sep + sourceFName + " to: " + tFile,
	)
}

func Test_ProcessExpand_GO_CancelOverwriteTargetForceNoVerbose(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	const targetFName = "doc.go"

	setupExpandGlobals(
		chk, expandGlobals{forceOverwrite: true, verboseLevel: 0},
	)
	chk.NoErr(setupExpandDirs(true, targetFName))

	chk.SetStdinData("N\n")

	chk.NoErr(expand.Process(templateName(chk, targetFName)))

	_, got, wnt, err := getExpandFiles(targetFName)
	chk.NoErr(err)
	chk.StrSlice(got, wnt)
}

func Test_ProcessExpand_GO_CancelOverwriteForceVerbose(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	const (
		targetFName = "doc.go"
		sourceFName = ".doc.gtm.go"
	)

	setupExpandGlobals(
		chk, expandGlobals{forceOverwrite: true, verboseLevel: 6},
	)
	chk.NoErr(setupExpandDirs(true, targetFName))

	chk.SetStdinData("N\n")

	chk.NoErr(expand.Process(templateName(chk, targetFName)))

	tFile, got, wnt, err := getExpandFiles(targetFName)
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Stdout(
		"Expanding " + tstpkgPath + sep + sourceFName + " to: " + tFile,
	)
}

func Test_ProcessExpand_GO_OverwriteTargetForceNoVerbose(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	const targetFName = "doc.go"

	setupExpandGlobals(
		chk, expandGlobals{forceOverwrite: true, verboseLevel: 0},
	)
	chk.NoErr(setupExpandDirs(true, targetFName))

	chk.SetStdinData("Y\n")

	chk.NoErr(expand.Process(templateName(chk, targetFName)))

	_, got, wnt, err := getExpandFiles(targetFName)
	chk.NoErr(err)
	chk.StrSlice(got, wnt)
}

func Test_ProcessExpand_GO_OverwriteForceVerbose(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	const (
		targetFName = "doc.go"
		sourceFName = ".doc.gtm.go"
	)

	setupExpandGlobals(
		chk, expandGlobals{forceOverwrite: true, verboseLevel: 6},
	)
	chk.NoErr(setupExpandDirs(true, targetFName))

	chk.SetStdinData("Y\n")

	chk.NoErr(expand.Process(templateName(chk, targetFName)))

	wFile, got, wnt, err := getExpandFiles(targetFName)
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Stdout(
		"Expanding " + tstpkgPath + sep + sourceFName + " to: " + wFile,
	)
}
