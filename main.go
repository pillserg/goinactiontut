package main

import (
	_ "goinaction/chapter2/matchers"
	"goinaction/chapter2/search"
	"log"
	"os"
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
