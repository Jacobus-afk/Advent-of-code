package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	// "unicode"
	"regexp"
)

var alphanumeric = regexp.MustCompile("^[a-zA-Z0-9_]*$")

func calculateAntinodeDirs(refRecordPos, compareRecordPos [2]int) ([2]int, [2]int) {
	xdir := compareRecordPos[0] - refRecordPos[0]
	ydir := compareRecordPos[1] - refRecordPos[1]

	return [2]int{xdir, ydir}, [2]int{xdir * (-1), ydir * (-1)}
}

func getNodePos(pos, dir, gridLen [2]int) ([2]int, error) {
	xgridlen := gridLen[0]
	ygridlen := gridLen[1]
	xpos := pos[0]
	ypos := pos[1]
	xdir := dir[0]
	ydir := dir[1]

	xnodepos := xpos + xdir
	ynodepos := ypos + ydir

	if xnodepos < 0 || xnodepos >= xgridlen || ynodepos < 0 || ynodepos >= ygridlen {
		return [2]int{}, errors.New("node out of bounds")
	}

	return [2]int{xnodepos, ynodepos}, nil
}

func calculateAntinodePositions(
	refRecordPos, compareRecordPos, antinodeDirForRef, antinodeDirForCompare, gridLen [2]int,
	antinodeMap map[[2]int]string, char rune,
) {
	for {
		refNode, err := getNodePos(refRecordPos, antinodeDirForRef, gridLen)
		if err == nil {
			antinodeMap[refNode] = string(char)
			copy(refRecordPos[:], refNode[:])
		} else {
			break
		}
	}

	for {
    refNode, err := getNodePos(compareRecordPos, antinodeDirForCompare, gridLen)
		if err == nil {
			antinodeMap[refNode] = string(char)
      copy(compareRecordPos[:], refNode[:])
		} else {
			break
		}

	}
}

func mapAntinodes(nodeRecord [][2]int, antinodeMap map[[2]int]string, gridLen [2]int, char rune) {
	for num, refRecordPos := range nodeRecord[:len(nodeRecord)-1] {
		for _, compareRecordPos := range nodeRecord[num+1:] {
			// fmt.Println(refRecordPos, compareRecordPos)
			antinodeDirForCompare, antinodeDirForRef := calculateAntinodeDirs(
				refRecordPos,
				compareRecordPos,
			)

			calculateAntinodePositions(
				refRecordPos,
				compareRecordPos,
				antinodeDirForRef,
				antinodeDirForCompare,
				gridLen,
				antinodeMap,
				char,
			)

		}
	}
}

func findAntinodes(
	partGrid []string,
	char rune,
	charPos, gridLen [2]int,
	antinodeMap map[[2]int]string,
) {
	nodeRecord := [][2]int{charPos}

	// part2
	antinodeMap[charPos] = string(char)

	for subypos, line := range partGrid {
		// this wont work if there are more than one of the same char on the same line..
		for xpos, potentialAntenna := range line {
			if potentialAntenna != char || (subypos == 0 && xpos == charPos[0]) {
				// don't want to add first antenna again
				continue
			}
			// xpos := strings.Index(line, string(char))
			// if xpos < 0 {
			// 	continue
			// }
			ypos := charPos[1] + subypos
			// fmt.Println(xpos, ypos)

			nodeRecord = append(nodeRecord, [2]int{xpos, ypos})
			// part 2
			antinodeMap[[2]int{xpos, ypos}] = string(char)
			// part 2 done
			mapAntinodes(nodeRecord, antinodeMap, gridLen, char)
		}
	}
}

func AntinodeCreation(grid []string) int {
	antinodeMap := make(map[[2]int]string)
	charHandledMap := make(map[rune]bool)
	// fmt.Println(len(grid[0]), len(grid))
	gridLen := [2]int{len(grid[0]), len(grid)}

	for ypos, line := range grid {
		for xpos, char := range line {

			// if _, ok := charHandledMap[char]; ok || !unicode.IsLetter(char) {
			if _, ok := charHandledMap[char]; ok || !alphanumeric.MatchString(string(char)) {
				continue
			}
			// fmt.Println(string(char), xpos, ypos)
			charHandledMap[char] = true

			findAntinodes(grid[ypos:], char, [2]int{xpos, ypos}, gridLen, antinodeMap)
		}
	}
	// for key, antiNode := range antinodeMap {
	// 	if antiNode == "J" {
	// 		fmt.Println(key, antiNode)
	// 	}
	// }

	return len(antinodeMap)
}

func main() {
	fmt.Println("start")
	grid := []string{}

	file, _ := os.Open("./data")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)
	}
	// fmt.Println(grid)

	nodes := AntinodeCreation(grid)
	fmt.Println(nodes)
}
