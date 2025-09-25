// Provides a simple example to test go run functionality.
package main

import (
	"fmt"
	"os"
)

//nolint:forbidigo // Printing is ok.
func main() {
	fmt.Printf("Running with %d arguments\n", len(os.Args)-1)

	for _, a := range os.Args[1:] {
		fmt.Printf("%s\n", a)
	}
}
