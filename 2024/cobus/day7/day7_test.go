package main

import (
	"fmt"
	"testing"
)

func TestCheckEquation(t *testing.T) {
	t.Run("valid equation", func(t *testing.T) {
		equations := []string{
			"190: 10 19",
		}

		answer, numbers := extractEquationInfo(equations[0])
		// fmt.Println(answer)
		// fmt.Println(numbers)
		got := checkEquation(answer, numbers)

		want := true

		if got != want {
			t.Errorf("got %t, want %t", got, want)
		}
	})

	t.Run("valid more than 2 equation", func(t *testing.T) {
		equations := []string{
			"292: 11 6 16 20",
		}

		answer, numbers := extractEquationInfo(equations[0])
		// fmt.Println(answer)
		// fmt.Println(numbers)
		got := checkEquation(answer, numbers)

		want := true

		if got != want {
			t.Errorf("got %t, want %t", got, want)
		}
	})

	t.Run("border case max", func(t *testing.T) {
		equations := []string{
			"21120: 11 6 16 20",
		}

		answer, numbers := extractEquationInfo(equations[0])
		// fmt.Println(answer)
		// fmt.Println(numbers)
		got := checkEquation(answer, numbers)

		want := true

		if got != want {
			t.Errorf("got %t, want %t", got, want)
		}
	})

	t.Run("border case min", func(t *testing.T) {
		equations := []string{
			"53: 11 6 16 20",
		}

		answer, numbers := extractEquationInfo(equations[0])
		// fmt.Println(answer)
		// fmt.Println(numbers)
		got := checkEquation(answer, numbers)

		want := true

		if got != want {
			t.Errorf("got %t, want %t", got, want)
		}
	})

}

func TestTotalCalibrationResult(t *testing.T) {
	equations := []string{
		"190: 10 19",
		"3267: 81 40 27",
		"83: 17 5",
		"156: 15 6",
		"7290: 6 8 6 15",
		"161011: 16 10 13",
		"192: 17 8 14",
		"21037: 9 7 18 13",
		"292: 11 6 16 20",
	}

  got := TotalCalibrationResult(equations)

  want := uint64(11387)

  if got != want {
    t.Errorf("got %d, want %d", got, want)
  }

  t1 := int64(1229588697000)
  fmt.Println(t1+1)
}
