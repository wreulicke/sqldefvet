package main

import (
	"log"
)

func mainInternal() error {
	return New().Execute()
}

func main() {
	if err := mainInternal(); err != nil {
		log.Fatal(err)
	}
}
