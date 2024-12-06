package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
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

func addDirectionToMap(position, direction [2]int, boundariesMap map[[2]int][][2]int)  {
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
	// potentialObstructions := false
	// loopCheck := false

	posx := startPosition[0]
	posy := startPosition[1]
	tmpLine := grid[posy]
	grid[posy] = tmpLine[:posx] + "X" + tmpLine[posx+1:]

	for {
		switch grid[posy][posx] {
		// task2
		//  case 'O':
		//    // fmt.Printf("got an obstruction [%d %d]\n", posx, posy)
		//    if loopCheck{
		//      // fmt.Println("got the obstruction again")
		//      return movements, true
		//    }
		//    loopCheck = true
		// // move a step back
		// posx -= direction[0]
		// posy -= direction[1]
		// // change direction 90 degrees right
		// tmpPos := direction[1]
		// if tmpPos != 0 {
		// 	tmpPos *= -1
		// }
		// direction[1] = direction[0]
		// direction[0] = tmpPos

		case '.':
			movements += 1
			tmpLine = grid[posy]
			grid[posy] = tmpLine[:posx] + "X" + tmpLine[posx+1:]
			// case 'O':
			//   fmt.Printf("hit a previous obstacle [%d %d]", posx, posy)
			//   return movements, true
		case '#':
			// step 2 flag point as visited
			// tmpLine = grid[posy]
			// grid[posy] = tmpLine[:posx] + "O" + tmpLine[posx+1:]
			// fmt.Printf("updated grid: %s\n", grid)
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

			// fmt.Printf("direction change [%d %d]\n", posx, posy)
			// if grid[posy][posx] == 'X' {
			//      return movements, true
			// }

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
				// tmpLine := line
				scratchGrid[ypos] = line[:xpos] + "#" + line[xpos+1:]
				// fmt.Printf("line: %s\n", scratchGrid[ypos])

				_, potentialObstruction := TrackMovement(scratchGrid, startPosition, direction)
				if potentialObstruction {
					loopCount++
				}
				copy(scratchGrid, grid)
				// grid[ypos] = tmpLine

			}
		}
		// fmt.Printf("\t\tdone with line %d\n\n", ypos)
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
