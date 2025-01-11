package main

import (
	"fmt"
	"testing"
)

func TestMazeCreation(t *testing.T) {
	t.Run("maze created", func(t *testing.T) {
		data := []string{
			"###############",
			"#.......#....E#",
			"#.#.###.#.###.#",
			"#.....#.#...#.#",
			"#.###.#####.#.#",
			"#.#.#.......#.#",
			"#.#.#####.###.#",
			"#...........#.#",
			"###.#.#####.#.#",
			"#...#.....#.#.#",
			"#.#.#.###.#.#.#",
			"#.....#...#.#.#",
			"#.###.#.#.#.#.#",
			"#S..#.....#...#",
			"###############",
		}

		maze, _, dimensions := createMaze(data)

		for posy := range dimensions[1] {
			for posx := range dimensions[0] {
				fmt.Print(maze[[2]int{posx, posy}])
			}
			fmt.Println("")
		}
	})

	t.Run("find paths", func(t *testing.T) {
		data := []string{
			"###############",
			"#.......#....E#",
			"#.#.###.#.###.#",
			"#.....#.#...#.#",
			"#.###.#####.#.#",
			"#.#.#.......#.#",
			"#.#.#####.###.#",
			"#...........#.#",
			"###.#.#####.#.#",
			"#...#.....#.#.#",
			"#.#.#.###.#.#.#",
			"#.....#...#.#.#",
			"#.###.#.#.#.#.#",
			"#S..#.....#...#",
			"###############",
		}

		maze, start, dimensions := createMaze(data)

		TraverseMaze(maze, start, dimensions)

		// for posy := range dimensions[1] {
		//   for posx := range dimensions[0] {
		//     fmt.Print(maze[[2]int{posx, posy}])
		//   }
		//   fmt.Println("")
		// }
	})

	t.Run("find paths", func(t *testing.T) {
		data := []string{
			"#################",
			"#...#...#...#..E#",
			"#.#.#.#.#.#.#.#.#",
			"#.#.#.#...#...#.#",
			"#.#.#.#.###.#.#.#",
			"#...#.#.#.....#.#",
			"#.#.#.#.#.#####.#",
			"#.#...#.#.#.....#",
			"#.#.#####.#.###.#",
			"#.#.#.......#...#",
			"#.#.###.#####.###",
			"#.#.#...#.....#.#",
			"#.#.#.#####.###.#",
			"#.#.#.........#.#",
			"#.#.#.#########.#",
			"#S#.............#",
			"#################",
		}

		maze, start, dimensions := createMaze(data)

    got := TraverseMaze(maze, start, dimensions)
    want := 11048

    if got != want {
      t.Errorf("got %d, want %d", got, want)
    }

		// for posy := range dimensions[1] {
		//   for posx := range dimensions[0] {
		//     fmt.Print(maze[[2]int{posx, posy}])
		//   }
		//   fmt.Println("")
		// }
	})
}
