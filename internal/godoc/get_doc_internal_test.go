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

// Program runs a specific go test transforming the output
// to github compatible markdown.  This is used within this
// project to help automate keeping the README.md up to date
// when an example changes.

package godoc

import (
	"os"
	"strings"
	"testing"

	"github.com/dancsecs/gotomd/internal/errs"
	"github.com/dancsecs/sztestlog"
)

const (
	pkgLabel   = "package"
	sep        = string(os.PathSeparator)
	tstpkg     = "tstpkg"
	tstpkgPath = "." + sep + "testdata" + sep + tstpkg + sep
)

func Test_CmdParse_ParseCmd_InvalidDir(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	cmd := "." + sep + "INVALID_DIR" + sep + "action"

	str, err := GetDoc(cmd)
	chk.Err(
		err,
		chk.ErrChain(
			errs.ErrInvalidDirectory,
			"\"."+sep+"INVALID_DIR\"",
		),
	)
	chk.Str(
		str,
		"",
	)
}

func Test_GetDoc_MarkGoCode(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	chk.Str(
		markGoCode("ABC"),
		"```go\nABC\n```",
	)

	chk.Str(
		markGoCode("ABC\n"),
		"```go\nABC\n```",
	)
}

func Test_GetDoc_MarkBashCode(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	chk.Str(
		markBashCode("ABC"),
		"```bash\nABC\n```",
	)

	chk.Str(
		markBashCode("ABC\n"),
		"```bash\nABC\n```",
	)
}

func Test_GetDoc_GetGoDcl_NoItems(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	s, err := GetDocDecl(tstpkgPath)
	chk.Err(err, errs.ErrMissingAction.Error())
	chk.Str(s, "")
}

func Test_GetDoc_GetGoDcl_Package(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	s, err := GetDocDecl(tstpkgPath + pkgLabel)
	chk.NoErr(err)
	chk.Str(
		s,
		markGoCode(pkgLabel+" "+tstpkg+"\n"),
	)

	chk.Stdout(
		"Loading package info for: ./testdata/tstpkg",
		"getInfo(\"package\")",
	)
}

func Test_GetDoc_GetGoDcl_InvalidItem(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	s, err := GetDocDecl(tstpkgPath + "unknownItem")
	chk.Err(err, errs.ErrUnknownObject.Error()+": unknownItem")
	chk.Str(s, "")

	chk.Stdout(
		"getInfo(\"unknownItem\")",
	)
}

func Test_GetDoc_GetGoDcl_OneItem(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	chk.AddSub(
		pkgLabel+` .*$`,
		pkgLabel+" ."+string(os.PathSeparator)+tstpkg,
	)

	s, err := GetDocDecl(tstpkgPath + "TimesTwo")
	chk.NoErr(err)
	chk.Str(
		s,
		markGoCode("func TimesTwo(i int) int\n"),
	)

	chk.Stdout(
		"getInfo(\"TimesTwo\")",
	)
}

func Test_GetDoc_GetGoDcl_TwoItems(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	chk.AddSub(
		pkgLabel+` .*$`,
		pkgLabel+" ."+string(os.PathSeparator)+tstpkg,
	)

	s, err := GetDocDecl(tstpkgPath + "TimesTwo TimesThree")
	chk.NoErr(err)
	chk.Str(
		s,
		markGoCode("func TimesTwo(i int) int\nfunc TimesThree(i int) int\n"),
	)

	chk.Stdout(
		"getInfo(\"TimesTwo\")",
		"getInfo(\"TimesThree\")",
	)
}

func Test_GetDoc_GetGoDclSingle_NoItems(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	s, err := GetDocDeclSingle(tstpkgPath)
	chk.Err(err, errs.ErrMissingAction.Error())
	chk.Str(s, "")
}

func Test_GetDoc_GetGoDclSingle_PackageNoItems(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	s, err := GetDocDeclSingle(tstpkgPath + pkgLabel)
	chk.NoErr(err)
	chk.Str(
		s,
		markGoCode(pkgLabel+" "+tstpkg+"\n"),
	)

	chk.Stdout(
		"getInfo(\"package\")",
	)
}

func Test_GetDoc_GetGoDclSingle_InvalidItem(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	s, err := GetDocDeclSingle(tstpkgPath + "unknownItem")
	chk.Err(err, errs.ErrUnknownObject.Error()+": unknownItem")
	chk.Str(s, "")

	chk.Stdout(
		"getInfo(\"unknownItem\")",
	)
}

func Test_GetDoc_GetGoDclSingle_OneItem(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	line, err := GetDocDeclSingle(tstpkgPath + "TimesTwo")

	chk.AddSub(
		pkgLabel+` .*$`,
		pkgLabel+" ."+string(os.PathSeparator)+tstpkg,
	)
	chk.NoErr(err)
	chk.Str(
		line,
		markGoCode("func TimesTwo(i int) int\n"),
	)

	chk.Stdout(
		"getInfo(\"TimesTwo\")",
	)
}

func Test_GetDoc_GetGoDclSingle_TwoItems(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	chk.AddSub(
		pkgLabel+` .*$`,
		pkgLabel+" ."+string(os.PathSeparator)+tstpkg,
	)

	s, err := GetDocDeclSingle(tstpkgPath + "TimesTwo TimesThree")
	chk.NoErr(err)
	chk.Str(
		s,
		markGoCode("func TimesTwo(i int) int\nfunc TimesThree(i int) int\n"),
	)

	chk.Stdout(
		"getInfo(\"TimesTwo\")",
		"getInfo(\"TimesThree\")",
	)
}

func Test_GetDoc_GetGoDclNatural_InvalidItem(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	s, err := GetDocDeclNatural(tstpkgPath + "unknownItem")
	chk.Err(err, errs.ErrUnknownObject.Error()+": unknownItem")
	chk.Str(s, "")

	chk.Stdout(
		"getInfo(\"unknownItem\")",
	)
}

func Test_GetDoc_GetGoDclNatural_OneItem(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	line, err := GetDocDeclNatural(tstpkgPath + "TimesTwo")

	chk.AddSub(
		pkgLabel+` .*$`,
		pkgLabel+" ."+string(os.PathSeparator)+tstpkg,
	)
	chk.NoErr(err)
	chk.Str(
		line,
		markGoCode(
			"// TimesTwo returns the value times two.\n"+
				"func TimesTwo(i int) int",
		),
	)

	chk.Stdout(
		"getInfo(\"TimesTwo\")",
	)
}

func Test_GetDoc_GetGoDclNatural_TwoItems(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	line, err := GetDocDeclNatural(tstpkgPath + "TimesTwo TimesThree")

	chk.AddSub(
		pkgLabel+` .*$`,
		pkgLabel+" ."+string(os.PathSeparator)+tstpkg,
	)
	chk.NoErr(err)
	chk.Str(
		line,
		markGoCode(
			"// TimesTwo returns the value times two.\n"+
				"func TimesTwo(i int) int\n"+
				"\n"+
				"// TimesThree returns the value times three.\n"+
				"func TimesThree(i int) int",
		),
	)

	chk.Stdout(
		"getInfo(\"TimesTwo\")",
		"getInfo(\"TimesThree\")",
	)
}

func Test_GetDoc_GetDoc_TwoItems(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	chk.AddSub(
		pkgLabel+` .*$`,
		pkgLabel+" ."+string(os.PathSeparator)+tstpkg,
	)

	s, err := GetDoc(tstpkgPath + "TimesTwo TimesThree")
	chk.NoErr(err)
	chk.Str(
		s,
		""+
			markGoCode("func TimesTwo(i int) int")+"\n\n"+
			"TimesTwo returns the value times two.\n"+
			"\n"+
			markGoCode("func TimesThree(i int) int")+"\n\n"+
			"TimesThree returns the value times three.",
	)

	chk.Stdout(
		"getInfo(\"TimesTwo\")",
		"getInfo(\"TimesThree\")",
	)
}

func Test_GetDoc_GetDocConstantBlock_InvalidItem(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	s, err := GetDocDeclConstantBlock(tstpkgPath + "unknownItem")
	chk.Err(err, errs.ErrUnknownObject.Error()+": unknownItem")
	chk.Str(s, "")

	chk.Stdout(
		"getInfo(\"unknownItem\")",
	)
}

func Test_GetDoc_GetDocConstantBlockOne(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	chk.AddSub(
		pkgLabel+` .*$`,
		pkgLabel+" ."+string(os.PathSeparator)+tstpkg,
	)

	s, err := GetDocDeclConstantBlock(tstpkgPath + "ConstantGroup1")
	chk.NoErr(err)
	chk.StrSlice(
		strings.Split(s, "\n"),
		[]string{
			"```go",
			"// Here is a typed constant block.  " +
				"All constants are reported as a group.",
			"const (",
			"    // ConstantGroup1 is a constant defined in a group.",
			"    ConstantGroup1 ConstGroupType = iota",
			"",
			"    // ConstantGroup2 is a constant defined in a group.",
			"    ConstantGroup2",
			")",
			"```",
		},
	)

	chk.Stdout(
		"getInfo(\"ConstantGroup1\")",
	)
}

func Test_GetDoc_GetDocConstantBlockTwo(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	s, err := GetDocDeclConstantBlock(tstpkgPath +
		"ConstantGroup1 ConstantGroupA",
	)
	chk.NoErr(err)
	chk.StrSlice(
		strings.Split(s, "\n"),
		[]string{
			"```go",
			"// Here is a typed constant block.  " +
				"All constants are reported as a group.",
			"const (",
			"    // ConstantGroup1 is a constant defined in a group.",
			"    ConstantGroup1 ConstGroupType = iota",
			"",
			"    // ConstantGroup2 is a constant defined in a group.",
			"    ConstantGroup2",
			")",
			"",
			"",
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
		"getInfo(\"ConstantGroup1\")",
		"getInfo(\"ConstantGroupA\")",
	)
}
