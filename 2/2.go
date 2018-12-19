package main

import (
	"bufio"
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

	twoTotal := 0
	threeTotal := 0

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		countMap := make(map[rune]int)
		for _, currRune := range line {
			if _, present := countMap[currRune]; present {
				countMap[currRune] += 1
			} else {
				countMap[currRune] = 1
			}
		}
		twoFlag := true
		threeFlag := true
		for _, v := range countMap {
			if v == 2 && twoFlag{
				twoTotal += 1
				twoFlag = false
			}
			if v == 3 && threeFlag{
				threeTotal += 1
				threeFlag = false
			}
		}
	}
	fmt.Println(twoTotal * threeTotal)
}

func dayTwo() {
	for col := 0; col < 26; col += 1 {
		cutLines := make(map[string]bool)
		f, err := os.Open("input.txt")
		if err != nil {
			log.Fatal(err)
		}
		s :=  bufio.NewScanner(f)
		for s.Scan() {
			cutLine := s.Text()[:col] + s.Text()[col+1:]
			if _, present := cutLines[cutLine]; present {
				fmt.Println(cutLine)
				f.Close()
				return
			}
			cutLines[cutLine] = true
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

