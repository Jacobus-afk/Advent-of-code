package main

import "testing"

func TestWordSearch(t *testing.T)  {
  t.Run("finds word horizontally", func(t *testing.T) {
    got := WordSearch([]string{"XMAS.SX"})

    want := 1

    if got != want {
      t.Errorf("got %d, want %d", got, want)
    }

  })

  t.Run("finds word horizontally reversed", func(t *testing.T) {
    got := WordSearch([]string{"XMAS.SX.SAMX."})

    want := 2

    if got != want {
      t.Errorf("got %d, want %d", got, want)
    }

  })

  t.Run("find vertical and diagonal words", func(t *testing.T) {
    got := WordSearch([]string{
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
