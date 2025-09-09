<!--- gotomd::Auto:: See github.com/dancsecs/gotomd **DO NOT MODIFY** -->

# Package example1

This project is used by the Szerszam utility function to test its markdown
update methods against an independent standalone project. All features
will be tested against this file so it will be updated and changed often.

The following will be replaced by the go package documentation

<!--- gotomd::Bgn::doc::./package -->
```go
package example1
```

Package example1 exists in order to test various go to git
markdown (gToMD) extraction utilities.  Various object will be defined that
exhibit the various comment and declaration options permitted by gofmt.

# Heading

This paragraph will demonstrating further documentation under a "markdown"
header.

Declarations can be single-line or multi-line blocks or constructions.  Each
type will be included here for complete testing.
<!--- gotomd::End::doc::./package -->

Here we will add function documentation:

<!--- gotomd::Bgn::doc::./TimesTwo -->
```go
func TimesTwo(i int) int
```

TimesTwo returns the value times two.
<!--- gotomd::End::doc::./TimesTwo -->

and another:

<!--- gotomd::Bgn::doc::./TimesThree -->
```go
func TimesThree(i int) int
```

TimesThree returns the value times three.
<!--- gotomd::End::doc::./TimesThree -->

and the defined interface:

<!--- gotomd::Bgn::doc::./InterfaceType -->
```go
type InterfaceType interface {
    func(int) int
}
```

InterfaceType tests the documentation of interfaces.
<!--- gotomd::End::doc::./InterfaceType -->

and the defined structure:

<!--- gotomd::Bgn::doc::./StructureType -->
```go
type StructureType struct {
    // F1 is the first test field of the structure.
    F1 string
    // F2 is the second test field of the structure.
    F2 int
}
```

StructureType tests the documentation of structures.
<!--- gotomd::End::doc::./StructureType -->

and run a specific test

<!--- gotomd::Bgn::tst::./Test_PASS_Example1 -->
```bash
go test -v -cover -run Test_PASS_Example1 .
```

$\small{\texttt{===&#xA0;&#x34F;&#xA0;&#x34F;RUN&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;PASS&#xA0;&#x332;&#xA0;&#x332;Example1}}$
<br>
$\small{\texttt{‒‒‒&#xA0;&#x34F;&#xA0;&#x34F;PASS:&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;PASS&#xA0;&#x332;&#xA0;&#x332;Example1&#xA0;&#x34F;&#xA0;&#x34F;(0.0s)}}$
<br>
$\small{\texttt{PASS}}$
<br>
$\small{\texttt{coverage:&#xA0;&#x34F;&#xA0;&#x34F;100.0&#xFE6A;&#xA0;&#x34F;&#xA0;&#x34F;of&#xA0;&#x34F;&#xA0;&#x34F;statements}}$
<br>
$\small{\texttt{ok&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;github.com/dancsecs/gotomd/example1&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;coverage:&#xA0;&#x34F;&#xA0;&#x34F;100.0&#xFE6A;&#xA0;&#x34F;&#xA0;&#x34F;of&#xA0;&#x34F;&#xA0;&#x34F;statements}}$
<br>
<!--- gotomd::End::tst::./Test_PASS_Example1 -->

or run all tests in a package:

<!--- gotomd::Bgn::tst::./package -->
```bash
go test -v -cover .
```

$\small{\texttt{===&#xA0;&#x34F;&#xA0;&#x34F;RUN&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;PASS&#xA0;&#x332;&#xA0;&#x332;Example1}}$
<br>
$\small{\texttt{‒‒‒&#xA0;&#x34F;&#xA0;&#x34F;PASS:&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;PASS&#xA0;&#x332;&#xA0;&#x332;Example1&#xA0;&#x34F;&#xA0;&#x34F;(0.0s)}}$
<br>
$\small{\texttt{===&#xA0;&#x34F;&#xA0;&#x34F;RUN&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;FAIL&#xA0;&#x332;&#xA0;&#x332;Example1}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;example1&#xA0;&#x332;&#xA0;&#x332;test.go:29:&#xA0;&#x34F;&#xA0;&#x34F;unexpected&#xA0;&#x34F;&#xA0;&#x34F;int:}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\emph{2+2=5&#xA0;&#x34F;&#xA0;&#x34F;(is&#xA0;&#x34F;&#xA0;&#x34F;true&#xA0;&#x34F;&#xA0;&#x34F;for&#xA0;&#x34F;&#xA0;&#x34F;big&#xA0;&#x34F;&#xA0;&#x34F;values&#xA0;&#x34F;&#xA0;&#x34F;of&#xA0;&#x34F;&#xA0;&#x34F;two)}}:}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{magenta}{GOT:&#xA0;&#x34F;&#xA0;&#x34F;}}{\color{darkturquoise}{4}}}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{cyan}{WNT:&#xA0;&#x34F;&#xA0;&#x34F;}}{\color{darkturquoise}{5}}}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;example1&#xA0;&#x332;&#xA0;&#x332;test.go:30:&#xA0;&#x34F;&#xA0;&#x34F;unexpected&#xA0;&#x34F;&#xA0;&#x34F;string:}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{magenta}{GOT:&#xA0;&#x34F;&#xA0;&#x34F;}}{\color{green}{New&#xA0;&#x34F;&#xA0;&#x34F;in&#xA0;&#x34F;&#xA0;&#x34F;Got}}&#xA0;&#x34F;&#xA0;&#x34F;Similar&#xA0;&#x34F;&#xA0;&#x34F;in&#xA0;&#x34F;&#xA0;&#x34F;({\color{darkturquoise}{1}})&#xA0;&#x34F;&#xA0;&#x34F;both}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{cyan}{WNT:&#xA0;&#x34F;&#xA0;&#x34F;}}&#xA0;&#x34F;&#xA0;&#x34F;Similar&#xA0;&#x34F;&#xA0;&#x34F;in&#xA0;&#x34F;&#xA0;&#x34F;({\color{darkturquoise}{2}})&#xA0;&#x34F;&#xA0;&#x34F;both{\color{red}{,&#xA0;&#x34F;&#xA0;&#x34F;new&#xA0;&#x34F;&#xA0;&#x34F;in&#xA0;&#x34F;&#xA0;&#x34F;Wnt}}}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;example1&#xA0;&#x332;&#xA0;&#x332;test.go:36:&#xA0;&#x34F;&#xA0;&#x34F;Unexpected&#xA0;&#x34F;&#xA0;&#x34F;stdout&#xA0;&#x34F;&#xA0;&#x34F;Entry:&#xA0;&#x34F;&#xA0;&#x34F;got&#xA0;&#x34F;&#xA0;&#x34F;(1&#xA0;&#x34F;&#xA0;&#x34F;lines)&#xA0;&#x34F;&#xA0;&#x34F;-&#xA0;&#x34F;&#xA0;&#x34F;want&#xA0;&#x34F;&#xA0;&#x34F;(1&#xA0;&#x34F;&#xA0;&#x34F;lines)}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{darkturquoise}{0}}:{\color{darkturquoise}{0}}&#xA0;&#x34F;&#xA0;&#x34F;This&#xA0;&#x34F;&#xA0;&#x34F;output&#xA0;&#x34F;&#xA0;&#x34F;line&#xA0;&#x34F;&#xA0;&#x34F;{\color{red}{is}}{\color{yellow}{/}}{\color{green}{will&#xA0;&#x34F;&#xA0;&#x34F;be}}&#xA0;&#x34F;&#xA0;&#x34F;different}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;example1&#xA0;&#x332;&#xA0;&#x332;test.go:40:&#xA0;&#x34F;&#xA0;&#x34F;unexpected&#xA0;&#x34F;&#xA0;&#x34F;string:}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{magenta}{GOT:&#xA0;&#x34F;&#xA0;&#x34F;}}{\color{darkturquoise}{Total}}:&#xA0;&#x34F;&#xA0;&#x34F;6}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{cyan}{WNT:&#xA0;&#x34F;&#xA0;&#x34F;}}{\color{darkturquoise}{Sum}}:&#xA0;&#x34F;&#xA0;&#x34F;6}}$
<br>
$\small{\texttt{‒‒‒&#xA0;&#x34F;&#xA0;&#x34F;FAIL:&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;FAIL&#xA0;&#x332;&#xA0;&#x332;Example1&#xA0;&#x34F;&#xA0;&#x34F;(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:&#xA0;&#x34F;&#xA0;&#x34F;100.0&#xFE6A;&#xA0;&#x34F;&#xA0;&#x34F;of&#xA0;&#x34F;&#xA0;&#x34F;statements}}$
<br>
$\small{\texttt{FAIL&#xA0;&#x34F;&#xA0;&#x34F;github.com/dancsecs/gotomd/example1&#xA0;&#x34F;&#xA0;&#x34F;0.0s}}$
<br>
$\small{\texttt{FAIL}}$
<br>
<!--- gotomd::End::tst::./package -->

or include a file

<!--- gotomd::Bgn::file::./example1.go -->
```bash
cat ./example1.go
```

```go
// Package example1 exists in order to test various go to git
// markdown (gToMD) extraction utilities.  Various object will be defined that
// exhibit the various comment and declaration options permitted by gofmt.
//
// # Heading
//
// This paragraph will demonstrating further documentation under a "markdown"
// header.
//
// Declarations can be single-line or multi-line blocks or constructions.  Each
// type will be included here for complete testing.
package example1

import "strconv"

// ConstDeclSingleCmtSingle has a single-line comment.
const ConstDeclSingleCmtSingle = "single-line declaration and comment"

// ConstDeclSingleCmtMulti has a multiline
// comment.
const ConstDeclSingleCmtMulti = "single-line declaration and comment"

// ConstDeclMultiCmtSingle has a single-line comment with a multiline decl.
const ConstDeclMultiCmtSingle = `multiline constant
definition
`

// ConstDeclMultiCmtMulti has a multiline comment with
// a multiline decl.
const ConstDeclMultiCmtMulti = `multiline constant
definition
`

// ConstDeclConstrCmtSingle has a single-line comment with a multiline decl.
const ConstDeclConstrCmtSingle = `multiline constant` + "\n" +
    ConstDeclMultiCmtSingle + " including other constants: \n" +
    ConstDeclSingleCmtSingle + "\n" + `
=========end of constant=============
`

// ConstDeclConstrCmtMulti has a multiline comment with
// a multiline decl.
const ConstDeclConstrCmtMulti = `multiline constant` + "\n" +
    ConstDeclMultiCmtSingle + " including other constants: \n" +
    ConstDeclSingleCmtSingle + "\n" + `
=========end of constant=============
`

// ConstantSingleLine tests single line constant definitions.
const ConstantSingleLine = "this is defined on a single-line"

// ConstantMultipleLines1 test a multiline comment with string addition.
// Also with longer:
//
// multiline comments with spacing.
const ConstantMultipleLines1 = "this constant" +
    "is defined on multiple " +
    "lines"

// ConstantMultipleLines2 tests a multiline comment with go multiline string.
const ConstantMultipleLines2 = `this constant
is defined on multiple
          lines
`

// Here is a constant block.  All constants are reported as a group.
const (
    // ConstantGroup1 is a constant defined in a group.
    ConstantGroup1 = "constant 1"

    // ConstantGroup2 is a constant defined in a group.
    ConstantGroup2 = "constant 2"
)

// InterfaceType tests the documentation of interfaces.
type InterfaceType interface {
    func(int) int
}

// StructureType tests the documentation of structures.
type StructureType struct {
    // F1 is the first test field of the structure.
    F1 string
    // F2 is the second test field of the structure.
    F2 int
}

// GetF1 is a method to a structure.
func (s *StructureType) GetF1(
    a, b, c int,
) string {
    const base10 = 10

    t := a + c + b

    return s.F1 + strconv.FormatInt(int64(t), base10)
}

// TimesTwo returns the value times two.
func TimesTwo(i int) int {
    return i + i
}

// TimesThree returns the value times three.
func TimesThree(i int) int {
    return i + i + i
}
```
<!--- gotomd::End::file::./example1.go -->

or a single declaration:

<!--- gotomd::Bgn::dcl::./TimesTwo -->
```go
func TimesTwo(i int) int
```
<!--- gotomd::End::dcl::./TimesTwo -->

or a multiple declarations:

<!--- gotomd::Bgn::dcl::./TimesTwo TimesThree -->
```go
func TimesTwo(i int) int
func TimesThree(i int) int
```
<!--- gotomd::End::dcl::./TimesTwo TimesThree -->

or a single declaration on a single-line:

<!--- gotomd::Bgn::dcls::./TimesTwo -->
```go
func TimesTwo(i int) int
```
<!--- gotomd::End::dcls::./TimesTwo -->

or a multiple declarations on a single-line:

<!--- gotomd::Bgn::dcls::./TimesTwo TimesThree -->
```go
func TimesTwo(i int) int
func TimesThree(i int) int
```
<!--- gotomd::End::dcls::./TimesTwo TimesThree -->

or a natural declaration:

<!--- gotomd::Bgn::dcln::./TimesTwo -->
```go
// TimesTwo returns the value times two.
func TimesTwo(i int) int
```
<!--- gotomd::End::dcln::./TimesTwo -->

or a multiple natural declarations:

<!--- gotomd::Bgn::dcln::./TimesTwo TimesThree -->
```go
// TimesTwo returns the value times two.
func TimesTwo(i int) int

// TimesThree returns the value times three.
func TimesThree(i int) int
```
<!--- gotomd::End::dcln::./TimesTwo TimesThree -->
