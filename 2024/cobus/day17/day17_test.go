package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestComputerInit(t *testing.T) {

	t.Run("test bst", func(t *testing.T) {
		data := []string{
			"Register A: 0",
			"Register B: 0",
			"Register C: 9",
			"",
			"Program: 2,6",
		}

    computer := initComputer(data)
    runInstructions(&computer)
    fmt.Println(computer)
	})

	t.Run("test out", func(t *testing.T) {
		data := []string{
			"Register A: 10",
			"Register B: 0",
			"Register C: 0",
			"",
			"Program: 5,0,5,1,5,4",
		}

    computer := initComputer(data)
    got := runInstructions(&computer)
    fmt.Println(computer)
    want := []int{0,1,2}

    if !reflect.DeepEqual(got, want) {
      t.Errorf("got %d, want %d", got, want)
    }
	})

	t.Run("test multiple", func(t *testing.T) {
		data := []string{
			"Register A: 2024",
			"Register B: 0",
			"Register C: 0",
			"",
			"Program: 0,1,5,4,3,0",
		}

    computer := initComputer(data)
    got := runInstructions(&computer)
    fmt.Println(computer)
    want := []int{4,2,5,6,7,7,7,7,3,1,0}

    if !reflect.DeepEqual(got, want) {
      t.Errorf("got %d, want %d", got, want)
    }
	})

	t.Run("test bxl", func(t *testing.T) {
		data := []string{
			"Register A: 0",
			"Register B: 29",
			"Register C: 0",
			"",
			"Program: 1,7",
		}

    computer := initComputer(data)
    runInstructions(&computer)
    fmt.Println(computer)
    // want := []int{4,2,5,6,7,7,7,7,3,1,0}

    // if !reflect.DeepEqual(got, want) {
    //   t.Errorf("got %d, want %d", got, want)
    // }
	})

	t.Run("test bxc", func(t *testing.T) {
		data := []string{
			"Register A: 0",
			"Register B: 2024",
			"Register C: 43690",
			"",
			"Program: 4,0",
		}

    computer := initComputer(data)
    runInstructions(&computer)
    fmt.Println(computer)
    // want := []int{4,2,5,6,7,7,7,7,3,1,0}

    // if !reflect.DeepEqual(got, want) {
    //   t.Errorf("got %d, want %d", got, want)
    // }
	})

	t.Run("More complex ex", func(t *testing.T) {
		data := []string{
			"Register A: 729",
			"Register B: 0",
			"Register C: 0",
			"",
			"Program: 0,1,5,4,3,0",
		}

    computer := initComputer(data)
    got := runInstructions(&computer)

    want := []int{4,6,3,5,6,3,5,2,1,0}

    if !reflect.DeepEqual(got, want) {
      t.Errorf("got %d, want %d", got, want)
    }
    // fmt.Println(computer)
	})

	t.Run("Trying to understand part2", func(t *testing.T) {
		data := []string{
			"Register A: 117440",
			"Register B: 0",
			"Register C: 0",
			"",
			"Program: 0,3,5,4,3,0",
		}

    computer := initComputer(data)
    output := runInstructions(&computer)

    // want := []int{4,6,3,5,6,3,5,2,1,0}

    // if !reflect.DeepEqual(got, want) {
    //   t.Errorf("got %d, want %d", got, want)
    // }
    fmt.Println(computer, output)
	})

}
