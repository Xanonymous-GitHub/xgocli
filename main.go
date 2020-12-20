package main

import (
	"flag"
	"log"
)

func main() {
	var name string

	flag.StringVar(&name, "name", "X Go Cli", "description")
	flag.StringVar(&name, "n", "X Go Cli", "description")

	flag.Parse()

	log.Printf("name: %s", name)
}