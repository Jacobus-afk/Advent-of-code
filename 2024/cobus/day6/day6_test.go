package main

import (
	// "fmt"
	"testing"
)

func TestGuardPositions(t *testing.T) {
	t.Run("horizontal movement counts the right value", func(t *testing.T) {
		grid := []string{">........#"}
		startPosition, direction := getStartingPosition(grid)
		// fmt.Printf("xpos: %d, direction: %d\n", xpos, direction)
		// startPosition := [2]int{xpos, 0}

		got, _ := TrackMovement(grid, startPosition, direction)

		want := 9

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("horizontal and vertical movement", func(t *testing.T) {
		grid := []string{
			"....#.....",
			".........#",
			"..........",
			"..#.......",
			".......#..",
			"..........",
			".#..^.....",
			"........#.",
			"#.........",
			"......#...",
		}

	   startPosition, direction := getStartingPosition(grid)
		// fmt.Printf("pos: %d, direction: %d\n", startPosition, direction)

		got, _ := TrackMovement(grid, startPosition, direction)

		want := 41

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	// t.Run("potential obstruction placement", func(t *testing.T) {
	// 	grid := []string{
	// 		"....#.....",
	// 		".........#",
	// 		"..........",
	// 		"..#.......",
	// 		".......#..",
	// 		"..........",
	// 		".#..^.....",
	// 		"........#.",
	// 		"#.........",
	// 		"......#...",
	// 	}
	//
	//    startPosition, direction := getStartingPosition(grid)
	// 	// fmt.Printf("pos: %d, direction: %d\n", startPosition, direction)
	//
	// 	_, got := TrackMovement(grid, startPosition, direction)
	//
	// 	want := 41
	//
	// 	if got != want {
	// 		t.Errorf("got %d, want %d", got, want)
	// 	}
	// })
}

func TestFindLoops(t *testing.T) {
	grid := []string{
		"....#.....",
		".........#",
		"..........",
		"..#.......",
		".......#..",
		"..........",
		".#..^.....",
		"........#.",
		"#.........",
		"......#...",
	}

  startPosition, direction := getStartingPosition(grid)

  got := FindLoops(grid, startPosition, direction)

  want := 8

  if got != want {
    t.Errorf("got %d, want %d", got, want)
  }
}
