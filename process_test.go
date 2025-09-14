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

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/dancsecs/szlog"
	"github.com/dancsecs/sztest"
	"github.com/dancsecs/sztestlog"
)

func setupTest(
	chk *sztest.Chk,
	tForceOverwrite bool,
	tVerbose szlog.VerboseLevel,
) {
	chk.T().Helper()

	origOutputDir := outputDir
	origCWD, err := os.Getwd()
	origForceOverwrite := forceOverwrite
	origVerboseLevel := szlog.Verbose()

	forceOverwrite = tForceOverwrite

	szlog.SetVerbose(tVerbose)

	if chk.NoErr(err) {
		outputDir = chk.CreateTmpDir()
		chk.PushPostReleaseFunc(func() error {
			outputDir = origOutputDir
			forceOverwrite = origForceOverwrite

			szlog.SetVerbose(origVerboseLevel)

			return os.Chdir(origCWD)
		})
	}
}

func Test_Process_ConfirmOverwrite(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	data := "The data."

	dir := chk.CreateTmpDir()

	ok, err := confirmOverwrite(filepath.Join(dir, "noFile.dat"), data)
	chk.NoErr(err)
	chk.True(ok)

	fPath := chk.CreateTmpFile([]byte(data))
	ok, err = confirmOverwrite(fPath, data)
	chk.NoErr(err)
	chk.False(ok)

	chk.Stdout(
		"No change: " + fPath,
	)
}

func Test_Process_AskToOverwrite(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	chk.SetStdinData("invalid\nr\nN\n")

	overwrite, err := askToOverwrite("file", "abc\n", "def")
	chk.NoErr(err)
	chk.False(overwrite)

	chk.Stdout(
		fmt.Sprintf(confirmMsg, "file")+
			fmt.Sprintf(confirmUnknown[:len(confirmUnknown)-1], "invalid"),
		fmt.Sprintf(confirmMsg, "file")+"--- Old_file",
		"+++ New_file",
		"@@ -1 +1 @@",
		"-abc",
		"+def",
		"",
		fmt.Sprintf(confirmMsg, "file")+
			confirmCancelled[:len(confirmCancelled)-1],
	)
}
