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
	"slices"
	"strings"
	"testing"

	"github.com/dancsecs/sztestlog"
)

//nolint:goCheckNoGlobals // Ok.
var (
	offTags = []string{
		BoldDimOff,
		ItalicOff,
		UnderlineOff,
		ReverseOff,
		BlinkOff,
		HiddenOff,
		StrikeoutOff,
	}
	//
	validStyles = []string{
		Bold, Italic, Dim, Blink,
		Underline, Reverse, Hidden, Strikeout,
	}
	//
	validColors = []string{
		Black, Red, Green, Yellow,
		Blue, Magenta, Cyan, White,
		//
		HiBlack, HiRed, HiGreen, HiYellow,
		HiBlue, HiMagenta, HiCyan, HiWhite,
		//
		BkBlack, BkRed, BkGreen, BkYellow,
		BkBlue, BkMagenta, BkCyan, BkWhite,
		//
		BkHiBlack, BkHiRed, BkHiGreen, BkHiYellow,
		BkHiBlue, BkHiMagenta, BkHiCyan, BkHiWhite,
	}
	validTags = append(validStyles, validColors...)
)

// func allStylesExcept(except ...string) []string {
// 	tags := make([]string, 0, len(validStyles))

// 	for _, t := range validStyles {
// 		if !slices.Contains(except, t) {
// 			tags = append(tags, t)
// 		}
// 	}

// 	return tags
// }

// func allColorsExcept(except ...string) []string {
// 	tags := make([]string, 0, len(validColors))

// 	for _, t := range validColors {
// 		if !slices.Contains(except, t) {
// 			tags = append(tags, t)
// 		}
// 	}

// 	return tags
// }

func allTagsExcept(except ...string) []string {
	tags := make([]string, 0, len(validTags))

	for _, t := range validTags {
		if !slices.Contains(except, t) {
			tags = append(tags, t)
		}
	}

	return tags
}

func TestAnsi_RemoveOpen_Empty(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	openTags = nil

	chk.False(removeOpen("abc"))

	chk.StrSlice(openTags, nil)
}

func TestAnsi_RemoveOpenTwo_NotFound(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	openTags = []string{"abc", "def"}

	chk.False(removeOpen("ghi"))

	chk.StrSlice(openTags, []string{"abc", "def"})
}

func TestAnsi_RemoveOpenThree_NotFound(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	openTags = []string{"abc", "def", "ghi"}

	chk.False(removeOpen("jkl"))

	chk.StrSlice(openTags, []string{"abc", "def", "ghi"})
}

func TestAnsi_RemoveOpenOne_Found(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	openTags = []string{"abc"}

	chk.True(removeOpen("abc"))

	chk.StrSlice(openTags, nil)
}

func TestAnsi_RemoveOpenTwo_FoundFirst(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	openTags = []string{"abc", "def"}

	chk.True(removeOpen("abc"))

	chk.StrSlice(openTags, []string{"def"})
}

func TestAnsi_RemoveOpenTwo_FoundLast(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	openTags = []string{"abc", "def"}

	chk.True(removeOpen("def"))

	chk.StrSlice(openTags, []string{"abc"})
}

func TestAnsi_RemoveOpenThree_FoundFirst(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	openTags = []string{"abc", "def", "ghi"}

	chk.True(removeOpen("abc"))

	chk.StrSlice(openTags, []string{"def", "ghi"})
}

func TestAnsi_RemoveOpenThree_FoundMiddle(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	openTags = []string{"abc", "def", "ghi"}

	chk.True(removeOpen("def"))

	chk.StrSlice(openTags, []string{"abc", "ghi"})
}

func TestAnsi_RemoveOpenThree_FoundLast(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	openTags = []string{"abc", "def", "ghi"}

	chk.True(removeOpen("ghi"))

	chk.StrSlice(openTags, []string{"abc", "def"})
}

func TestAnsi_RemoveOpenThree_FoundFistAndLast(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	openTags = []string{"abc", "def", "abc"}

	chk.True(removeOpen("abc"))

	chk.StrSlice(openTags, []string{"def"})
}

func TestAnsi_ResolveTag_Normal(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	openTags = nil

	str := resolveTag("abc")

	chk.Str(str, "abc")

	chk.StrSlice(openTags, []string{"abc"})
}

func TestAnsi_ResolveTag_NoOpenTags_Off(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	openTags = nil

	str := resolveTag(Off)
	chk.Str(str, "")
	chk.StrSlice(openTags, nil)

	for i, t := range offTags {
		str = resolveTag(t)
		chk.Str(str, "", "Index:", i)
		chk.StrSlice(openTags, nil, "index: .i")
	}
}

func TestAnsi_ResolveTag_Off(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	for i, t := range validTags {
		openTags = append([]string(nil), t)
		str := resolveTag(Off)
		chk.Str(str, Off, "Index: ", i)
		chk.StrSlice(openTags, nil, "Index: ", i)
	}

	openTags = allTagsExcept()

	str := resolveTag(Off)
	chk.Str(str, strings.Repeat(Off, len(validTags)))
	chk.StrSlice(openTags, nil)
}

func TestAnsi_ResolveTag_BoldDimOff(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	openTags = allTagsExcept(Bold, Dim)

	str := resolveTag(BoldDimOff)
	chk.Str(str, "")
	chk.StrSlice(openTags, allTagsExcept(Bold, Dim))

	openTags = append(openTags, Bold)
	str = resolveTag(BoldDimOff)
	chk.Str(
		str,
		strings.Repeat(Off, len(validTags)-1)+
			strings.Join(allTagsExcept(Bold, Dim), ""),
	)
	chk.StrSlice(openTags, allTagsExcept(Bold, Dim))

	openTags = append(openTags, Dim)
	str = resolveTag(BoldDimOff)
	chk.Str(
		str,
		strings.Repeat(Off, len(validTags)-1)+
			strings.Join(allTagsExcept(Bold, Dim), ""),
	)
	chk.StrSlice(openTags, allTagsExcept(Bold, Dim))

	openTags = append(openTags, Bold, Dim)
	str = resolveTag(BoldDimOff)
	chk.Str(
		str,
		strings.Repeat(Off, len(validTags))+
			strings.Join(allTagsExcept(Bold, Dim), ""),
	)
	chk.StrSlice(openTags, allTagsExcept(Bold, Dim))
}

func TestAnsi_ResolveTag_ItalicOff(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	openTags = allTagsExcept(Italic)

	str := resolveTag(ItalicOff)
	chk.Str(str, "")
	chk.StrSlice(openTags, allTagsExcept(Italic))

	openTags = append(openTags, Italic)
	str = resolveTag(ItalicOff)
	chk.Str(
		str,
		strings.Repeat(Off, len(validTags))+
			strings.Join(allTagsExcept(Italic), ""),
	)
	chk.StrSlice(openTags, allTagsExcept(Italic))
}

func TestAnsi_ResolveTag_UnderlineOff(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	openTags = allTagsExcept(Underline)

	str := resolveTag(UnderlineOff)
	chk.Str(str, "")
	chk.StrSlice(openTags, allTagsExcept(Underline))

	openTags = append(openTags, Underline)
	str = resolveTag(UnderlineOff)
	chk.Str(
		str,
		strings.Repeat(Off, len(validTags))+
			strings.Join(allTagsExcept(Underline), ""),
	)
	chk.StrSlice(openTags, allTagsExcept(Underline))
}

func TestAnsi_ResolveTag_ReverseOff(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	openTags = allTagsExcept(Reverse)

	str := resolveTag(ReverseOff)
	chk.Str(str, "")
	chk.StrSlice(openTags, allTagsExcept(Reverse))

	openTags = append(openTags, Reverse)
	str = resolveTag(ReverseOff)
	chk.Str(
		str,
		strings.Repeat(Off, len(validTags))+
			strings.Join(allTagsExcept(Reverse), ""),
	)
	chk.StrSlice(openTags, allTagsExcept(Reverse))
}

func TestAnsi_ResolveTag_BlinkOff(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	openTags = allTagsExcept(Blink)

	str := resolveTag(BlinkOff)
	chk.Str(str, "")
	chk.StrSlice(openTags, allTagsExcept(Blink))

	openTags = append(openTags, Blink)
	str = resolveTag(BlinkOff)
	chk.Str(
		str,
		strings.Repeat(Off, len(validTags))+
			strings.Join(allTagsExcept(Blink), ""),
	)
	chk.StrSlice(openTags, allTagsExcept(Blink))
}

func TestAnsi_ResolveTag_HiddenOff(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	openTags = allTagsExcept(Hidden)

	str := resolveTag(HiddenOff)
	chk.Str(str, "")
	chk.StrSlice(openTags, allTagsExcept(Hidden))

	openTags = append(openTags, Hidden)
	str = resolveTag(HiddenOff)
	chk.Str(
		str,
		strings.Repeat(Off, len(validTags))+
			strings.Join(allTagsExcept(Hidden), ""),
	)
	chk.StrSlice(openTags, allTagsExcept(Hidden))
}

func TestAnsi_ResolveTag_StrikeoutOff(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	openTags = allTagsExcept(Strikeout)

	str := resolveTag(StrikeoutOff)
	chk.Str(str, "")
	chk.StrSlice(openTags, allTagsExcept(Strikeout))

	openTags = append(openTags, Strikeout)
	str = resolveTag(StrikeoutOff)
	chk.Str(
		str,
		strings.Repeat(Off, len(validTags))+
			strings.Join(allTagsExcept(Strikeout), ""),
	)
	chk.StrSlice(openTags, allTagsExcept(Strikeout))
}

func TestAnsi_NextTag_NoTags(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	openTags = nil

	origLine := "\x1bThere are no tags"

	newLine, resolvedTag := nextTag(origLine)
	chk.Str(newLine, "There are no tags")
	chk.Str(resolvedTag, "\\x1b")

	chk.Stdout(
		"unknown escape sequence: '" + origLine + "'",
	)
}

func TestAnsi_NextTag_SingleTag(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	openTags = nil

	origLine := Bold + "There is just a single Bold tag."

	newLine, resolvedTag := nextTag(origLine)
	chk.Str(newLine, "There is just a single Bold tag.")
	chk.Str(resolvedTag, Bold)
	chk.StrSlice(openTags, []string{Bold})
}

func TestAnsi_NextTag_TwoTags(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	openTags = nil

	origLine := "\x1b[1;2mThere are two embedded tags."

	newLine, resolvedTag := nextTag(origLine)
	chk.Str(newLine, Bold+Dim+"There are two embedded tags.")
	chk.Str(resolvedTag, "")
	chk.StrSlice(openTags, nil)
}

func TestAnsi_ResolveNextTag_ThreeTags(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	openTags = nil

	origLine := "\x1b[1;2;3mThere are three embedded tags."

	newLine, resolvedTag := nextTag(origLine)
	chk.Str(newLine, Bold+Dim+Italic+"There are three embedded tags.")
	chk.Str(resolvedTag, "")
	chk.StrSlice(openTags, nil)
}

func TestAnsi_ReconcileLine_NoTags(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	openTags = nil

	origLine := "There are no open or embedded tags."

	reconciledLine := reconcileLine(origLine)
	chk.Str(reconciledLine, origLine)
	chk.StrSlice(openTags, nil)
}

func TestAnsi_ReconcileLine_OneOpenTag(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	openTags = []string{Blue}

	origLine := "There is one open tags and no embedded tags."

	reconciledLine := reconcileLine(origLine)
	chk.Str(reconciledLine, Blue+origLine+Off)
	chk.StrSlice(openTags, []string{Blue})
}

func TestAnsi_ReconcileLine_TwoOpenTags(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	openTags = []string{Bold, Blue}

	origLine := "There are two open and one embedded open and close (" +
		Italic + "italic" + ItalicOff + ") tag."

	reconciledLine := reconcileLine(origLine)
	chk.Str(
		reconciledLine,
		Bold+Blue+"There are two open and one embedded open and close ("+
			Italic+"italic"+Off+Off+Off+Bold+Blue+") tag."+Off+Off,
	)
	chk.StrSlice(openTags, []string{Bold, Blue})
}

func TestAnsi_ReconcileMarkers_OpenTagsOverALine(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	openTags = nil

	origLines := Blue + "The blue tag\n" +
		"will " + Bold + "continue" + BoldDimOff + " to the next line\n" +
		"but must be " + Italic + "closed" + ItalicOff + " on the\n" +
		"previous line."

	chk.StrSlice(
		strings.Split(reconcileMarkers(origLines), "\n"),
		[]string{
			Blue + "The blue tag" + Off,
			Blue + "will " + Bold + "continue" + Off + Off + Blue +
				" to the next line" + Off,
			Blue + "but must be " + Italic + "closed" + Off + Off + Blue +
				" on the" + Off,
			Blue + "previous line." + Off,
		},
	)
}
