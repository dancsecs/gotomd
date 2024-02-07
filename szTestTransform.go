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
	"fmt"
	"math"
	"strings"

	"github.com/dancsecs/sztest"
)

const (
	markInsOn  = "<{INS_ON}>"
	markInsOff = "<{INS_OFF}>"
	markDelOn  = "<{DEL_ON}>"
	markDelOff = "<{DEL_OFF}>"
	markChgOn  = "<{CHG_ON}>"
	markChgOff = "<{CHG_OFF}>"
	markWntOn  = "<{WNT_ON}>"
	markWntOff = "<{WNT_OFF}>"
	markGotOn  = "<{GOT_ON}>"
	markGotOff = "<{GOT_OFF}>"
	markSepOn  = "<{SEP_ON}>"
	markSepOff = "<{SEP_OFF}>"
	markMsgOn  = "<{MSG_ON}>"
	markMsgOff = "<{MSG_OFF}>"
)

const (
	internalTestMarkDelOn  = `\color{red}`
	internalTestMarkDelOff = `\color{default}`
	internalTestMarkInsOn  = `\color{green}`
	internalTestMarkInsOff = `\color{default}`
	internalTestMarkChgOn  = `\color{darkturquoise}`
	internalTestMarkChgOff = `\color{default}`
	internalTestMarkSepOn  = `\color{yellow}`
	internalTestMarkSepOff = `\color{default}`
	internalTestMarkWntOn  = `\color{cyan}`
	internalTestMarkWntOff = `\color{default}`
	internalTestMarkGotOn  = `\color{magenta}`
	internalTestMarkGotOff = `\color{default}`
	internalTestMarkMsgOn  = `\emph{`
	internalTestMarkMsgOff = `}`
)

// findNextMark searches the string for all known marks.
//
//nolint:funlen // Ok.
func findNextMark(s, expectedClose string) (int, string, string, string) {
	if s == "" {
		return -1, "", "", ""
	}

	markOpenIndex := math.MaxInt
	markOpen := ""
	markOpenInternal := ""
	markOpenExpectedInternal := ""

	findOnMark := func(eOpenMark, iOpenMark, iCloseMark string) {
		tmpIndex := strings.Index(s, eOpenMark)
		if tmpIndex >= 0 && tmpIndex < markOpenIndex {
			markOpenIndex = tmpIndex
			markOpen = eOpenMark
			markOpenInternal = iOpenMark
			markOpenExpectedInternal = iCloseMark
		}
	}

	findOnMark(sztest.SettingMarkInsOn(), markInsOn, markInsOff)
	findOnMark(sztest.SettingMarkDelOn(), markDelOn, markDelOff)
	findOnMark(sztest.SettingMarkChgOn(), markChgOn, markChgOff)
	findOnMark(sztest.SettingMarkWntOn(), markWntOn, markWntOff)
	findOnMark(sztest.SettingMarkGotOn(), markGotOn, markGotOff)
	findOnMark(sztest.SettingMarkSepOn(), markSepOn, markSepOff)
	findOnMark(sztest.SettingMarkMsgOn(), markMsgOn, markMsgOff)

	markCloseIndex := math.MaxInt
	markClose := ""
	markCloseInternal := ""

	findOffMark := func(mark, internalMark string) {
		tmpIndex := strings.Index(s, mark)
		if tmpIndex >= 0 &&
			tmpIndex < markOpenIndex &&
			tmpIndex <= markCloseIndex {
			if tmpIndex == markCloseIndex && markCloseInternal == expectedClose {
				return
			}
			markCloseIndex = tmpIndex
			markClose = mark
			markCloseInternal = internalMark
		}
	}

	findOffMark(sztest.SettingMarkInsOff(), markInsOff)
	findOffMark(sztest.SettingMarkDelOff(), markDelOff)
	findOffMark(sztest.SettingMarkChgOff(), markChgOff)
	findOffMark(sztest.SettingMarkWntOff(), markWntOff)
	findOffMark(sztest.SettingMarkGotOff(), markGotOff)
	findOffMark(sztest.SettingMarkSepOff(), markSepOff)
	findOffMark(sztest.SettingMarkMsgOff(), markMsgOff)

	if markOpenIndex < math.MaxInt || markCloseIndex < math.MaxInt {
		if markOpenIndex < markCloseIndex {
			return markOpenIndex,
				markOpen,
				markOpenInternal,
				markOpenExpectedInternal
		} else {
			return markCloseIndex, markClose, markCloseInternal, ""
		}
	}
	return -1, "", "", ""
}

func translateToTestSymbols(s string) string {
	s = strings.ReplaceAll(s, markDelOn, internalTestMarkDelOn)
	s = strings.ReplaceAll(s, markDelOff, internalTestMarkDelOff)
	s = strings.ReplaceAll(s, markInsOn, internalTestMarkInsOn)
	s = strings.ReplaceAll(s, markInsOff, internalTestMarkInsOff)
	s = strings.ReplaceAll(s, markChgOn, internalTestMarkChgOn)
	s = strings.ReplaceAll(s, markChgOff, internalTestMarkChgOff)
	s = strings.ReplaceAll(s, markSepOn, internalTestMarkSepOn)
	s = strings.ReplaceAll(s, markSepOff, internalTestMarkSepOff)
	s = strings.ReplaceAll(s, markWntOn, internalTestMarkWntOn)
	s = strings.ReplaceAll(s, markWntOff, internalTestMarkWntOff)
	s = strings.ReplaceAll(s, markGotOn, internalTestMarkGotOn)
	s = strings.ReplaceAll(s, markGotOff, internalTestMarkGotOff)
	s = strings.ReplaceAll(s, markMsgOn, internalTestMarkMsgOn)
	s = strings.ReplaceAll(s, markMsgOff, internalTestMarkMsgOff)
	return s
}

func translateToBlankSymbols(s string) string {
	s = strings.ReplaceAll(s, markDelOn, "")
	s = strings.ReplaceAll(s, markDelOff, "")
	s = strings.ReplaceAll(s, markInsOn, "")
	s = strings.ReplaceAll(s, markInsOff, "")
	s = strings.ReplaceAll(s, markChgOn, "")
	s = strings.ReplaceAll(s, markChgOff, "")
	s = strings.ReplaceAll(s, markSepOn, "")
	s = strings.ReplaceAll(s, markSepOff, "")
	s = strings.ReplaceAll(s, markWntOn, "")
	s = strings.ReplaceAll(s, markWntOff, "")
	s = strings.ReplaceAll(s, markGotOn, "")
	s = strings.ReplaceAll(s, markGotOff, "")
	s = strings.ReplaceAll(s, markMsgOn, "")
	s = strings.ReplaceAll(s, markMsgOff, "")
	return s
}

func marksToMarkdownHTML(source string) (string, error) {
	iCloseMarkExpected := ""
	newS := ""
	for {
		i, eNextMark, iNextMark, iNextCloseMark :=
			findNextMark(source, iCloseMarkExpected)

		// If no more marks are present then we are done.  Either return the
		// translated string with the all marks reversed or an error if we are
		// expecting a close mark.
		if i < 0 {
			if iCloseMarkExpected != "" {
				return "", fmt.Errorf(
					"no closing mark found for %q in %q",
					iCloseMarkExpected,
					source,
				)
			}

			if szColorize {
				return translateToTestSymbols(newS + source), nil
			}
			return translateToBlankSymbols(newS + source), nil
		}

		// Otherwise we found a Mark.  Move all text up to the next mark from
		// the string to the translated string.
		if i > 0 {
			newS += source[:i]
			source = source[i:]
		}

		// Add the internal representation, replacing the resolved marks.
		newS += iNextMark

		// Remove the resolved Mark from the source string
		source = source[len(eNextMark):]

		if iCloseMarkExpected != "" {
			// There is an open mark that needs to be closed.
			if iNextMark != iCloseMarkExpected {
				return "", fmt.Errorf(
					"unexpected closing mark: Got: %q  Want: %q",
					iNextMark,
					iCloseMarkExpected,
				)
			}
			iCloseMarkExpected = ""
		} else {
			iCloseMarkExpected = iNextCloseMark
		}
	}
}
