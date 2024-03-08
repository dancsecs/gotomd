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

package main

import (
	"os"
	"testing"

	"github.com/dancsecs/sztest"
)

func Test_GetDoc_MarkGoCode(t *testing.T) {
	chk := sztest.CaptureNothing(t)
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
	chk := sztest.CaptureNothing(t)
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
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	s, err := getDocDecl(example1Path)
	chk.Err(err, ErrMissingAction.Error())
	chk.Str(s, "")
}

func Test_GetDoc_GetGoDcl_Package(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	s, err := getDocDecl(example1Path + pkgLabel)
	chk.NoErr(err)
	chk.Str(
		s,
		markGoCode(pkgLabel+" "+example1+"\n"),
	)
}

func Test_GetDoc_GetGoDcl_InvalidItem(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	s, err := getDocDecl(example1Path + "unknownItem")
	chk.Err(err, ErrUnknownObject.Error()+": unknownItem")
	chk.Str(s, "")
}

func Test_GetDoc_GetGoDcl_OneItem(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	chk.AddSub(
		pkgLabel+` .*$`,
		pkgLabel+" ."+string(os.PathSeparator)+example1,
	)

	s, err := getDocDecl(example1Path + "TimesTwo")
	chk.NoErr(err)
	chk.Str(
		s,
		markGoCode("func TimesTwo(i int) int\n"),
	)
}

func Test_GetDoc_GetGoDcl_TwoItems(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	chk.AddSub(
		pkgLabel+` .*$`,
		pkgLabel+" ."+string(os.PathSeparator)+example1,
	)

	s, err := getDocDecl(example1Path + "TimesTwo TimesThree")
	chk.NoErr(err)
	chk.Str(
		s,
		markGoCode("func TimesTwo(i int) int\nfunc TimesThree(i int) int\n"),
	)
}

func Test_GetDoc_GetGoDclSingle_NoItems(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	s, err := getDocDeclSingle(example1Path)
	chk.Err(err, ErrMissingAction.Error())
	chk.Str(s, "")
}

func Test_GetDoc_GetGoDclSingle_PackageNoItems(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	s, err := getDocDeclSingle(example1Path + pkgLabel)
	chk.NoErr(err)
	chk.Str(
		s,
		markGoCode(pkgLabel+" "+example1+"\n"),
	)
}

func Test_GetDoc_GetGoDclSingle_InvalidItem(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	s, err := getDocDeclSingle(example1Path + "unknownItem")
	chk.Err(err, ErrUnknownObject.Error()+": unknownItem")
	chk.Str(s, "")
}

func Test_GetDoc_GetGoDclSingle_OneItem(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	s, err := getDocDeclSingle(example1Path + "TimesTwo")

	chk.AddSub(
		pkgLabel+` .*$`,
		pkgLabel+" ."+string(os.PathSeparator)+example1,
	)
	chk.NoErr(err)
	chk.Str(
		s,
		markGoCode("func TimesTwo(i int) int\n"),
	)
}

func Test_GetDoc_GetGoDclSingle_TwoItems(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	chk.AddSub(
		pkgLabel+` .*$`,
		pkgLabel+" ."+string(os.PathSeparator)+example1,
	)

	s, err := getDocDeclSingle(example1Path + "TimesTwo TimesThree")
	chk.NoErr(err)
	chk.Str(
		s,
		markGoCode("func TimesTwo(i int) int\nfunc TimesThree(i int) int\n"),
	)
}

func Test_GetDoc_GetGoDclNatural_InvalidItem(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	s, err := getDocDeclNatural(example1Path + "unknownItem")
	chk.Err(err, ErrUnknownObject.Error()+": unknownItem")
	chk.Str(s, "")
}

func Test_GetDoc_GetGoDclNatural_OneItem(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	s, err := getDocDeclNatural(example1Path + "TimesTwo")

	chk.AddSub(
		pkgLabel+` .*$`,
		pkgLabel+" ."+string(os.PathSeparator)+example1,
	)
	chk.NoErr(err)
	chk.Str(
		s,
		markGoCode(
			"// TimesTwo returns the value times two.\n"+
				"func TimesTwo(i int) int",
		),
	)
}

func Test_GetDoc_GetGoDclNatural_TwoItems(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	s, err := getDocDeclNatural(example1Path + "TimesTwo TimesThree")

	chk.AddSub(
		pkgLabel+` .*$`,
		pkgLabel+" ."+string(os.PathSeparator)+example1,
	)
	chk.NoErr(err)
	chk.Str(
		s,
		markGoCode(
			"// TimesTwo returns the value times two.\n"+
				"func TimesTwo(i int) int\n"+
				"\n"+
				"// TimesThree returns the value times three.\n"+
				"func TimesThree(i int) int",
		),
	)
}

func Test_GetDoc_GetDoc_TwoItems(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	chk.AddSub(
		pkgLabel+` .*$`,
		pkgLabel+" ."+string(os.PathSeparator)+example1,
	)

	s, err := getDoc(example1Path + "TimesTwo TimesThree")
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
}
