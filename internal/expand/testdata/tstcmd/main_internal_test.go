package main

import (
	"strings"
	"testing"

	"github.com/dancsecs/sztestlog"
)

func Test_PASS_Run(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	main()

	chk.Stdout("Run complete.")
}

func Test_PASS_Help(t *testing.T) {
	chk := sztestlog.CaptureStdout(t)
	defer chk.Release()

	chk.SetArgs("tstcmd", "--help")

	main()

	chk.Stdout(
		strings.TrimRight(usageInfo[1:], "\n"),
		"",
	)
}
