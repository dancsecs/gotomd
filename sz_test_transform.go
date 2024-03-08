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

//nolint:goCheckNoGlobals // Ok.
var szEnvSetup = []string{
	sztest.EnvMarkWntOn + "=" + markWntOn,
	sztest.EnvMarkWntOff + "=" + markWntOff,
	sztest.EnvMarkGotOn + "=" + markGotOn,
	sztest.EnvMarkGotOff + "=" + markGotOff,
	sztest.EnvMarkMsgOn + "=" + markMsgOn,
	sztest.EnvMarkMsgOff + "=" + markMsgOff,
	sztest.EnvMarkInsOn + "=" + markInsOn,
	sztest.EnvMarkInsOff + "=" + markInsOff,
	sztest.EnvMarkDelOn + "=" + markDelOn,
	sztest.EnvMarkDelOff + "=" + markDelOff,
	sztest.EnvMarkChgOn + "=" + markChgOn,
	sztest.EnvMarkChgOff + "=" + markChgOff,
	sztest.EnvMarkSepOn + "=" + markSepOn,
	sztest.EnvMarkSepOff + "=" + markSepOff,
}

func translateToTestSymbols(lines string) string {
	lines = strings.ReplaceAll(lines, markDelOn, internalTestMarkDelOn)
	lines = strings.ReplaceAll(lines, markDelOff, internalTestMarkDelOff)
	lines = strings.ReplaceAll(lines, markInsOn, internalTestMarkInsOn)
	lines = strings.ReplaceAll(lines, markInsOff, internalTestMarkInsOff)
	lines = strings.ReplaceAll(lines, markChgOn, internalTestMarkChgOn)
	lines = strings.ReplaceAll(lines, markChgOff, internalTestMarkChgOff)
	lines = strings.ReplaceAll(lines, markSepOn, internalTestMarkSepOn)
	lines = strings.ReplaceAll(lines, markSepOff, internalTestMarkSepOff)
	lines = strings.ReplaceAll(lines, markWntOn, internalTestMarkWntOn)
	lines = strings.ReplaceAll(lines, markWntOff, internalTestMarkWntOff)
	lines = strings.ReplaceAll(lines, markGotOn, internalTestMarkGotOn)
	lines = strings.ReplaceAll(lines, markGotOff, internalTestMarkGotOff)
	lines = strings.ReplaceAll(lines, markMsgOn, internalTestMarkMsgOn)
	lines = strings.ReplaceAll(lines, markMsgOff, internalTestMarkMsgOff)

	return lines
}

func translateToBlankSymbols(lines string) string {
	lines = strings.ReplaceAll(lines, markDelOn, "")
	lines = strings.ReplaceAll(lines, markDelOff, "")
	lines = strings.ReplaceAll(lines, markInsOn, "")
	lines = strings.ReplaceAll(lines, markInsOff, "")
	lines = strings.ReplaceAll(lines, markChgOn, "")
	lines = strings.ReplaceAll(lines, markChgOff, "")
	lines = strings.ReplaceAll(lines, markSepOn, "")
	lines = strings.ReplaceAll(lines, markSepOff, "")
	lines = strings.ReplaceAll(lines, markWntOn, "")
	lines = strings.ReplaceAll(lines, markWntOff, "")
	lines = strings.ReplaceAll(lines, markGotOn, "")
	lines = strings.ReplaceAll(lines, markGotOff, "")
	lines = strings.ReplaceAll(lines, markMsgOn, "")
	lines = strings.ReplaceAll(lines, markMsgOff, "")

	return lines
}
