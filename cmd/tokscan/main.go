package main

import (
	"log"

	tokscan "github.com/Qs-F/titech-advanced-software-engineering"
)

func main() {
	if err := tokscan.Run(); err != nil {
		log.Fatal(err)
	}
}
