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

package sample1_test

import (
	"testing"

	"github.com/dancsecs/gotomd/internal/gopkg/sample1"
	"github.com/dancsecs/sztestlog"
)

func Test_StructureType(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	structureType := &sample1.StructureType{
		F1: "F1-",
		F2: 4,
	}

	chk.Str(
		structureType.GetF1(1, 2, 3),
		"F1-6",
	)

	chk.Int(
		sample1.TimesTwo(2),
		4,
	)

	chk.Int(
		sample1.TimesThree(3),
		9,
	)
}
