package main

import (
	// "fmt"
	"testing"
)

func TestGetTotalPrice(t *testing.T) {
	t.Run("calcs price, simple ex", func(t *testing.T) {
		garden := []string{
			"AAAA",
			"BBCD",
			"BBCC",
			"EEEC",
		}

		got, _ := PlotGarden(garden)
		// fmt.Println(plantMap)
		want := 140

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("calcs price", func(t *testing.T) {
		garden := []string{
			"OOOOO",
			"OXOXO",
			"OOOOO",
			"OXOXO",
			"OOOOO",
		}

		got, _ := PlotGarden(garden)
		// fmt.Println(plantMap)
		want := 772

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("calcs price", func(t *testing.T) {
		garden := []string{
			"RRRRIICCFF",
			"RRRRIICCCF",
			"VVRRRCCFFF",
			"VVRCCCJFFF",
			"VVVVCJJCFE",
			"VVIVCCJJEE",
			"VVIIICJJEE",
			"MIIIIIJJEE",
			"MIIISIJEEE",
			"MMMISSJEEE",
		}

		got, _ := PlotGarden(garden)
		// fmt.Println(plantMap)
		want := 1930

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
