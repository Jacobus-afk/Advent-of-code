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
	for _, dir := range PossibleDirections {
		localScore := score
		xpos := dir[0] + currPos[0]
		ypos := dir[1] + currPos[1]
		newPos := [2]int{xpos, ypos}

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
			continue
		}

		tryDirection(
			maze,
			newPos,
			dir,
			localScore,
			dimensions,
			nodePoints,
		)

	}
}

func TraverseMaze(maze map[[2]int]string, start, end, dimensions [2]int) int {
	nodePoints := map[[2]int]int{}

	tryDirection(maze, start, [2]int{1, 0}, 0, dimensions, nodePoints)
	fmt.Println(nodePoints)
	fmt.Println(start)
	fmt.Print("")

	return nodePoints[end]
}

func main() {
	data := aoc.ReadFileByLine("./data")

	maze, start, end, dimensions := createMaze(data)

	minScore := TraverseMaze(maze, start, end, dimensions)

	fmt.Println(minScore)
}
