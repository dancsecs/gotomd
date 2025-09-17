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

package gopkg_test

import (
	"strings"
	"testing"

	"github.com/dancsecs/gotomd/internal/gopkg"
	"github.com/dancsecs/sztestlog"
)

type docInfoTest struct {
	action  string
	header  []string
	body    []string
	doc     []string
	oneLine string
}

func Test_GoPackage_GetInfo_InvalidDirectory(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	_, err := gopkg.Info("INVALID_DIRECTORY", "TimesTwo")
	chk.Err(
		err,
		gopkg.ErrInvalidPackage.Error(),
	)

	chk.Stdout(
		"Loading package info for: INVALID_DIRECTORY",
	)
}

func Test_GoPackage_GetInfo_InvalidObject(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	_, err := gopkg.Info("./sample1", "DOES_NOT_EXIST")
	chk.Err(err, gopkg.ErrUnknownObject.Error()+": DOES_NOT_EXIST")

	chk.Stdout(
		"Loading package info for: ./sample1",
		"getInfo(\"DOES_NOT_EXIST\")",
	)
}

//nolint:funlen // Ok.
func Test_GoPackage_DocInfo_PackageInfo(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	gopkg.Reset()

	data, err := gopkg.Info("./sample1", "package")

	chk.NoErr(err)

	chk.Str(
		data.OneLine(),
		"package sample1",
	)

	chk.Str(
		data.Declaration(),
		"package sample1",
	)

	chk.StrSlice(
		strings.Split(data.NaturalComment(), "\n"),
		[]string{
			"// Package sample1 exists in order to test various go to git",
			"// markdown (gToMD) extraction utilities.  Various object " +
				"will be defined that",
			"// exhibit the various comment and declaration options " +
				"permitted by gofmt.",
			"// ",
			"// # Heading",
			"// ",
			"// This paragraph will demonstrating further documentation " +
				"under a \"markdown\"",
			"// header.",
			"// ",
			"// Declarations can be single-line or multi-line blocks or " +
				"constructions.  Each",
			"// type will be included here for complete testing.",
		},
	)

	chk.StrSlice(
		strings.Split(data.Comment(), "\n"),
		[]string{
			"Package sample1 exists in order to test various go to git",
			"markdown (gToMD) extraction utilities.  Various object will be " +
				"defined that",
			"exhibit the various comment and declaration options permitted " +
				"by gofmt.",
			"",
			"# Heading",
			"",
			"This paragraph will demonstrating further documentation under " +
				"a \"markdown\"",
			"header.",
			"",
			"Declarations can be single-line or multi-line blocks or " +
				"constructions.  Each",
			"type will be included here for complete testing.",
		},
	)

	chk.Stdout(
		"Loading package info for: ./sample1",
		"getInfo(\"package\")",
	)
}

func Test_GoPackage_DocInfo_TypeInfo(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	data, err := gopkg.Info("./sample1", "ConstGroupType")

	chk.NoErr(err)

	chk.Str(
		data.OneLine(),
		"type ConstGroupType int",
	)

	chk.StrSlice(
		data.Doc(),
		[]string{
			"ConstGroupType set the type of the constant.",
		},
	)

	chk.Stdout(
		"getInfo(\"ConstGroupType\")",
	)
}

func Test_GoPackage_DocInfo_ConstantBlock(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	data, err := gopkg.Info("./sample1", "ConstantGroup1")

	chk.NoErr(err)

	chk.StrSlice(
		strings.Split(data.ConstantBlock(), "\n"),
		[]string{
			"// Here is a typed constant block.  All constants are " +
				"reported as a group.",
			"const (",
			"    // ConstantGroup1 is a constant defined in a group.",
			"    ConstantGroup1 ConstGroupType = iota",
			"",
			"    // ConstantGroup2 is a constant defined in a group.",
			"    ConstantGroup2",
			")",
		},
	)

	chk.Stdout(
		"getInfo(\"ConstantGroup1\")",
	)
}

//nolint:funlen,lll // Ok.
func Test_GoPackage_DocInfo_RunTests(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	docInfoTests := []docInfoTest{
		//  ----------------------------------------------------------------------
		{
			action: "TimesTwo",
			header: []string{
				"func TimesTwo(i int) int",
			},
			body: []string{
				"func TimesTwo(i int) int {",
				"    return i + i",
				"}",
			},
			doc: []string{
				"TimesTwo returns the value times two.",
			},
			oneLine: "func TimesTwo(i int) int",
		},
		//  ----------------------------------------------------------------------
		{
			action: "TimesThree",
			header: []string{
				"func TimesThree(i int) int",
			},
			body: []string{
				"func TimesThree(i int) int {",
				"    return i + i + i",
				"}",
			},
			doc: []string{
				"TimesThree returns the value times three.",
			},
			oneLine: "func TimesThree(i int) int",
		},
		//  ----------------------------------------------------------------------
		{
			action: "ConstDeclSingleCmtSingle",
			header: nil,
			body: []string{
				"const ConstDeclSingleCmtSingle = " +
					"\"single-line declaration and comment\"",
			},
			doc: []string{
				"ConstDeclSingleCmtSingle has a single-line comment.",
			},
			oneLine: "" +
				"const ConstDeclSingleCmtSingle = " +
				"\"single-line declaration and comment\"",
		},
		//  ----------------------------------------------------------------------
		{
			action: "ConstDeclMultiCmtSingle",
			header: nil,
			body: []string{
				"const ConstDeclMultiCmtSingle = `multiline constant",
				"definition",
				"`",
			},
			doc: []string{
				"ConstDeclMultiCmtSingle has a single-line comment with a multiline decl.",
			},
			oneLine: "" +
				"const ConstDeclMultiCmtSingle = `multiline constant ...",
		},
		//  ----------------------------------------------------------------------
		{
			action: "ConstDeclConstrCmtSingle",
			header: nil,
			body: []string{
				"const ConstDeclConstrCmtSingle = `multiline constant` + \"\n\" +",
				"    ConstDeclMultiCmtSingle + \" including other constants: \n\" +",
				"    ConstDeclSingleCmtSingle + \"\n\" + `",
				"=========end of constant=============",
				"`",
			},
			doc: []string{
				"ConstDeclConstrCmtSingle has a single-line comment with a multiline decl.",
			},
			oneLine: "" +
				"const ConstDeclConstrCmtSingle = `multiline constant` + \"\n\" + ...",
		},
		//  ----------------------------------------------------------------------
		{
			action: "ConstDeclConstrCmtMulti",
			header: nil,
			body: []string{
				"const ConstDeclConstrCmtMulti = `multiline constant` + \"\n\" +",
				"    ConstDeclMultiCmtSingle + \" including other constants: \n\" +",
				"    ConstDeclSingleCmtSingle + \"\n\" + `",
				"=========end of constant=============",
				"`",
			},
			doc: []string{
				"ConstDeclConstrCmtMulti has a multiline comment with",
				"a multiline decl.",
			},
			oneLine: "const ConstDeclConstrCmtMulti =" +
				" `multiline constant` + \"\n\" + ...",
		},
		//  ----------------------------------------------------------------------
		{
			action: "StructureType.GetF1",
			header: []string{
				"func (s *StructureType) GetF1(",
				"    a, b, c int,",
				") string",
			},
			body: []string{
				"func (s *StructureType) GetF1(",
				"    a, b, c int,",
				") string {",
				"    const base10 = 10",
				"",
				"    t := a + c + b",
				"",
				"    return s.F1 + strconv.FormatInt(int64(t), base10)",
				"}",
			},
			doc: []string{
				"GetF1 is a method to a structure.",
			},
			oneLine: "func (s *StructureType) GetF1(a, b, c int) string",
		},
	}

	for _, tst := range docInfoTests {
		dInfo, err := gopkg.Info("./sample1", tst.action)
		chk.NoErr(err)
		chk.StrSlice(dInfo.Header(), tst.header, "HEADER For action: ", tst.action)
		chk.StrSlice(dInfo.Body(), tst.body, "BODY For action: ", tst.action)
		chk.StrSlice(dInfo.Doc(), tst.doc, "DOC For action: ", tst.action)
		chk.Str(dInfo.OneLine(), tst.oneLine, "OneLine For action: ", tst.action)
	}

	chk.Stdout(
		"getInfo(\"TimesTwo\")",
		"getInfo(\"TimesThree\")",
		"getInfo(\"ConstDeclSingleCmtSingle\")",
		"getInfo(\"ConstDeclMultiCmtSingle\")",
		"getInfo(\"ConstDeclConstrCmtSingle\")",
		"getInfo(\"ConstDeclConstrCmtMulti\")",
		"getInfo(\"StructureType.GetF1\")",
	)
}
