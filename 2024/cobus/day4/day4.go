package main

import (
	// "fmt"
	"bufio"
	"fmt"
	"os"
)

type wordCoords struct {
	character   rune
	coordinates [2]int
}

var wordsCoordsXMAS = [8][3]wordCoords{
	{
		wordCoords{
			character:   'M',
			coordinates: [2]int{0, -1},
		},
		wordCoords{
			character:   'A',
			coordinates: [2]int{0, -2},
		},
		wordCoords{
			character:   'S',
			coordinates: [2]int{0, -3},
		},
	},

	{
		wordCoords{
			character:   'M',
			coordinates: [2]int{-1, -1},
		},
		wordCoords{
			character:   'A',
			coordinates: [2]int{-2, -2},
		},
		wordCoords{
			character:   'S',
			coordinates: [2]int{-3, -3},
		},
	},

	{
		wordCoords{
			character:   'M',
			coordinates: [2]int{1, -1},
		},
		wordCoords{
			character:   'A',
			coordinates: [2]int{2, -2},
		},
		wordCoords{
			character:   'S',
			coordinates: [2]int{3, -3},
		},
	},

	{
		wordCoords{
			character:   'M',
			coordinates: [2]int{0, 1},
		},
		wordCoords{
			character:   'A',
			coordinates: [2]int{0, 2},
		},
		wordCoords{
			character:   'S',
			coordinates: [2]int{0, 3},
		},
	},

	{
		wordCoords{
			character:   'M',
			coordinates: [2]int{1, 1},
		},
		wordCoords{
			character:   'A',
			coordinates: [2]int{2, 2},
		},
		wordCoords{
			character:   'S',
			coordinates: [2]int{3, 3},
		},
	},

	{
		wordCoords{
			character:   'M',
			coordinates: [2]int{-1, 1},
		},
		wordCoords{
			character:   'A',
			coordinates: [2]int{-2, 2},
		},
		wordCoords{
			character:   'S',
			coordinates: [2]int{-3, 3},
		},
	},

	{
		wordCoords{
			character:   'M',
			coordinates: [2]int{1, 0},
		},
		wordCoords{
			character:   'A',
			coordinates: [2]int{2, 0},
		},
		wordCoords{
			character:   'S',
			coordinates: [2]int{3, 0},
		},
	},

	{
		wordCoords{
			character:   'M',
			coordinates: [2]int{-1, 0},
		},
		wordCoords{
			character:   'A',
			coordinates: [2]int{-2, 0},
		},
		wordCoords{
			character:   'S',
			coordinates: [2]int{-3, 0},
		},
	},
}

func checkWordCoords(xpos, ypos int, grid []string, coordReg [3]wordCoords) bool {
	xcoord := -1
	ycoord := -1
	line := ""

	for _, element := range coordReg {
		xcoord = xpos + element.coordinates[0]
		ycoord = ypos + element.coordinates[1]
		// fmt.Printf("xpos: %d, xcoord: %d, char: %c\n", xpos, xcoord, element.character)
		// fmt.Printf("ypos: %d, ycoord: %d, gridlen: %d\n", ypos, ycoord, len(grid))
		if ycoord < 0 || ycoord > len(grid)-1 {
			return false
		}
		line = grid[ycoord]
		if xcoord < 0 || xcoord > len(line)-1 {
			return false
		}
		if line[xpos+element.coordinates[0]] != byte(element.character) {
			return false
		}
		// fmt.Printf("VALID xcoord: %d, ycoord: %d, char: %c\n", xcoord, ycoord, element.character)
	}
	return true
}

func findWordXMAS(xpos, ypos int, grid []string) int {
	foundWords := 0
	// wordCordList := [2][3]wordCoords{wordsCoords[6], wordsCoords[7]}

	for _, coordReg := range wordsCoordsXMAS {
		if foundWord := checkWordCoords(xpos, ypos, grid, coordReg); foundWord {
			foundWords++
		}
	}

	return foundWords
}

func WordSearchXMAS(grid []string) int {
	totalWordsFound := 0

	for ypos, line := range grid {
		for xpos := 0; xpos < len(line); xpos++ {
			if line[xpos] == 'X' {
				if foundWords := findWordXMAS(xpos, ypos, grid); foundWords > 0 {
					totalWordsFound += foundWords
				}
			}
		}
	}
	return totalWordsFound
}

func main() {
	lineReg := []string{}

	file, _ := os.Open("./data")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineReg = append(lineReg, line)
	}

  totalWords := WordSearchXMAS(lineReg)
  fmt.Println(totalWords)
}
