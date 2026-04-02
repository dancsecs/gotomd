package main

import (
	"fmt"
	"os"
)

const usageInfo = `
# Usage
	tstcmd [--help]
	
	--help
		Display usage information
`

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--help" {
		fmt.Fprintln(os.Stdout, usageInfo[1:])
	} else {
		fmt.Fprintln(os.Stdout, "Run complete.")
	}
}
