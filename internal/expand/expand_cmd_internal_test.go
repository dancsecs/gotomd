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

package expand

import (
	"strings"
	"testing"

	"github.com/dancsecs/gotomd/internal/format"
	"github.com/dancsecs/gotomd/internal/gopkg"
	"github.com/dancsecs/sztestlog"
)

func TestInternalExpand_ExpandCmd_InlineRun_ForGoDoc(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	gopkg.Reset()
	format.ForGoDoc()

	lines := []string{
		"line:0",
		"<!--- gotomd::irun::./testdata/tstcmd/. --help -->",
		"line:2",
	}

	i := 1
	cmdIndex, cmdLength, err := isCmd(lines[i])
	chk.NoErr(err)

	block, i, err := expandCmd(i, cmdIndex, cmdLength, lines)

	chk.NoErr(err)
	chk.Int(i, 1)
	chk.StrSlice(
		strings.Split(block, "\n"),
		[]string{
			"\t# Usage",
			"",
			"\ttstcmd [--help]",
			"",
			"\t    --help",
			"\t        Display usage information",
		},
	)
}

func TestInternalExpand_ExpandCmd_InlineRun_ForMarkdown(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	gopkg.Reset()
	format.ForMarkdown()

	lines := []string{
		"line:0",
		"<!--- gotomd::irun::./testdata/tstcmd/. --help -->",
		"line:2",
	}

	i := 1
	cmdIndex, cmdLength, err := isCmd(lines[i])
	chk.NoErr(err)

	block, i, err := expandCmd(i, cmdIndex, cmdLength, lines)

	chk.NoErr(err)
	chk.Int(i, 1)
	chk.StrSlice(
		strings.Split(block, "\n"),
		[]string{
			"```",
			"# Usage",
			"",
			"tstcmd [--help]",
			"",
			"\t--help",
			"\t    Display usage information",
			"```",
		},
	)
}

func TestInternalExpand_ExpandCmd_Snippet_ForGoDoc(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	gopkg.Reset()
	format.ForGoDoc()

	lines := []string{
		"line:0",
		"<!--- gotomd::snip::./testdata/tstpkg/.sharedTemplate.sds.md",
		" # START SNIPPET -->",
		"line:3",
	}

	i := 1
	cmdIndex, cmdLength, err := isCmd(lines[i])
	chk.NoErr(err)

	block, i, err := expandCmd(
		i,
		cmdIndex,
		cmdLength,
		lines,
	)

	chk.NoErr(err)
	chk.Int(i, 2)
	chk.StrSlice(
		strings.Split(block, "\n"),
		[]string{
			"# Common Snippet Inclusion",
			"",
			"\t#!/bin/bash",
			"\techo \"Hello, world.\"",
		},
	)
}

func TestInternalExpand_ExpandCmd_Snippet_ForMarkdown(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	gopkg.Reset()
	format.ForMarkdown()

	lines := []string{
		"line:0",
		"<!--- gotomd::snip::./testdata/tstpkg/.sharedTemplate.sds.md",
		" # START SNIPPET",
		"-->",
		"line:4",
	}

	i := 1
	cmdIndex, cmdStart, err := isCmd(lines[i])
	chk.NoErr(err)

	block, i, err := expandCmd(
		i,
		cmdIndex,
		cmdStart,
		lines,
	)

	chk.NoErr(err)
	chk.Int(i, 3)
	chk.StrSlice(
		strings.Split(block, "\n"),
		[]string{
			"# Common Snippet Inclusion",
			"",
			"```bash",
			"#!/bin/bash",
			"echo \"Hello, world.\"",
			"```",
		},
	)
}

func TestInternalExpand_ExpandCmd_DocPackage_ForGoDoc(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	gopkg.Reset()
	format.ForGoDoc()

	lines := []string{
		"line:0",
		"<!--- gotomd::doc::./testdata/tstpkg/package -->",
		"line:2",
	}

	i := 1
	cmdIndex, cmdLength, err := isCmd(lines[i])
	chk.NoErr(err)

	block, i, err := expandCmd(i, cmdIndex, cmdLength, lines)

	chk.NoErr(err)
	chk.Int(i, 1)
	chk.StrSlice(
		strings.Split(block, "\n"),
		[]string{
			"\tpackage example1",
			"",
			"Package example1 demonstrates various template options.",
			"",
			"# MarkDown Headings can be used in go docs.",
			"",
			"# Markdown code formatting may be used in go doc templates.",
			"",
			"It will be translated to go doc format (tabbed) when processed.",
			"",
			"\t#!/bin/bash",
			"\techo \"Hello, world.\"",
			"",
			"# Include (and expand) Shared Snippet From .doc.gtm.go Template",
			"# Common Snippet Inclusion",
			"",
			"\t#!/bin/bash",
			"\techo \"Hello, world.\"",
		},
	)

	chk.Stdout(
		"Loading package info for: ./testdata/tstpkg",
		"getInfo(\"package\")",
	)
}

func TestInternalExpand_ExpandCmd_DocPackage_ForMarkdown(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	gopkg.Reset()
	format.ForMarkdown()

	lines := []string{
		"line:0",
		"<!--- gotomd::doc::./testdata/tstpkg/package -->",
		"line:2",
	}

	i := 1
	cmdIndex, cmdLength, err := isCmd(lines[i])
	chk.NoErr(err)

	block, i, err := expandCmd(i, cmdIndex, cmdLength, lines)

	chk.NoErr(err)
	chk.Int(i, 1)
	chk.StrSlice(
		strings.Split(block, "\n"),
		[]string{
			"```go",
			"package example1",
			"```",
			"",
			"Package example1 demonstrates various template options.",
			"",
			"# MarkDown Headings can be used in go docs.",
			"",
			"# Markdown code formatting may be used in go doc templates.",
			"",
			"It will be translated to go doc format (tabbed) when processed.",
			"",
			"\t#!/bin/bash",
			"\techo \"Hello, world.\"",
			"",
			"# Include (and expand) Shared Snippet From .doc.gtm.go Template",
			"# Common Snippet Inclusion",
			"",
			"\t#!/bin/bash",
			"\techo \"Hello, world.\"",
		},
	)
	chk.Stdout(
		"Loading package info for: ./testdata/tstpkg",
		"getInfo(\"package\")",
	)
}

func TestInternalExpand_ExpandCmd_Declaration_ForGoDoc(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	gopkg.Reset()
	format.ForGoDoc()

	lines := []string{
		"line:0",
		"<!--- gotomd::dcl::./testdata/tstpkg/ConstGroupType -->",
		"line:2",
	}

	i := 1
	cmdIndex, cmdLength, err := isCmd(lines[i])
	chk.NoErr(err)

	block, i, err := expandCmd(i, cmdIndex, cmdLength, lines)

	chk.NoErr(err)
	chk.Int(i, 1)
	chk.StrSlice(
		strings.Split(block, "\n"),
		[]string{
			"\ttype ConstGroupType int",
		},
	)

	chk.Stdout(
		"Loading package info for: ./testdata/tstpkg",
		"getInfo(\"ConstGroupType\")",
	)
}

func TestInternalExpand_ExpandCmd_Declaration_ForMarkdown(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	gopkg.Reset()
	format.ForMarkdown()

	lines := []string{
		"line:0",
		"<!--- gotomd::dcl::./testdata/tstpkg/ConstGroupType -->",
		"line:2",
	}

	i := 1
	cmdIndex, cmdLength, err := isCmd(lines[i])
	chk.NoErr(err)

	block, i, err := expandCmd(i, cmdIndex, cmdLength, lines)

	chk.NoErr(err)
	chk.Int(i, 1)
	chk.StrSlice(
		strings.Split(block, "\n"),
		[]string{
			"```go",
			"type ConstGroupType int",
			"```",
		},
	)
	chk.Stdout(
		"Loading package info for: ./testdata/tstpkg",
		"getInfo(\"ConstGroupType\")",
	)
}

func TestInternalExpand_ExpandCmd_DeclarationGroup_ForGoDoc(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	gopkg.Reset()
	format.ForGoDoc()

	lines := []string{
		"line:0",
		"<!--- gotomd::dclg::./testdata/tstpkg/ConstantGroupA -->",
		"line:2",
	}

	i := 1
	cmdIndex, cmdLength, err := isCmd(lines[i])
	chk.NoErr(err)

	block, i, err := expandCmd(i, cmdIndex, cmdLength, lines)

	chk.NoErr(err)
	chk.Int(i, 1)
	chk.StrSlice(
		strings.Split(block, "\n"),
		[]string{
			"\t// Here is a second constant block.  " +
				"All constants are reported as a group.",
			"\tconst (",
			"\t    // ConstantGroupA is a constant defined in a group.",
			"\t    ConstantGroupA = \"constant A\"",
			"",
			"\t    // ConstantGroupB is a constant defined in a group.",
			"\t    ConstantGroupB = \"constant B\"",
			"\t)",
		},
	)

	chk.Stdout(
		"Loading package info for: ./testdata/tstpkg",
		"getInfo(\"ConstantGroupA\")",
	)
}

func TestInternalExpand_ExpandCmd_DeclarationGroup_ForMarkdown(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	gopkg.Reset()
	format.ForMarkdown()

	lines := []string{
		"line:0",
		"<!--- gotomd::dclg::./testdata/tstpkg/ConstantGroupA -->",
		"line:2",
	}

	i := 1
	cmdIndex, cmdLength, err := isCmd(lines[i])
	chk.NoErr(err)

	block, i, err := expandCmd(i, cmdIndex, cmdLength, lines)

	chk.NoErr(err)
	chk.Int(i, 1)
	chk.StrSlice(
		strings.Split(block, "\n"),
		[]string{
			"```go",
			"// Here is a second constant block.  " +
				"All constants are reported as a group.",
			"const (",
			"    // ConstantGroupA is a constant defined in a group.",
			"    ConstantGroupA = \"constant A\"",
			"",
			"    // ConstantGroupB is a constant defined in a group.",
			"    ConstantGroupB = \"constant B\"",
			")",
			"```",
		},
	)
	chk.Stdout(
		"Loading package info for: ./testdata/tstpkg",
		"getInfo(\"ConstantGroupA\")",
	)
}

func TestInternalExpand_ExpandCmd_DeclarationNatural_ForGoDoc(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	gopkg.Reset()
	format.ForGoDoc()

	lines := []string{
		"line:0",
		"<!--- gotomd::dcln::./testdata/tstpkg/StructureType.GetF1 -->",
		"line:2",
	}

	i := 1
	cmdIndex, cmdLength, err := isCmd(lines[i])
	chk.NoErr(err)

	block, i, err := expandCmd(i, cmdIndex, cmdLength, lines)

	chk.NoErr(err)
	chk.Int(i, 1)
	chk.StrSlice(
		strings.Split(block, "\n"),
		[]string{
			"\t// GetF1 is a method to a structure.",
			"\tfunc (s *StructureType) GetF1(a, b, c int) string",
		},
	)

	chk.Stdout(
		"Loading package info for: ./testdata/tstpkg",
		"getInfo(\"StructureType.GetF1\")",
	)
}

func TestInternalExpand_ExpandCmd_DeclarationNatural_ForMarkdown(
	t *testing.T,
) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	gopkg.Reset()
	format.ForMarkdown()

	lines := []string{
		"line:0",
		"<!--- gotomd::dcln::./testdata/tstpkg/StructureType.GetF1 -->",
		"line:2",
	}

	i := 1
	cmdIndex, cmdLength, err := isCmd(lines[i])
	chk.NoErr(err)

	block, i, err := expandCmd(i, cmdIndex, cmdLength, lines)

	chk.NoErr(err)
	chk.Int(i, 1)
	chk.StrSlice(
		strings.Split(block, "\n"),
		[]string{
			"```go",
			"// GetF1 is a method to a structure.",
			"func (s *StructureType) GetF1(a, b, c int) string",
			"```",
		},
	)

	chk.Stdout(
		"Loading package info for: ./testdata/tstpkg",
		"getInfo(\"StructureType.GetF1\")",
	)
}

func TestInternalExpand_ExpandCmd_DeclarationSingle_ForGoDoc(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	gopkg.Reset()
	format.ForGoDoc()

	lines := []string{
		"line:0",
		"<!--- gotomd::dcls::./testdata/tstpkg/StructureType.GetF1 -->",
		"line:2",
	}

	i := 1
	cmdIndex, cmdLength, err := isCmd(lines[i])
	chk.NoErr(err)

	block, i, err := expandCmd(i, cmdIndex, cmdLength, lines)

	chk.NoErr(err)
	chk.Int(i, 1)
	chk.StrSlice(
		strings.Split(block, "\n"),
		[]string{
			"\tfunc (s *StructureType) GetF1(a, b, c int) string",
		},
	)

	chk.Stdout(
		"Loading package info for: ./testdata/tstpkg",
		"getInfo(\"StructureType.GetF1\")",
	)
}

func TestInternalExpand_ExpandCmd_DeclarationSingle_ForMarkdown(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	gopkg.Reset()
	format.ForMarkdown()

	lines := []string{
		"line:0",
		"<!--- gotomd::dcls::./testdata/tstpkg/StructureType.GetF1 -->",
		"line:2",
	}

	i := 1
	cmdIndex, cmdLength, err := isCmd(lines[i])
	chk.NoErr(err)

	block, i, err := expandCmd(i, cmdIndex, cmdLength, lines)

	chk.NoErr(err)
	chk.Int(i, 1)
	chk.StrSlice(
		strings.Split(block, "\n"),
		[]string{
			"```go",
			"func (s *StructureType) GetF1(a, b, c int) string",
			"```",
		},
	)

	chk.Stdout(
		"Loading package info for: ./testdata/tstpkg",
		"getInfo(\"StructureType.GetF1\")",
	)
}

func TestInternalExpand_ExpandCmd_Run_ForGoDoc(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	gopkg.Reset()
	format.ForGoDoc()

	lines := []string{
		"line:0",
		"<!--- gotomd::run::./testdata/tstcmd/. abc -->",
		"line:2",
	}

	i := 1
	cmdIndex, cmdLength, err := isCmd(lines[i])
	chk.NoErr(err)

	block, i, err := expandCmd(i, cmdIndex, cmdLength, lines)

	chk.NoErr(err)
	chk.Int(i, 1)
	chk.StrSlice(
		strings.Split(block, "\n"),
		[]string{
			strings.Repeat("-", 78),
			"\tgo run ./testdata/tstcmd abc",
			"",
			"\tRun complete.",
			strings.Repeat("-", 78),
		},
	)
}

func TestInternalExpand_ExpandCmd_Run_ForMarkdown(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	gopkg.Reset()
	format.ForMarkdown()

	lines := []string{
		"line:0",
		"<!--- gotomd::run::./testdata/tstcmd/. abc -->",
		"line:2",
	}

	i := 1
	cmdIndex, cmdLength, err := isCmd(lines[i])
	chk.NoErr(err)

	block, i, err := expandCmd(i, cmdIndex, cmdLength, lines)

	chk.NoErr(err)
	chk.Int(i, 1)
	chk.StrSlice(
		strings.Split(block, "\n"),
		[]string{
			"---",
			"```bash",
			"go run ./testdata/tstcmd abc",
			"```",
			"",
			"```",
			"Run complete.",
			"```",
			"---",
		},
	)
}

func TestInternalExpand_ExpandCmd_Irun_ForGoDoc(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	gopkg.Reset()
	format.ForGoDoc()

	lines := []string{
		"line:0",
		"<!--- gotomd::irun::./testdata/tstcmd/. abc -->",
		"line:2",
	}

	i := 1
	cmdIndex, cmdLength, err := isCmd(lines[i])
	chk.NoErr(err)

	block, i, err := expandCmd(i, cmdIndex, cmdLength, lines)

	chk.NoErr(err)
	chk.Int(i, 1)
	chk.StrSlice(
		strings.Split(block, "\n"),
		[]string{
			"\tRun complete.",
		},
	)
}

func TestInternalExpand_ExpandCmd_Irun_ForMarkdown(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	gopkg.Reset()
	format.ForMarkdown()

	lines := []string{
		"line:0",
		"<!--- gotomd::irun::./testdata/tstcmd/. abc -->",
		"line:2",
	}

	i := 1
	cmdIndex, cmdLength, err := isCmd(lines[i])
	chk.NoErr(err)

	block, i, err := expandCmd(i, cmdIndex, cmdLength, lines)

	chk.NoErr(err)
	chk.Int(i, 1)
	chk.StrSlice(
		strings.Split(block, "\n"),
		[]string{
			"```",
			"Run complete.",
			"```",
		},
	)
}

func TestInternalExpand_ExpandCmd_Src_ForGoDoc(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	gopkg.Reset()
	format.ForGoDoc()

	lines := []string{
		"line:0",
		"<!--- gotomd::src::./testdata/tstcmd/crumb.go -->",
		"line:2",
	}

	i := 1
	cmdIndex, cmdLength, err := isCmd(lines[i])
	chk.NoErr(err)

	block, i, err := expandCmd(i, cmdIndex, cmdLength, lines)

	chk.NoErr(err)
	chk.Int(i, 1)
	chk.StrSlice(
		strings.Split(block, "\n"),
		[]string{
			"\tcat ./testdata/tstcmd/crumb.go",
			"",
			"\t// This is an empty crumb file.",
			"",
			"\tpackage main",
		},
	)
}

func TestInternalExpand_ExpandCmd_Src_ForMarkdown(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	gopkg.Reset()
	format.ForMarkdown()

	lines := []string{
		"line:0",
		"<!--- gotomd::src::./testdata/tstcmd/crumb.go -->",
		"line:2",
	}

	i := 1
	cmdIndex, cmdLength, err := isCmd(lines[i])
	chk.NoErr(err)

	block, i, err := expandCmd(i, cmdIndex, cmdLength, lines)

	chk.NoErr(err)
	chk.Int(i, 1)
	chk.StrSlice(
		strings.Split(block, "\n"),
		[]string{
			"```bash",
			"cat ./testdata/tstcmd/crumb.go",
			"```",
			"",
			"```go",
			"// This is an empty crumb file.",
			"",
			"package main",
			"```",
		},
	)
}

func TestInternalExpand_ExpandCmd_Snip_ForGoDoc(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	gopkg.Reset()
	format.ForGoDoc()

	lines := []string{
		"line:0",
		"<!--- gotomd::snip::./testdata/tstpkg/.sharedTemplate.sds.md " +
			"# START SNIPPET -->",
		"line:2",
	}

	i := 1
	cmdIndex, cmdLength, err := isCmd(lines[i])
	chk.NoErr(err)

	block, i, err := expandCmd(i, cmdIndex, cmdLength, lines)

	chk.NoErr(err)
	chk.Int(i, 1)
	chk.StrSlice(
		strings.Split(block, "\n"),
		[]string{
			"# Common Snippet Inclusion",
			"",
			"\t#!/bin/bash",
			"\techo \"Hello, world.\"",
		},
	)
}

func TestInternalExpand_ExpandCmd_Snip_ForMarkdown(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	gopkg.Reset()
	format.ForMarkdown()

	lines := []string{
		"line:0",
		"<!--- gotomd::snip::./testdata/tstpkg/.sharedTemplate.sds.md " +
			"# START SNIPPET -->",
		"line:2",
	}

	i := 1
	cmdIndex, cmdLength, err := isCmd(lines[i])
	chk.NoErr(err)

	block, i, err := expandCmd(i, cmdIndex, cmdLength, lines)

	chk.NoErr(err)
	chk.Int(i, 1)
	chk.StrSlice(
		strings.Split(block, "\n"),
		[]string{
			"# Common Snippet Inclusion",
			"",
			"```bash",
			"#!/bin/bash",
			"echo \"Hello, world.\"",
			"```",
		},
	)
}

// Change <PRE> to ``` if it does impact colorize.
// func TestInternalExpand_ExpandCmd_Tst_ForGoDoc(t *testing.T) {
// 	chk := sztestlog.CaptureNothing(t)
// 	defer chk.Release()

// 	gopkg.Reset()
// 	format.ForGoDoc()

// 	lines := []string{
// 		"line:0",
// 		"<!--- gotomd::tst::./testdata/tstcmd/. -->",
// 		"line:2",
// 	}

// 	i := 1
// 	cmdIndex, cmdLength, err := isCmd(lines[i])
// 	chk.NoErr(err)

// 	block, i, err := expandCmd(i, cmdIndex, cmdLength, lines)

// 	chk.NoErr(err)
// 	chk.Int(i, 1)
// 	chk.StrSlice(
// 		strings.Split(block, "\n"),
// 		[]string{
// 			"\tRun complete.",
// 		},
// 	)
// }

// func TestInternalExpand_ExpandCmd_Tst_ForMarkdown(t *testing.T) {
// 	chk := sztestlog.CaptureNothing(t)
// 	defer chk.Release()

// 	gopkg.Reset()
// 	format.ForMarkdown()

// 	lines := []string{
// 		"line:0",
// 		"<!--- gotomd::tst::./testdata/tstcmd/. -->",
// 		"line:2",
// 	}

// 	i := 1
// 	cmdIndex, cmdLength, err := isCmd(lines[i])
// 	chk.NoErr(err)

// 	block, i, err := expandCmd(i, cmdIndex, cmdLength, lines)

// 	chk.NoErr(err)
// 	chk.Int(i, 1)
// 	chk.StrSlice(
// 		strings.Split(block, "\n"),
// 		[]string{
// 			"```",
// 			"Run complete.",
// 			"```",
// 		},
// 	)
// }
