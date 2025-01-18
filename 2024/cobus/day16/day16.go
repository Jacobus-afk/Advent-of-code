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

// func updateNodeTreeMap(
// 	node Node,
// 	nodeTreeMap map[[3]int]map[[3]int]int,
// 	// maze map[[2]int]string,
// 	// newPos [2]int,
// ) {
// 	// newPosKey := [3]int{newPos[0], newPos[1], node.dir}
// 	nodePosKey := [3]int{node.pos[0], node.pos[1], node.dir}
// 	if node.parent == nil {
// 		nodeTreeMap[nodePosKey] = nil
// 		return
// 	}
// 	parentNode := *(node.parent)
// 	// fmt.Println("parentNode in fn", parentNode)
// 	nodeParentKey := [3]int{parentNode.pos[0], parentNode.pos[1], parentNode.dir}
// 	// prevPos := [2]int{node.pos[0] - dir[0], node.pos[1] - dir[1]}
// 	// if prevBlock := maze[prevPos]; prevBlock == "#" {
// 	// 	return
// 	// }
// 	// fmt.Println("adding parent node", nodeParentKey, "to node", nodePosKey)
// 	// fmt.Println("adding prev node", prevPos, "to", node)
// 	// fmt.Println("")
// 	nodeTree, ok := nodeTreeMap[nodePosKey]
// 	if ok {
// 		// fmt.Println(nodeTree)
// 		nodeTree[nodeParentKey] = parentNode.score
// 		nodeTreeMap[nodePosKey] = nodeTree
// 		// fmt.Println(node)
// 		// fmt.Println("newPos", newPos, "nodeTree", nodeTree)
// 	} else {
// 		nodeTreeMap[nodePosKey] = map[[3]int]int{nodeParentKey: parentNode.score}
// 	}
// }

func loopAttempt(start [2]int, maze map[[2]int]string) (int, []Node) {
	node := Node{dir: 1, pos: start, score: 0}
	// nodeTree := NodeTree{node: node}
	// nodeTreeMap := map[[3]int]map[[3]int]int{}
	nodeList := []Node{node}
	// nodeReg := map[[2]int][][2]int{}
	possibleScores := []int{}
	possiblePaths := []Node{}
	dirScores := map[[3]int]int{}

	for len(nodeList) > 0 {
		// fmt.Println("trying nodes", nodeList)
		// fmt.Println(nodeReg)
		node, nodeList = nodeList[0], nodeList[1:]
		// fmt.Println("trying node", node)
		// fmt.Println(nodeTreeMap)
		// fmt.Println("")

		dir := PossibleDirections[node.dir]
		// score := node.score + 1
		// newBlock := maze[newPos]

		dsKey := [3]int{node.pos[0], node.pos[1], node.dir}
		dScore, exists := dirScores[dsKey]
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

		currNode := node
		if currBlock == "E" {
			possibleScores = append(possibleScores, node.score)
			possiblePaths = append(possiblePaths, currNode)
			// fmt.Println("GOT END", node.score, node.path)
			// updateNodeTreeMap(node, nodeTreeMap)
			// updateNodeTreeMap(node, nodeTreeMap, maze, dir)
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
		// fmt.Println("added nodes for checking", newNode, newNode90, newNode270)
		// fmt.Println("newNode's parent", *(newNode.parent))
		// if node.parent != nil {
		// 	fmt.Println("node's parent", *(currNode.parent))
		// }
		nodeList = append(nodeList, newNode, newNode90, newNode270)

		// updateNodeTreeMap(currNode, nodeTreeMap)
	}
	minScore := math.MaxInt64
	for _, score := range possibleScores {
		if score < minScore {
			minScore = score
		}
	}
	fmt.Println(possibleScores)
	// fmt.Println(nodeTreeMap)
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

	// for i := range 4 {
	// 	endPtKey := [3]int{end[0], end[1], i}
	// 	endPoints := nodeTreeMap[endPtKey]
	// 	// fmt.Println("endPoints", endPoints)
	//
	// 	for point, score := range endPoints {
	// 		if score > minScore {
	// 			continue
	// 		}
	// 		bestTileCount := map[[2]int]bool{end: true}
	//
	// 		currPoint := point
	// 		pointList := [][3]int{currPoint}
	//
	// 		for len(pointList) > 0 {
	// 			currPoint, pointList = pointList[0], pointList[1:]
	// 			block := maze[[2]int{currPoint[0], currPoint[1]}]
	// 			if block == "S" {
	// 				break
	// 			}
	// 			currPointMap := nodeTreeMap[currPoint]
	// 			// fmt.Println("currPoint", currPoint, "currPointMap", currPointMap)
	// 			for pt := range currPointMap {
	// 				// if score > minScore {
	// 				//   continue
	// 				// }
	// 				pointList = append(pointList, pt)
	// 				bestTileCount[[2]int{pt[0], pt[1]}] = true
	// 				// fmt.Println(pt)
	// 			}
	//
	// 		}
	// 		fmt.Println(len(bestTileCount) + 1)
	// 	}
	//
	// }

	return minScore
}

func main() {
	data := aoc.ReadFileByLine("./data")

	maze, start, end, dimensions := createMaze(data)

	minScore := TraverseMaze(maze, start, end, dimensions)

	fmt.Println(minScore)
}
