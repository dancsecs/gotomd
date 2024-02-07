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
	"testing"

	"github.com/dancsecs/sztest"
)

func Test_CmdParse_ParseCmd(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	cmd := ""

	dir, action, err := parseCmd(cmd)
	chk.Err(err, "relative directory must be specified in cmd: \""+cmd+"\"")
	chk.Str(dir, "")
	chk.Str(action, "")

	cmd = string(os.PathSeparator) + "action"

	dir, action, err = parseCmd(cmd)
	chk.Err(err, "relative directory must be specified in cmd: \""+cmd+"\"")
	chk.Str(dir, "")
	chk.Str(action, "")

	cmd = "." + string(os.PathSeparator)

	dir, action, err = parseCmd(cmd)
	chk.Err(
		err,
		"invalid action: a non-blank action is required",
	)
	chk.Str(dir, "")
	chk.Str(action, "")

	cmd = sampleGoProjectOne + string(os.PathSeparator) + "action"

	dir, action, err = parseCmd(cmd)
	chk.Err(
		err,
		"relative directory must be specified in cmd: \""+cmd+"\"",
	)
	chk.Str(dir, "")
	chk.Str(action, "")

	cmd = sampleGoProjectOnePath + "action"
	dir, action, err = parseCmd(cmd)
	chk.NoErr(err)
	chk.Str(dir, "."+string(os.PathSeparator)+sampleGoProjectOne)
	chk.Str(action, "action")
}

func Test_CmdParse_ParseCmds1(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	cmd := ""

	dirs, actions, err := parseCmds(cmd)
	chk.Nil(dirs)
	chk.Nil(actions)
	chk.Err(err, "relative directory must be specified in cmd: \""+cmd+"\"")

	cmd = string(os.PathSeparator) + "action"

	dirs, actions, err = parseCmds(cmd)
	chk.Nil(dirs)
	chk.Nil(actions)
	chk.Err(err, "relative directory must be specified in cmd: \""+cmd+"\"")

	cmd = "." + string(os.PathSeparator)

	dirs, actions, err = parseCmds(cmd)
	chk.Nil(dirs)
	chk.Nil(actions)
	chk.Err(err, "invalid action: a non-blank action is required")

	cmd = sampleGoProjectOne + string(os.PathSeparator) + "action"

	dirs, actions, err = parseCmds(cmd)
	chk.Nil(dirs)
	chk.Nil(actions)
	chk.Err(
		err,
		"relative directory must be specified in cmd: \""+cmd+"\"",
	)

	cmd = sampleGoProjectOnePath + "action"
	expDir := "." + string(os.PathSeparator) + sampleGoProjectOne

	dirs, actions, err = parseCmds(cmd)
	chk.NoErr(err)
	chk.StrSlice(dirs, []string{expDir})
	chk.StrSlice(actions, []string{"action"})
}

func Test_CmdParse_ParseCmds2(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	file1 := sampleGoProjectOnePath + "action"
	file2 := string(os.PathSeparator) + "action2"

	dirs, actions, err := parseCmds(file1 + " " + file2)
	chk.Nil(dirs)
	chk.Nil(actions)
	chk.Err(err, "relative directory must be specified in cmd: \""+file2+"\"")

	file2 = sampleGoProjectOne + string(os.PathSeparator) + "action"
	dirs, actions, err = parseCmds(file1 + " " + file2)
	chk.Nil(dirs)
	chk.Nil(actions)
	chk.Err(err, "relative directory must be specified in cmd: \""+file2+"\"")

	file2 = "action2"
	expDir := "." + string(os.PathSeparator) + sampleGoProjectOne

	dirs, actions, err = parseCmds(file1 + " " + file2)
	chk.NoErr(err)
	chk.StrSlice(dirs, []string{expDir, expDir})
	chk.StrSlice(actions, []string{"action", "action2"})

	file2 = sampleGoProjectOnePath + "action2"
	dirs, actions, err = parseCmds(file1 + " " + file2)
	chk.NoErr(err)
	chk.StrSlice(dirs, []string{expDir, expDir})
	chk.StrSlice(actions, []string{"action", "action2"})
}
