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

//nolint:lll // Ok LaTeX examples are longer.
package ansi_test

import (
	"fmt"
	"testing"

	"github.com/dancsecs/gotomd/internal/ansi"
	"github.com/dancsecs/sztestlog"
)

//nolint:goCheckNoGlobals // Ok.
var testStrings = []string{
	"This is " + ansi.Bold + "bold" + ansi.Off + ".",
	"This is " + ansi.Italic + "italic" + ansi.Off + ".",
	"This is " + ansi.Underline + "underline not implemented" + ansi.Off + ".",
	"This is " + ansi.Reverse + "reverse not implemented" + ansi.Off + ".",
	"This is " + ansi.Blink + "blink not implemented" + ansi.Off + ".",
	"This is " + ansi.Hidden + "hidden not implemented" + ansi.Off + ".",
	"This is " + ansi.Strikeout + "strikeout not implemented" + ansi.Off + ".",

	"This is " + ansi.Black + "black" + ansi.Off + ".",
	"This is " + ansi.Red + "red" + ansi.Off + ".",
	"This is " + ansi.Green + "green" + ansi.Off + ".",
	"This is " + ansi.Yellow + "yellow" + ansi.Off + ".",
	"This is " + ansi.Blue + "blue" + ansi.Off + ".",
	"This is " + ansi.Magenta + "magenta" + ansi.Off + ".",
	"This is " + ansi.Cyan + "cyan" + ansi.Off + ".",
	"This is " + ansi.White + "white" + ansi.Off + ".",

	"This is " + ansi.HiBlack + "hiBlack not implemented" + ansi.Off + ".",
	"This is " + ansi.HiRed + "hiRed not implemented" + ansi.Off + ".",
	"This is " + ansi.HiGreen + "hiGreen not implemented" + ansi.Off + ".",
	"This is " + ansi.HiYellow + "hiYellow not implemented" + ansi.Off + ".",
	"This is " + ansi.HiBlue + "hiBlue not implemented" + ansi.Off + ".",
	"This is " + ansi.HiMagenta + "hiMagenta not implemented" + ansi.Off + ".",
	"This is " + ansi.HiCyan + "hiCyan not implemented" + ansi.Off + ".",
	"This is " + ansi.HiWhite + "hiWhite not implemented" + ansi.Off + ".",

	"This is " + ansi.BkBlack + "bkBlack not implemented" + ansi.Off + ".",
	"This is " + ansi.BkRed + "bkRed not implemented" + ansi.Off + ".",
	"This is " + ansi.BkGreen + "bkGreen not implemented" + ansi.Off + ".",
	"This is " + ansi.BkYellow + "bkYellow not implemented" + ansi.Off + ".",
	"This is " + ansi.BkBlue + "bkBlue not implemented" + ansi.Off + ".",
	"This is " + ansi.BkMagenta + "bkMagenta not implemented" + ansi.Off + ".",
	"This is " + ansi.BkCyan + "bkCyan not implemented" + ansi.Off + ".",
	"This is " + ansi.BkWhite + "bkWhite not implemented" + ansi.Off + ".",

	"This is " + ansi.BkHiBlack + "bkHiBlack not implemented" + ansi.Off + ".",
	"This is " + ansi.BkHiRed + "bkHiRed not implemented" + ansi.Off + ".",
	"This is " + ansi.BkHiGreen + "bkHiGreen not implemented" + ansi.Off + ".",
	"This is " + ansi.BkHiYellow + "bkHiYellow not implemented" + ansi.Off + ".",
	"This is " + ansi.BkHiBlue + "bkHiBlue not implemented" + ansi.Off + ".",
	"This is " + ansi.BkHiMagenta + "bkHiMagenta not implemented" + ansi.Off + ".",
	"This is " + ansi.BkHiCyan + "bkHiCyan not implemented" + ansi.Off + ".",
	"This is " + ansi.BkHiWhite + "bkHiWhite not implemented" + ansi.Off + ".",
}

//nolint:goCheckNoGlobals,lll // LaTeX strings get long.
var latexResults = []string{
	"$\\small{\\texttt{This&#xA0;&#x34F;&#xA0;&#x34F;is&#xA0;&#x34F;&#xA0;&#x34F;{\\bf{bold}}.}}$\n<br>",
	"$\\small{\\texttt{This&#xA0;&#x34F;&#xA0;&#x34F;is&#xA0;&#x34F;&#xA0;&#x34F;{\\emph{italic}}.}}$\n<br>",
	"$\\small{\\texttt{This&#xA0;&#x34F;&#xA0;&#x34F;is&#xA0;&#x34F;&#xA0;&#x34F;{{underline&#xA0;&#x34F;&#xA0;&#x34F;not&#xA0;&#x34F;&#xA0;&#x34F;implemented}}.}}$\n<br>",
	"$\\small{\\texttt{This&#xA0;&#x34F;&#xA0;&#x34F;is&#xA0;&#x34F;&#xA0;&#x34F;{{reverse&#xA0;&#x34F;&#xA0;&#x34F;not&#xA0;&#x34F;&#xA0;&#x34F;implemented}}.}}$\n<br>",
	"$\\small{\\texttt{This&#xA0;&#x34F;&#xA0;&#x34F;is&#xA0;&#x34F;&#xA0;&#x34F;{{blink&#xA0;&#x34F;&#xA0;&#x34F;not&#xA0;&#x34F;&#xA0;&#x34F;implemented}}.}}$\n<br>",
	"$\\small{\\texttt{This&#xA0;&#x34F;&#xA0;&#x34F;is&#xA0;&#x34F;&#xA0;&#x34F;{{hidden&#xA0;&#x34F;&#xA0;&#x34F;not&#xA0;&#x34F;&#xA0;&#x34F;implemented}}.}}$\n<br>",
	"$\\small{\\texttt{This&#xA0;&#x34F;&#xA0;&#x34F;is&#xA0;&#x34F;&#xA0;&#x34F;{{strikeout&#xA0;&#x34F;&#xA0;&#x34F;not&#xA0;&#x34F;&#xA0;&#x34F;implemented}}.}}$\n<br>",

	"$\\small{\\texttt{This&#xA0;&#x34F;&#xA0;&#x34F;is&#xA0;&#x34F;&#xA0;&#x34F;{\\color{black}{black}}.}}$\n<br>",
	"$\\small{\\texttt{This&#xA0;&#x34F;&#xA0;&#x34F;is&#xA0;&#x34F;&#xA0;&#x34F;{\\color{red}{red}}.}}$\n<br>",
	"$\\small{\\texttt{This&#xA0;&#x34F;&#xA0;&#x34F;is&#xA0;&#x34F;&#xA0;&#x34F;{\\color{green}{green}}.}}$\n<br>",
	"$\\small{\\texttt{This&#xA0;&#x34F;&#xA0;&#x34F;is&#xA0;&#x34F;&#xA0;&#x34F;{\\color{yellow}{yellow}}.}}$\n<br>",
	"$\\small{\\texttt{This&#xA0;&#x34F;&#xA0;&#x34F;is&#xA0;&#x34F;&#xA0;&#x34F;{\\color{blue}{blue}}.}}$\n<br>",
	"$\\small{\\texttt{This&#xA0;&#x34F;&#xA0;&#x34F;is&#xA0;&#x34F;&#xA0;&#x34F;{\\color{magenta}{magenta}}.}}$\n<br>",
	"$\\small{\\texttt{This&#xA0;&#x34F;&#xA0;&#x34F;is&#xA0;&#x34F;&#xA0;&#x34F;{\\color{cyan}{cyan}}.}}$\n<br>",
	"$\\small{\\texttt{This&#xA0;&#x34F;&#xA0;&#x34F;is&#xA0;&#x34F;&#xA0;&#x34F;{\\color{white}{white}}.}}$\n<br>",

	"$\\small{\\texttt{This&#xA0;&#x34F;&#xA0;&#x34F;is&#xA0;&#x34F;&#xA0;&#x34F;{\\color{black}{hiBlack&#xA0;&#x34F;&#xA0;&#x34F;not&#xA0;&#x34F;&#xA0;&#x34F;implemented}}.}}$\n<br>",
	"$\\small{\\texttt{This&#xA0;&#x34F;&#xA0;&#x34F;is&#xA0;&#x34F;&#xA0;&#x34F;{\\color{red}{hiRed&#xA0;&#x34F;&#xA0;&#x34F;not&#xA0;&#x34F;&#xA0;&#x34F;implemented}}.}}$\n<br>",
	"$\\small{\\texttt{This&#xA0;&#x34F;&#xA0;&#x34F;is&#xA0;&#x34F;&#xA0;&#x34F;{\\color{green}{hiGreen&#xA0;&#x34F;&#xA0;&#x34F;not&#xA0;&#x34F;&#xA0;&#x34F;implemented}}.}}$\n<br>",
	"$\\small{\\texttt{This&#xA0;&#x34F;&#xA0;&#x34F;is&#xA0;&#x34F;&#xA0;&#x34F;{\\color{yellow}{hiYellow&#xA0;&#x34F;&#xA0;&#x34F;not&#xA0;&#x34F;&#xA0;&#x34F;implemented}}.}}$\n<br>",
	"$\\small{\\texttt{This&#xA0;&#x34F;&#xA0;&#x34F;is&#xA0;&#x34F;&#xA0;&#x34F;{\\color{blue}{hiBlue&#xA0;&#x34F;&#xA0;&#x34F;not&#xA0;&#x34F;&#xA0;&#x34F;implemented}}.}}$\n<br>",
	"$\\small{\\texttt{This&#xA0;&#x34F;&#xA0;&#x34F;is&#xA0;&#x34F;&#xA0;&#x34F;{\\color{magenta}{hiMagenta&#xA0;&#x34F;&#xA0;&#x34F;not&#xA0;&#x34F;&#xA0;&#x34F;implemented}}.}}$\n<br>",
	"$\\small{\\texttt{This&#xA0;&#x34F;&#xA0;&#x34F;is&#xA0;&#x34F;&#xA0;&#x34F;{\\color{cyan}{hiCyan&#xA0;&#x34F;&#xA0;&#x34F;not&#xA0;&#x34F;&#xA0;&#x34F;implemented}}.}}$\n<br>",
	"$\\small{\\texttt{This&#xA0;&#x34F;&#xA0;&#x34F;is&#xA0;&#x34F;&#xA0;&#x34F;{\\color{white}{hiWhite&#xA0;&#x34F;&#xA0;&#x34F;not&#xA0;&#x34F;&#xA0;&#x34F;implemented}}.}}$\n<br>",

	"$\\small{\\texttt{This&#xA0;&#x34F;&#xA0;&#x34F;is&#xA0;&#x34F;&#xA0;&#x34F;{\\color{black}{bkBlack&#xA0;&#x34F;&#xA0;&#x34F;not&#xA0;&#x34F;&#xA0;&#x34F;implemented}}.}}$\n<br>",
	"$\\small{\\texttt{This&#xA0;&#x34F;&#xA0;&#x34F;is&#xA0;&#x34F;&#xA0;&#x34F;{\\color{red}{bkRed&#xA0;&#x34F;&#xA0;&#x34F;not&#xA0;&#x34F;&#xA0;&#x34F;implemented}}.}}$\n<br>",
	"$\\small{\\texttt{This&#xA0;&#x34F;&#xA0;&#x34F;is&#xA0;&#x34F;&#xA0;&#x34F;{\\color{green}{bkGreen&#xA0;&#x34F;&#xA0;&#x34F;not&#xA0;&#x34F;&#xA0;&#x34F;implemented}}.}}$\n<br>",
	"$\\small{\\texttt{This&#xA0;&#x34F;&#xA0;&#x34F;is&#xA0;&#x34F;&#xA0;&#x34F;{\\color{yellow}{bkYellow&#xA0;&#x34F;&#xA0;&#x34F;not&#xA0;&#x34F;&#xA0;&#x34F;implemented}}.}}$\n<br>",
	"$\\small{\\texttt{This&#xA0;&#x34F;&#xA0;&#x34F;is&#xA0;&#x34F;&#xA0;&#x34F;{\\color{blue}{bkBlue&#xA0;&#x34F;&#xA0;&#x34F;not&#xA0;&#x34F;&#xA0;&#x34F;implemented}}.}}$\n<br>",
	"$\\small{\\texttt{This&#xA0;&#x34F;&#xA0;&#x34F;is&#xA0;&#x34F;&#xA0;&#x34F;{\\color{magenta}{bkMagenta&#xA0;&#x34F;&#xA0;&#x34F;not&#xA0;&#x34F;&#xA0;&#x34F;implemented}}.}}$\n<br>",
	"$\\small{\\texttt{This&#xA0;&#x34F;&#xA0;&#x34F;is&#xA0;&#x34F;&#xA0;&#x34F;{\\color{cyan}{bkCyan&#xA0;&#x34F;&#xA0;&#x34F;not&#xA0;&#x34F;&#xA0;&#x34F;implemented}}.}}$\n<br>",
	"$\\small{\\texttt{This&#xA0;&#x34F;&#xA0;&#x34F;is&#xA0;&#x34F;&#xA0;&#x34F;{\\color{white}{bkWhite&#xA0;&#x34F;&#xA0;&#x34F;not&#xA0;&#x34F;&#xA0;&#x34F;implemented}}.}}$\n<br>",

	"$\\small{\\texttt{This&#xA0;&#x34F;&#xA0;&#x34F;is&#xA0;&#x34F;&#xA0;&#x34F;{\\color{black}{bkHiBlack&#xA0;&#x34F;&#xA0;&#x34F;not&#xA0;&#x34F;&#xA0;&#x34F;implemented}}.}}$\n<br>",
	"$\\small{\\texttt{This&#xA0;&#x34F;&#xA0;&#x34F;is&#xA0;&#x34F;&#xA0;&#x34F;{\\color{red}{bkHiRed&#xA0;&#x34F;&#xA0;&#x34F;not&#xA0;&#x34F;&#xA0;&#x34F;implemented}}.}}$\n<br>",
	"$\\small{\\texttt{This&#xA0;&#x34F;&#xA0;&#x34F;is&#xA0;&#x34F;&#xA0;&#x34F;{\\color{green}{bkHiGreen&#xA0;&#x34F;&#xA0;&#x34F;not&#xA0;&#x34F;&#xA0;&#x34F;implemented}}.}}$\n<br>",
	"$\\small{\\texttt{This&#xA0;&#x34F;&#xA0;&#x34F;is&#xA0;&#x34F;&#xA0;&#x34F;{\\color{yellow}{bkHiYellow&#xA0;&#x34F;&#xA0;&#x34F;not&#xA0;&#x34F;&#xA0;&#x34F;implemented}}.}}$\n<br>",
	"$\\small{\\texttt{This&#xA0;&#x34F;&#xA0;&#x34F;is&#xA0;&#x34F;&#xA0;&#x34F;{\\color{blue}{bkHiBlue&#xA0;&#x34F;&#xA0;&#x34F;not&#xA0;&#x34F;&#xA0;&#x34F;implemented}}.}}$\n<br>",
	"$\\small{\\texttt{This&#xA0;&#x34F;&#xA0;&#x34F;is&#xA0;&#x34F;&#xA0;&#x34F;{\\color{magenta}{bkHiMagenta&#xA0;&#x34F;&#xA0;&#x34F;not&#xA0;&#x34F;&#xA0;&#x34F;implemented}}.}}$\n<br>",
	"$\\small{\\texttt{This&#xA0;&#x34F;&#xA0;&#x34F;is&#xA0;&#x34F;&#xA0;&#x34F;{\\color{cyan}{bkHiCyan&#xA0;&#x34F;&#xA0;&#x34F;not&#xA0;&#x34F;&#xA0;&#x34F;implemented}}.}}$\n<br>",
	"$\\small{\\texttt{This&#xA0;&#x34F;&#xA0;&#x34F;is&#xA0;&#x34F;&#xA0;&#x34F;{\\color{white}{bkHiWhite&#xA0;&#x34F;&#xA0;&#x34F;not&#xA0;&#x34F;&#xA0;&#x34F;implemented}}.}}$\n<br>",
}

//nolint:goCheckNoGlobals // Ok.
var cleanResults = []string{
	"This is bold.",
	"This is italic.",
	"This is underline not implemented.",
	"This is reverse not implemented.",
	"This is blink not implemented.",
	"This is hidden not implemented.",
	"This is strikeout not implemented.",

	"This is black.",
	"This is red.",
	"This is green.",
	"This is yellow.",
	"This is blue.",
	"This is magenta.",
	"This is cyan.",
	"This is white.",

	"This is hiBlack not implemented.",
	"This is hiRed not implemented.",
	"This is hiGreen not implemented.",
	"This is hiYellow not implemented.",
	"This is hiBlue not implemented.",
	"This is hiMagenta not implemented.",
	"This is hiCyan not implemented.",
	"This is hiWhite not implemented.",

	"This is bkBlack not implemented.",
	"This is bkRed not implemented.",
	"This is bkGreen not implemented.",
	"This is bkYellow not implemented.",
	"This is bkBlue not implemented.",
	"This is bkMagenta not implemented.",
	"This is bkCyan not implemented.",
	"This is bkWhite not implemented.",

	"This is bkHiBlack not implemented.",
	"This is bkHiRed not implemented.",
	"This is bkHiGreen not implemented.",
	"This is bkHiYellow not implemented.",
	"This is bkHiBlue not implemented.",
	"This is bkHiMagenta not implemented.",
	"This is bkHiCyan not implemented.",
	"This is bkHiWhite not implemented.",
}

func TestColorize_NoAnsi(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	chk.Str(
		ansi.Colorize("This string has no ansi escape codes.", true),
		`$\small{\texttt{This&#xA0;&#x34F;&#xA0;&#x34F;string`+
			`&#xA0;&#x34F;&#xA0;&#x34F;has&#xA0;&#x34F;&#xA0;&#x34F;no`+
			`&#xA0;&#x34F;&#xA0;&#x34F;ansi&#xA0;&#x34F;&#xA0;&#x34F;`+
			`escape&#xA0;&#x34F;&#xA0;&#x34F;codes.}}$\n<br>`,
	)
}

func TestColorize_ToLatex(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	chk.AddSub(`\&\#xA0\;\&\#x34F\;\&\#xA0\;\&\#x34F\;`, " ")

	for i, ln := range testStrings {
		fmt.Println(ln)
		chk.Str(
			ansi.Colorize(ln, true),
			latexResults[i],
			"Index: ", i,
		)
	}
}

func TestColorize_ToClean(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	for i, ln := range testStrings {
		chk.Str(
			ansi.Colorize(ln, false),
			cleanResults[i],
			"Index: ", i,
		)
	}
}

func TestColorize_MultilineToLatex(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	chk.AddSub(`\&\#xA0\;\&\#x34F\;\&\#xA0\;\&\#x34F\;`, " ")
	chk.Str(
		ansi.Colorize(
			"This is "+ansi.Bold+"bold"+ansi.Off+".\n"+
				"This is "+ansi.Italic+"italic"+ansi.Off+".\n",
			true,
		),
		"$\\small{\\texttt{This is {\\bf{bold}}.}}$\n<br>\n"+
			"$\\small{\\texttt{This is {\\emph{italic}}.}}$\n<br>",
	)
}

func TestColorize_MultiAnsiAttributesToLatex(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	chk.AddSub(`\&\#xA0\;\&\#x34F\;\&\#xA0\;\&\#x34F\;`, " ")
	chk.Str(
		ansi.Colorize(
			"This is "+ansi.Bold+ansi.Italic+"bold and italic"+ansi.Off+".\n",
			true,
		),
		`$\small{\texttt{This&#xA0;&#x34F;&#xA0;&#x34F;is&#xA0;&#x34F;&#xA0;&#x34F;{\bf{{\emph{bold&#xA0;&#x34F;&#xA0;&#x34F;and&#xA0;&#x34F;&#xA0;&#x34F;italic}}}}.}}$`+"\n<br>",
	)
}
