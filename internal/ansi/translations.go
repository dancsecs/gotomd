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

// GitHub LaTeX supported color/style codes.
const (
	LaTeXOff       = `}}`
	LaTeXBold      = `{\bf{`
	LaTeXItalic    = `{\emph{`
	LaTeXDim       = "{{" // Not implemented by GitHub.
	LaTeXUnderline = "{{" // Not implemented by GitHub.
	LaTeXBlink     = "{{" // Not implemented by GitHub.
	LaTeXReverse   = "{{" // Not implemented by GitHub.
	LateXHidden    = "{{" // Not implemented by GitHub.
	LaTeXStrikeout = "{{" // Not implemented by GitHub.

	LaTeXBlack   = `{\color{black}{`
	LaTeXRed     = `{\color{red}{`
	LaTeXGreen   = `{\color{green}{`
	LaTeXYellow  = `{\color{yellow}{`
	LaTeXBlue    = `{\color{blue}{`
	LaTeXMagenta = `{\color{magenta}{`
	LaTeXCyan    = `{\color{cyan}{`
	LaTeXWhite   = `{\color{white}{`

	LaTeXHiBlack   = LaTeXBlack   // Not implemented by GitHub.
	LaTeXHiRed     = LaTeXRed     // Not implemented by GitHub.
	LaTeXHiGreen   = LaTeXGreen   // Not implemented by GitHub.
	LaTeXHiYellow  = LaTeXYellow  // Not implemented by GitHub.
	LaTeXHiBlue    = LaTeXBlue    // Not implemented by GitHub.
	LaTeXHiMagenta = LaTeXMagenta // Not implemented by GitHub.
	LaTeXHiCyan    = LaTeXCyan    // Not implemented by GitHub.
	LaTeXHiWhite   = LaTeXWhite   // Not implemented by GitHub.

	LaTeXBkBlack   = LaTeXBlack   // Not implemented by GitHub.
	LaTeXBkRed     = LaTeXRed     // Not implemented by GitHub.
	LaTeXBkGreen   = LaTeXGreen   // Not implemented by GitHub.
	LaTeXBkYellow  = LaTeXYellow  // Not implemented by GitHub.
	LaTeXBkBlue    = LaTeXBlue    // Not implemented by GitHub.
	LaTeXBkMagenta = LaTeXMagenta // Not implemented by GitHub.
	LaTeXBkCyan    = LaTeXCyan    // Not implemented by GitHub.
	LaTeXBkWhite   = LaTeXWhite   // Not implemented by GitHub.

	LaTeXBkHiBlack   = LaTeXBkBlack   // Not implemented by GitHub.
	LaTeXBkHiRed     = LaTeXBkRed     // Not implemented by GitHub.
	LaTeXBkHiGreen   = LaTeXBkGreen   // Not implemented by GitHub.
	LaTeXBkHiYellow  = LaTeXBkYellow  // Not implemented by GitHub.
	LaTeXBkHiBlue    = LaTeXBkBlue    // Not implemented by GitHub.
	LaTeXBkHiMagenta = LaTeXBkMagenta // Not implemented by GitHub.
	LaTeXBkHiCyan    = LaTeXBkCyan    // Not implemented by GitHub.
	LaTeXBkHiWhite   = LaTeXBkWhite   // Not implemented by GitHub.

)

// ANSI terminal color/style escape codes.
const (
	Off          = "\x1b[0m"
	Bold         = "\x1b[1m"
	Dim          = "\x1b[2m"
	BoldDimOff   = "\x1b[22m"
	Italic       = "\x1b[3m"
	ItalicOff    = "\x1b[23m"
	Underline    = "\x1b[4m"
	UnderlineOff = "\x1b[24m"
	Blink        = "\x1b[5m"
	BlinkOff     = "\x1b[25m"
	Reverse      = "\x1b[7m"
	ReverseOff   = "\x1b[27m"
	Hidden       = "\x1b[8m"
	HiddenOff    = "\x1b[28m"
	Strikeout    = "\x1b[9m"
	StrikeoutOff = "\x1b[29m"

	// Basic foreground colors.
	Black   = "\x1b[30m"
	Red     = "\x1b[31m"
	Green   = "\x1b[32m"
	Yellow  = "\x1b[33m"
	Blue    = "\x1b[34m"
	Magenta = "\x1b[35m"
	Cyan    = "\x1b[36m"
	White   = "\x1b[37m"

	// Bright (intense) foreground colors.
	HiBlack   = "\x1b[90m"
	HiRed     = "\x1b[91m"
	HiGreen   = "\x1b[92m"
	HiYellow  = "\x1b[93m"
	HiBlue    = "\x1b[94m"
	HiMagenta = "\x1b[95m"
	HiCyan    = "\x1b[96m"
	HiWhite   = "\x1b[97m"

	// Basic background colors.
	BkBlack   = "\x1b[40m"
	BkRed     = "\x1b[41m"
	BkGreen   = "\x1b[42m"
	BkYellow  = "\x1b[43m"
	BkBlue    = "\x1b[44m"
	BkMagenta = "\x1b[45m"
	BkCyan    = "\x1b[46m"
	BkWhite   = "\x1b[47m"

	// Bright (intense) background colors.
	BkHiBlack   = "\x1b[100m"
	BkHiRed     = "\x1b[101m"
	BkHiGreen   = "\x1b[102m"
	BkHiYellow  = "\x1b[103m"
	BkHiBlue    = "\x1b[104m"
	BkHiMagenta = "\x1b[105m"
	BkHiCyan    = "\x1b[106m"
	BkHiWhite   = "\x1b[107m"
)

//nolint:goCheckNoGlobals // Ok.
var (
	translations = map[string]string{
		Off:       LaTeXOff,
		Bold:      LaTeXBold,
		Italic:    LaTeXItalic,
		Dim:       LaTeXDim,
		Blink:     LaTeXBlink,
		Underline: LaTeXUnderline,
		Reverse:   LaTeXReverse,
		Hidden:    LateXHidden,
		Strikeout: LaTeXStrikeout,

		Black:   LaTeXBlack,
		Red:     LaTeXRed,
		Green:   LaTeXGreen,
		Yellow:  LaTeXYellow,
		Blue:    LaTeXBlue,
		Magenta: LaTeXMagenta,
		Cyan:    LaTeXCyan,
		White:   LaTeXWhite,

		HiBlack:   LaTeXHiBlack,
		HiRed:     LaTeXHiRed,
		HiGreen:   LaTeXHiGreen,
		HiYellow:  LaTeXHiYellow,
		HiBlue:    LaTeXHiBlue,
		HiMagenta: LaTeXHiMagenta,
		HiCyan:    LaTeXHiCyan,
		HiWhite:   LaTeXHiWhite,

		BkBlack:   LaTeXBkBlack,
		BkRed:     LaTeXBkRed,
		BkGreen:   LaTeXBkGreen,
		BkYellow:  LaTeXBkYellow,
		BkBlue:    LaTeXBkBlue,
		BkMagenta: LaTeXBkMagenta,
		BkCyan:    LaTeXBkCyan,
		BkWhite:   LaTeXBkWhite,

		BkHiBlack:   LaTeXBkHiBlack,
		BkHiRed:     LaTeXBkHiRed,
		BkHiGreen:   LaTeXBkHiGreen,
		BkHiYellow:  LaTeXBkHiYellow,
		BkHiBlue:    LaTeXBkHiBlue,
		BkHiMagenta: LaTeXBkHiMagenta,
		BkHiCyan:    LaTeXBkHiCyan,
		BkHiWhite:   LaTeXBkHiWhite,
	}
)
