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

Synchronize Go package and GitHub style README.md documentation files by
embedding Go documentation, source code, test and command output directly from
the Go codebase. This ensures that program documentation is kept in one
place—the Go code itself—while keeping the README and package documentation
automatically up to date. It does this by processing template files containing
markdown formatting and replacing embedded directives with content generated
directly from your Go codebase. This ensures your documentation is always
accurate and in sync with the source.

    gotomd [-v | --verbose ...] [-l | --license] [-h | --help]
           [-f | --force] [-z | --colorize] [-o | --output <dir>]
           [-p | --permission <perm>] --uptodate [path ...]

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


    --uptodate
        Returns 0 if all template expansions are up to date.  No changes are
        made.  It is not compatible with the flags:-f, -o <dir>, -p <perm>.


    [path ...]
        Specific template files (named like '.*.gtm.md' or '.*.gtm.go') or a
        directory which will be searched for all matching template files.  All
		subdirectories may be searched by using the special './...' path.
        It defaults to search the current directory: '.'



# Directives

gotomd processes template files (.*.gtm.md and .*.gtm.go) into their
respective *.md and *.go expanding included directives.

Directives are written inside HTML-style comments:

```html
<!--- gotomd::ACTION::PARAMETERS -->
```
When processing the file, gotomd replaces each directive with the
corresponding generated content.

## Actions

Each directive's ACTION determines what is included:

```html
<!--- gotomd::doc::./relativeDirectory/goObject -->
```
Runs go doc on the specified object from the given directory.

The parameter format is: ./path/to/package/ObjectName

A special object name 'package' includes the package-level comments.

```html
<!--- gotomd::docConstGrp::./relativeDirectory/goConstName ListOfConstNames -->
```
Runs go doc on the specified constant block(s) from the given directory.

```html
<!--- gotomd::dcls::./relativeDirectory/declaredObject ListOfDeclaredGoObjects -->
```
Inserts each listed declaration as a single line, regardless of how it is
declared in the source. No comments are included.

Example: functions, methods, constants.

```html
<!---
gotomd::dcln::./relativeDirectory/declaredObject ListOfDeclaredGoObjects
-->
```

Inserts each listed declaration exactly as in the source, including any
leading comments.

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

Runs go test in the given directory, targeting the specified test(s) or
package, and includes the output.

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

Runs go run on the package in the given directory (assumes main) with the
provided arguments, including the output.

<!--- gotomd::inline-run::./. --help -->

# Dedication

This project is dedicated to Reem. Your brilliance, courage, and quiet
strength continue to inspire me. Every line is written in gratitude for the
light and hope you brought into my life.

NOTE: Documentation reviewed and polished with the assistance of ChatGPT from
OpenAI.
*/
package main

import (
	"strings"

	"github.com/dancsecs/szlog"
)

const copyrightMessage = `
//<<<! gotomd::file::~/.copyright >>
`

// Copyright writes the copyright message to os.Stdout.
func Copyright() {
	szlog.Say0(strings.Trim(copyrightMessage, " \t\n") + "\n")
}
