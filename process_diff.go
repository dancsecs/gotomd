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
	"github.com/dancsecs/szlog"
	"github.com/hexops/gotextdiff"
	"github.com/hexops/gotextdiff/myers"
	"github.com/hexops/gotextdiff/span"
)

func diffFile(fileName string, oldData, newData string) {
	// Create an edit script using the Myers diff algorithm
	edits := myers.ComputeEdits(span.URI(fileName), oldData, newData)

	// Format the edits into a standard unified diff
	unifiedDiff := gotextdiff.ToUnified(
		"Old_"+fileName, "New_"+fileName,
		oldData,
		edits,
	)

	szlog.Say0(unifiedDiff, "\n")
}
