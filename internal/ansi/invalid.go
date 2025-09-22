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

import "regexp"

// GitHub LaTeX unsupported ansi escape codes.

//nolint:goCheckNoGlobals // Ok.
var reUnsupported = []*regexp.Regexp{
	// Cursor movements
	// -----------------------------------------------------------------
	// ESC[H 	moves cursor to home position (0, 0)
	// ESC[{line};{column}H
	// ESC[{line};{column}f 	moves cursor to line #, column #
	// ESC[#A 	moves cursor up # lines
	// ESC[#B 	moves cursor down # lines
	// ESC[#C 	moves cursor right # columns
	// ESC[#D 	moves cursor left # columns
	// ESC[#E 	moves cursor to beginning of next line, # lines down
	// ESC[#F 	moves cursor to beginning of previous line, # lines up
	// ESC[#G 	moves cursor to column #
	// ESC[6n 	request cursor position (reports as ESC[#;#R)
	// ESC M 	moves cursor one line up, scrolling if needed
	// ESC 7 	save cursor position (DEC)
	// ESC 8 	restores the cursor to the last saved position (DEC)
	// ESC[s 	save cursor position (SCO)
	// ESC[u 	restores the cursor to the last saved position (SCO)
	//
	regexp.MustCompile("\x1b\\[[Hsu]"),
	regexp.MustCompile("\x1b\\[\\{\\d+\\};\\{\\d+\\}[Hf]"),
	regexp.MustCompile("\x1b\\[\\d+[ABCDEFG]"),
	regexp.MustCompile("\x1b\\[6n"),
	regexp.MustCompile("\x1b [M78]"),

	// Erase escape codes:
	// -----------------------------------------------------------------
	// ESC[J 	erase in display (same as ESC[0J)
	// ESC[0J 	erase from cursor until end of screen
	// ESC[1J 	erase from cursor to beginning of screen
	// ESC[2J 	erase entire screen
	// ESC[3J 	erase saved lines
	// ESC[K 	erase in line (same as ESC[0K)
	// ESC[0K 	erase from cursor to end of line
	// ESC[1K 	erase start of line to the cursor
	// ESC[2K 	erase the entire line
	regexp.MustCompile("\x1b" + `\[\d{0,1}[JKjk]`),

	// Color codes not supported.
	// -----------------------------------------------------------------
	// ESC[38;5;{ID}m 	Set foreground color.
	// ESC[48;5;{ID}m 	Set background color.
	// ESC[38;2;{r};{g};{b}m 	Set foreground color as RGB.
	// ESC[48;2;{r};{g};{b}m 	Set background color as RGB.
	regexp.MustCompile("\x1b" + `\[(?:38|48);5;\{[A-Za-z]+?\}m`),
	regexp.MustCompile("\x1b" + `\[(?:38|48);2;\{\d+\};\{\d+\};\{\d+\}m`),

	// Screen Modes
	// -----------------------------------------------------------------
	// ESC[={value}h 	Changes the screen width or mode type.
	// ESC[=0h 	40 x 25 monochrome (text)
	// ESC[=1h 	40 x 25 color (text)
	// ESC[=2h 	80 x 25 monochrome (text)
	// ESC[=3h 	80 x 25 color (text)
	// ESC[=4h 	320 x 200 4-color (graphics)
	// ESC[=5h 	320 x 200 monochrome (graphics)
	// ESC[=6h 	640 x 200 monochrome (graphics)
	// ESC[=7h 	Enables line wrapping
	// ESC[=13h 	320 x 200 color (graphics)
	// ESC[=14h 	640 x 200 color (16-color graphics)
	// ESC[=15h 	640 x 350 monochrome (2-color graphics)
	// ESC[=16h 	640 x 350 color (16-color graphics)
	// ESC[=17h 	640 x 480 monochrome (2-color graphics)
	// ESC[=18h 	640 x 480 color (16-color graphics)
	// ESC[=19h 	320 x 200 color (256-color graphics)
	// ESC[={value}l Resets the mode by using the same values as Set Mode.
	regexp.MustCompile("\x1b" + `\[\=\{\d+?\}[hl]`),
	regexp.MustCompile("\x1b" + `\[\=\d{1,2}h`),

	// Common Private Modes
	// -----------------------------------------------------------------
	// ESC[?25l 	make cursor invisible
	// ESC[?25h 	make cursor visible
	// ESC[?47l 	restore screen
	// ESC[?47h 	save screen
	// ESC[?1049h 	enables the alternative buffer
	// ESC[?1049l 	disables the alternative buffer
	regexp.MustCompile("\x1b" + `\[\?\d+?[hl]`),
}

func removeUnsupported(str string) string {
	for _, re := range reUnsupported {
		str = re.ReplaceAllString(str, "")
	}

	return str
}
