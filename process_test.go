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
	"testing"

	"github.com/dancsecs/szlog"
	"github.com/dancsecs/sztest"
	"github.com/dancsecs/sztestlog"
)

func setupTest(
	chk *sztest.Chk,
	tCleanOnly, tReplace, tForceOverwrite bool,
	tVerbose szlog.LogLevel,
) {
	chk.T().Helper()

	origOutputDir := outputDir
	origCWD, err := os.Getwd()
	origCleanOnly := cleanOnly
	origReplace := replace
	origForceOverwrite := forceOverwrite
	origLogLevel := szlog.Level()

	cleanOnly = tCleanOnly
	replace = tReplace
	forceOverwrite = tForceOverwrite

	szlog.SetLevel(tVerbose)

	if chk.NoErr(err) {
		outputDir = chk.CreateTmpDir()
		chk.PushPostReleaseFunc(func() error {
			outputDir = origOutputDir
			cleanOnly = origCleanOnly
			replace = origReplace
			forceOverwrite = origForceOverwrite

			szlog.SetLevel(origLogLevel)

			return os.Chdir(origCWD)
		})
	}
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

//  func Test_Process_Stop(t *testing.T) {
//  	t.Fatal("STOPPING")
//  }

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
