package main

import "testing"

func TestAntinodeCreation(t *testing.T) {
	// task 1 tests

	// t.Run("1 antenna creates no nodes", func(t *testing.T) {
	// 	grid := []string{
	// 		"..........",
	// 		"..........",
	// 		"..........",
	// 		"....A.....",
	// 		"..........",
	// 		"..........",
	// 		"..........",
	// 		"..........",
	// 		"..........",
	// 		"..........",
	// 	}
	//
	// 	got := AntinodeCreation(grid)
	//
	// 	want := 0
	// 	if got != want {
	// 		t.Errorf("got %d, want %d", got, want)
	// 	}
	// })
	//
	// t.Run("3 antennas create 6 antinodes", func(t *testing.T) {
	// 	grid := []string{
	// 		"..........",
	// 		"..........",
	// 		"..........",
	// 		"....a.a...",
	// 		"..........",
	// 		".....a....",
	// 		"..........",
	// 		"..........",
	// 		"..........",
	// 		"..........",
	// 	}
	//
	// 	got := AntinodeCreation(grid)
	//
	// 	want := 6
	// 	if got != want {
	// 		t.Errorf("got %d, want %d", got, want)
	// 	}
	// })
	//
	// t.Run("antinodes off grid", func(t *testing.T) {
	// 	grid := []string{
	// 		"..........",
	// 		"...#......",
	// 		"#.........",
	// 		"....a.....",
	// 		"........a.",
	// 		".....a....",
	// 		"..#.......",
	// 		"......A...",
	// 		"..........",
	// 		"..........",
	// 	}
	//
	// 	got := AntinodeCreation(grid)
	//
	// 	want := 4
	// 	if got != want {
	// 		t.Errorf("got %d, want %d", got, want)
	// 	}
	// })
	//
	// t.Run("multiple antenna types", func(t *testing.T) {
	// 	grid := []string{
	// 		"..0.........",
	// 		"........0...",
	// 		".....0......",
	// 		".......0f...",
	// 		"f...0.......",
	// 		"......A...f.",
	// 		"f...........",
	// 		"...........f",
	// 		"........A...",
	// 		".........A..",
	// 		"............",
	// 		"............",
	// 	}
	//
	// 	got := AntinodeCreation(grid)
	//
	// 	want := 18
	// 	if got != want {
	// 		t.Errorf("got %d, want %d", got, want)
	// 	}
	// })

	t.Run("multiple antenna types", func(t *testing.T) {
		grid := []string{
			"............",
			"........0...",
			".....0......",
			".......0....",
			"....0.......",
			"......A.....",
			"............",
			"............",
			"........A...",
			".........A..",
			"............",
			"............",
		}

		got := AntinodeCreation(grid)

		want := 18
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
