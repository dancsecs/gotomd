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

package ansi

import (
	"testing"

	"github.com/dancsecs/sztestlog"
)

//nolint:goCheckNoGlobals // Ok.
var testUnsupported = []string{
	// Cursor movements
	"Remove this ->\x1b[H<- code but not this \x1b[0m.",
	"Remove this ->\x1b[{1};{22}H<- code but not this \x1b[0m.",
	"Remove this ->\x1b[{31};{12}f<- code but not this \x1b[0m.",
	"Remove this ->\x1b[1A<- code but not this \x1b[0m.",
	"Remove this ->\x1b[2A<- code but not this \x1b[0m.",
	"Remove this ->\x1b[3A<- code but not this \x1b[0m.",
	"Remove this ->\x1b[30A<- code but not this \x1b[0m.",
	"Remove this ->\x1b[4B<- code but not this \x1b[0m.",
	"Remove this ->\x1b[41B<- code but not this \x1b[0m.",
	"Remove this ->\x1b[5C<- code but not this \x1b[0m.",
	"Remove this ->\x1b[55C<- code but not this \x1b[0m.",
	"Remove this ->\x1b[6D<- code but not this \x1b[0m.",
	"Remove this ->\x1b[66D<- code but not this \x1b[0m.",
	"Remove this ->\x1b[7E<- code but not this \x1b[0m.",
	"Remove this ->\x1b[78E<- code but not this \x1b[0m.",
	"Remove this ->\x1b[8F<- code but not this \x1b[0m.",
	"Remove this ->\x1b[89F<- code but not this \x1b[0m.",
	"Remove this ->\x1b[9G<- code but not this \x1b[0m.",
	"Remove this ->\x1b[23G<- code but not this \x1b[0m.",
	"Remove this ->\x1b[6n<- code but not this \x1b[0m.",
	"Remove this ->\x1b M<- code but not this \x1b[0m.",
	"Remove this ->\x1b 7<- code but not this \x1b[0m.",
	"Remove this ->\x1b 8<- code but not this \x1b[0m.",
	"Remove this ->\x1b[s<- code but not this \x1b[0m.",
	"Remove this ->\x1b[u<- code but not this \x1b[0m.",

	// Erase Ansi codes.
	"Remove this ->\x1b[J<- code but not this \x1b[0m.",
	"Remove this ->\x1b[0J<- code but not this \x1b[0m.",
	"Remove this ->\x1b[1J<- code but not this \x1b[0m.",
	"Remove this ->\x1b[2J<- code but not this \x1b[0m.",
	"Remove this ->\x1b[3J<- code but not this \x1b[0m.",
	"Remove this ->\x1b[K<- code but not this \x1b[0m.",
	"Remove this ->\x1b[0K<- code but not this \x1b[0m.",
	"Remove this ->\x1b[1K<- code but not this \x1b[0m.",
	"Remove this ->\x1b[2K<- code but not this \x1b[0m.",

	// Color codes not supported.
	"Remove this ->\x1b[38;5;{blue}m<- code but not this \x1b[0m.",
	"Remove this ->\x1b[48;5;{green}m<- code but not this \x1b[0m.",
	"Remove this ->\x1b[38;2;{1};{12};{123}m<- code but not this \x1b[0m.",
	"Remove this ->\x1b[48;2;{4};{45};{255}m<- code but not this \x1b[0m.",

	// Screen Modes
	"Remove this ->\x1b[={12}h<- code but not this \x1b[0m.",
	"Remove this ->\x1b[=0h<- code but not this \x1b[0m.",
	"Remove this ->\x1b[=1h<- code but not this \x1b[0m.",
	"Remove this ->\x1b[=2h<- code but not this \x1b[0m.",
	"Remove this ->\x1b[=3h<- code but not this \x1b[0m.",
	"Remove this ->\x1b[=4h<- code but not this \x1b[0m.",
	"Remove this ->\x1b[=5h<- code but not this \x1b[0m.",
	"Remove this ->\x1b[=6h<- code but not this \x1b[0m.",
	"Remove this ->\x1b[=7h<- code but not this \x1b[0m.",
	"Remove this ->\x1b[=13h<- code but not this \x1b[0m.",
	"Remove this ->\x1b[=14h<- code but not this \x1b[0m.",
	"Remove this ->\x1b[=15h<- code but not this \x1b[0m.",
	"Remove this ->\x1b[=16h<- code but not this \x1b[0m.",
	"Remove this ->\x1b[=17h<- code but not this \x1b[0m.",
	"Remove this ->\x1b[=18h<- code but not this \x1b[0m.",
	"Remove this ->\x1b[=19h<- code but not this \x1b[0m.",
	"Remove this ->\x1b[={19}l<- code but not this \x1b[0m.",

	// Common Private Modes
	// -----------------------------------------------------------------
	"Remove this ->\x1b[?25l<- code but not this \x1b[0m.",
	"Remove this ->\x1b[?25h<- code but not this \x1b[0m.",
	"Remove this ->\x1b[?47l<- code but not this \x1b[0m.",
	"Remove this ->\x1b[?47h<- code but not this \x1b[0m.",
	"Remove this ->\x1b[?1049l<- code but not this \x1b[0m.",
	"Remove this ->\x1b[?1049h<- code but not this \x1b[0m.",
}

const cleanStr = "Remove this -><- code but not this \x1b[0m."

func TestAnsi_UnsupportedCodes(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	for i, line := range testUnsupported {
		chk.Str(
			removeUnsupported(line),
			cleanStr,
			"Failed idx: ", i,
		)
	}
}
