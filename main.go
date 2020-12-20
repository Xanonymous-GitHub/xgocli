package main

import (
	"log"
	"xgocli/cmd"
)

func main() {
	err := cmd.Execute()

	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
