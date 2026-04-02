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
	"github.com/dancsecs/gotomd/internal/errs"
	"github.com/dancsecs/gotomd/internal/expand"
	"github.com/dancsecs/gotomd/internal/gopkg"
	"github.com/dancsecs/szlog"
	"github.com/dancsecs/sztest"
	"github.com/dancsecs/sztestlog"
)

const (
	sep         = string(os.PathSeparator)
	tstpkg      = "tstpkg"
	tstcmd      = "tstcmd"
	tstpkgPaths = "." + sep + "testdata" + sep + tstpkg
	tstcmdPaths = "." + sep + "testdata" + sep + tstcmd
)

type expandGlobals struct {
	forceOverwrite bool
	verboseLevel   szlog.VerboseLevel
}

func templateName(chk *sztest.Chk, fPath, fName string) string {
	rootFName, isGo := strings.CutSuffix(fName, ".go")
	if isGo {
		return fPath + sep + "." + rootFName + ".gtm.go"
	}

	rootFName, isMD := strings.CutSuffix(fName, ".md")
	if isMD {
		return fPath + sep + "." + rootFName + ".gtm.md"
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

func setupExpandDirs(makeTarget bool, fPath, fName string) error {
	var (
		err   error
		tFile string
		fData []byte
	)

	if makeTarget {
		//nolint:gosec // Ok.
		fData, err = os.ReadFile(filepath.Join(fPath, fName))
		if err == nil {
			tFile = filepath.Join(args.OutputDir(), fName)
			err = os.WriteFile(tFile, fData, args.Perm())
		}
	}

	return err
}

func getExpandFiles(fPath, fName string) (string, []string, []string, error) {
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
		wntBytes, err = os.ReadFile(fPath + sep + fName)
	}

	if err != nil {
		return "", nil, nil, err
	}

	return targetPath,
		strings.Split(string(gotBytes), "\n"),
		strings.Split(string(wntBytes), "\n"),
		nil
}

func Test_ProcessExpand_UnknownTemplate(t *testing.T) {
	chk := sztestlog.CaptureLog(t)
	defer chk.Release()

	chk.Err(
		expand.Process("this.unknownTemplate"),
		chk.ErrChain(
			errs.ErrUnknownTemplate,
		),
	)

	chk.Log()
}

func Test_ProcessExpand_MDPkg_NoTargetNoForceNoVerbose(t *testing.T) {
	chk := sztestlog.CaptureLog(t)
	defer chk.Release()

	const (
		tPath = tstpkgPaths
		tName = "README.md"
	)

	setupExpandGlobals(
		chk, expandGlobals{forceOverwrite: false, verboseLevel: 0},
	)
	chk.NoErr(setupExpandDirs(false, tPath, tName))

	chk.NoErr(expand.Process(templateName(chk, tPath, tName)))

	_, got, wnt, err := getExpandFiles(tPath, tName)
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Log()
}

func Test_ProcessExpand_MDPkg_NoTargetForceNoVerbose(t *testing.T) {
	chk := sztestlog.CaptureLog(t)
	defer chk.Release()

	const (
		tPath = tstpkgPaths
		tName = "README.md"
	)

	setupExpandGlobals(
		chk, expandGlobals{forceOverwrite: true, verboseLevel: 0},
	)
	chk.NoErr(setupExpandDirs(false, tPath, tName))

	chk.NoErr(expand.Process(templateName(chk, tPath, tName)))

	_, got, wnt, err := getExpandFiles(tPath, tName)
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Log()
}

func Test_ProcessExpand_MDPkg_NoTargetNoForceVerbose(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	const (
		tPath       = tstpkgPaths
		tName       = "README.md"
		sourceFName = ".README.gtm.md"
	)

	setupExpandGlobals(
		chk,
		expandGlobals{forceOverwrite: false, verboseLevel: 6},
	)
	chk.NoErr(setupExpandDirs(false, tPath, tName))

	chk.NoErr(expand.Process(templateName(chk, tPath, tName)))

	tFile, got, wnt, err := getExpandFiles(tPath, tName)
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Stdout(
		"Expanding "+tPath+sep+sourceFName+" to: "+tFile,
		"Loading package info for: .",
		"getInfo(\"package\")",
	)
}

func Test_ProcessExpand_MDPkg_NoTargetForceVerbose(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	const (
		tPath       = tstpkgPaths
		tName       = "README.md"
		sourceFName = ".README.gtm.md"
	)

	setupExpandGlobals(
		chk,
		expandGlobals{forceOverwrite: true, verboseLevel: 6},
	)
	chk.NoErr(setupExpandDirs(false, tPath, tName))

	chk.NoErr(expand.Process(templateName(chk, tPath, tName)))

	tFile, got, wnt, err := getExpandFiles(tPath, tName)
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Stdout(
		"Expanding "+tPath+sep+sourceFName+" to: "+tFile,
		"Loading package info for: .",
		"getInfo(\"package\")",
	)
}

func Test_ProcessExpand_MDPkg_CancelOverwriteTargetForceNoVerbose(
	t *testing.T,
) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	const (
		tPath = tstpkgPaths
		tName = "README.md"
	)

	setupExpandGlobals(
		chk, expandGlobals{forceOverwrite: true, verboseLevel: 0},
	)
	chk.NoErr(setupExpandDirs(true, tPath, tName))

	chk.SetStdinData("N\n")

	chk.NoErr(expand.Process(templateName(chk, tPath, tName)))

	_, got, wnt, err := getExpandFiles(tPath, tName)
	chk.NoErr(err)
	chk.StrSlice(got, wnt)
}

func Test_ProcessExpand_MDPkg_CancelOverwriteForceVerbose(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	const (
		tPath       = tstpkgPaths
		tName       = "README.md"
		sourceFName = ".README.gtm.md"
	)

	setupExpandGlobals(
		chk, expandGlobals{forceOverwrite: true, verboseLevel: 6},
	)
	chk.NoErr(setupExpandDirs(true, tPath, tName))

	chk.SetStdinData("N\n")

	chk.NoErr(expand.Process(templateName(chk, tPath, tName)))

	tFile, got, wnt, err := getExpandFiles(tPath, tName)
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Stdout(
		"Expanding "+tPath+sep+sourceFName+" to: "+tFile,
		"Loading package info for: .",
		"getInfo(\"package\")",
	)
}

func Test_ProcessExpand_MDPkg_OverwriteTargetForceNoVerbose(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	const (
		tPath = tstpkgPaths
		tName = "README.md"
	)

	setupExpandGlobals(
		chk, expandGlobals{forceOverwrite: true, verboseLevel: 0},
	)
	chk.NoErr(setupExpandDirs(true, tPath, tName))

	chk.SetStdinData("Y\n")

	chk.NoErr(expand.Process(templateName(chk, tPath, tName)))

	_, got, wnt, err := getExpandFiles(tPath, tName)
	chk.NoErr(err)
	chk.StrSlice(got, wnt)
}

func Test_ProcessExpand_MDPkg_OverwriteForceVerbose(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	const (
		tPath       = tstpkgPaths
		tName       = "README.md"
		sourceFName = ".README.gtm.md"
	)

	setupExpandGlobals(
		chk, expandGlobals{forceOverwrite: true, verboseLevel: 6},
	)
	chk.NoErr(setupExpandDirs(true, tPath, tName))

	chk.SetStdinData("Y\n")

	chk.NoErr(expand.Process(templateName(chk, tPath, tName)))

	wFile, got, wnt, err := getExpandFiles(tPath, tName)
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Stdout(
		"Expanding "+tPath+sep+sourceFName+" to: "+wFile,
		"Loading package info for: .",
		"getInfo(\"package\")",
	)
}

//

func Test_ProcessExpand_GOPkg_NoTargetNoForceNoVerbose(t *testing.T) {
	chk := sztestlog.CaptureLog(t)
	defer chk.Release()

	const (
		tPath = tstpkgPaths
		tName = "doc.go"
	)

	setupExpandGlobals(
		chk, expandGlobals{forceOverwrite: false, verboseLevel: 0},
	)
	chk.NoErr(setupExpandDirs(false, tPath, tName))

	chk.NoErr(expand.Process(templateName(chk, tPath, tName)))

	_, got, wnt, err := getExpandFiles(tPath, tName)
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Log()
}

func Test_ProcessExpand_GOPkg_NoTargetForceNoVerbose(t *testing.T) {
	chk := sztestlog.CaptureLog(t)
	defer chk.Release()

	const (
		tPath = tstpkgPaths
		tName = "doc.go"
	)

	setupExpandGlobals(
		chk, expandGlobals{forceOverwrite: true, verboseLevel: 0},
	)
	chk.NoErr(setupExpandDirs(false, tPath, tName))

	chk.NoErr(expand.Process(templateName(chk, tPath, tName)))

	_, got, wnt, err := getExpandFiles(tPath, tName)
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Log()
}

func Test_ProcessExpand_GOPkg_NoTargetNoForceVerbose(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	const (
		tPath       = tstpkgPaths
		tName       = "doc.go"
		sourceFName = ".doc.gtm.go"
	)

	setupExpandGlobals(
		chk,
		expandGlobals{forceOverwrite: false, verboseLevel: 8},
	)
	chk.NoErr(setupExpandDirs(false, tPath, tName))

	chk.NoErr(expand.Process(templateName(chk, tPath, tName)))

	tFile, got, wnt, err := getExpandFiles(tPath, tName)
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Stdout(
		"Expanding " + tPath + sep + sourceFName + " to: " + tFile,
	)
}

func Test_ProcessExpand_GOPkg_NoTargetForceVerbose(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	const (
		tPath       = tstpkgPaths
		tName       = "doc.go"
		sourceFName = ".doc.gtm.go"
	)

	setupExpandGlobals(
		chk,
		expandGlobals{forceOverwrite: true, verboseLevel: 6},
	)
	chk.NoErr(setupExpandDirs(false, tPath, tName))

	chk.NoErr(expand.Process(templateName(chk, tPath, tName)))

	tFile, got, wnt, err := getExpandFiles(tPath, tName)
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Stdout(
		"Expanding " + tPath + sep + sourceFName + " to: " + tFile,
	)
}

func Test_ProcessExpand_GOPkg_CancelOverwriteTargetForceNoVerbose(
	t *testing.T,
) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	const (
		tPath = tstpkgPaths
		tName = "doc.go"
	)

	setupExpandGlobals(
		chk, expandGlobals{forceOverwrite: true, verboseLevel: 0},
	)
	chk.NoErr(setupExpandDirs(true, tPath, tName))

	chk.SetStdinData("N\n")

	chk.NoErr(expand.Process(templateName(chk, tPath, tName)))

	_, got, wnt, err := getExpandFiles(tPath, tName)
	chk.NoErr(err)
	chk.StrSlice(got, wnt)
}

func Test_ProcessExpand_GOPkg_CancelOverwriteForceVerbose(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	const (
		tPath       = tstpkgPaths
		tName       = "doc.go"
		sourceFName = ".doc.gtm.go"
	)

	setupExpandGlobals(
		chk, expandGlobals{forceOverwrite: true, verboseLevel: 6},
	)
	chk.NoErr(setupExpandDirs(true, tPath, tName))

	chk.SetStdinData("N\n")

	chk.NoErr(expand.Process(templateName(chk, tPath, tName)))

	tFile, got, wnt, err := getExpandFiles(tPath, tName)
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Stdout(
		"Expanding " + tPath + sep + sourceFName + " to: " + tFile,
	)
}

func Test_ProcessExpand_GOPkg_OverwriteTargetForceNoVerbose(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	const (
		tPath = tstpkgPaths
		tName = "doc.go"
	)

	setupExpandGlobals(
		chk, expandGlobals{forceOverwrite: true, verboseLevel: 0},
	)
	chk.NoErr(setupExpandDirs(true, tPath, tName))

	chk.SetStdinData("Y\n")

	chk.NoErr(expand.Process(templateName(chk, tPath, tName)))

	_, got, wnt, err := getExpandFiles(tPath, tName)
	chk.NoErr(err)
	chk.StrSlice(got, wnt)
}

func Test_ProcessExpand_GOPkg_OverwriteForceVerbose(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	const (
		tPath       = tstpkgPaths
		tName       = "doc.go"
		sourceFName = ".doc.gtm.go"
	)

	setupExpandGlobals(
		chk, expandGlobals{forceOverwrite: true, verboseLevel: 6},
	)
	chk.NoErr(setupExpandDirs(true, tPath, tName))

	chk.SetStdinData("Y\n")

	chk.NoErr(expand.Process(templateName(chk, tPath, tName)))

	wFile, got, wnt, err := getExpandFiles(tPath, tName)
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Stdout(
		"Expanding " + tPath + sep + sourceFName + " to: " + wFile,
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
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//

func Test_ProcessExpand_MDCmd_NoTargetNoForceNoVerbose(t *testing.T) {
	chk := sztestlog.CaptureLog(t)
	defer chk.Release()

	const (
		tPath = tstcmdPaths
		tName = "README.md"
	)

	setupExpandGlobals(
		chk, expandGlobals{forceOverwrite: false, verboseLevel: 0},
	)
	chk.NoErr(setupExpandDirs(false, tPath, tName))

	chk.NoErr(expand.Process(templateName(chk, tPath, tName)))

	_, got, wnt, err := getExpandFiles(tPath, tName)
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Log()
}

func Test_ProcessExpand_MDCmd_NoTargetForceNoVerbose(t *testing.T) {
	chk := sztestlog.CaptureLog(t)
	defer chk.Release()

	const (
		tPath = tstcmdPaths
		tName = "README.md"
	)

	setupExpandGlobals(
		chk, expandGlobals{forceOverwrite: true, verboseLevel: 0},
	)
	chk.NoErr(setupExpandDirs(false, tPath, tName))

	chk.NoErr(expand.Process(templateName(chk, tPath, tName)))

	_, got, wnt, err := getExpandFiles(tPath, tName)
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Log()
}

func Test_ProcessExpand_MDCmd_NoTargetNoForceVerbose(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	const (
		tPath       = tstcmdPaths
		tName       = "README.md"
		sourceFName = ".README.gtm.md"
	)

	setupExpandGlobals(
		chk,
		expandGlobals{forceOverwrite: false, verboseLevel: 6},
	)
	chk.NoErr(setupExpandDirs(false, tPath, tName))

	chk.NoErr(expand.Process(templateName(chk, tPath, tName)))

	tFile, got, wnt, err := getExpandFiles(tPath, tName)
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Stdout(
		"Expanding " + tPath + sep + sourceFName + " to: " + tFile,
	)
}

func Test_ProcessExpand_MDCmd_NoTargetForceVerbose(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	const (
		tPath       = tstcmdPaths
		tName       = "README.md"
		sourceFName = ".README.gtm.md"
	)

	setupExpandGlobals(
		chk,
		expandGlobals{forceOverwrite: true, verboseLevel: 6},
	)
	chk.NoErr(setupExpandDirs(false, tPath, tName))

	chk.NoErr(expand.Process(templateName(chk, tPath, tName)))

	tFile, got, wnt, err := getExpandFiles(tPath, tName)
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Stdout(
		"Expanding " + tPath + sep + sourceFName + " to: " + tFile,
	)
}

func Test_ProcessExpand_MDCmd_CancelOverwriteTargetForceNoVerbose(
	t *testing.T,
) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	const (
		tPath = tstcmdPaths
		tName = "README.md"
	)

	setupExpandGlobals(
		chk, expandGlobals{forceOverwrite: true, verboseLevel: 0},
	)
	chk.NoErr(setupExpandDirs(true, tPath, tName))

	chk.SetStdinData("N\n")

	chk.NoErr(expand.Process(templateName(chk, tPath, tName)))

	_, got, wnt, err := getExpandFiles(tPath, tName)
	chk.NoErr(err)
	chk.StrSlice(got, wnt)
}

func Test_ProcessExpand_MDCmd_CancelOverwriteForceVerbose(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	const (
		tPath       = tstcmdPaths
		tName       = "README.md"
		sourceFName = ".README.gtm.md"
	)

	setupExpandGlobals(
		chk, expandGlobals{forceOverwrite: true, verboseLevel: 6},
	)
	chk.NoErr(setupExpandDirs(true, tPath, tName))

	chk.SetStdinData("N\n")

	chk.NoErr(expand.Process(templateName(chk, tPath, tName)))

	tFile, got, wnt, err := getExpandFiles(tPath, tName)
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Stdout(
		"Expanding " + tPath + sep + sourceFName + " to: " + tFile,
	)
}

func Test_ProcessExpand_MDCmd_OverwriteTargetForceNoVerbose(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	const (
		tPath = tstcmdPaths
		tName = "README.md"
	)

	setupExpandGlobals(
		chk, expandGlobals{forceOverwrite: true, verboseLevel: 0},
	)
	chk.NoErr(setupExpandDirs(true, tPath, tName))

	chk.SetStdinData("Y\n")

	chk.NoErr(expand.Process(templateName(chk, tPath, tName)))

	_, got, wnt, err := getExpandFiles(tPath, tName)
	chk.NoErr(err)
	chk.StrSlice(got, wnt)
}

func Test_ProcessExpand_MDCmd_OverwriteForceVerbose(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	const (
		tPath       = tstcmdPaths
		tName       = "README.md"
		sourceFName = ".README.gtm.md"
	)

	setupExpandGlobals(
		chk, expandGlobals{forceOverwrite: true, verboseLevel: 6},
	)
	chk.NoErr(setupExpandDirs(true, tPath, tName))

	chk.SetStdinData("Y\n")

	chk.NoErr(expand.Process(templateName(chk, tPath, tName)))

	wFile, got, wnt, err := getExpandFiles(tPath, tName)
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Stdout(
		"Expanding " + tPath + sep + sourceFName + " to: " + wFile,
	)
}

//

func Test_ProcessExpand_GOCmd_NoTargetNoForceNoVerbose(t *testing.T) {
	chk := sztestlog.CaptureLog(t)
	defer chk.Release()

	const (
		tPath = tstcmdPaths
		tName = "doc.go"
	)

	setupExpandGlobals(
		chk, expandGlobals{forceOverwrite: false, verboseLevel: 0},
	)
	chk.NoErr(setupExpandDirs(false, tPath, tName))

	chk.NoErr(expand.Process(templateName(chk, tPath, tName)))

	_, got, wnt, err := getExpandFiles(tPath, tName)
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Log()
}

func Test_ProcessExpand_GOCmd_NoTargetForceNoVerbose(t *testing.T) {
	chk := sztestlog.CaptureLog(t)
	defer chk.Release()

	const (
		tPath = tstcmdPaths
		tName = "doc.go"
	)

	setupExpandGlobals(
		chk, expandGlobals{forceOverwrite: true, verboseLevel: 0},
	)
	chk.NoErr(setupExpandDirs(false, tPath, tName))

	chk.NoErr(expand.Process(templateName(chk, tPath, tName)))

	_, got, wnt, err := getExpandFiles(tPath, tName)
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Log()
}

func Test_ProcessExpand_GOCmd_NoTargetNoForceVerbose(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	const (
		tPath       = tstcmdPaths
		tName       = "doc.go"
		sourceFName = ".doc.gtm.go"
	)

	setupExpandGlobals(
		chk,
		expandGlobals{forceOverwrite: false, verboseLevel: 8},
	)
	chk.NoErr(setupExpandDirs(false, tPath, tName))

	chk.NoErr(expand.Process(templateName(chk, tPath, tName)))

	tFile, got, wnt, err := getExpandFiles(tPath, tName)
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Stdout(
		"Expanding " + tPath + sep + sourceFName + " to: " + tFile,
	)
}

func Test_ProcessExpand_GOCmd_NoTargetForceVerbose(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	const (
		tPath       = tstcmdPaths
		tName       = "doc.go"
		sourceFName = ".doc.gtm.go"
	)

	setupExpandGlobals(
		chk,
		expandGlobals{forceOverwrite: true, verboseLevel: 6},
	)
	chk.NoErr(setupExpandDirs(false, tPath, tName))

	chk.NoErr(expand.Process(templateName(chk, tPath, tName)))

	tFile, got, wnt, err := getExpandFiles(tPath, tName)
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Stdout(
		"Expanding " + tPath + sep + sourceFName + " to: " + tFile,
	)
}

func Test_ProcessExpand_GOCmd_CancelOverwriteTargetForceNoVerbose(
	t *testing.T,
) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	const (
		tPath = tstcmdPaths
		tName = "doc.go"
	)

	setupExpandGlobals(
		chk, expandGlobals{forceOverwrite: true, verboseLevel: 0},
	)
	chk.NoErr(setupExpandDirs(true, tPath, tName))

	chk.SetStdinData("N\n")

	chk.NoErr(expand.Process(templateName(chk, tPath, tName)))

	_, got, wnt, err := getExpandFiles(tPath, tName)
	chk.NoErr(err)
	chk.StrSlice(got, wnt)
}

func Test_ProcessExpand_GOCmd_CancelOverwriteForceVerbose(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	const (
		tPath       = tstcmdPaths
		tName       = "doc.go"
		sourceFName = ".doc.gtm.go"
	)

	setupExpandGlobals(
		chk, expandGlobals{forceOverwrite: true, verboseLevel: 6},
	)
	chk.NoErr(setupExpandDirs(true, tPath, tName))

	chk.SetStdinData("N\n")

	chk.NoErr(expand.Process(templateName(chk, tPath, tName)))

	tFile, got, wnt, err := getExpandFiles(tPath, tName)
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Stdout(
		"Expanding " + tPath + sep + sourceFName + " to: " + tFile,
	)
}

func Test_ProcessExpand_GOCmd_OverwriteTargetForceNoVerbose(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	const (
		tPath = tstcmdPaths
		tName = "doc.go"
	)

	setupExpandGlobals(
		chk, expandGlobals{forceOverwrite: true, verboseLevel: 0},
	)
	chk.NoErr(setupExpandDirs(true, tPath, tName))

	chk.SetStdinData("Y\n")

	chk.NoErr(expand.Process(templateName(chk, tPath, tName)))

	_, got, wnt, err := getExpandFiles(tPath, tName)
	chk.NoErr(err)
	chk.StrSlice(got, wnt)
}

func Test_ProcessExpand_GOCmd_OverwriteForceVerbose(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	const (
		tPath       = tstcmdPaths
		tName       = "doc.go"
		sourceFName = ".doc.gtm.go"
	)

	setupExpandGlobals(
		chk, expandGlobals{forceOverwrite: true, verboseLevel: 6},
	)
	chk.NoErr(setupExpandDirs(true, tPath, tName))

	chk.SetStdinData("Y\n")

	chk.NoErr(expand.Process(templateName(chk, tPath, tName)))

	wFile, got, wnt, err := getExpandFiles(tPath, tName)
	chk.NoErr(err)
	chk.StrSlice(got, wnt)

	chk.Stdout(
		"Expanding " + tPath + sep + sourceFName + " to: " + wFile,
	)
}
