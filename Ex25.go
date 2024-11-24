package main

import (
	"fmt"
	"log"
)

func main() {
	n := 0
	fmt.Print("Enter integer: ")
	_. err := fmt.Scan(&n)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("You have enter: %d\n", n)
}

