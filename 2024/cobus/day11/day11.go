package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Blinks(blinkCount int, stones []string) []string {
	memoizations := map[string]string{}
	for t := 0; t < blinkCount; t++ {
		for i := 0; i < len(stones); i++ {
			stoneLen := len(stones[i])
			halfStone := stoneLen / 2
			if stones[i] == "0" {
				stones[i] = "1"
			} else if stoneLen%2 == 0 {
				part1 := stones[i][:halfStone]
				part2 := stones[i][halfStone:]
				number, _ := strconv.Atoi(part2)
				part2 = strconv.Itoa(number)
				// fmt.Println(part1, part2)

				stones[i] = part1
				stones = append(stones[:i+1], stones[i:]...)
				stones[i+1] = part2
				i++
			} else {
				answer, ok := memoizations[stones[i]]
				if !ok {
          number, _ := strconv.Atoi(stones[i])
          multiplied := number * 2024
          answer = strconv.Itoa(multiplied)
					memoizations[stones[i]] = answer
				}
				stones[i] = answer
			}
			// fmt.Println(stones)

		}
	}
  fmt.Println(memoizations)
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

	stoneReg := Blinks(25, stones)
	fmt.Println(len(stoneReg))
}
