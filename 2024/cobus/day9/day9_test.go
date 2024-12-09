package main

import (
	"reflect"
	"testing"
)

func TestCreateChecksum(t *testing.T) {
	// t.Run("disk map translates to blocks", func(t *testing.T) {
	// 	got, blocksSlice := translateDiskMap("12345")
	//
	// 	want := "0..111....22222"
	//
	// 	fmt.Println(blocksSlice)
	// 	if got != want {
	// 		t.Errorf("got %s, want %s", got, want)
	// 	}
	// })

	// t.Run("disk map translates to blocks", func(t *testing.T) {
	// 	got := translateDiskMap("2333133121414131402")
	//
	// 	want := "00...111...2...333.44.5555.6666.777.888899"
	//
	// 	if got != want {
	// 		t.Errorf("got %s, want %s", got, want)
	// 	}
	// })
	//
	t.Run("frees up space", func(t *testing.T) {
		got := []int{0, -1, -1, 1, 1, 1, -1, -1, -1, -1, 2, 2, 2, 2, 2}
		freeUpSpaceFragmented(got)

		want := []int{0, 2, 2, 1, 1, 1, 2, 2, 2, -1, -1, -1, -1, -1, -1}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	//
	// t.Run("frees up space", func(t *testing.T) {
	// 	got := freeUpSpace("00...111...2...333.44.5555.6666.777.888899")
	//
	// 	want := "0099811188827773336446555566.............."
	//
	// 	if got != want {
	// 		t.Errorf("got %s, want %s", got, want)
	// 	}
	// })
	//
	t.Run("calcs checksum", func(t *testing.T) {
	  got := CalcCheckSum("2333133121414131402")

	  want := 2858

	  if got != want {
	    t.Errorf("got %d, want %d", got, want)
	  }
	})
}
