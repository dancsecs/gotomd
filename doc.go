/*
   Golang To Github Markdown Utility: gotomd
   Copyright (C) 2023-2025 Leslie Dancsecs

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
Package gotomd maintains GitHub-style README.md files by embedding Go
documentation, source code, test output, and command output directly
from the Go codebase. This ensures that program documentation is kept
in one place—the Go code itself—while keeping the README automatically
up to date.

## How it works

gotomd processes a Markdown file in one of two ways:

 1. **Template mode** — You maintain a separate template file named `*.gtm.md`.
 2. **In-place mode** — You insert directives directly into an existing
    `*.md` file, and gotomd replaces them in place.

Directives are written inside HTML-style comments:

```html
<!---
gotomd::ACTION::PARAMETERS
-->
```

When processing the file, gotomd replaces each directive with the
corresponding generated content, enclosed in matching "Bgn" and "End"
markers.

## Actions

Each directive's ACTION determines what is included:

```html
<!---
gotomd::doc::./relativeDirectory/goObject
-->
```
Runs go doc on the specified object from the given directory.

The parameter format is: ./path/to/package/ObjectName

A special object name package includes the package-level comments.

```html
<!---
gotomd::docConstGrp::./relativeDirectory/goConstName ListOfConstNames
-->
```
Runs go doc on the specified constant block(s) from the given directory.

```html
<!---
gotomd::dcls::./relativeDirectory/declaredObject ListOfDeclaredGoObjects
-->
```

Inserts each listed declaration as a single line, regardless of
how it is declared in the source. No comments are included.

Example: functions, methods, constants.

```html
<!---
gotomd::dcln::./relativeDirectory/declaredObject ListOfDeclaredGoObjects
-->
```

Inserts each listed declaration exactly as in the source, including
any leading comments.

```html
<!---
gotomd::dcl::./relativeDirectory/declaredObject ListOfDeclaredGoObjects
-->
```

Inserts each listed declaration exactly as in the source, but without
comments.

```html
<!---
gotomd::tst::goTest::./relativeDirectory/testName
-->
```

Runs go test in the given directory, targeting the specified test(s)
or package, and includes the output.

```html
<!---
gotomd::file::./relativeDirectory/fName
-->
```

Inserts the contents of the specified file into a fenced code block.

```html
<!---
gotomd::run::./relativeDirectory [args ...]
-->
```

Runs go run on the package in the given directory (assumes main)
with the provided arguments, including the output.

## Output Markers

Generated content is wrapped between markers in the target file:

	const sztestBgnPrefix = sztestPrefix + "Bgn::"
	const sztestEndPrefix = sztestPrefix + "End::"

Additionally, an auto-generated section header is prefixed with:

	const szAutoPrefix = sztestPrefix + "Auto::"

This header is followed by a blank line. If operating in template mode
(not in-place), a DO NOT MODIFY warning is also inserted.

# Usage: gotomd

Synchronize GitHub README.md files with Go source code,
documentation, tests, and command output. gotomd processes
Markdown templates or existing README files, replacing special
directives with content generated directly from your Go
codebase. This ensures your documentation is always accurate
and in sync with the source.

	gotomd [-v | --verbose ...] [-l | --license] [-h | --help]
	       [-f | --force] [-z | --colorize] [-o | --output <dir>]
	       [-p | --permission <perm>] [path ...]

	[-v | --verbose ...]
	    Increase the verbose level for each v provided.


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


	[-p | --permission <perm>]
	    Permissions to use when creating new file.

	    (can only set RW bits)


	[path ...]
	    A specific gotomd file template with the extension '*.gtm.md' or a
	    directory which will be searched for all matching template
	    '*.gtm.md' files.  It defaults to the current directory: '.'

# Dedication

This project is dedicated to Reem.
Your brilliance, courage, and quiet strength continue to inspire me.
Every line is written in gratitude for the light and hope you brought into my
life.

NOTE: Documentation reviewed and polished with the assistance of ChatGPT from
OpenAI.
*/
package main
