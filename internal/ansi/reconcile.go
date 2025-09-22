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
	"regexp"
	"strings"

	"github.com/dancsecs/szlog"
)

//nolint:goCheckNoGlobals // Ok.
var (
	openTags []string
	reMulti  = regexp.MustCompile("^\x1b" + `\[([0-9;]+)m`)
)

func removeOpen(tag string) bool {
	var (
		tIdx  int
		found bool
	)

	for fIdx, t := range openTags {
		if t != tag {
			openTags[tIdx] = openTags[fIdx]
			tIdx++
		} else {
			found = true
		}
	}

	openTags = openTags[:tIdx]

	return found
}

func resolveTag(tag string) string {
	var (
		newLine string
		found   bool
		numTags = len(openTags)
	)

	switch tag {
	case BoldDimOff:
		foundBold := removeOpen(Bold)
		foundDim := removeOpen(Dim)
		found = foundBold || foundDim
	case ItalicOff,
		ReverseOff,
		UnderlineOff,
		BlinkOff,
		HiddenOff,
		StrikeoutOff:
		// Just remove the 2x from code.
		found = removeOpen("\x1b[" + tag[3:])
	case Off:
		found = true
		openTags = openTags[:0]
	default:
		newLine = tag
		openTags = append(openTags, tag)
	}

	if found {
		newLine = strings.Repeat(Off, numTags) + strings.Join(openTags, "")
	}

	return newLine
}

// nextTag extracts the ansi tag prefixing the provided line resolving the
// tag, returning yhe updated line and resolved tag.  If the tag is unknown
// then the leading escape character is removed from the line and a stringified
// version of it returned.  If the tag is a multi tag then it is expanded and
// the prepended to the line without resolving the tag.  Finally if it is a
// recognized tag it is removed from the line and a resolved tag is returned.
func nextTag(line string) (string, string) {
	matches := reMulti.FindStringSubmatch(line)
	if matches == nil {
		szlog.Say0("unknown escape sequence: '", line, "'\n")

		return line[1:], "\\x1b"
	}

	codes := strings.Split(matches[1], ";")
	if len(codes) > 1 {
		var linePrefix string

		for _, tagID := range codes {
			linePrefix += "\x1b[" + tagID + "m"
		}

		return linePrefix + line[len(matches[0]):], ""
	}

	return line[len(matches[0]):], resolveTag(matches[0])
}

func reconcileLine(line string) string {
	var resolvedTag string

	newLine := strings.Join(openTags, "")

	for len(line) > 0 {
		escIdx := strings.IndexByte(line, '\x1b')

		if escIdx < 0 {
			newLine += line

			break
		}

		newLine += line[:escIdx]
		line, resolvedTag = nextTag(line[escIdx:])
		newLine += resolvedTag
	}

	return newLine + strings.Repeat(Off, len(openTags))
}

// ReconcileMarkers adds additional resets to account for multiple ansi tags
// being terminated simultaneously.
func reconcileMarkers(str string) string {
	result := ""

	for _, l := range strings.Split(str, "\n") {
		if result != "" {
			result += "\n"
		}

		result += reconcileLine(l)
	}

	return result
}
