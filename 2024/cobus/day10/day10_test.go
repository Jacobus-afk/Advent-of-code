package main

import "testing"

func TestPathFinder(t *testing.T) {
	t.Run("finds path 0 to 9", func(t *testing.T) {
		trailString := []string{
			"0123",
			"1234",
			"8765",
			"9876",
		}
    trail := buildMap(trailString)
		got := findPath(trail)

		want := 1

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	// t.Run("find multiple paths", func(t *testing.T) {
	// 	trail := [][]int{
	// 		{8, 9, 0, 1, 0, 1, 2, 3},
	// 		{7, 8, 1, 2, 1, 8, 7, 4},
	// 		{8, 7, 4, 3, 0, 9, 6, 5},
	// 		{9, 6, 5, 4, 9, 8, 7, 4},
	// 		{4, 5, 6, 7, 8, 9, 0, 3},
	// 		{3, 2, 0, 1, 9, 0, 1, 2},
	// 		{0, 1, 3, 2, 9, 8, 0, 1},
	// 		{1, 0, 4, 5, 6, 7, 3, 2},
	// 	}
	//
	// 	got := findPath(trail)
	//
	// 	want := 36
	//
	// 	if got != want {
	// 		t.Errorf("got %d, want %d", got, want)
	// 	}
	// })
}
