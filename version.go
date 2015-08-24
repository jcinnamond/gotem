package main

import (
	"fmt"
	"os"
	"path"
)

// Returns the version
func version() string {
	return "1.0.0"
}

// Prints the version information and exits
func printVersion() {
	fmt.Printf("%s v%s\n", path.Base(os.Args[0]), version())
	os.Exit(0)
}
