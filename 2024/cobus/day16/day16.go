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
}

func dirPresent(dirs [][2]int, dir [2]int) bool {
	for _, way := range dirs {
		if reflect.DeepEqual(way, dir) {
			return true
		}
	}
	return false
}

func loopAttempt(start [2]int, maze map[[2]int]string) int {
	nodeList := []Node{{dir: 1, pos: start, score: 0}}
	node := Node{}
	// nodeReg := map[[2]int][][2]int{}
	possibleScores := []int{}
  dirScores := map[[3]int]int{}

	for len(nodeList) > 0 {
		// fmt.Println("trying nodes", nodeList)
		// fmt.Println(nodeReg)
		node, nodeList = nodeList[0], nodeList[1:]
		// fmt.Println("trying node", node)

		dir := PossibleDirections[node.dir]
		// score := node.score + 1
		// newBlock := maze[newPos]

    dsKey := [3]int{node.pos[0], node.pos[1], node.dir}
    dScore , exists := dirScores[dsKey]
    if exists {
      if node.score > dScore {
        continue
      }
    }
    dirScores[dsKey] = node.score

		// dirs, exists := nodeReg[node.pos]
		// if exists {
		// 	if dirPresent(dirs, dir) {
		// 		continue
		// 	}
		// }
		// nodeReg[node.pos] = append(nodeReg[node.pos], dir)

		currBlock := maze[node.pos]

		if currBlock == "#" {
			continue
		}

		if currBlock == "E" {
			possibleScores = append(possibleScores, node.score)
			continue
		}

		newPos := [2]int{node.pos[0] + dir[0], node.pos[1] + dir[1]}
		newNode := Node{dir: node.dir, pos: newPos, score: node.score + 1}
		newNode90 := Node{dir: (node.dir + 1) % 4, pos: node.pos, score: node.score + 1000}
		newNode270 := Node{dir: (node.dir + 3) % 4, pos: node.pos, score: node.score + 1000}
		nodeList = append(nodeList, newNode, newNode90, newNode270)

		// fmt.Println(nodeList, node)
		// for _, dir := range PossibleDirections {
		// 	score := node.score + 1
		// 	xpos := node.pos[0] + dir[0]
		// 	ypos := node.pos[1] + dir[1]
		// 	newPos := [2]int{xpos, ypos}
		// 	newBlock := maze[newPos]
		//
		// 	dirs, exists := nodeReg[newPos]
		// 	if exists {
		//       if dirPresent(dirs, dir) {
		//         continue
		//       }
		// 	}
		// 	// fmt.Println("nodeReg", nodeReg)
		//
		// 	nodeReg[newPos] = append(nodeReg[newPos], dir)
		//
		// 	// fmt.Println("trying node", newBlock, newPos, dir)
		// 	if newBlock == "#" {
		// 		continue
		// 	}
		//
		// 	if newBlock == "E" {
		// 		possibleScores = append(possibleScores, score)
		// 		continue
		// 	}
		//
		// 	if !reflect.DeepEqual(dir, node.dir) {
		// 		score += 1000
		// 	}
		//
		// 	newNode := Node{dir: dir, pos: newPos, score: score}
		// 	nodeList = append(nodeList, newNode)
		//
		// }
	}
	minScore := math.MaxInt64
	for _, score := range possibleScores {
		if score < minScore {
			minScore = score
		}
	}
	fmt.Println(possibleScores)
	return minScore
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

	return loopAttempt(start, maze)
}

func main() {
	data := aoc.ReadFileByLine("./data")

	maze, start, end, dimensions := createMaze(data)

	minScore := TraverseMaze(maze, start, end, dimensions)

	fmt.Println(minScore)
}
