/*
   Golang To Github Markdown Utility: gotomd
   Copyright (C) 2025 Leslie Dancsecs

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

package update_test

import (
	"path/filepath"
	"testing"

	"github.com/dancsecs/gotomd/internal/update"
	"github.com/dancsecs/sztestlog"
)

const perm = 0o0600

func Test_Process_Invalid(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	dir := chk.CreateTmpDir()

	result, err := update.File(dir, false, "", perm)

	chk.Err(
		err,
		chk.ErrChain(
			update.ErrFileUpdate,
			"'"+dir+"'",
			update.ErrInvalidFileType,
		),
		``)

	chk.Int(int(result), int(update.Failed))

	chk.Stdout()
}

func Test_Process_Create(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	dir := chk.CreateTmpDir()

	file := filepath.Join(dir, "fileToCreate.md")

	result, err := update.File(file, false, "", perm)

	chk.NoErr(err)

	chk.Int(int(result), int(update.Created))

	chk.Stdout()
}

func Test_Process_Unchanged(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	file := chk.CreateTmpFile([]byte("abc"))

	result, err := update.File(file, true, "abc", perm)

	chk.NoErr(err)

	chk.Int(int(result), int(update.Unchanged))

	chk.Stdout(
		"No change: " + file,
	)
}

func Test_Process_Forced_Update(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	file := chk.CreateTmpFile([]byte("abc"))

	result, err := update.File(file, true, "def", perm)

	chk.NoErr(err)

	chk.Int(int(result), int(update.Updated))

	chk.Stdout()
}

func Test_Process_Confirmed_Update(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	file := chk.CreateTmpFile([]byte("abc"))

	chk.SetStdinData("Y\n")

	result, err := update.File(file, false, "def", perm)

	chk.NoErr(err)

	chk.Int(int(result), int(update.Updated))

	chk.Stdout(
		"Confirm overwrite of: "+file,
		"(Y to overwrite; [N/n] to Cancel; [R/r] to review)? ",
	)
}

func Test_Process_Cancelled_Update(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	file := chk.CreateTmpFile([]byte("abc"))

	chk.SetStdinData("n\n")

	result, err := update.File(file, false, "def", perm)

	chk.NoErr(err)

	chk.Int(int(result), int(update.Cancelled))

	chk.Stdout(
		"Confirm overwrite of: "+file,
		"(Y to overwrite; [N/n] to Cancel; [R/r] to review)? "+
			"Overwrite cancelled",
		"",
	)
}

// func setupTest(
// 	chk *sztest.Chk,
// 	tForceOverwrite bool,
// 	tVerbose szlog.VerboseLevel,
// ) {
// 	chk.T().Helper()

// 	origOutputDir := outputDir
// 	origCWD, err := os.Getwd()
// 	origForceOverwrite := forceOverwrite
// 	origVerboseLevel := szlog.Verbose()

// 	forceOverwrite = tForceOverwrite

// 	szlog.SetVerbose(tVerbose)

// 	if chk.NoErr(err) {
// 		outputDir = chk.CreateTmpDir()
// 		chk.PushPostReleaseFunc(func() error {
// 			outputDir = origOutputDir
// 			forceOverwrite = origForceOverwrite

// 			szlog.SetVerbose(origVerboseLevel)

// 			return os.Chdir(origCWD)
// 		})
// 	}
// }

// func Test_Process_AskToOverwrite(t *testing.T) {
// 	chk := sztestlog.CaptureStdout(t)
// 	defer chk.Release()

// 	chk.SetStdinData("invalid\nr\nN\n")

// 	overwrite, err := askToOverwrite("file", "abc\n", "def")
// 	chk.NoErr(err)
// 	chk.False(overwrite)

// 	chk.Stdout(
// 		fmt.Sprintf(confirmMsg, "file")+
// 			fmt.Sprintf(confirmUnknown[:len(confirmUnknown)-1], "invalid"),
// 		fmt.Sprintf(confirmMsg, "file")+"--- Old_file",
// 		"+++ New_file",
// 		"@@ -1 +1 @@",
// 		"-abc",
// 		"+def",
// 		"",
// 		fmt.Sprintf(confirmMsg, "file")+
// 			confirmCancelled[:len(confirmCancelled)-1],
// 	)
// }
