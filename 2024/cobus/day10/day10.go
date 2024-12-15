package main

import (
	"bufio"
	"fmt"
	"os"
)

// import "fmt"

func buildMap(trailMap []string) [][]int {
	trail := [][]int{}

	for _, line := range trailMap {
		row := []int{}
		for _, char := range line {
			if char == '.' {
				row = append(row, -1)
			} else {
				row = append(row, int(char-'0'))
			}
		}
		trail = append(trail, row)
	}

	return trail
}

var PossibleDirections = [4][2]int{
	{0, 1}, {1, 0}, {-1, 0}, {0, -1},
}

func chartPath(endMap map[[2]int]bool, height int, position [2]int, trail [][]int, rating *int) {
	rowLen := len(trail)
	colLen := len(trail[0])
	for _, dir := range PossibleDirections {
		posx := position[0] + dir[0]
		posy := position[1] + dir[1]

		if posx < 0 || posx >= colLen || posy < 0 || posy >= rowLen {
			continue
		}

		if trail[posy][posx] == height {
			newPos := [2]int{posx, posy}
			if height == 9 {
        *rating += 1
				endMap[newPos] = true
			} else {
				chartPath(endMap, height+1, newPos, trail, rating)
			}
		}
	}
}

func FindPath(trail [][]int) int {
	tally := 0
  rating := 0
	for posy, line := range trail {
		for posx, height := range line {
			position := [2]int{posx, posy}
			if height == 0 {
				endMap := map[[2]int]bool{}
				chartPath(endMap, height+1, position, trail, &rating)
				tally += len(endMap)
			}
		}
	}
  fmt.Println(rating)
	return tally
}

func main() {
	grid := []string{}

	file, _ := os.Open("./data")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)
	}
  trail := buildMap(grid)

  trailheads := FindPath(trail)
  fmt.Println(trailheads)
}
