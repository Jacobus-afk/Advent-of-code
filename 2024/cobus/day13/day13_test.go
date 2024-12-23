package main

import (
	// "fmt"
	"reflect"
	"testing"
)

func TestDataIsExtractedFromInput(t *testing.T) {
	data := []string{
		"Button A: X+94, Y+34",
		"Button B: X+22, Y+67",
		"Prize: X=8400, Y=5400",
		"",
		"Button A: X+26, Y+66",
		"Button B: X+67, Y+21",
		"Prize: X=12748, Y=12176",
		"",
		"Button A: X+17, Y+86",
		"Button B: X+84, Y+37",
		"Prize: X=7870, Y=6450",
		"",
		"Button A: X+69, Y+23",
		"Button B: X+27, Y+71",
		"Prize: X=18641, Y=10279",
	}

	machinesDetails := getMachineDetails(data)

	// fmt.Println(machinesDetails)
	got := machinesDetails[3].prize

	want := [2]int{18641, 10279}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestFindPossiblePrizes(t *testing.T) {
	t.Run("finds possible prize", func(t *testing.T) {
		data := []string{
			"Button A: X+94, Y+34",
			"Button B: X+22, Y+67",
			"Prize: X=8400, Y=5400",
		}

		machinesDetails := getMachineDetails(data)

		permutations := findPrizePermutations(machinesDetails[0])
	   // fmt.Println("got", permutations)
		got := permutations[0]
		want := [2]int{80, 40}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("finds cheapest way to win", func(t *testing.T) {
		data := []string{
			"Button A: X+17, Y+86",
			"Button B: X+84, Y+37",
			"Prize: X=7870, Y=6450",
		}

		machinesDetails := getMachineDetails(data)

		permutations := findPrizePermutations(machinesDetails[0])
		got := permutations[0]
		want := [2]int{38, 86}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("find total tokens", func(t *testing.T) {
		data := []string{
			"Button A: X+94, Y+34",
			"Button B: X+22, Y+67",
			"Prize: X=8400, Y=5400",
			"",
			"Button A: X+26, Y+66",
			"Button B: X+67, Y+21",
			"Prize: X=12748, Y=12176",
			"",
			"Button A: X+17, Y+86",
			"Button B: X+84, Y+37",
			"Prize: X=7870, Y=6450",
			"",
			"Button A: X+69, Y+23",
			"Button B: X+27, Y+71",
			"Prize: X=18641, Y=10279",
		}
    got := FindPossiblePrizes(data)

    want := 480

    if got != want {
      t.Errorf("got %d, want %d", got, want)
    }
	})
}
