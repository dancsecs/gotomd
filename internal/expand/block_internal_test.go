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

package expand

import (
	"testing"

	"github.com/dancsecs/gotomd/internal/errs"
	"github.com/dancsecs/sztestlog"
)

func Test_ExpandGetBlockInitialNotTerminated(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	lines := []string{
		"<!--- gotomd::cmd::./path/cmd arg1 arg2 ->",
	}

	i, res, err := getBlock(
		1,
		len("<!--- gotomd::cmd::"),
		lines, true, "-->", " -<>", " ",
	)

	chk.Int(i, 1)
	chk.Err(err, errs.ErrBlockNotTerminated.Error())
	chk.Str(res, "")
}

func Test_ExpandGetBlockNotTerminated(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	var (
		res   string
		lines = []string{
			"<!--- gotomd::cmd::./path/cmd arg1 arg2 ->",
		}
	)

	i, res, err := getBlock(
		0,
		len("<!--- gotomd::cmd::"),
		lines, true, "-->", " -<>", " ",
	)

	chk.Int(i, 1)
	chk.Err(err, errs.ErrBlockNotTerminated.Error())
	chk.Str(res, "./path/cmd arg1 arg2 ->")
}

func Test_ExpandGetBlockOneLine(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	var (
		res   string
		lines = []string{
			"<!--- gotomd::cmd::./path/cmd arg1 arg2 -->",
		}
	)

	i, res, err := getBlock(
		0,
		len("<!--- gotomd::cmd::"),
		lines, true, "-->", " -<>", " ",
	)

	chk.Int(i, 0)
	chk.NoErr(err)
	chk.Str(res, "./path/cmd arg1 arg2")
}

func Test_ExpandGetBlockTwoLines1(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	var (
		res   string
		lines = []string{
			"<!--- gotomd::cmd::./path/cmd arg1 arg2",
			"-->",
		}
	)

	i, res, err := getBlock(
		0,
		len("<!--- gotomd::cmd::"),
		lines, true, "-->", " -<>", " ",
	)

	chk.Int(i, 1)
	chk.NoErr(err)
	chk.Str(res, "./path/cmd arg1 arg2")
}

func Test_ExpandGetBlockTwoLines2(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	var (
		res   string
		lines = []string{
			"<!--- gotomd::cmd::./path/cmd arg1",
			" arg2 -->",
		}
	)

	i, res, err := getBlock(
		0,
		len("<!--- gotomd::cmd::"),
		lines, true, "-->", " -<>", " ",
	)

	chk.Int(i, 1)
	chk.NoErr(err)
	chk.Str(res, "./path/cmd arg1 arg2")
}

func Test_ExpandGetBlockTwoLines3(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	var (
		res   string
		lines = []string{
			"<!--- gotomd::cmd::./path/cmd",
			" arg1",
			"arg2 -->",
		}
	)

	i, res, err := getBlock(
		0,
		len("<!--- gotomd::cmd::"),
		lines, true, "-->", " -<>", " ",
	)

	chk.Int(i, 2)
	chk.NoErr(err)
	chk.Str(res, "./path/cmd arg1 arg2")
}

func Test_ExpandGetBlockTwoLines4(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	var (
		res   string
		lines = []string{
			"<!--- gotomd::cmd::./path/cmd",
			" arg1",
			"arg2 ",
			"-->",
		}
	)

	i, res, err := getBlock(
		0,
		len("<!--- gotomd::cmd::"),
		lines, true, "-->", " -<>", " ",
	)

	chk.Int(i, 3)
	chk.NoErr(err)
	chk.Str(res, "./path/cmd arg1 arg2")
}
