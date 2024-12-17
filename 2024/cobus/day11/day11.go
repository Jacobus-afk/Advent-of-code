package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Blinks(blinkCount int, stones []string) []string {
	// memoizations := map[string]string{}
	// stoneCount := map[string]int{}

	tally := 0
  stoneMap := map[string]int{}
  for _, stone := range stones {
    stoneMap[stone]++
  }

	for t := 0; t < blinkCount; t++ {
		stoneReg := map[string]int{}
		for stone, count := range stoneMap {
			if stone == "0" {
				stoneReg["1"]+=count
			} else if len(stone)%2 == 0 {
				half := len(stone) / 2
				part1 := stone[:half]
				part2 := stone[half:]
				number, _ := strconv.Atoi(part2)
				part2 = strconv.Itoa(number)
				stoneReg[part1]+=count
				stoneReg[part2]+=count
			} else {
				number, _ := strconv.Atoi(stone)
				multiplied := number * 2024
				answer := strconv.Itoa(multiplied)
				stoneReg[answer]+=count
			}
		}
    // fmt.Println(stoneReg)
    stoneMap = stoneReg
	}

  for _, count := range stoneMap {
    tally += count
  }

	fmt.Println(tally)

	return stones
}

func main() {
	stones := []string{}
	file, _ := os.Open("./data")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		stones = strings.Fields(line)
	}

	stoneReg := Blinks(75, stones)
	fmt.Println(len(stoneReg))
}
