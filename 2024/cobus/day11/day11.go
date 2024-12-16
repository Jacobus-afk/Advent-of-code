package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func blinkChecks(
	stone string,
	multMemos map[string]string,
	splitMemos map[string][2]string,
) (string, string) {
	if stone == "0" {
		return "1", ""
	}
	if stoneLen := len(stone); stoneLen%2 == 0 {
		parts, ok := splitMemos[stone]
		if !ok {
			halfStone := stoneLen / 2
			part1 := stone[:halfStone]
			part2 := stone[halfStone:]
			number, _ := strconv.Atoi(part2)
			part2 = strconv.Itoa(number)
			parts = [2]string{part1, part2}
			splitMemos[stone] = parts
		}

		return parts[0], parts[1]

	}
	answer, ok := multMemos[stone]
	if !ok {
		number, _ := strconv.Atoi(stone)
		multiplied := number * 2024
		answer = strconv.Itoa(multiplied)
		multMemos[stone] = answer
	}
	return answer, ""
}

func Blinks(blinkCount int, stones []string) []string {
	multMemos := map[string]string{}
  splitMemos := map[string][2]string{}
	for t := 0; t < blinkCount; t++ {
		fmt.Println(t)
		for i := 0; i < len(stones); i++ {
      part1, part2 := blinkChecks(stones[i], multMemos, splitMemos)
      stones[i] = part1
      if part2 != "" {
				stones = append(stones[:i+1], stones[i:]...)
        stones[i+1] = part2
        i++
      }
			// fmt.Println(stones)

		}
	}
	// fmt.Println(memoizations)
	return stones
}

func main() {
	stones := make([]string, 0, 9000000)
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
