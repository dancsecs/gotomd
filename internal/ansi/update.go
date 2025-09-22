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
	"strings"
)

const (
	// TabSpaces need to be converted to hard spaces.
	TabSpaces = "    "
	// HardSpace is using a no break space \u00A0 with a "Combining Grapheme
	// Joiner" \u034F which has no visible display but prevents GitHub's LaTeX
	// from merging consecutive spaces.  Two in a row is the same width as a
	// fixed LaTeX font character.
	HardSpace = "&#xA0;&#x34F;&#xA0;&#x34F;" //
	// HardUnderscore is using a no break space \u00A0 with a "Combining Low
	// Line" \u0332 to simulate an underscore which GitHub LaTeX only permits
	// in math mode. Two in a row is the same width as a fixed LaTeX font
	// character.
	HardUnderscore = "&#xA0;&#x332;&#xA0;&#x332;"
	// HardPercent is using a "SMALL PERCENT SIGN" \uFE6A in place of a
	// regular percent sign GitHub markdown processes normal percent signs.
	HardPercent = "&#xFE6A;"
	// Dashes replacing hyphens with an 'FIGURE DASH' u2012 as a regular
	// hyphen in LaTeX is too short (compared to a = used in the
	// corresponding '=== RUN' test bracket.)
	Dashes = "\u2012\u2012\u2012"
)

// Colorize either removes or converts ansi colors into LaTeX text to be
// displayed properly in a GitHub README.md file.
func Colorize(raw string, colorIt bool) string {
	var res string

	openTags = nil

	res, _ = strings.CutSuffix(raw, "\n")
	res = removeUnsupported(res)

	if colorIt {
		res = reconcileMarkers(res)

		res = strings.ReplaceAll(res, "\t", TabSpaces)
		res = strings.ReplaceAll(res, "---", Dashes)
		res = strings.ReplaceAll(res, "%", HardPercent)
		res = strings.ReplaceAll(res, " ", HardSpace)
		res = strings.ReplaceAll(res, "_", HardUnderscore)

		for k := range translations {
			res = strings.ReplaceAll(res, k, translations[k])
		}

		latexRes := ""
		lines := strings.Split(res, "\n")

		for _, line := range lines {
			if latexRes != "" {
				latexRes += "\n"
			}

			latexRes += "$\\small{\\texttt{" +
				line +
				"}}$\n<br>"
		}

		res = latexRes
	} else {
		for k := range translations {
			res = strings.ReplaceAll(res, k, "")
		}
	}

	return res
}
