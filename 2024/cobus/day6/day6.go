package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getStartingPosition(grid []string) ([2]int, [2]int) {
	// probably not needed
	// fmt.Printf("line: %s\n", line)
	for ypos, line := range grid {
		if xpos := strings.Index(line, ">"); xpos >= 0 {
			return [2]int{xpos, ypos}, [2]int{1, 0}
		}
		// if pos := strings.Index(line, "<"); pos > 0 {
		//   return pos, [2]int{-1, 0}
		// }
		if xpos := strings.Index(line, "^"); xpos >= 0 {
			return [2]int{xpos, ypos}, [2]int{0, -1}
		}
	}
	return [2]int{-1, -1}, [2]int{-1, -1}
}

func TrackMovement(grid []string, startPosition, direction [2]int) int {
	movements := 1

	posx := startPosition[0]
	posy := startPosition[1]
	tmpLine := grid[posy]
	grid[posy] = tmpLine[:posx] + "X" + tmpLine[posx+1:]

	for {
		switch grid[posy][posx] {
		case '.':
			movements += 1
			tmpLine = grid[posy]
			grid[posy] = tmpLine[:posx] + "X" + tmpLine[posx+1:]
		case '#':
			// move a step back
			posx -= direction[0]
			posy -= direction[1]
			// change direction 90 degrees right
			tmpPos := direction[1]
			if tmpPos != 0 {
				tmpPos *= -1
			}
			direction[1] = direction[0]
			direction[0] = tmpPos
		}
		posx += direction[0]
		posy += direction[1]
		if posy < 0 || posy >= len(grid) {
			break
		}

		if posx < 0 || posx >= len(grid[posy]) {
			break
		}
	}

	return movements
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

	startPosition, direction := getStartingPosition(grid)
	fmt.Printf("pos: %d, direction: %d\n", startPosition, direction)

	distinctPositions := TrackMovement(grid, startPosition, direction)
	fmt.Println(distinctPositions)
}
