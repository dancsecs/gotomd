/*
   Golang To Github Markdown Utility: gotomd
   Copyright (C) 2026 Leslie Dancsecs

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

const (
	programFunction = `
Synchronize Go package and GitHub style README.md documentation by
embedding Go documentation, source code, test and command output directly
from the Go codebase. This ensures that program documentation is kept in
one place—the Go code itself—while keeping the README and package
documentation automatically up to date. It does this by processing template
files containing markdown formatting and replacing embedded directives with
content generated directly from your Go codebase. This ensures your
documentation is always accurate and in sync with the source.
`
	directiveFlag = "[-d | --directive]"
	directiveDesc = `
Display directive documentation.
`

	licenseFlag = "[-l | --license]"
	licenseDesc = `
Display license before program exits.
`

	helpFlag = "[-h | --help]"
	helpDesc = `
Display program usage information.
`

	forceFlag = "[-f | --force]"
	forceDesc = `
Do not confirm overwrite of destination.
`

	outputDirFlag = "[-o | --output <dir>]"
	outputDirDesc = `
Direct all output to the specified directory.
`

	upToDateFlag = "[-u | --uptodate]"
	upToDateDesc = `
Returns 0 if no changes would have been made. No writes are performed.
`

	permFlag = "[-p | --permission <perm>]"
	permDesc = `
Permissions to use when creating new file.

(can only set RW bits).
`

	pathArg  = "[path ...]"
	pathDesc = `
Specific template files (named like '.*.gtm.md' or '.*.gtm.go') or
a directory which will be searched for all matching template files.
All subdirectories may be searched by using the special './...' path.
It defaults to search the current directory: '.'
`
)
