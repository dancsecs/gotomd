<!--- gotomd::Auto:: See github.com/dancsecs/gotomd **DO NOT MODIFY** -->

# Package example2

This project is used by the Szerszam utility function to test its markdown
update methods against an independent standalone project. All features
will be tested against this file so it will be updated and changed often.

The following will be replaced by the go package documentation

<!--- gotomd::Bgn::doc::./package -->
```go
package example2
```

Package example2 exists in order to test various go to git
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

<!--- gotomd::Bgn::tst::./Test_PASS_Example2 -->
```bash
go test -v -cover -run Test_PASS_Example2 .
```

$\small{\texttt{===&#xa0;RUN&#xa0;&#xa0;&#xa0;Test&#x332;PASS&#x332;Example2}}$
<br>
$\small{\texttt{---&#xa0;PASS:&#xa0;Test&#x332;PASS&#x332;Example2&#xa0;(0.0s)}}$
<br>
$\small{\texttt{PASS}}$
<br>
$\small{\texttt{coverage:&#xa0;100.0&#xFE6A;&#xa0;of&#xa0;statements}}$
<br>
$\small{\texttt{ok&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;github.com/dancsecs/gotomd/example2&#xa0;&#xa0;&#xa0;&#xa0;coverage:&#xa0;100.0&#xFE6A;&#xa0;of&#xa0;statements}}$
<br>
<!--- gotomd::End::tst::./Test_PASS_Example2 -->

or run all tests in a package:

<!--- gotomd::Bgn::tst::./package -->
```bash
go test -v -cover .
```

$\small{\texttt{===&#xa0;RUN&#xa0;&#xa0;&#xa0;Test&#x332;PASS&#x332;Example2}}$
<br>
$\small{\texttt{---&#xa0;PASS:&#xa0;Test&#x332;PASS&#x332;Example2&#xa0;(0.0s)}}$
<br>
$\small{\texttt{===&#xa0;RUN&#xa0;&#xa0;&#xa0;Test&#x332;FAIL&#x332;Example2}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;example2&#x332;test.go:29:&#xa0;unexpected&#xa0;int:}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\emph{2+2=5&#xa0;(is&#xa0;true&#xa0;for&#xa0;big&#xa0;values&#xa0;of&#xa0;two)}:}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{magenta}GOT:&#xa0;\color{default}\color{darkturquoise}4\color{default}}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{cyan}WNT:&#xa0;\color{default}\color{darkturquoise}5\color{default}}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;example2&#x332;test.go:30:&#xa0;unexpected&#xa0;string:}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{magenta}GOT:&#xa0;\color{default}\color{green}New&#xa0;in&#xa0;Got\color{default}&#xa0;Similar&#xa0;in&#xa0;(\color{darkturquoise}1\color{default})&#xa0;both}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{cyan}WNT:&#xa0;\color{default}&#xa0;Similar&#xa0;in&#xa0;(\color{darkturquoise}2\color{default})&#xa0;both\color{red},&#xa0;new&#xa0;in&#xa0;Wnt\color{default}}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;example2&#x332;test.go:36:&#xa0;Unexpected&#xa0;stdout&#xa0;Entry:&#xa0;got&#xa0;(1&#xa0;lines)&#xa0;-&#xa0;want&#xa0;(1&#xa0;lines)}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{darkturquoise}0\color{default}:\color{darkturquoise}0\color{default}&#xa0;This&#xa0;output&#xa0;line&#xa0;\color{red}is\color{default}\color{yellow}/\color{default}\color{green}will&#xa0;be\color{default}&#xa0;different}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;example2&#x332;test.go:40:&#xa0;unexpected&#xa0;string:}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{magenta}GOT:&#xa0;\color{default}\color{darkturquoise}Total\color{default}:&#xa0;6}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{cyan}WNT:&#xa0;\color{default}\color{darkturquoise}Sum\color{default}:&#xa0;6}}$
<br>
$\small{\texttt{---&#xa0;FAIL:&#xa0;Test&#x332;FAIL&#x332;Example2&#xa0;(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:&#xa0;100.0&#xFE6A;&#xa0;of&#xa0;statements}}$
<br>
$\small{\texttt{FAIL&#xa0;github.com/dancsecs/gotomd/example2&#xa0;0.0s}}$
<br>
$\small{\texttt{FAIL}}$
<br>
<!--- gotomd::End::tst::./package -->

or include a file

<!--- gotomd::Bgn::file::./example2.go -->
```bash
cat ./example2.go
```

```go
// Package example2 exists in order to test various go to git
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
package example2

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
<!--- gotomd::End::file::./example2.go -->

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
