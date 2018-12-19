package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func dayOne() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	total := 0
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		if string(line[0]) == "+" {
			addVal, err := strconv.ParseInt(string(line[1:]), 10, 0)
			if err != nil {
				log.Fatal(err)
			}
			total += int(addVal)
		} else if string(line[0]) == "-"{
			subVal, err := strconv.ParseInt(string(line[1:]), 10, 0)
			if err != nil {
				log.Fatal(err)
			}
			total -= int(subVal)
		} else {
			fmt.Println("ERROR!")
			fmt.Println(line[0])
		}
	}
	fmt.Println(total)
}

func dayTwo() {

	total := 0
	seenValues := make(map[int]bool)
	for {
		f, err := os.Open("input.txt")
		if err != nil {
			log.Fatal(err)
		}
		s := bufio.NewScanner(f)
		for s.Scan() {
			line := s.Text()
			if string(line[0]) == "+" {
				addVal, err := strconv.ParseInt(string(line[1:]), 10, 0)
				if err != nil {
					log.Fatal(err)
				}
				total += int(addVal)
				if _, present := seenValues[total]; present {
					fmt.Println(total)
					f.Close()
					return
				}
				seenValues[total] = true
			} else if string(line[0]) == "-" {
				subVal, err := strconv.ParseInt(string(line[1:]), 10, 0)
				if err != nil {
					log.Fatal(err)
				}
				total -= int(subVal)
				if _, present := seenValues[total]; present {
					fmt.Println(total)
					f.Close()
					return
				}
				seenValues[total] = true
			} else {
				fmt.Println("ERROR!")
				fmt.Println(line[0])
			}
		}
		f.Close()
	}
}

func main() {
	if len(os.Args) != 2 {
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
