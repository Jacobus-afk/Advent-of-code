package main

import (
	aoc "aoc-24/lib"
	"fmt"
	"math"
	"reflect"
)

var PossibleDirections = [4][2]int{
	{0, 1}, {1, 0}, {0, -1}, {-1, 0},
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

		if !reflect.DeepEqual(prevDir, dir) {
			// rotated
			localScore += 1000
		}
		localScore += 1

		currentNodeScore, exists := nodePoints[newPos]
		if exists && localScore >= currentNodeScore {
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

type Node struct {
	dir   int
	pos   [2]int
	score int
	path  [][2]int
}

func dirPresent(dirs [][2]int, dir [2]int) bool {
	for _, way := range dirs {
		if reflect.DeepEqual(way, dir) {
			return true
		}
	}
	return false
}

func loopAttempt(start [2]int, maze map[[2]int]string) (int, []Node) {
	node := Node{dir: 1, pos: start, score: 0}
	nodeList := []Node{node}
	possibleScores := []int{}
	possiblePaths := []Node{}
	dirScores := map[[3]int]int{}

	for len(nodeList) > 0 {
		node, nodeList = nodeList[0], nodeList[1:]

		dir := PossibleDirections[node.dir]

		dsKey := [3]int{node.pos[0], node.pos[1], node.dir}
		dScore, exists := dirScores[dsKey]
		if exists {
			if node.score > dScore {
				continue
			}
		}
		dirScores[dsKey] = node.score

		currBlock := maze[node.pos]

		if currBlock == "#" {
			continue
		}

		currNode := node
		if currBlock == "E" {
			possibleScores = append(possibleScores, node.score)
			possiblePaths = append(possiblePaths, currNode)
			continue
		}

		currNodePath := append(currNode.path, currNode.pos)
		newPos := [2]int{node.pos[0] + dir[0], node.pos[1] + dir[1]}
		newNode := Node{
			dir:   node.dir,
			pos:   newPos,
			score: node.score + 1,
			path:  currNodePath,
		}
		newNode90 := Node{
			dir:   (node.dir + 1) % 4,
			pos:   node.pos,
			score: node.score + 1000,
			path:  node.path,
		}
		newNode270 := Node{
			dir:   (node.dir + 3) % 4,
			pos:   node.pos,
			score: node.score + 1000,
			path:  node.path,
		}
		nodeList = append(nodeList, newNode, newNode90, newNode270)

	}
	minScore := math.MaxInt64
	for _, score := range possibleScores {
		if score < minScore {
			minScore = score
		}
	}
	fmt.Println(possibleScores)
	return minScore, possiblePaths
}

func TraverseMaze(maze map[[2]int]string, start, end, dimensions [2]int) int {
	// nodePoints := map[[2]int]int{}
	//
	// tryDirection(maze, start, [2]int{1, 0}, 0, dimensions, nodePoints)
	// fmt.Println(nodePoints)
	// fmt.Println(start)
	// fmt.Print("")
	//
	// return nodePoints[end]
	minScore, possiblePaths := loopAttempt(start, maze)

	pathTally := map[[2]int]bool{}
	for _, pathMap := range possiblePaths {
		if pathMap.score > minScore {
			continue
		}
		for _, path := range pathMap.path {
			pathTally[path] = true
		}
	}
	fmt.Println(len(pathTally) + 1)

	return minScore
}

func main() {
	data := aoc.ReadFileByLine("./data")

	maze, start, end, dimensions := createMaze(data)

	minScore := TraverseMaze(maze, start, end, dimensions)

	fmt.Println(minScore)
}
