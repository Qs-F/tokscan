package main

import (
	"fmt"
	"log"

	tokscan "github.com/Qs-F/titech-advanced-software-engineering"
)

func main() {
	output, err := tokscan.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(output)
}
