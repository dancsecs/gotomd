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

/*
# Usage: gotomd

Synchronize GitHub README.md files with Go source code,
documentation, tests, and command output. gotomd processes
Markdown templates or existing README files, replacing special
directives with content generated directly from your Go
codebase. This ensures your documentation is always accurate
and in sync with the source.

	gotomd [-v | --verbose ...] [--quiet] [--log <level | (levels)>]
	       [--language <lang>] [--long-labels] [-c | --clean]
	       [-r | --replace] [-l | --license] [-h | --help] [-f | --force]
	       [-z | --colorize] [-o | --output <dir>]
	       [-u | --usage <filename>] [-p | --permission <perm>] [path ...]

	[-v | --verbose ...]
	    Increase the verbose level for each v provided.


	[--quiet]
	    Sets the verbose level to -1 squashing all (non-logged) output.


	[--log <level | (levels)>]
	    Set the level to log (or a custom combination of levels).


	[--language <lang>]
	    Sets the local language used for formatting.


	[--long-labels]
	    Use long labels in log output.


	[-c | --clean]
	    Reverse operation and remove generated markdown

	    (Cannot be used with the [-r | --replace] option).


	[-r | --replace]
	    Replace the *.MD in place

	    (Cannot be used with the [-c | --clean] option).


	[-l | --license]
	    Display license before program exits.


	[-h | --help]
	    Display program usage information.


	[-f | --force]
	    Do not confirm overwrite of destination.


	[-z | --colorize]
	    Colorize go test output.


	[-o | --output <dir>]
	    Direct all output to the specified directory.


	[-u | --usage <filename>]
	    Replace the usage section in the given Go source file using
	    content from standard input.  The section is identified as the
	    text between the first occurrence of '^\n/*\n# Usage .*$' and the
	    following package declaration.  This allows keeping command-line
	    usage output (e.g., from --help) synchronized with the package
	    documentation.


	[-p | --permission <perm>]
	    Permissions to use when creating new file.

	    (can only set RW bits)


	[path ...]
	    A specific gotomd file template with the extension '*.gtm.md' or a
	    directory which will be searched for all matching template
	    '*.gtm.md' files.  It defaults to the current directory: '.'
*/
package main
