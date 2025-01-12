package main

import (
	aoc "aoc-24/lib"
	"fmt"
	"reflect"
)

var PossibleDirections = [4][2]int{
	{0, 1}, {1, 0}, {-1, 0}, {0, -1},
}

func createMaze(data []string) (map[[2]int]string, [2]int, [2]int, [2]int) {
	maze := map[[2]int]string{}
	start := [2]int{-1, -1}
	end := [2]int{-1, -1}
	dimensions := [2]int{len(data[0]), len(data)}

	for posy, line := range data {
		for posx, char := range line {
			maze[[2]int{posx, posy}] = string(char)
			if char == 'S' {
				start = [2]int{posx, posy}
			}
			if char == 'E' {
				end = [2]int{posx, posy}
			}
		}
	}
	return maze, start, end, dimensions
}

func tryDirection(
	maze map[[2]int]string,
	currPos, prevDir [2]int,
	score int,
	dimensions [2]int,
	nodePoints map[[2]int]int,
) {
	// updatedMap := map[[2]int]string{}
	// for k, v := range maze {
	// 	updatedMap[k] = v
	// }
	// possibleMoves := [][2]int{}
	// localPoints := map[[2]int]bool{}
	// for k,v := range dejavuPoints {
	//   localPoints[k] = v
	// }
	for _, dir := range PossibleDirections {
		localScore := score
		xpos := dir[0] + currPos[0]
		ypos := dir[1] + currPos[1]
		newPos := [2]int{xpos, ypos}

		// if point := localPoints[newPos]; point {
		//   // discard, we're going in circles
		//   continue
		// }
		//
		// localPoints[newPos] = true

		// if xpos == prevPos[0] && ypos == prevPos[1] {
		// 	// discard, we're not moving back on path
		// 	continue
		// }

		nextMazeBlock := maze[newPos]
		// fmt.Println("current pos", currPos, "next block", nextMazeBlock, xpos, ypos, dir)

		if nextMazeBlock == "#" || nextMazeBlock == "X" {
			// discard, we're not moving through hedges
			continue
		}

		if reflect.DeepEqual(prevDir, dir) {
			// in same direction
			localScore += 1
		} else {
			// rotated
			localScore += 1001
		}

		currentNodeScore, exists := nodePoints[newPos]
		if exists && localScore > currentNodeScore {
			// discard, already a path with smaller score at this point
			continue
		}
		nodePoints[newPos] = localScore

		if nextMazeBlock == "E" {
			// pathScore[localScore] = true

			// fmt.Println("")
			// for posy := range dimensions[1] {
			//   for posx := range dimensions[0] {
			//     fmt.Print(maze[[2]int{posx, posy}])
			//   }
			//   fmt.Println("")
			// }
			// fmt.Println("")
			//    fmt.Println(pathScore)

			continue
		}
		// fmt.Println("updated score", currPos, localScore)

		// updatedMap[currPos] = "X"
		tryDirection(
			maze,
			newPos,
			dir,
			localScore,
			dimensions,
			nodePoints,
		)

		// possibleMoves = append(possibleMoves, [2]int{xpos, ypos})

	}
	// fmt.Println("dead end")
	// not needed i think
	// if len(possibleMoves) == 0 {
	// 	// this path is a dead end
	// 	return
	// }
}

func TraverseMaze(maze map[[2]int]string, start, end, dimensions [2]int) int {
	nodePoints := map[[2]int]int{}
	// pathScore := map[int]bool{}
	// minScore := math.MaxInt64

	tryDirection(maze, start, [2]int{1, 0}, 0, dimensions, nodePoints)
	// fmt.Println(dejavuPoints[[2]int{1,13}], dejavuPoints[[2]int{1,9}], dejavuPoints[[2]int{3,9}], dejavuPoints[[2]int{3,7}], dejavuPoints[[2]int{11,7}])
	fmt.Println(nodePoints)
	// fmt.Println(pathScore)
	fmt.Println(start)
	fmt.Print("")

	// for score := range pathScore {
	// 	if minScore > score {
	// 		minScore = score
	// 	}
	// }

	return nodePoints[end]
}

func main() {
	data := aoc.ReadFileByLine("./data")

	maze, start, end, dimensions := createMaze(data)

	minScore := TraverseMaze(maze, start, end, dimensions)

	fmt.Println(minScore)
}
