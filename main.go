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
	if len(os.Args) != 2 {
		log.Fatal("search term not provided")
		os.Exit(1)
	}
	searchTerm := os.Args[1]
	log.Printf("searching for [%s]", os.Args[1])
	search.Run(searchTerm)
}
