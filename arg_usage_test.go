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
	"testing"

	"github.com/dancsecs/sztest"
)

func Test_ArgUsage_SampleNoArgsDefaultsToCWD(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	dir := chk.CreateTmpDir()

	chk.SetArgs(
		"programName",
		"-o",
		dir,
	)

	main()
}

func Test_ArgUsage_SampleInvalidFile(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	fPath := chk.CreateTmpFile(nil)
	chk.SetArgs(
		"programName",
		fPath,
	)

	chk.Panic(
		main,
		ErrUnexpectedExtension.Error()+": expected - .gtm.md",
	)
}

func Test_ArgUsage_InvalidDefaultPermissions(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	fPath := chk.CreateTmpFile(nil)
	chk.SetArgs(
		"programName",
		"-p", "0744",
		fPath,
	)

	chk.AddSub(`(?s)\n.*$`, "\\nUsage Information")
	chk.Panic(
		main,
		ErrInvalidDefPerm.Error()+"\\nUsage Information",
	)
}

func Test_ArgUsage_InvalidCleanAndReplace(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	fPath := chk.CreateTmpFile(nil)
	chk.SetArgs(
		"programName",
		"-r",
		"-c",
		fPath,
	)

	chk.AddSub(`(?s)\n.*$`, "\\nUsage Information")
	chk.Panic(
		main,
		ErrInvalidOptionsRC.Error()+"\\nUsage Information",
	)
}

func Test_ArgUsage_InvalidOutDirectory(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	fPath := chk.CreateTmpFile(nil)
	chk.SetArgs(
		"programName",
		"-o", "DIRECTORY_DOES_NOT_EXIST",
		fPath,
	)

	chk.AddSub(`(?s)\n.*$`, "\\nUsage Information")
	chk.Panic(
		main,
		ErrInvalidOutputDir.Error()+": "+
			"'DIRECTORY_DOES_NOT_EXIST'\\nUsage Information",
	)
}
