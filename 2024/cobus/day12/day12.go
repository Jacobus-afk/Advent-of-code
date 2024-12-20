package main

import (
	"bufio"
	"fmt"
	"os"
)

type Plot struct {
	label       rune
	area        int
	perimeter   int
	sides       int
	positions   [][2]int
	posHash     map[[2]int]bool
	leftSides   map[[2]int]bool
	rightSides  map[[2]int]bool
	topSides    map[[2]int]bool
	bottomSides map[[2]int]bool
}

var PossibleDirections = [4][2]int{
	{0, 1}, {1, 0}, {-1, 0}, {0, -1},
}

func PlotGarden(garden []string) (int, int) {
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
				label:       plant,
				area:        0,
				perimeter:   0,
				sides:       0,
				positions:   [][2]int{},
				posHash:     make(map[[2]int]bool),
				leftSides:   make(map[[2]int]bool),
				rightSides:  make(map[[2]int]bool),
				topSides:    make(map[[2]int]bool),
				bottomSides: make(map[[2]int]bool),
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
            if dir[0] == 1 {
              plot.rightSides[newPos] = true
            }
            if dir[0] == -1 {
              plot.leftSides[newPos] = true
            }
            if dir[1] == 1 {
              plot.bottomSides[newPos] = true
            }
            if dir[1] == -1 {
              plot.topSides[newPos] = true
            }
						// if dir[0] == 0 {
						// 	plot.leftSides[newPos] = true
						// } else {
						// 	plot.topSides[newPos] = true
						// }

					}
				}
			}
			GardenMap = append(GardenMap, plot)
		}
	}

	for idx := range GardenMap {
		plot := &GardenMap[idx]
		// fmt.Println(plot.label, "left", plot.leftSides)
		// fmt.Println(plot.label, "right", plot.rightSides)
		// fmt.Println(plot.label, "top", plot.topSides)
		// fmt.Println(plot.label, "bottom", plot.bottomSides)

		for len(plot.topSides) > 0 {
			topSideList := [][2]int{}
			posy := -99

			for side := range plot.topSides {
				topSideList = append(topSideList, side)
				posy = side[1]
				plot.sides++
				delete(plot.topSides, side)
				break
			}

			for i := 0; i < len(topSideList); i++ {
        // fmt.Println("top", topSideList)
				pos := topSideList[i]
				for _, dir := range [2]int{-1, 1} {
					posx := pos[0] + dir
					newPos := [2]int{posx, posy}
          // fmt.Println("newPos", newPos)

					if _, ok := plot.topSides[newPos]; ok {
						topSideList = append(topSideList, newPos)
						delete(plot.topSides, newPos)
					}
				}
			}
		}

		for len(plot.bottomSides) > 0 {
			bottomSideList := [][2]int{}
			posy := -99

			for side := range plot.bottomSides {
				bottomSideList = append(bottomSideList, side)
				posy = side[1]
				plot.sides++
				delete(plot.bottomSides, side)
				break
			}

			for i := 0; i < len(bottomSideList); i++ {
				pos := bottomSideList[i]
				for _, dir := range [2]int{-1, 1} {
					posx := pos[0] + dir
					newPos := [2]int{posx, posy}

					if _, ok := plot.bottomSides[newPos]; ok {
						bottomSideList = append(bottomSideList, newPos)
						delete(plot.bottomSides, newPos)
					}
				}
			}
		}

		for len(plot.leftSides) > 0 {
			leftSideList := [][2]int{}
			posx := -99

			for side := range plot.leftSides {
				leftSideList = append(leftSideList, side)
				posx = side[0]
				plot.sides++
				// fmt.Println("hor sides:", plot.sides)
				delete(plot.leftSides, side)
				break
			}

			for i := 0; i < len(leftSideList); i++ {
				// fmt.Println(horizontalSideList)
				pos := leftSideList[i]
				for _, dir := range [2]int{-1, 1} {
					posy := pos[1] + dir
					newPos := [2]int{posx, posy}

					if _, ok := plot.leftSides[newPos]; ok {
						leftSideList = append(leftSideList, newPos)
						delete(plot.leftSides, newPos)
					}
				}

			}
		}

		for len(plot.rightSides) > 0 {
			rightSideList := [][2]int{}
			posx := -99

			for side := range plot.rightSides {
				rightSideList = append(rightSideList, side)
				posx = side[0]
				plot.sides++
				// fmt.Println("hor sides:", plot.sides)
				delete(plot.rightSides, side)
				break
			}

			for i := 0; i < len(rightSideList); i++ {
				// fmt.Println(horizontalSideList)
				pos := rightSideList[i]
				for _, dir := range [2]int{-1, 1} {
					posy := pos[1] + dir
					newPos := [2]int{posx, posy}

					if _, ok := plot.rightSides[newPos]; ok {
						rightSideList = append(rightSideList, newPos)
						delete(plot.rightSides, newPos)
					}
				}

			}
		}

	}

	regionPrice := 0
	sidePrice := 0

	for _, plot := range GardenMap {
		// fmt.Println(plot.label, plot.sides)
		regionPrice += plot.area * plot.perimeter
		sidePrice += plot.area * plot.sides
	}
	// fmt.Println(GardenMap[0].label, GardenMap[0].sides)
	// fmt.Println(sidePrice)
	return regionPrice, sidePrice
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

	regionPrice, sidePrice := PlotGarden(garden)
	fmt.Println(regionPrice)
	fmt.Println(sidePrice)
}
