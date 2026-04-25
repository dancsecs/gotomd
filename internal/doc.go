//                *****  AUTO GENERATED:  DO NOT MODIFY  *****.
//                   MODIFY TEMPLATE: 'internal/.doc.gtm.go'.
//                  See: 'https://github.com/dancsecs/gotomd'.

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

/*
Package internal implements the command's mainline `Main`.
*/
//
//nolint:lll // Ok.
package internal

// LicenseCopyright provides for program access.
const LicenseCopyright = "" +
	"Golang To Github Markdown: gotomd." + "\n" +
	"Copyright (C) 2023-2026 Leslie Dancsecs" + "\n" +
	"" + "\n" +
	"This program is free software: you can redistribute it and/or modify" + "\n" +
	"it under the terms of the GNU General Public License as published by" + "\n" +
	"the Free Software Foundation, either version 3 of the License, or" + "\n" +
	"(at your option) any later version." + "\n" +
	"" + "\n" +
	"This program is distributed in the hope that it will be useful," + "\n" +
	"but WITHOUT ANY WARRANTY; without even the implied warranty of" + "\n" +
	"MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the" + "\n" +
	"GNU General Public License for more details." + "\n" +
	"" + "\n" +
	"You should have received a copy of the GNU General Public License" + "\n" +
	"along with this program.  If not, see <https://www.gnu.org/licenses/>." + "\n" +
	""

// DirectiveHowTo provides for program access.
const DirectiveHowTo = "" +
	"# Directives" + "\n" +
	"" + "\n" +
	"The gotomd utility processes template files (`*.gtm.md` and `*.gtm.go`)" + "\n" +
	"into their corresponding `*.md` and `*.go` files, expanding any embedded" + "\n" +
	"directives." + "\n" +
	"" + "\n" +
	"Directives are written inside HTML-style comments:" + "\n" +
	"" + "\n" +
	"\t<!--- gotomd::ACTION::OBJECT [OPTIONAL ...] -->" + "\n" +
	"" + "\n" +
	"They may also span multiple lines:" + "\n" +
	"" + "\n" +
	"\t<!--- gotomd::ACTION::OBJECT [OPTIONAL...]" + "\n" +
	"\t[OPTIONAL ...] -->" + "\n" +
	"" + "\n" +
	"\t<!--- gotomd::ACTION::OBJECT [OPTIONAL...]" + "\n" +
	"\t[OPTIONAL ...]" + "\n" +
	"\t..." + "\n" +
	"\t-->" + "\n" +
	"" + "\n" +
	"The `OPTIONAL` elements may be additional objects or parameters, depending on" + "\n" +
	"the `ACTION` being used." + "\n" +
	"" + "\n" +
	"When processing the file, `gotomd` replaces each directive with the generated" + "\n" +
	"content corresponding to that directive." + "\n" +
	"" + "\n" +
	"## Actions" + "\n" +
	"" + "\n" +
	"Available actions are:" + "\n" +
	"" + "\n" +
	"   - `doc`   runs and embeds output from `go doc` for a package object" + "\n" +
	"   - `dcl`   inserts the declaration of package objects" + "\n" +
	"   - `dclg`  inserts the declaration group for package objects (IE `const` blocks)" + "\n" +
	"   - `dcln`  inserts the declaration exactly as defined in source including comments" + "\n" +
	"   - `dcls`  inserts the declaration formatted as a single line" + "\n" +
	"   - `irun`  runs the package and inserts the output without decorations" + "\n" +
	"   - `run`   runs the package and frames the output with the command executed" + "\n" +
	"   - `snip`  includes an external snippet expanding any embedded directives" + "\n" +
	"   - `src`   includes a Go source file" + "\n" +
	"   - `tst`   runs a Go test (or all tests) in a package" + "\n" +
	"   - `tstc`  Runs a go test (or all tests) and converts output to TeX to preserve formatting" + "\n" +
	"" + "\n" +
	"### Action: doc" + "\n" +
	"" + "\n" +
	"Runs `go doc` on the specified object in the given relative package" + "\n" +
	"directory." + "\n" +
	"" + "\n" +
	"The required argument is the relative package path followed by the name of the" + "\n" +
	"object to document." + "\n" +
	"" + "\n" +
	"A special object name, `package`, includes the package-level comments." + "\n" +
	"" + "\n" +
	"Additional objects may be specified as optional arguments, with or without a" + "\n" +
	"relative directory. If no directory is provided, the most recently specified" + "\n" +
	"directory is used." + "\n" +
	"" + "\n" +
	"\t<!--- gotomd::doc::./directory/goObject -->" + "\n" +
	"" + "\n" +
	"\t<!--- gotomd::doc::./directory/goObject1 goObject2 -->" + "\n" +
	"" + "\n" +
	"\t<!--- gotomd::doc::./directory/goObject1 goObject2" + "\n" +
	"\t./differentDirectory/goObject3 goObject4" + "\n" +
	"\t./anotherDifferentDirectory/package" + "\n" +
	"\t-->" + "\n" +
	"" + "\n" +
	"There are four additional directives all similar to `doc` but focused on object" + "\n" +
	"declaration formatting as follows:" + "\n" +
	"" + "\n" +
	"   | directive | formatting              | comments | typical use                     |" + "\n" +
	"   | --------- | ----------------------- | -------: | ------------------------------- |" + "\n" +
	"   | `dcl`     | preserves source layout |       no | show declaration cleanly        |" + "\n" +
	"   | `dclg`    | exact source layout     |      yes | show as written in group block  |" + "\n" +
	"   | `dcln`    | exact source layout     |      yes | show source as written          |" + "\n" +
	"   | `dcls`    | single line             |       no | compact summaries / inline docs |" + "\n" +
	"" + "\n" +
	"See individual Action sections for more detail." + "\n" +
	"" + "\n" +
	"### Action: dcl" + "\n" +
	"" + "\n" +
	"Similar to the `doc` directive, `dcl` inserts the declaration of the specified" + "\n" +
	"object as formatted in the source file, excluding comments." + "\n" +
	"" + "\n" +
	"This preserves the original multi-line source layout." + "\n" +
	"" + "\n" +
	"\t<!--- gotomd::dcl::./directory/goObject" + "\n" +
	"\t[[./directory/]goObject...]" + "\n" +
	"\t..." + "\n" +
	"\t-->" + "\n" +
	"" + "\n" +
	"### Action: dclg" + "\n" +
	"" + "\n" +
	"Similar to the `doc` directive, `dclg` inserts the grouped declaration group" + "\n" +
	"containing the specified object exactly as formatted in the source files," + "\n" +
	"including comments." + "\n" +
	"" + "\n" +
	"This is limited to grouped `const (...)` and `var (...)` blocks." + "\n" +
	"" + "\n" +
	"\t<!--- gotomd::dclg::./directory/goObject" + "\n" +
	"\t[[./directory/]goObject...]" + "\n" +
	"\t..." + "\n" +
	"\t-->" + "\n" +
	"" + "\n" +
	"### Action: dcln" + "\n" +
	"" + "\n" +
	"Similar to the `doc` directive, `dcln` inserts the declaration exactly as it" + "\n" +
	"appears in the source file, including all associated comments." + "\n" +
	"" + "\n" +
	"This is the closest representation of the original source code." + "\n" +
	"" + "\n" +
	"\t<!--- gotomd::dcln::./directory/goObject" + "\n" +
	"\t[[./directory/]goObject..." + "\n" +
	"\t..." + "\n" +
	"\t-->" + "\n" +
	"" + "\n" +
	"### Action: dcls" + "\n" +
	"" + "\n" +
	"Similar to the `doc` directive, `dcls` inserts the declaration of the" + "\n" +
	"specified object reformatted onto a single line." + "\n" +
	"" + "\n" +
	"Comments are not included." + "\n" +
	"" + "\n" +
	"\t<!--- gotomd::dcls::./directory/goObject" + "\n" +
	"\t[[./directory/]goObject..." + "\n" +
	"\t..." + "\n" +
	"\t-->" + "\n" +
	"" + "\n" +
	"### Action: irun" + "\n" +
	"" + "\n" +
	"Runs `go run` on the package in the specified directory (assumes `main`) with" + "\n" +
	"the provided arguments." + "\n" +
	"" + "\n" +
	"The output is embedded as preformatted text without decorations." + "\n" +
	"" + "\n" +
	"\t<!--- gotomd::irun::./directory/. [args ...] -->" + "\n" +
	"" + "\n" +
	"### Action: run" + "\n" +
	"" + "\n" +
	"Runs `go run` on the package in the specified directory (assumes `main`) with" + "\n" +
	"the provided arguments." + "\n" +
	"" + "\n" +
	"The output is framed together with the command that was executed." + "\n" +
	"" + "\n" +
	"\t<!--- gotomd::run::./directory/. [args ...] -->" + "\n" +
	"" + "\n" +
	"### Action: snip" + "\n" +
	"" + "\n" +
	"Loads the referenced snippet and expands any embedded directives." + "\n" +
	"" + "\n" +
	"If the optional first parameter (the word `string`) is present then the output" + "\n" +
	"will be a series of concatenated quote terminated escaped strings suitable of" + "\n" +
	"embedding in code." + "\n" +
	"" + "\n" +
	"If the optional [`startAfter`] argument is supplied, only content appearing" + "\n" +
	"after the first line matching `startAfter` is included." + "\n" +
	"" + "\n" +
	"\t<!--- gotomd::snip::./directory/fileName [string ][startAfter] -->" + "\n" +
	"" + "\n" +
	"### Action: src" + "\n" +
	"" + "\n" +
	"Inserts the contents of the specified Go source file, formatted as Go code." + "\n" +
	"" + "\n" +
	"\t<!--- gotomd::src::./directory/fileName.go -->" + "\n" +
	"" + "\n" +
	"### Action: tst" + "\n" +
	"" + "\n" +
	"Runs the specified Go test." + "\n" +
	"" + "\n" +
	"A `.` represents all tests." + "\n" +
	"" + "\n" +
	"\t<!--- gotomd::tst::./directory/testName -->" + "\n" +
	"" + "\n" +
	"\t<!--- gotomd::tst::./directory/. -->" + "\n" +
	"" + "\n" +
	"### Action: tstc" + "\n" +
	"" + "\n" +
	"Runs the specified Go test." + "\n" +
	"" + "\n" +
	"A `.` represents all tests." + "\n" +
	"" + "\n" +
	"The output is converted to TeX format to preserve text decorations when" + "\n" +
	"displayed on the GitHub website." + "\n" +
	"" + "\n" +
	"\t<!--- gotomd::tstc::./directory/testName -->" + "\n" +
	"" + "\n" +
	"\t<!--- gotomd::tstc::./directory/. -->" + "\n" +
	""
