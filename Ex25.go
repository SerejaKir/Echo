package main

import (
	"fmt"
	"log"
)

func main() {
	n := "" 
	fmt.Print("Enter integer: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("You have enter: %v\n", n)
}

