package main

import "testing"

func TestIncDecCheck(t *testing.T) {
	t.Run("test decrement", func(t *testing.T) {
		got := incDecCheck([]int{7, 6, 4, 2, 1})

		want := true
		if got != want {
			t.Errorf("got %t, want %t", got, want)
		}
	})

	t.Run("test increment", func(t *testing.T) {
		got := incDecCheck([]int{1, 3, 6, 7, 9})

		want := true
		if got != want {
			t.Errorf("got %t, want %t", got, want)
		}
	})

	t.Run("test decrement of more than 2", func(t *testing.T) {
		got := incDecCheck([]int{9, 7, 6, 2, 1})

		want := false
		if got != want {
			t.Errorf("got %t, want %t", got, want)
		}
	})

	t.Run("test increment of more than 2", func(t *testing.T) {
		got := incDecCheck([]int{1, 2, 7, 8, 9})

		want := false
		if got != want {
			t.Errorf("got %t, want %t", got, want)
		}
	})

	t.Run("test increase then decrease", func(t *testing.T) {
		got := incDecCheck([]int{1, 3, 2, 4, 5})

		want := false
		if got != want {
			t.Errorf("got %t, want %t", got, want)
		}
	})

	t.Run("test neither increase nor decrease", func(t *testing.T) {
		got := incDecCheck([]int{8, 6, 4, 4, 1})

		want := false
		if got != want {
			t.Errorf("got %t, want %t", got, want)
		}
	})
}

func TestGetSafeReports(t *testing.T) {
	reportList := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}

	got := GetSafeReports(reportList)

	want := 2
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
