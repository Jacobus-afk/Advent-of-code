package main

import (
	"bufio"
	"fmt"
	"os"
)

type RobotDetails struct {
	position [2]int
	velocity [2]int
}

func buildGrid(dimensions [2]int) [][]int {
	grid := make([][]int, dimensions[1])
	for i := range grid {
		grid[i] = make([]int, dimensions[0])
	}
	fmt.Println("")
	return grid
}

func positionRobots(grid [][]int, robotsInitInfo []string) []RobotDetails {
	robots := []RobotDetails{}
	for _, line := range robotsInitInfo {
		var posx, posy, velx, vely int
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &posx, &posy, &velx, &vely)
		grid[posy][posx]++
		robot := RobotDetails{position: [2]int{posx, posy}, velocity: [2]int{velx, vely}}
		robots = append(robots, robot)
	}
	return robots
}

func handleVelocity(posx, posy, velx, vely, width, height int) (int, int) {
	newposx := posx + velx
	newposy := posy + vely
	if newposx < 0 {
		newposx += width
	}
	if newposx >= width {
		newposx -= width
	}
	if newposy < 0 {
		newposy += height
	}
	if newposy >= height {
		newposy -= height
	}

	return newposx, newposy
}

func moveRobots(seconds int, grid [][]int, robots []RobotDetails) {
	for range seconds {
		// for _, field := range grid {
		// 	fmt.Println(field)
		// }
		// fmt.Println("")
		for idx := range robots {
			robot := &robots[idx]
			posx := robot.position[0]
			posy := robot.position[1]
			velx := robot.velocity[0]
			vely := robot.velocity[1]
			grid[posy][posx]--

			newposx, newposy := handleVelocity(posx, posy, velx, vely, len(grid[0]), len(grid))
			robot.position = [2]int{newposx, newposy}
			grid[newposy][newposx]++
		}
	}
}

func splitGrid(grid [][]int) [4][][]int {
	verticalSplit := len(grid) / 2
	horizontalSplit := len(grid[0]) / 2
	// fmt.Println(verticalSplit, horizontalSplit)
	quadrants := [4][][]int{}
	topLeft := make([][]int, verticalSplit)
	topRight := make([][]int, verticalSplit)
	bottomLeft := make([][]int, verticalSplit)
	bottomRight := make([][]int, verticalSplit)

	for i := range verticalSplit {
		topLeft[i] = make([]int, horizontalSplit)
		topRight[i] = make([]int, horizontalSplit)
		bottomLeft[i] = make([]int, horizontalSplit)
		bottomRight[i] = make([]int, horizontalSplit)
		copy(topLeft[i], grid[i][:horizontalSplit])
		copy(topRight[i], grid[i][horizontalSplit+1:])
		copy(bottomLeft[i], grid[i+verticalSplit+1][:horizontalSplit])
		copy(bottomRight[i], grid[i+verticalSplit+1][horizontalSplit+1:])
	}
	// fmt.Println(topLeft)
	quadrants[0] = topLeft
	quadrants[1] = topRight
	quadrants[2] = bottomLeft
	quadrants[3] = bottomRight

	return quadrants
}

func CalculateSafetyFactor(seconds int, gridDimensions [2]int, robotsInitInfo []string) int {
	safetyFactor := 1
	grid := buildGrid(gridDimensions)
	robots := positionRobots(grid, robotsInitInfo)
	moveRobots(seconds, grid, robots)
	quadrants := splitGrid(grid)
	for _, quadrant := range quadrants {
		count := 0
		for _, field := range quadrant {
			for _, num := range field {
				count += num
			}
			// fmt.Println(field)
		}
		// fmt.Println("")
		safetyFactor *= count
	}

	return safetyFactor
}

func checkForTreeBottom(seconds int, grid [][]int) {
  bottomCount := 0
  for _, line := range grid {
    bottomCount = 0
    for _, entry := range line {
      if entry == 0 {
        bottomCount = 0
      } else {
        bottomCount++
      }
      if bottomCount > 10 {
        for _, line := range grid {
          for _, entry := range line {
            if entry == 0 {
              fmt.Print(".")
            } else {
              fmt.Print("#")
            }
          }
          fmt.Println("")
        }
        fmt.Println(seconds + 1)
        fmt.Println("")
        return
      }
    }
  }

}

func findPossibleTrees(seconds int, grid [][]int, robots []RobotDetails) {
  for sec := range seconds {
		for idx := range robots {
			robot := &robots[idx]
			posx := robot.position[0]
			posy := robot.position[1]
			velx := robot.velocity[0]
			vely := robot.velocity[1]
			grid[posy][posx]--

			newposx, newposy := handleVelocity(posx, posy, velx, vely, len(grid[0]), len(grid))
			robot.position = [2]int{newposx, newposy}
			grid[newposy][newposx]++
		}
    gridCopy := make([][]int, len(grid))
    copy(gridCopy, grid)
    checkForTreeBottom(sec, gridCopy)
	}

}

func FindChristmasTree(seconds int, gridDimensions [2]int, robotsInitInfo []string) {
	grid := buildGrid(gridDimensions)
	robots := positionRobots(grid, robotsInitInfo)
  findPossibleTrees(seconds, grid, robots)
}

func main() {
	robotsInitInfo := []string{}

	file, _ := os.Open("./data")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		robotsInitInfo = append(robotsInitInfo, line)
	}

  safetyFactor := CalculateSafetyFactor(100, [2]int{101,103}, robotsInitInfo)
  fmt.Println(safetyFactor)
  FindChristmasTree(1000000, [2]int{101,103}, robotsInitInfo)
}
