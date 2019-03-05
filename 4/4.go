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

	s := bufio.NewScanner(f)
	allGuardStats := make(map[int][]int)
	currentGuard := 0
	CGCurrentStats := make([]int, 60)
	CGSleepTimes := make([]int, 0)
	CGWakeTimes := make([]int, 0)

	for s.Scan() {
		line := strings.Fields(s.Text())
		if len(line) == 6 {
			if currentGuard != 0 {
				for i := range CGSleepTimes {
					for time := CGSleepTimes[i]; time < CGWakeTimes[i]; i++ {
						CGCurrentStats[time] = 1
					}
				}
				if _, ok := allGuardStats[currentGuard]; ok {
					for i, val := range CGCurrentStats {
						allGuardStats[currentGuard][i] += val
					}
				} else {
					allGuardStats[currentGuard] = CGCurrentStats
				}
			}
			currentGuard, _ = strconv.Atoi(strings.TrimPrefix(line[3], "#"))
		} else {
			if line[2] == "wakes" {

			}
		}
	}
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

