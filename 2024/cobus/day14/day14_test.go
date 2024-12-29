package main

import (
	"fmt"
	"testing"
)

func TestBuildRobotStartingGrid(t *testing.T) {
	t.Run("builds grid with robots in correct starting position", func(t *testing.T) {
		grid := buildGrid([2]int{11, 7})

		robots := positionRobots(grid, []string{
			"p=0,4 v=3,-3",
			"p=6,3 v=-1,-3",
			"p=10,3 v=-1,2",
			"p=2,0 v=2,-1",
			"p=0,0 v=1,3",
			"p=3,0 v=-2,-2",
			"p=7,6 v=-1,-3",
			"p=3,0 v=-1,-2",
			"p=9,3 v=2,3",
			"p=7,3 v=-1,2",
			"p=2,4 v=2,-3",
			"p=9,5 v=-3,-3",
		})

		for _, field := range grid {
			fmt.Println(field)
		}

		fmt.Println(robots)
	})
}

func TestRobotMovement(t *testing.T) {
	t.Run("robot moves correctly", func(t *testing.T) {
		grid := buildGrid([2]int{11, 7})

		robots := positionRobots(grid, []string{
			"p=2,4 v=2,-3",
		})

		moveRobots(5, grid, robots)

		for _, field := range grid {
			fmt.Println(field)
		}

		fmt.Println(robots)
	})

	t.Run("robots move correctly", func(t *testing.T) {
		grid := buildGrid([2]int{11, 7})

		robots := positionRobots(grid, []string{
			"p=0,4 v=3,-3",
			"p=6,3 v=-1,-3",
			"p=10,3 v=-1,2",
			"p=2,0 v=2,-1",
			"p=0,0 v=1,3",
			"p=3,0 v=-2,-2",
			"p=7,6 v=-1,-3",
			"p=3,0 v=-1,-2",
			"p=9,3 v=2,3",
			"p=7,3 v=-1,2",
			"p=2,4 v=2,-3",
			"p=9,5 v=-3,-3",
		})

		moveRobots(100, grid, robots)

		for _, field := range grid {
			fmt.Println(field)
		}

		fmt.Println(robots)
	})

	t.Run("calculates safety factor", func(t *testing.T) {
		gridDimensions := [2]int{11, 7}

		got := CalculateSafetyFactor(100, gridDimensions, []string{
			"p=0,4 v=3,-3",
			"p=6,3 v=-1,-3",
			"p=10,3 v=-1,2",
			"p=2,0 v=2,-1",
			"p=0,0 v=1,3",
			"p=3,0 v=-2,-2",
			"p=7,6 v=-1,-3",
			"p=3,0 v=-1,-2",
			"p=9,3 v=2,3",
			"p=7,3 v=-1,2",
			"p=2,4 v=2,-3",
			"p=9,5 v=-3,-3",
		})


    want := 12

    if got != want {
      t.Errorf("got %d, want %d", got, want)
    }
	})
}
