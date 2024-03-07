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
