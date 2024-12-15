package main

import "fmt"

type Slope struct {
	height   int
	position [2]int
}

type MapList struct {
	startPosition [2]int
	branches      []Slope
}

var PossibleDirections = [4][2]int{
	{0, 1}, {1, 0}, {-1, 0}, {0, -1},
}

func (m *MapList) branchOff(pos [2]int, height int) {
	for branchIndex, branch := range m.branches {
		currentSlopeHeight := branch.height
		if currentSlopeHeight == 9 {
			continue
		}
		for _, dir := range PossibleDirections {
			if branch.position[0]+dir[0] != pos[0] || branch.position[1]+dir[1] != pos[1] {
				continue
			}
			if currentSlopeHeight + 1 != height {
				continue
			}
			// new branch
			if branch.height != currentSlopeHeight {
				slope := Slope{height: height, position: pos}
        fmt.Println("slope")
        fmt.Println(slope.height)
				m.branches = append(m.branches, slope)
        fmt.Println(m.branches)
			} else {
        fmt.Println("updating branch", height, pos)
				m.branches[branchIndex].height = height
				m.branches[branchIndex].position = pos
			}
		}
	}
  // fmt.Println(m.branches)
}

func buildMap(trailMap []string) [][]int {
  trail := [][]int{}

  for _, line := range trailMap {
    row := []int{}
    for _, char := range line {
      if char == '.' {
        row = append(row, -1)
      } else {
        row = append(row, int(char - '0'))
      }
    }
    trail = append(trail, row)
  }

  return trail
}

func findPath(trail [][]int) int {
	endMap := map[[2]int]map[[2]int]bool{}
	startMaps := []MapList{}
	tally := 0
	for ypos, line := range trail {
		for xpos, height := range line {
			pos := [2]int{xpos, ypos}
			if height == 0 {
        slope := Slope{height: height, position: pos}
				mapList := MapList{branches: []Slope{slope}, startPosition: pos}
				startMaps = append(startMaps, mapList)
				continue
			}
			for _, startMap := range startMaps {
				startMap.branchOff(pos, height)
				if height == 9 {
					if _, ok := endMap[pos]; !ok {
						endMap[pos] = make(map[[2]int]bool)
					}
					if _, ok := endMap[pos][startMap.startPosition]; !ok {
						tally++
					} else {
						endMap[pos][startMap.startPosition] = true
					}
				}
			}
		}
	}
	// for ypos := len(trail) - 1; ypos >= 0; ypos-- {
	// 	for xpos, height := range trail[ypos] {
	// 		pos := [2]int{xpos, ypos}
	// 		if height == 0 {
	// 			continue
	// 		}
	// 		for _, startMap := range startMaps {
	// 			startMap.branchOff(pos, height)
	// 			if height == 9 {
	// 				if _, ok := endMap[pos]; !ok {
	// 					endMap[pos] = make(map[[2]int]bool)
	// 				}
	// 				if _, ok := endMap[pos][startMap.startPosition]; !ok {
	// 					tally++
	// 				} else {
	// 					endMap[pos][startMap.startPosition] = true
	// 				}
	// 			}
	// 		}
	// 	}
	// }
  //
	// for _, sm := range startMaps {
	// 	fmt.Println(sm.startPosition)
	// 	for _, branch := range sm.branches {
	// 		fmt.Println(branch)
	// 	}
	// }
	// for _, em := range endMap {
	// 	fmt.Println(em)
	// }
  fmt.Println(startMaps)
	return tally
}
