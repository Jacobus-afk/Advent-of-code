package main

import (
	aoc "aoc-24/lib"
	"fmt"
	"math"
	"reflect"
)

var PossibleDirections = [4][2]int{
	{0, 1}, {1, 0}, {-1, 0}, {0, -1},
}

func createMaze(data []string) (map[[2]int]string, [2]int, [2]int) {
	maze := map[[2]int]string{}
	start := [2]int{-1, -1}
	dimensions := [2]int{len(data[0]), len(data)}

	for posy, line := range data {
		for posx, char := range line {
			maze[[2]int{posx, posy}] = string(char)
			if char == 'S' {
				start = [2]int{posx, posy}
			}
		}
	}
	return maze, start, dimensions
}

func tryDirection(
	maze map[[2]int]string,
	currPos, prevPos, prevDir [2]int,
	score int,
	pathScore map[int]bool,
	dimensions [2]int,
) {
	updatedMap := map[[2]int]string{}
	for k, v := range maze {
		updatedMap[k] = v
	}
	// possibleMoves := [][2]int{}
	for _, dir := range PossibleDirections {
		localScore := score
		xpos := dir[0] + currPos[0]
		ypos := dir[1] + currPos[1]

		if xpos == prevPos[0] && ypos == prevPos[1] {
			// discard, we're not moving back on path
			continue
		}

		nextMazeBlock := updatedMap[[2]int{xpos, ypos}]
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

		if nextMazeBlock == "E" {
			pathScore[localScore] = true

			fmt.Println("")
			for posy := range dimensions[1] {
			  for posx := range dimensions[0] {
			    fmt.Print(updatedMap[[2]int{posx, posy}])
			  }
			  fmt.Println("")
			}
			fmt.Println("")
      fmt.Println(pathScore)

			return
		}
		// fmt.Println("updated score", currPos, localScore)

		updatedMap[currPos] = "X"
		tryDirection(
			updatedMap,
			[2]int{xpos, ypos},
			currPos,
			dir,
			localScore,
			pathScore,
			dimensions,
		)

		// possibleMoves = append(possibleMoves, [2]int{xpos, ypos})

	}
	fmt.Println("dead end")
	// not needed i think
	// if len(possibleMoves) == 0 {
	// 	// this path is a dead end
	// 	return
	// }
}

func TraverseMaze(maze map[[2]int]string, start, dimensions [2]int) int {
	pathScore := map[int]bool{}
	minScore := math.MaxInt64

	tryDirection(maze, start, [2]int{-1, -1}, [2]int{1, 0}, 0, pathScore, dimensions)
	fmt.Print("")

	for score := range pathScore {
		if minScore > score {
			minScore = score
		}
	}

	return minScore
}

func main() {
	data := aoc.ReadFileByLine("./data")

	maze, start, dimensions := createMaze(data)

	minScore := TraverseMaze(maze, start, dimensions)

	fmt.Println(minScore)
}
