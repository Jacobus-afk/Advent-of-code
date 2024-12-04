package main

import "testing"

func TestWordSearchXMAS(t *testing.T) {
	t.Run("finds word horizontally", func(t *testing.T) {
		got := WordSearchXMAS([]string{"XMAS.SX"})

		want := 1

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("finds word horizontally reversed", func(t *testing.T) {
		got := WordSearchXMAS([]string{"XMAS.SX.SAMX."})

		want := 2

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("find vertical and diagonal words", func(t *testing.T) {
		got := WordSearchXMAS([]string{
			"MMMSXXMASM",
			"MSAMXMSMSA",
			"AMXSXMAAMM",
			"MSAMASMSMX",
			"XMASAMXAMM",
			"XXAMMXXAMA",
			"SMSMSASXSS",
			"SAXAMASAAA",
			"MAMMMXMMMM",
			"MXMXAXMASX",
		})

		want := 18

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestWordSearchMAS(t *testing.T) {
	t.Run("finds words all directions", func(t *testing.T) {
    got := WordSearchMAS([]string{
			"MMMSXXMASM",
			"MSAMXMSMSA",
			"AMXSXMAAMM",
			"MSAMASMSMX",
			"XMASAMXAMM",
			"XXAMMXXAMA",
			"SMSMSASXSS",
			"SAXAMASAAA",
			"MAMMMXMMMM",
			"MXMXAXMASX",
    })

    want := 9 
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

}
