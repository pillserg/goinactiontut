package main

import (
	"log"
	"os"

	"github.com/goinaction/code/chapter2/sample/search"
	_ "github.com/goinaction/code/chapter2/sample/search"
)

// init is called prior to main
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

func main() {
	// Perform the search for the specified term.
	search.Run("president")
}
