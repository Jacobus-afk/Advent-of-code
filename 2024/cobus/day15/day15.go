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

func initializeAltWarehouse(data []string) (map[[2]int]string, string, [2]int, [2]int) {
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
			warehouseDimensions = [2]int{len(data[posy-1]) * 2, posy}
			continue
		}
		for idx, char := range line {
			posx := idx * 2
			switch char {
			case '@':
				robot = [2]int{posx, posy}
				warehouse[[2]int{posx, posy}] = "@"
				warehouse[[2]int{posx + 1, posy}] = "."
			case 'O':
				warehouse[[2]int{posx, posy}] = "["
				warehouse[[2]int{posx + 1, posy}] = "]"
			case '.':
				warehouse[[2]int{posx, posy}] = "."
				warehouse[[2]int{posx + 1, posy}] = "."
			default:
				warehouse[[2]int{posx, posy}] = "#"
				warehouse[[2]int{posx + 1, posy}] = "#"
			}
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

func handleVerticalMovement(
	robot, movementVector [2]int,
	warehouse map[[2]int]string,
	warehouseDimensions [2]int,
) [2]int {
	// fmt.Println("")
	// for posy := range warehouseDimensions[1] {
	// 	for posx := range warehouseDimensions[0] {
	// 		fmt.Print(warehouse[[2]int{posx, posy}])
	// 	}
	// 	fmt.Println("")
	// }
	posx := robot[0]
	posy := robot[1]
	// potentialMove := warehouse[[2]int{posx, posy}]
	// knockOnMoves := map[[2]int]string{{posx, posy}: potentialMove}
	knockOnMoves := map[[2]int]string{robot: "@"}
	movesList := []map[[2]int]string{}

	for loop := true; loop; {
		tmpMap := map[[2]int]string{}
		// posx += movementVector[0]
		// posy += movementVector[1]
		// have to have loop false here so that we can break out if it's all '.'s
		loop = false
		for k, v := range knockOnMoves {
			// fmt.Println("trying..", k, v)
			posx = k[0] + movementVector[0]
			posy = k[1] + movementVector[1]
			// if warehouse[[2]int{posx, posy}] == "#" {
			//   fmt.Println("got #", posx, posy, knockOnMoves)
			//   return robot
			// }
			switch v {
			case "@":
				// posx = movementVector[0]
				// posy = movementVector[1]
				tmpMap[[2]int{posx, posy}] = warehouse[[2]int{posx, posy}]
				// fmt.Println("got @ updated", tmpMap)
				loop = true
			case "[":
				// posx += movementVector[0]
				// posy += movementVector[1]
				tmpMap[[2]int{posx, posy}] = warehouse[[2]int{posx, posy}]
				tmpMap[[2]int{posx + 1, posy}] = warehouse[[2]int{posx + 1, posy}]
				// fmt.Println("got [ updated", tmpMap)
				loop = true
			case "]":
				// posx += movementVector[0]
				// posy += movementVector[1]
				tmpMap[[2]int{posx, posy}] = warehouse[[2]int{posx, posy}]
				tmpMap[[2]int{posx - 1, posy}] = warehouse[[2]int{posx - 1, posy}]
				// fmt.Println("got ] updated", tmpMap)
				loop = true
			case "#":
				// fmt.Println("got #")
				return robot
			}
		}
		movesList = append(movesList, tmpMap)
		knockOnMoves = tmpMap
	}

	for i := len(movesList) - 1; i >= 0; i-- {
		moves := movesList[i]
		// fmt.Println("rewinding..", moves)
		for k := range moves {
			posx = k[0] - movementVector[0]
			posy = k[1] - movementVector[1]
			replacement := warehouse[[2]int{posx, posy}]
			// fmt.Printf(
			// 	"replacing %s%d with %s[%d,%d]\n",
			// 	warehouse[k],
			// 	k,
			// 	replacement,
			// 	posx,
			// 	posy,
			// )
			warehouse[k] = replacement
			warehouse[[2]int{posx, posy}] = "."
		}
		// fmt.Println("")
	}

	posx = robot[0] + movementVector[0]
	posy = robot[1] + movementVector[1]
	warehouse[robot] = "."
	robot = [2]int{posx, posy}
	// warehouse[robot] = "@"

	// for loop := true; loop; {
	// 	potentialMove = warehouse[[2]int{posx, posy}]
	// 	if posx < 0 || posy < 0 {
	// 		break
	// 	}
	// 	switch potentialMove {
	// 	case "[":
	// 		posx += movementVector[0]
	// 		posy += movementVector[1]
	// 		posxAlt = posx + 1
	// 	case "]":
	// 		posx += movementVector[0]
	// 		posy += movementVector[1]
	// 		posxAlt = posx - 1
	// 	}
	// }

	return robot
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
		// case "O":
		// 	posx += movementVector[0]
		// 	posy += movementVector[1]
		case ".":
			loop = false
		default:
			posx += movementVector[0]
			posy += movementVector[1]
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

func moveAltRobot(
	robot [2]int,
	movements string,
	warehouse map[[2]int]string,
	warehouseDimensions [2]int,
) [2]int {
	for _, movement := range movements {
		movementVector := getMovementVector(movement)
		// fmt.Println("movement", string(movement), movementVector)

		// for posy := range warehouseDimensions[1] {
		// 	for posx := range warehouseDimensions[0] {
		// 		fmt.Print(warehouse[[2]int{posx, posy}])
		// 	}
		// 	fmt.Println("")
		// }

		if movementVector[1] == 0 {
			robot = handleMovement(robot, movementVector, warehouse)
		} else {
			robot = handleVerticalMovement(robot, movementVector, warehouse, warehouseDimensions)
		}

		// fmt.Println(robot)
	}

	fmt.Println("")
	return robot
}

func CalcGPSCoordinates(data []string) (int, int) {
	sumGPS := 0
	sumAltGPS := 0

	warehouse, movements, robot, _ := initializeFromData(data)
	moveRobot(robot, movements, warehouse)

	for k, v := range warehouse {
		if v == "O" {
			gps := k[1]*100 + k[0]
			sumGPS += gps
		}
	}

	warehouseAlt, movementsAlt, robotAlt, warehouseDimensions := initializeAltWarehouse(data)
	moveAltRobot(robotAlt, movementsAlt, warehouseAlt, warehouseDimensions)

	for posy := range warehouseDimensions[1] {
		for posx := range warehouseDimensions[0] {
			fmt.Print(warehouseAlt[[2]int{posx, posy}])
		}
		fmt.Println("")
	}

	for k, v := range warehouseAlt {
		if v == "[" {
			gps := k[1]*100 + k[0]
			sumAltGPS += gps
		}
	}

	return sumGPS, sumAltGPS
}

func main() {
	data := aoc.ReadFileByLine("./data")

	sumGPS, sumAltGPS := CalcGPSCoordinates(data)
	fmt.Println(sumGPS, sumAltGPS)
}
