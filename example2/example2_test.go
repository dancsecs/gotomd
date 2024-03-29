package example2_test

import (
	"fmt"
	"testing"

	"github.com/dancsecs/gotomd/example2"
	"github.com/dancsecs/sztest"
)

func Test_PASS_Example2(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	chk.Int(example2.TimesTwo(2), 4)
	chk.Int(example2.TimesThree(222222222), 666666666)

	structure := new(example2.StructureType)
	structure.F1 = "Total: "
	chk.Str(structure.GetF1(1, 2, 3), "Total: 6")
}

func Test_FAIL_Example2(t *testing.T) {
	chk := sztest.CaptureStdout(t)
	defer chk.Release()

	chk.FailFast(false) // Run all tests before exiting function.

	chk.Int(example2.TimesTwo(2), 5, "2+2=5 (is true for big values of two)")
	chk.Str(
		"New in Got"+" Similar in (1) both",
		" Similar in (2) both"+", new in Wnt",
	)

	fmt.Println("This output line will be different")
	chk.Stdout("This output line is different")

	structure := new(example2.StructureType)
	structure.F1 = "Total: "
	chk.Str(structure.GetF1(1, 2, 3), "Sum: 6")
}
