package main

import (
	// "fmt"
	"testing"
)

// func TestInitialization(t *testing.T) {
// 	t.Run("creates warehouse and robot movements", func(t *testing.T) {
// 		data := []string{
// 			"########",
// 			"#..O.O.#",
// 			"##@.O..#",
// 			"#...O..#",
// 			"#.#.O..#",
// 			"#...O..#",
// 			"#......#",
// 			"########",
// 			"",
// 			"<^^>>>vv<v>>v<<",
// 		}
//
// 		warehouse, movements, robot, warehouseDimensions := initializeFromData(data)
// 		fmt.Println(movements, robot)
// 		fmt.Println("")
// 		for posy := range warehouseDimensions[1] {
// 			for posx := range warehouseDimensions[0] {
// 				fmt.Print(warehouse[[2]int{posx, posy}])
// 			}
// 			fmt.Println("")
// 		}
// 	})
//
// 	t.Run("creates warehouse and robot movements", func(t *testing.T) {
// 		data := []string{
// 			"##########",
// 			"#..O..O.O#",
// 			"#......O.#",
// 			"#.OO..O.O#",
// 			"#..O@..O.#",
// 			"#O#..O...#",
// 			"#O..O..O.#",
// 			"#.OO.O.OO#",
// 			"#....O...#",
// 			"##########",
// 			"",
// 			"<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^",
// 			"vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v",
// 			"><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<",
// 			"<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^",
// 			"^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><",
// 			"^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^",
// 			">^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^",
// 			"<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>",
// 			"^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>",
// 			"v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^",
// 		}
//
// 		warehouse, movements, robot, warehouseDimensions := initializeFromData(data)
// 		fmt.Println(movements, robot)
// 		fmt.Println("")
// 		for posy := range warehouseDimensions[1] {
// 			for posx := range warehouseDimensions[0] {
// 				fmt.Print(warehouse[[2]int{posx, posy}])
// 			}
// 			fmt.Println("")
// 		}
// 	})
//
// 	t.Run("creates alt warehouse and robot movements", func(t *testing.T) {
// 		data := []string{
// 			"##########",
// 			"#..O..O.O#",
// 			"#......O.#",
// 			"#.OO..O.O#",
// 			"#..O@..O.#",
// 			"#O#..O...#",
// 			"#O..O..O.#",
// 			"#.OO.O.OO#",
// 			"#....O...#",
// 			"##########",
// 			"",
// 			"<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^",
// 			"vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v",
// 			"><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<",
// 			"<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^",
// 			"^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><",
// 			"^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^",
// 			">^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^",
// 			"<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>",
// 			"^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>",
// 			"v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^",
// 		}
//
// 		warehouse, movements, robot, warehouseDimensions := initializeAltWarehouse(data)
// 		fmt.Println(movements, robot)
// 		fmt.Println("")
// 		for posy := range warehouseDimensions[1] {
// 			for posx := range warehouseDimensions[0] {
// 				fmt.Print(warehouse[[2]int{posx, posy}])
// 			}
// 			fmt.Println("")
// 		}
// 	})
// }

func TestRobotMovement(t *testing.T) {
	// t.Run("small warehouse robot movement", func(t *testing.T) {
	// 	data := []string{
	// 		"########",
	// 		"#..O.O.#",
	// 		"##@.O..#",
	// 		"#...O..#",
	// 		"#.#.O..#",
	// 		"#...O..#",
	// 		"#......#",
	// 		"########",
	// 		"",
	// 		"<^^>>>vv<v>>v<<",
	// 	}
	//
	// 	warehouse, movements, robot, warehouseDimensions := initializeFromData(data)
	//
	// 	robot = moveRobot(robot, movements, warehouse)
	//
	// 	fmt.Println(movements, robot)
	// 	for posy := range warehouseDimensions[1] {
	// 		for posx := range warehouseDimensions[0] {
	// 			fmt.Print(warehouse[[2]int{posx, posy}])
	// 		}
	// 		fmt.Println("")
	// 	}
	// })
	//
	// t.Run("large warehouse robot movement", func(t *testing.T) {
	// 	data := []string{
	// 		"##########",
	// 		"#..O..O.O#",
	// 		"#......O.#",
	// 		"#.OO..O.O#",
	// 		"#..O@..O.#",
	// 		"#O#..O...#",
	// 		"#O..O..O.#",
	// 		"#.OO.O.OO#",
	// 		"#....O...#",
	// 		"##########",
	// 		"",
	// 		"<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^",
	// 		"vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v",
	// 		"><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<",
	// 		"<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^",
	// 		"^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><",
	// 		"^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^",
	// 		">^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^",
	// 		"<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>",
	// 		"^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>",
	// 		"v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^",
	// 	}
	//
	// 	warehouse, movements, robot, warehouseDimensions := initializeFromData(data)
	//
	// 	robot = moveRobot(robot, movements, warehouse)
	//
	// 	fmt.Println(movements, robot)
	// 	for posy := range warehouseDimensions[1] {
	// 		for posx := range warehouseDimensions[0] {
	// 			fmt.Print(warehouse[[2]int{posx, posy}])
	// 		}
	// 		fmt.Println("")
	// 	}
	// })
	//
	// t.Run("large alt warehouse robot movement", func(t *testing.T) {
	// 	data := []string{
	// 		"##########",
	// 		"#..O..O.O#",
	// 		"#......O.#",
	// 		"#.OO..O.O#",
	// 		"#..O@..O.#",
	// 		"#O#..O...#",
	// 		"#O..O..O.#",
	// 		"#.OO.O.OO#",
	// 		"#....O...#",
	// 		"##########",
	// 		"",
	// 		"<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^",
	// 		"vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v",
	// 		"><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<",
	// 		"<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^",
	// 		"^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><",
	// 		"^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^",
	// 		">^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^",
	// 		"<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>",
	// 		"^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>",
	// 		"v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^",
	// 	}
	//
	// 	warehouse, movements, robot, warehouseDimensions := initializeAltWarehouse(data)
	//
	// 	robot = moveRobot(robot, movements, warehouse)
	//
	// 	fmt.Println(movements, robot)
	// 	for posy := range warehouseDimensions[1] {
	// 		for posx := range warehouseDimensions[0] {
	// 			fmt.Print(warehouse[[2]int{posx, posy}])
	// 		}
	// 		fmt.Println("")
	// 	}
	// })
	//
	// t.Run("calcsGPS coords", func(t *testing.T) {
	// 	data := []string{
	// 		"##########",
	// 		"#..O..O.O#",
	// 		"#......O.#",
	// 		"#.OO..O.O#",
	// 		"#..O@..O.#",
	// 		"#O#..O...#",
	// 		"#O..O..O.#",
	// 		"#.OO.O.OO#",
	// 		"#....O...#",
	// 		"##########",
	// 		"",
	// 		"<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^",
	// 		"vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v",
	// 		"><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<",
	// 		"<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^",
	// 		"^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><",
	// 		"^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^",
	// 		">^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^",
	// 		"<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>",
	// 		"^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>",
	// 		"v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^",
	// 	}
	//
	// 	get, _ := CalcGPSCoordinates(data)
	//
	// 	want := 10092
	//
	// 	if get != want {
	// 		t.Errorf("get %d, want %d", get, want)
	// 	}
	// })

	t.Run("large alt warehouse gps calcs", func(t *testing.T) {
		data := []string{
			"##########",
			"#..O..O.O#",
			"#......O.#",
			"#.OO..O.O#",
			"#..O@..O.#",
			"#O#..O...#",
			"#O..O..O.#",
			"#.OO.O.OO#",
			"#....O...#",
			"##########",
			"",
			"<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^",
			"vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v",
			"><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<",
			"<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^",
			"^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><",
			"^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^",
			">^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^",
			"<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>",
			"^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>",
			"v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^",
		}

		_, get := CalcGPSCoordinates(data)

		want := 9021

		if get != want {
			t.Errorf("get %d, want %d", get, want)
		}
	})

	t.Run("small alt warehouse gps calcs", func(t *testing.T) {
		data := []string{
			"#######",
			"#...#.#",
			"#.....#",
			"#..OO@#",
			"#..O..#",
			"#.....#",
			"#######",
			"",
			"<vv<<^^<<^^",
		}

		_, get := CalcGPSCoordinates(data)

		want := 618

		if get != want {
			t.Errorf("get %d, want %d", get, want)
		}
	})
}
