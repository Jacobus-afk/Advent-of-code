package main

import (
  "fmt"

  aoc "aoc-24/lib"
)

func initializeFromData(data []string) (map[[2]int]string, string, [2]int, [2]int) {
	handleMovements := false
	warehouse := map[[2]int]string{}
	robot := [2]int{-1, -1}
	warehouseDimensions := [2]int{0, 0}
	movements := ""
	for posy, line := range data {
		if handleMovements {
			movements += line
			continue
		}
		if line == "" {
			handleMovements = true
			warehouseDimensions = [2]int{len(data[posy-1]), posy}
			continue
		}
		for posx, char := range line {
			if char == '@' {
				robot = [2]int{posx, posy}
			}
			warehouse[[2]int{posx, posy}] = string(char)
		}
	}
	return warehouse, movements, robot, warehouseDimensions
}

func getMovementVector(movement rune) [2]int {
	switch movement {
	case '<':
		return [2]int{-1, 0}
	case '>':
		return [2]int{1, 0}
	case '^':
		return [2]int{0, -1}
	case 'v':
		return [2]int{0, 1}
	}
	return [2]int{-99, -99}
}

func handleMovement(robot, movementVector [2]int, warehouse map[[2]int]string) [2]int {
	posx := robot[0] + movementVector[0]
	posy := robot[1] + movementVector[1]
	potentialMove := ""

	for loop := true; loop; {
		potentialMove = warehouse[[2]int{posx, posy}]
		// fmt.Println("potential move", posx, posy, potentialMove)
		if posx < 0 || posy < 0 {
			break
		}
		switch potentialMove {
		case "#":
			loop = false
		case "O":
			posx += movementVector[0]
			posy += movementVector[1]
		case ".":
			loop = false
		}
	}

	if potentialMove != "." {
		return robot
	}
	for potentialMove != "@" {
		// fmt.Println("tracking back", potentialMove, posx, posy)
		newposx := posx - movementVector[0]
		newposy := posy - movementVector[1]
		movedBlock := warehouse[[2]int{newposx, newposy}]
		warehouse[[2]int{posx, posy}] = movedBlock

		posx = newposx
		posy = newposy
		potentialMove = movedBlock
	}
	warehouse[[2]int{posx, posy}] = "."
	robot = [2]int{posx + movementVector[0], posy + movementVector[1]}
	return robot
}

func moveRobot(robot [2]int, movements string, warehouse map[[2]int]string) [2]int {
	for _, movement := range movements {
		movementVector := getMovementVector(movement)
		// fmt.Println("movement", string(movement), movementVector)
		robot = handleMovement(robot, movementVector, warehouse)
		// fmt.Println(robot)
	}

	fmt.Println("")
	return robot
}

func CalcGPSCoordinates(data []string) int {
  sumGPS := 0

  warehouse, movements, robot, _ := initializeFromData(data)
  robot = moveRobot(robot, movements, warehouse)

  for k, v := range warehouse {
    if v == "O" {
      gps := k[1] * 100 + k[0]
      sumGPS += gps
    }
  }

  return sumGPS
}

func main() {
  data := aoc.ReadFileByLine("./data")

  sumGPS := CalcGPSCoordinates(data)
  fmt.Println(sumGPS)
}
