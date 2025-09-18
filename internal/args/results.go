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

package args

import "os"

const defaultPerm = os.FileMode(0o0644)

//nolint:goCheckNoGlobals // Ok.
var (
	goFiles        []string
	mdFiles        []string
	usage          string
	forceOverwrite bool
	szColorize     bool
	outputDir      = "."
	perm           = defaultPerm
	showLicense    bool
	showHelp       bool

	alreadyIncluded = make(map[string]bool)
)

// Reset clears any existing results.
func Reset() {
	goFiles = nil
	mdFiles = nil
	usage = ""
	forceOverwrite = false
	szColorize = false
	outputDir = "."
	perm = defaultPerm
	showLicense = false
	showHelp = false

	for k := range alreadyIncluded {
		delete(alreadyIncluded, k)
	}
}

// GoFiles returns the found go file templates.
func GoFiles() []string {
	return goFiles
}

// MdFiles returns the found markdown file templates.
func MdFiles() []string {
	return mdFiles
}

// Usage returns the usage string.
func Usage() string {
	return usage
}

// Force returns the force setting.
func Force() bool {
	return forceOverwrite
}

// Colorize returns the colorize setting.
func Colorize() bool {
	return szColorize
}

// OutputDir returns the overridden output directory.
func OutputDir() string {
	return outputDir
}

// Perm return the default permission.
func Perm() os.FileMode {
	return perm
}

// ShowLicense returns the show license setting.
func ShowLicense() bool {
	return showLicense
}

// ShowHelp returns the show help setting.
func ShowHelp() bool {
	return showHelp
}
