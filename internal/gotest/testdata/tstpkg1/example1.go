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
package tstpkg1

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

// ConstGroupType set the type of the constant.
type ConstGroupType int

// Here is a typed constant block.  All constants are reported as a group.
const (
	// ConstantGroup1 is a constant defined in a group.
	ConstantGroup1 ConstGroupType = iota

	// ConstantGroup2 is a constant defined in a group.
	ConstantGroup2
)

// Here is a second constant block.  All constants are reported as a group.
const (
	// ConstantGroupA is a constant defined in a group.
	ConstantGroupA = "constant A"

	// ConstantGroupB is a constant defined in a group.
	ConstantGroupB = "constant B"
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
