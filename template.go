package main

import (
	"fmt"
	"log"
	"os"
)
func dayOne() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fmt.Println("TODO")
}

func dayTwo() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fmt.Println("TODO")
}

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s (one | two)", os.Args[0])
		return
	}

	if os.Args[1] == "one" {
		dayOne()
	} else if os.Args[1] == "two" {
		dayTwo()
	} else {
		fmt.Printf("Usage: %s (one | two)", os.Args[0])
	}


}

