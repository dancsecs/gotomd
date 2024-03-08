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

package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/dancsecs/sztest"
)

func Test_GetTest_GetGoTst(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	cmd := "TEST_DIRECTORY_DOES_NOT_EXIST" + string(os.PathSeparator)
	_, err := getGoTst(cmd)
	chk.Err(
		err,
		ErrInvalidRelativeDir.Error()+": \""+cmd+"\"",
	)

	_, err = getGoTst("./TEST_DOES_NOT_EXIST")
	chk.Err(err, ErrNoTestToRun.Error())
}

func Test_GetTest_RunTestNotDirectory(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	f := chk.CreateTmpFile(nil)
	chk.Panic(
		func() {
			_, _, _ = runTest(f, "")
		},
		"",
	)
}

//nolint:funlen // Ok.
func Test_GetTest_RunTestColorize(t *testing.T) {
	chk := sztest.CaptureStdout(t)
	defer chk.Release()

	szColorize = true
	defer func() {
		szColorize = false
	}()

	file1 := example1Path + pkgLabel
	file2 := example2Path + pkgLabel
	s, err := getGoTst(file1 + " " + file2)

	chk.NoErr(err)
	fmt.Printf("%s\n", s)

	chk.AddSub("{{file1}}", file1)
	chk.AddSub("{{file2}}", file2)
	chk.AddSub("{{insOn}}", internalTestMarkInsOn)
	chk.AddSub("{{insOff}}", internalTestMarkInsOff)
	chk.AddSub("{{delOn}}", internalTestMarkDelOn)
	chk.AddSub("{{delOff}}", internalTestMarkDelOff)
	chk.AddSub("{{chgOn}}", internalTestMarkChgOn)
	chk.AddSub("{{chgOff}}", internalTestMarkChgOff)
	chk.AddSub("{{gotOn}}", internalTestMarkGotOn)
	chk.AddSub("{{gotOff}}", internalTestMarkGotOff)
	chk.AddSub("{{wntOn}}", internalTestMarkWntOn)
	chk.AddSub("{{wntOff}}", internalTestMarkWntOff)
	chk.AddSub("{{msgOn}}", internalTestMarkMsgOn)
	chk.AddSub("{{msgOff}}", internalTestMarkMsgOff)
	chk.AddSub("{{sepOn}}", internalTestMarkSepOn)
	chk.AddSub("{{sepOff}}", internalTestMarkSepOff)
	chk.AddSub("{{latexOn}}", `$\small{\texttt{`)
	chk.AddSub("{{latexOff}}", `}}$`)
	chk.AddSub(`\t\d+\.\d+s`, "\t0.0s")
	chk.AddSub(` `, hardSpace)
	chk.AddSub(`_`, hardUnderscore)
	chk.Stdout("" +
		markBashCode(
			"go test -v -cover ."+
				string(os.PathSeparator)+example1) + "\n" +
		`
    {{latexOn}}=== RUN   Test_PASS_Example1{{latexOff}}
    <br>
    {{latexOn}}--- PASS: Test_PASS_Example1 (0.0s){{latexOff}}
    <br>
    {{latexOn}}=== RUN   Test_FAIL_Example1{{latexOff}}
    <br>
    {{latexOn}}    example1_test.go:28: unexpected int:{{latexOff}}
    <br>
    {{latexOn}}        {{msgOn}}2+2=5 (is true for big values of two){{msgOff}}:{{latexOff}}
    <br>
    {{latexOn}}        {{gotOn}}GOT: {{gotOff}}{{chgOn}}4{{chgOff}}{{latexOff}}
    <br>
    {{latexOn}}        {{wntOn}}WNT: {{wntOff}}{{chgOn}}5{{chgOff}}{{latexOff}}
    <br>
    {{latexOn}}    example1_test.go:29: unexpected string:{{latexOff}}
    <br>
    {{latexOn}}        {{gotOn}}GOT: {{gotOff}}{{insOn}}New in Got{{insOff}} Similar in ({{chgOn}}1{{chgOff}}) both{{latexOff}}
    <br>
    {{latexOn}}        {{wntOn}}WNT: {{wntOff}} Similar in ({{chgOn}}2{{chgOff}}) both{{delOn}}, new in Wnt{{delOff}}{{latexOff}}
    <br>
    {{latexOn}}    example1_test.go:35: Unexpected stdout Entry: got (1 lines) - want (1 lines){{latexOff}}
    <br>
    {{latexOn}}        {{chgOn}}0{{chgOff}}:{{chgOn}}0{{chgOff}} This output line {{delOn}}is{{delOff}}{{sepOn}}/{{sepOff}}{{insOn}}will be{{insOff}} different{{latexOff}}
    <br>
    {{latexOn}}    example1_test.go:39: unexpected string:{{latexOff}}
    <br>
    {{latexOn}}        {{gotOn}}GOT: {{gotOff}}{{chgOn}}Total{{chgOff}}: 6{{latexOff}}
    <br>
    {{latexOn}}        {{wntOn}}WNT: {{wntOff}}{{chgOn}}Sum{{chgOff}}: 6{{latexOff}}
    <br>
    {{latexOn}}--- FAIL: Test_FAIL_Example1 (0.0s){{latexOff}}
    <br>
    {{latexOn}}FAIL{{latexOff}}
    <br>
    {{latexOn}}coverage: 100.0&#xFE6A; of statements{{latexOff}}
    <br>
    {{latexOn}}FAIL github.com/dancsecs/gotomd/example1 0.0s{{latexOff}}
    <br>
    {{latexOn}}FAIL{{latexOff}}
    <br>

    ` +
		markBashCode(
			"go test -v -cover ."+
				string(os.PathSeparator)+example2) + "\n" +
		`
    {{latexOn}}=== RUN   Test_PASS_Example2{{latexOff}}
    <br>
    {{latexOn}}--- PASS: Test_PASS_Example2 (0.0s){{latexOff}}
    <br>
    {{latexOn}}=== RUN   Test_FAIL_Example2{{latexOff}}
    <br>
    {{latexOn}}    example2_test.go:28: unexpected int:{{latexOff}}
    <br>
    {{latexOn}}        {{msgOn}}2+2=5 (is true for big values of two){{msgOff}}:{{latexOff}}
    <br>
    {{latexOn}}        {{gotOn}}GOT: {{gotOff}}{{chgOn}}4{{chgOff}}{{latexOff}}
    <br>
    {{latexOn}}        {{wntOn}}WNT: {{wntOff}}{{chgOn}}5{{chgOff}}{{latexOff}}
    <br>
    {{latexOn}}    example2_test.go:29: unexpected string:{{latexOff}}
    <br>
    {{latexOn}}        {{gotOn}}GOT: {{gotOff}}{{insOn}}New in Got{{insOff}} Similar in ({{chgOn}}1{{chgOff}}) both{{latexOff}}
    <br>
    {{latexOn}}        {{wntOn}}WNT: {{wntOff}} Similar in ({{chgOn}}2{{chgOff}}) both{{delOn}}, new in Wnt{{delOff}}{{latexOff}}
    <br>
    {{latexOn}}    example2_test.go:35: Unexpected stdout Entry: got (1 lines) - want (1 lines){{latexOff}}
    <br>
    {{latexOn}}        {{chgOn}}0{{chgOff}}:{{chgOn}}0{{chgOff}} This output line {{delOn}}is{{delOff}}{{sepOn}}/{{sepOff}}{{insOn}}will be{{insOff}} different{{latexOff}}
    <br>
    {{latexOn}}    example2_test.go:39: unexpected string:{{latexOff}}
    <br>
    {{latexOn}}        {{gotOn}}GOT: {{gotOff}}{{chgOn}}Total{{chgOff}}: 6{{latexOff}}
    <br>
    {{latexOn}}        {{wntOn}}WNT: {{wntOff}}{{chgOn}}Sum{{chgOff}}: 6{{latexOff}}
    <br>
    {{latexOn}}--- FAIL: Test_FAIL_Example2 (0.0s){{latexOff}}
    <br>
    {{latexOn}}FAIL{{latexOff}}
    <br>
    {{latexOn}}coverage: 100.0&#xFE6A; of statements{{latexOff}}
    <br>
    {{latexOn}}FAIL github.com/dancsecs/gotomd/example2 0.0s{{latexOff}}
    <br>
    {{latexOn}}FAIL{{latexOff}}
    <br>
  `)
}

//nolint:funlen // Ok.
func Test_GetTest_RunTestNoColor(t *testing.T) {
	chk := sztest.CaptureStdout(t)
	defer chk.Release()

	file1 := example1Path + pkgLabel
	file2 := example2Path + pkgLabel
	s, err := getGoTst(file1 + " " + file2)

	chk.NoErr(err)
	fmt.Printf("%s\n", s)

	chk.Stdout("" +
		markBashCode(
			"go test -v -cover ."+
				string(os.PathSeparator)+example1) + "\n" +
		`
    <pre>
    === RUN   Test_PASS_Example1
    --- PASS: Test_PASS_Example1 (0.0s)
    === RUN   Test_FAIL_Example1
    \s   example1_test.go:28: unexpected int:
    \s       2+2=5 (is true for big values of two):
    \s       GOT: 4
    \s       WNT: 5
    \s   example1_test.go:29: unexpected string:
    \s       GOT: New in Got Similar in (1) both
    \s       WNT:  Similar in (2) both, new in Wnt
    \s   example1_test.go:35: Unexpected stdout Entry: got (1 lines) - want (1 lines)
    \s       0:0 This output line is/will be different
    \s   example1_test.go:39: unexpected string:
    \s       GOT: Total: 6
    \s       WNT: Sum: 6
    --- FAIL: Test_FAIL_Example1 (0.0s)
    FAIL
    coverage: 100.0&#xFE6A; of statements
    FAIL github.com/dancsecs/gotomd/example1 0.0s
    FAIL
    </pre>

    ` +
		markBashCode(
			"go test -v -cover ."+
				string(os.PathSeparator)+example2) + "\n" +
		`
    <pre>
    === RUN   Test_PASS_Example2
    --- PASS: Test_PASS_Example2 (0.0s)
    === RUN   Test_FAIL_Example2
    \s   example2_test.go:28: unexpected int:
    \s       2+2=5 (is true for big values of two):
    \s       GOT: 4
    \s       WNT: 5
    \s   example2_test.go:29: unexpected string:
    \s       GOT: New in Got Similar in (1) both
    \s       WNT:  Similar in (2) both, new in Wnt
    \s   example2_test.go:35: Unexpected stdout Entry: got (1 lines) - want (1 lines)
    \s       0:0 This output line is/will be different
    \s   example2_test.go:39: unexpected string:
    \s       GOT: Total: 6
    \s       WNT: Sum: 6
    --- FAIL: Test_FAIL_Example2 (0.0s)
    FAIL
    coverage: 100.0&#xFE6A; of statements
    FAIL github.com/dancsecs/gotomd/example2 0.0s
    FAIL
    </pre>
  `)
}

func Test_GetTest_BuildTestCmds(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	dirs := []string{}
	actions := []string{}
	d, a := buildTestCmds(dirs, actions)
	chk.StrSlice(d, nil)
	chk.StrSlice(a, nil)

	actions = append(actions, "D1A1")
	d, a = buildTestCmds(dirs, actions)
	chk.StrSlice(d, nil)
	chk.StrSlice(a, nil)

	dirs = append(dirs, "D1")
	d, a = buildTestCmds(dirs, actions)
	chk.StrSlice(d, []string{"D1"})
	chk.StrSlice(a, []string{"D1A1"})

	dirs = append(dirs, "D1")
	actions = append(actions, "D1A2")
	d, a = buildTestCmds(dirs, actions)
	chk.StrSlice(d, []string{"D1"})
	chk.StrSlice(a, []string{"D1A1 D1A2"})

	dirs = append(dirs, "D2")
	actions = append(actions, "D2A1")
	d, a = buildTestCmds(dirs, actions)
	chk.StrSlice(d, []string{"D1", "D2"})
	chk.StrSlice(a, []string{"D1A1 D1A2", "D2A1"})
}
