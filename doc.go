//                *****  AUTO GENERATED:  DO NOT MODIFY  *****.
//                       MODIFY TEMPLATE: '.doc.gtm.go'.
//                  See: 'https://github.com/dancsecs/gotomd'.

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

//nolint:lll // Ok.
/*
	usage: gotomd [-v | --verbose ...] [-d | --directive] [-l | --license]
	              [-h | --help] [-f | --force] [-u | --uptodate]
	              [-o | --output <dir>] [-p | --permission <perm>] [path ...]

	Synchronize Go package and GitHub style README.md documentation by embedding
	Go documentation, source code, test and command output directly from the Go
	codebase. This ensures that program documentation is kept in one place—the
	Go code itself—while keeping the README and package documentation
	automatically up to date. It does this by processing template files
	containing markdown formatting and replacing embedded directives with
	content generated directly from your Go codebase. This ensures your
	documentation is always accurate and in sync with the source.

	    [-v | --verbose ...]
	        Increase the verbose level for each v provided.

	    [-d | --directive]
	        Display directive documentation.

	    [-l | --license]
	        Display license before program exits.

	    [-h | --help]
	        Display program usage information.

	    [-f | --force]
	        Do not confirm overwrite of destination.

	    [-u | --uptodate]
	        Returns 0 if no changes would have been made. No writes are
	        performed.

	    [-o | --output <dir>]
	        Direct all output to the specified directory.

	    [-p | --permission <perm>]
	        Permissions to use when creating new file.

	        (can only set RW bits).

	    [path ...]
	        Specific template files (named like '.*.gtm.md' or '.*.gtm.go') or a
	        directory which will be searched for all matching template files.
	        All subdirectories may be searched by using the special './...'
	        path. It defaults to search the current directory: '.'

# Directives

The gotomd utility processes template files (`*.gtm.md` and `*.gtm.go`)
into their corresponding `*.md` and `*.go` files, expanding any embedded
directives.

Directives are written inside HTML-style comments:

	<!--- gotomd::ACTION::OBJECT [OPTIONAL ...] -->

They may also span multiple lines:

	<!--- gotomd::ACTION::OBJECT [OPTIONAL...]
	   [OPTIONAL ...] -->

	<!--- gotomd::ACTION::OBJECT [OPTIONAL...]
	   [OPTIONAL ...]
	    ...
	-->

The `OPTIONAL` elements may be additional objects or parameters, depending on
the `ACTION` being used.

When processing the file, `gotomd` replaces each directive with the generated
content corresponding to that directive.

## Actions

Available actions are:

   - `doc`   runs and embeds output from `go doc` for a package object
   - `dcl`   inserts the declaration of package objects
   - `dclg`  inserts the declaration group for package objects (IE `const` blocks)
   - `dcln`  inserts the declaration exactly as defined in source including comments
   - `dcls`  inserts the declaration formatted as a single line
   - `irun`  runs the package and inserts the output without decorations
   - `run`   runs the package and frames the output with the command executed
   - `snip`  includes an external snippet expanding any embedded directives
   - `src`   includes a Go source file
   - `tst`   runs a Go test (or all tests) in a package
   - `tstc`  Runs a go test (or all tests) and converts output to TeX to preserve formatting

### Action: doc

Runs `go doc` on the specified object in the given relative package
directory.

The required argument is the relative package path followed by the name of the
object to document.

A special object name, `package`, includes the package-level comments.

Additional objects may be specified as optional arguments, with or without a
relative directory. If no directory is provided, the most recently specified
directory is used.

	<!--- gotomd::doc::./directory/goObject -->

	<!--- gotomd::doc::./directory/goObject1 goObject2 -->

	<!--- gotomd::doc::./directory/goObject1 goObject2
	   ./differentDirectory/goObject3 goObject4
	   ./anotherDifferentDirectory/package
	-->

There are four additional directives all similar to `doc` but focused on object
declaration formatting as follows:

   | directive | formatting              | comments | typical use                     |
   | --------- | ----------------------- | -------: | ------------------------------- |
   | `dcl`     | preserves source layout |       no | show declaration cleanly        |
   | `dclg`    | exact source layout     |      yes | show as written in group block  |
   | `dcln`    | exact source layout     |      yes | show source as written          |
   | `dcls`    | single line             |       no | compact summaries / inline docs |

See individual Action sections for more detail.

### Action: dcl

Similar to the `doc` directive, `dcl` inserts the declaration of the specified
object as formatted in the source file, excluding comments.

This preserves the original multi-line source layout.

	<!--- gotomd::dcl::./directory/goObject
	   [[./directory/]goObject...]
	   ...
	-->

### Action: dclg

Similar to the `doc` directive, `dclg` inserts the grouped declaration group
containing the specified object exactly as formatted in the source files,
including comments.

This is limited to grouped `const (...)` and `var (...)` blocks.

	<!--- gotomd::dclg::./directory/goObject
	   [[./directory/]goObject...]
	   ...
	-->

### Action: dcln

Similar to the `doc` directive, `dcln` inserts the declaration exactly as it
appears in the source file, including all associated comments.

This is the closest representation of the original source code.

	<!--- gotomd::dcln::./directory/goObject
	   [[./directory/]goObject...
	   ...
	-->

### Action: dcls

Similar to the `doc` directive, `dcls` inserts the declaration of the
specified object reformatted onto a single line.

Comments are not included.

	<!--- gotomd::dcls::./directory/goObject
	   [[./directory/]goObject...
	   ...
	-->

### Action: irun

Runs `go run` on the package in the specified directory (assumes `main`) with
the provided arguments.

The output is embedded as preformatted text without decorations.

	<!--- gotomd::irun::./directory/. [args ...] -->

### Action: run

Runs `go run` on the package in the specified directory (assumes `main`) with
the provided arguments.

The output is framed together with the command that was executed.

	<!--- gotomd::run::./directory/. [args ...] -->

### Action: snip

Loads the referenced snippet and expands any embedded directives.

If the optional first parameter (the word `string`) is present then the output
will be a series of concatenated quote terminated escaped strings suitable of
embedding in code.

If the optional [`startAfter`] argument is supplied, only content appearing
after the first line matching `startAfter` is included.

	<!--- gotomd::snip::./directory/fileName [string ][startAfter] -->

### Action: src

Inserts the contents of the specified Go source file, formatted as Go code.

	<!--- gotomd::src::./directory/fileName.go -->

### Action: tst

Runs the specified Go test.

A `.` represents all tests.

	<!--- gotomd::tst::./directory/testName -->

	<!--- gotomd::tst::./directory/. -->

### Action: tstc

Runs the specified Go test.

A `.` represents all tests.

The output is converted to TeX format to preserve text decorations when
displayed on the GitHub website.

	<!--- gotomd::tstc::./directory/testName -->

	<!--- gotomd::tstc::./directory/. -->

# Dedication

	***************************************************************************
	**                                                                       **
	** This project is dedicated to Reem.                                    **
	** Your brilliance, courage, and quiet strength continue to inspire me.  **
	** Every line is written in gratitude for the light and hope you brought **
	** into my life.                                                         **
	**                                                                       **
	***************************************************************************

NOTE: Documentation reviewed and polished with the assistance of ChatGPT from
OpenAI.
*/
package main
