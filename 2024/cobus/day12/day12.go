package main

import (
	"bufio"
	"fmt"
	"os"
)

type Plot struct {
	label     rune
	area      int
	perimeter int
	positions [][2]int
	posHash   map[[2]int]bool
}

var PossibleDirections = [4][2]int{
	{0, 1}, {1, 0}, {-1, 0}, {0, -1},
}

func PlotGarden(garden []string) int {
	GardenMap := []Plot{}
	plantMap := map[rune]map[[2]int]bool{}
	for posy, line := range garden {
		for posx, plant := range line {
			plantPositions, ok := plantMap[plant]
			if !ok {
				plantPositions = make(map[[2]int]bool)
			}
			plantPositions[[2]int{posx, posy}] = true
			plantMap[plant] = plantPositions
		}
	}

	for plant, positions := range plantMap {
		for len(positions) > 0 {

			// initialize plot with a position
			plot := Plot{
				label:     plant,
				area:      0,
				perimeter: 0,
				positions: [][2]int{},
				posHash:   make(map[[2]int]bool),
			}
			for pos := range positions {
				plot.positions = append(plot.positions, pos)
				plot.posHash[pos] = true
				plot.area++
				delete(positions, pos)
				break
			}

			// for _, pos := range plot.positions {
			for i := 0; i < len(plot.positions); i++ {
				pos := plot.positions[i]
				// perimeterTally := 0
				for _, dir := range PossibleDirections {
					posx := pos[0] + dir[0]
					posy := pos[1] + dir[1]
					// if posx < 0 || posy < 0 || posx > len(garden[0]) - 1 || posy > len(garden) - 1 {
					//   continue
					// }
					newPos := [2]int{posx, posy}

					if _, ok := positions[newPos]; ok {
						plot.area++
						plot.positions = append(plot.positions, newPos)
						plot.posHash[newPos] = true
						delete(positions, newPos)
						// making sure it's not a boundary of already added plant
					} else if _, ok := plot.posHash[newPos]; !ok {
						// fmt.Println(plant, newPos)
						// fmt.Println(plot.posHash)
						plot.perimeter++
					}
				}
			}
			GardenMap = append(GardenMap, plot)

		}
	}
	price := 0

	for _, plot := range GardenMap {
		price += plot.area * plot.perimeter
	}
	// fmt.Println(GardenMap)
	fmt.Println("")
	return price
}

func main() {
	garden := []string{}

	file, _ := os.Open("./data")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		garden = append(garden, line)
	}

  price := PlotGarden(garden)
  fmt.Println(price)
}
