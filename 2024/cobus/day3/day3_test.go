package main

import (
	// "fmt"
	"reflect"
	"testing"
)

func TestFindInstructions(t *testing.T) {
	t.Run("find matched patterns of instructions", func(t *testing.T) {
		got := findInstructionPattern(
			"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
		)
		// fmt.Println(got)

		want := []string{"2,4", "5,5", "32,64]then(mul(11,8", "8,5"}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %s, want %s", got, want)
		}
	})

	t.Run("check for valid instructions", func(t *testing.T) {
		got := multiplyValidInstruction("2,4")

		want := 8

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("check for invalid instructions", func(t *testing.T) {
		got := multiplyValidInstruction("32,64]then(mul(11,8")

		want := -1

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestAddValidInstructions(t *testing.T) {
	got := AddValidInstructions(
		"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
	)

  want := 161

  if got != want {
    t.Errorf("got %d, want %d", got, want)
  }
}

func TestRemoveConditionalStatements(t *testing.T) {
  got := removeConditionalStatements("xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))")

  want := "xmul(2,4)&mul[3,7]!^?mul(8,5))"

  if got != want {
    t.Errorf("got %s, want %s", got, want)
  }
}
