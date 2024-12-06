package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

func getStartingPosition(grid []string) ([2]int, [2]int) {
	for ypos, line := range grid {
		if xpos := strings.Index(line, ">"); xpos >= 0 {
			return [2]int{xpos, ypos}, [2]int{1, 0}
		}
		if xpos := strings.Index(line, "^"); xpos >= 0 {
			return [2]int{xpos, ypos}, [2]int{0, -1}
		}
	}
	return [2]int{-1, -1}, [2]int{-1, -1}
}

func obstacleHasBeenReachedFromThisDirection(
	position, direction [2]int,
	boundariesMap map[[2]int][][2]int,
) bool {
	directionsApproached, ok := boundariesMap[position]

	if !ok {
		return false
	}

	for _, directionApproached := range directionsApproached {
		if reflect.DeepEqual(direction, directionApproached) {
			return true
		}
	}

	return false
}

func addDirectionToMap(position, direction [2]int, boundariesMap map[[2]int][][2]int) {
	directionsApproached, ok := boundariesMap[position]

	if !ok {
		boundariesMap[position] = [][2]int{direction}
	} else {
		directionsApproached = append(directionsApproached, direction)
		boundariesMap[position] = directionsApproached
	}
}

func TrackMovement(grid []string, startPosition, direction [2]int) (int, bool) {
	movements := 1
	boundariesMap := make(map[[2]int][][2]int)

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
			if obstacleHasBeenReachedFromThisDirection(
				[2]int{posx, posy},
				direction,
				boundariesMap,
			) {
				return movements, true
			}

			addDirectionToMap([2]int{posx, posy}, direction, boundariesMap)

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

	return movements, false
}

func FindLoops(grid []string, startPosition, direction [2]int) int {
	loopCount := 0
	scratchGrid := make([]string, len(grid))
	copy(scratchGrid, grid)
	for ypos, line := range grid {
		for xpos := 0; xpos < len(line); xpos++ {
			if line[xpos] == '.' {
				scratchGrid[ypos] = line[:xpos] + "#" + line[xpos+1:]

				_, potentialObstruction := TrackMovement(scratchGrid, startPosition, direction)
				if potentialObstruction {
					loopCount++
				}
				copy(scratchGrid, grid)
			}
		}
	}
	return loopCount
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
	// fmt.Println(grid)

	startPosition, direction := getStartingPosition(grid)
	fmt.Printf("pos: %d, direction: %d\n", startPosition, direction)

	potentialObstructions := FindLoops(grid, startPosition, direction)

	distinctPositions, _ := TrackMovement(grid, startPosition, direction)
	fmt.Println(distinctPositions)
	fmt.Println(potentialObstructions)
}
