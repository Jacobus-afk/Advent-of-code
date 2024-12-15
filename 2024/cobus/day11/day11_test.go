package main

import (
	"reflect"
	"testing"
)

func TestBlinks(t *testing.T) {
	t.Run("single blink", func(t *testing.T) {
		stones := []string{
			"0", "1", "10", "99", "999",
		}

		got := Blinks(1, stones)

		want := []string{"1", "2024", "1", "0", "9", "9", "2021976"}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %s, want %s", got, want)
		}
	})

	t.Run("multiple blinks", func(t *testing.T) {
		stones := []string{
			"125", "17",
		}

		got := Blinks(6, stones)

		want := []string{
			"2097446912", "14168", "4048", "2", "0", "2", "4", "40", "48", "2024", "40", "48",
			"80", "96", "2", "8", "6", "7", "6", "0", "3", "2",
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}
