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

package cmds_test

import (
	"os"
	"testing"

	"github.com/dancsecs/gotomd/internal/cmds"
	"github.com/dancsecs/gotomd/internal/errs"
	"github.com/dancsecs/sztestlog"
)

const (
	sep = string(os.PathSeparator)

	example1     = "example1"
	example1Path = "." + sep + "examples" + sep + example1

	example2     = "example2"
	example2Path = "." + sep + "examples" + sep + example2
)

func Test_CmdParse_ParseCmd_InvalidDir(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	cmd := "." + sep + "INVALID_DIR" + sep + "action"

	dir, action, err := cmds.ParseCmd(cmd)
	chk.Err(
		err,
		chk.ErrChain(
			errs.ErrInvalidDirectory,
			"\"."+sep+"INVALID_DIR\"",
		),
	)
	chk.Str(dir, "")
	chk.Str(action, "")
}

func Test_CmdParse_ParseCmd(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	tstDir := chk.CreateTmpDir()
	_ = chk.CreateTmpSubDir("examples", "example1")

	t.Chdir(tstDir)

	cmd := ""

	dir, action, err := cmds.ParseCmd(cmd)
	chk.Err(
		err,
		errs.ErrInvalidRelativeDir.Error()+": \""+cmd+"\"",
	)
	chk.Str(dir, "")
	chk.Str(action, "")

	cmd = sep + "action"

	dir, action, err = cmds.ParseCmd(cmd)
	chk.Err(
		err,
		errs.ErrInvalidRelativeDir.Error()+": \""+cmd+"\"",
	)
	chk.Str(dir, "")
	chk.Str(action, "")

	cmd = "." + sep

	dir, action, err = cmds.ParseCmd(cmd)
	chk.Err(
		err,
		errs.ErrMissingAction.Error(),
	)
	chk.Str(dir, "")
	chk.Str(action, "")

	cmd = "examples" + sep + example1 + sep + "action"

	dir, action, err = cmds.ParseCmd(cmd)
	chk.Err(
		err,
		errs.ErrInvalidRelativeDir.Error()+": \""+cmd+"\"",
	)
	chk.Str(dir, "")
	chk.Str(action, "")

	cmd = example1Path + sep + "action"
	dir, action, err = cmds.ParseCmd(cmd)
	chk.NoErr(err)
	chk.Str(
		dir,
		example1Path,
	)
	chk.Str(action, "action")
}

func Test_CmdParse_ParseCmds1(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	tstDir := chk.CreateTmpDir()
	_ = chk.CreateTmpSubDir("examples", "example1")

	t.Chdir(tstDir)

	cmd := ""

	dirs, actions, err := cmds.ParseCmds(cmd)
	chk.Nil(dirs)
	chk.Nil(actions)
	chk.Err(
		err,
		errs.ErrInvalidRelativeDir.Error()+": \""+cmd+"\"",
	)

	cmd = sep + "action"

	dirs, actions, err = cmds.ParseCmds(cmd)
	chk.Nil(dirs)
	chk.Nil(actions)
	chk.Err(
		err,
		errs.ErrInvalidRelativeDir.Error()+": \""+cmd+"\"",
	)

	cmd = "." + sep

	dirs, actions, err = cmds.ParseCmds(cmd)
	chk.Nil(dirs)
	chk.Nil(actions)
	chk.Err(err, errs.ErrMissingAction.Error())

	cmd = "examples" + sep + example1 + sep + "action"

	dirs, actions, err = cmds.ParseCmds(cmd)
	chk.Nil(dirs)
	chk.Nil(actions)
	chk.Err(
		err,
		errs.ErrInvalidRelativeDir.Error()+": \""+cmd+"\"",
	)

	cmd = example1Path + sep + "action"

	dirs, actions, err = cmds.ParseCmds(cmd)
	chk.NoErr(err)
	chk.StrSlice(dirs, []string{example1Path})
	chk.StrSlice(actions, []string{"action"})
}

func Test_CmdParse_ParseCmds2(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	tstDir := chk.CreateTmpDir()
	_ = chk.CreateTmpSubDir("examples", "example1")
	_ = chk.CreateTmpSubDir("examples", "example2")

	t.Chdir(tstDir)

	file1 := example1Path + sep + "action"
	file2 := sep + "action2"

	dirs, actions, err := cmds.ParseCmds(file1 + " " + file2)
	chk.Nil(dirs)
	chk.Nil(actions)
	chk.Err(
		err,
		errs.ErrInvalidRelativeDir.Error()+": \""+file2+"\"",
	)

	file2 = "examples" + sep + example1 + sep + "action"
	dirs, actions, err = cmds.ParseCmds(file1 + " " + file2)
	chk.Nil(dirs)
	chk.Nil(actions)
	chk.Err(
		err,
		errs.ErrInvalidRelativeDir.Error()+": \""+file2+"\"",
	)

	file2 = "action2"

	dirs, actions, err = cmds.ParseCmds(file1 + " " + file2)
	chk.NoErr(err)
	chk.StrSlice(dirs, []string{example1Path, example1Path})
	chk.StrSlice(actions, []string{"action", "action2"})

	file2 = example2Path + sep + "action2"
	dirs, actions, err = cmds.ParseCmds(file1 + " " + file2)
	chk.NoErr(err)
	chk.StrSlice(dirs, []string{example1Path, example2Path})
	chk.StrSlice(actions, []string{"action", "action2"})
}
