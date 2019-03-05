package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)
func dayOne() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Key = (row * column)
	// Value = number of fabrics that use this square
	overlapTracker := make(map[int]int)

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := strings.Fields(s.Text())
		yOffsetStr := strings.Split(line[2], ",")[1]
		yOffset, _ := strconv.Atoi(yOffsetStr[:len(yOffsetStr)-1])
		xOffset, _ := strconv.Atoi(strings.Split(line[2], ",")[0])
		xSize, _ := strconv.Atoi(strings.Split(line[3], "x")[0])
		ySize, _ := strconv.Atoi(strings.Split(line[3], "x")[1])
		for x := xOffset; x < (xOffset + xSize); x++ {
			for y := yOffset; y < (yOffset + ySize); y++ {
				if _, ok := overlapTracker[x * 1000 + y]; ok {
					overlapTracker[x * 1000 + y] += 1
				} else {
					overlapTracker[x * 1000 + y] = 1
				}
			}
		}
	}
	overLapCounter := 0
	for _, value := range overlapTracker {
		if value > 1 {
			overLapCounter += 1
		}
	}
	fmt.Println(overLapCounter)
}

func dayTwo() {
	// Key = (row * column)
	// Value = number of fabrics that use this square
	overlapTracker := make(map[int]int)

	// first pass - map out claims to overlap tracker
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := strings.Fields(s.Text())
		yOffsetStr := strings.Split(line[2], ",")[1]
		yOffset, _ := strconv.Atoi(yOffsetStr[:len(yOffsetStr)-1])
		xOffset, _ := strconv.Atoi(strings.Split(line[2], ",")[0])
		xSize, _ := strconv.Atoi(strings.Split(line[3], "x")[0])
		ySize, _ := strconv.Atoi(strings.Split(line[3], "x")[1])
		for x := xOffset; x < (xOffset + xSize); x++ {
			for y := yOffset; y < (yOffset + ySize); y++ {
				if _, ok := overlapTracker[x * 1000 + y]; ok {
					overlapTracker[x * 1000 + y] += 1
				} else {
					overlapTracker[x * 1000 + y] = 1
				}
			}
		}
	}
	f.Close()

	// second pass - find claim where every mapped square equals one in overlapTracker
	f, err = os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s = bufio.NewScanner(f)
	for s.Scan() {
		line := strings.Fields(s.Text())
		yOffsetStr := strings.Split(line[2], ",")[1]
		yOffset, _ := strconv.Atoi(yOffsetStr[:len(yOffsetStr)-1])
		xOffset, _ := strconv.Atoi(strings.Split(line[2], ",")[0])
		xSize, _ := strconv.Atoi(strings.Split(line[3], "x")[0])
		ySize, _ := strconv.Atoi(strings.Split(line[3], "x")[1])
		for x := xOffset; x < (xOffset + xSize); x++ {
			for y := yOffset; y < (yOffset + ySize); y++ {
				if overlapTracker[x * 1000 + y] != 1 {
					goto overLapFound
				}
			}
		}
		fmt.Println(line[0])
		overLapFound:
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

