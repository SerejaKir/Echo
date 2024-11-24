package main

import (
	"fmt"
	"log"
)

func main() {
	n := "" 
	fmt.Print("Enter data: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("You have enter data: %v\n", n)
}

