package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Machine struct {
	buttonA [2]int
	buttonB [2]int
	prize   [2]int
}

func getMachineDetails(data []string) []Machine {
	machinesDetails := []Machine{}
	machine := Machine{}
	for _, line := range data {
		if line == "" {
			machine = Machine{}
			continue
		}
		entries := strings.Fields(line)

		if strings.Contains(line, "Button A") {
			_, xstring, _ := strings.Cut(entries[2], "X+")
			_, ystring, _ := strings.Cut(entries[3], "Y+")
			xcount, _ := strconv.Atoi(xstring[:len(xstring)-1])
			ycount, _ := strconv.Atoi(ystring)
			machine.buttonA = [2]int{xcount, ycount}
		}

		if strings.Contains(line, "Button B") {
			_, xstring, _ := strings.Cut(entries[2], "X+")
			_, ystring, _ := strings.Cut(entries[3], "Y+")
			xcount, _ := strconv.Atoi(xstring[:len(xstring)-1])
			ycount, _ := strconv.Atoi(ystring)
			machine.buttonB = [2]int{xcount, ycount}
		}

		if strings.Contains(line, "Prize") {
			_, xstring, _ := strings.Cut(entries[1], "X=")
			_, ystring, _ := strings.Cut(entries[2], "Y=")
			xcount, _ := strconv.Atoi(xstring[:len(xstring)-1])
			ycount, _ := strconv.Atoi(ystring)
			machine.prize = [2]int{xcount + 10000000000000, ycount + 10000000000000}
			// machine.prize = [2]int{xcount, ycount}
			machinesDetails = append(machinesDetails, machine)
		}

	}

	return machinesDetails
}

const epsilon = 0.01 // or whatever small threshold you prefer

func shouldRoundToInt(f float64) (int, bool) {
	// Check if it's very close to a whole number
	nearestInt := math.Round(f)
	// fmt.Println("nearestInt", nearestInt)
	// fmt.Println("f-nearestInt", f-nearestInt)
	if math.Abs(f-nearestInt) < epsilon {
		return int(nearestInt), true
	}
	return -1, false
}

func findIntersection(machine Machine) int {
	ax := machine.buttonA[0]
	ay := machine.buttonA[1]
	bx := machine.buttonB[0]
	by := machine.buttonB[1]
	px := machine.prize[0]
	py := machine.prize[1]

	// px += 10000000000000
	// py += 10000000000000

	det := ax*by - ay*bx

	if det == 0 {
		return 0
	}

	aPress := (px*by - py*bx) / det
	bPress := (ax*py - ay*px) / det

	if ax*aPress+bx*bPress == px && ay*aPress+by*bPress == py && aPress >= 0 && bPress >= 0 {
		cost := aPress*3 + bPress
		fmt.Println(machine, [2]int{aPress, bPress})
		return cost
		// output += cost
	}
	return 0
}

func findPrizePermutations(machine Machine) [][2]int {
	prizePermutations := [][2]int{}

	// part2
	intersection1 := calcIntersection(machine.prize, machine.buttonA, machine.buttonB)
	// intersection1 := calcInter(machine.buttonA, machine.buttonB, machine.prize)
	// fmt.Println("intersection", intersection1)

	// sectX := int(math.Round(intersection[0]))
	// sectY := int(math.Round(intersection[1]))
	sectX, _ := shouldRoundToInt(intersection1[0])
	sectY, _ := shouldRoundToInt(intersection1[1])
	//
	// fmt.Println("sect", sectX, sectY)
	// if intersection1[0] < 0 || intersection1[1] < 0 {
	// 	return prizePermutations
	// }

	if sectX%machine.buttonA[0] == 0 && sectY%machine.buttonA[1] == 0 {
		aCountX := sectX / machine.buttonA[0]
		aCountY := sectY / machine.buttonA[1]
		// fmt.Println("aCount", aCountX, aCountY)
		// fmt.Println("")

		if aCountX == aCountY {
			bCountX := (machine.prize[0] - sectX) / machine.buttonB[0]
			// aCountXInt, _ := shouldRoundToInt(aCountX)
			// bCountXInt, _ := shouldRoundToInt(bCountX)
			prizePermutations = append(prizePermutations, [2]int{aCountX, bCountX})
		}
	}

	if machine.prize[0]%machine.buttonA[0] == 0 && machine.prize[1]%machine.buttonA[1] == 0 {
		aCount := machine.prize[0] / machine.buttonA[0]
		if machine.prize[1]/machine.buttonA[1] == aCount {
			prizePermutations = append(prizePermutations, [2]int{aCount, 0})
		}
	}

	if machine.prize[0]%machine.buttonB[0] == 0 && machine.prize[1]%machine.buttonB[1] == 0 {
		bCount := machine.prize[0] / machine.buttonB[0]
		if machine.prize[1]/machine.buttonB[1] == bCount {
			prizePermutations = append(prizePermutations, [2]int{0, bCount})
		}
	}

	// sectX, _ = shouldRoundToInt(intersection2[0])
	// sectY, _ = shouldRoundToInt(intersection2[1])
	//
	// if sectX < 0 || sectY < 0 {
	// 	return prizePermutations
	// }
	//
	// if sectX%machine.buttonB[0] == 0 && sectY%machine.buttonB[1] == 0 {
	// 	bCountX := sectX / machine.buttonB[0]
	// 	bCountY := sectY / machine.buttonB[1]
	//
	// 	if bCountX == bCountY {
	// 		aCountX := (machine.prize[0] - sectX) / machine.buttonA[0]
	// 		prizePermutations = append(prizePermutations, [2]int{aCountX, bCountX})
	//      fmt.Println(machine, prizePermutations)
	// 	}
	// }
	// }

	// part 1
	// aCount := 1
	// aXVal := machine.buttonA[0] * aCount
	// aYVal := machine.buttonA[1] * aCount
	// bXVal := machine.buttonB[0]
	// bYVal := machine.buttonB[1]
	// prizeRestX := machine.prize[0] - aXVal
	// prizeRestY := machine.prize[1] - aYVal
	//
	// for aXVal < machine.prize[0] && aYVal < machine.prize[1] {
	// 	// fmt.Println(aCount, aXVal, aYVal, bXVal, bYVal, prizeRestX, prizeRestY)
	// 	bModX := prizeRestX % bXVal
	// 	bModY := prizeRestY % bYVal
	// 	if bModX == 0 && bModY == 0 {
	// 		// fmt.Println("got here")
	// 		bCountX := prizeRestX / bXVal
	// 		bCountY := prizeRestY / bYVal
	// 		if bCountX == bCountY {
	//        prizePermutations = append(prizePermutations, [2]int{aCount, bCountX})
	// 		}
	// 	}
	//
	// 	aCount++
	// 	aXVal = machine.buttonA[0] * aCount
	// 	aYVal = machine.buttonA[1] * aCount
	// 	prizeRestX = machine.prize[0] - aXVal
	// 	prizeRestY = machine.prize[1] - aYVal
	// }
	return prizePermutations
}

func calcInter(buttonA, buttonB, prizePoint [2]int) [2]float64 {
	a1 := float64(buttonA[0])
	b1 := float64(buttonA[1])
	c1 := float64(0)
	a2 := float64(buttonB[0])
	b2 := float64(buttonB[1])
	c2 := float64(prizePoint[1]) - a2/b2*float64(prizePoint[0])
	fmt.Println(c2)
	x := (b1*c2 - b2*c1) / (a1*b2 - a2*b1)
	y := (c1*a2 - c2*a1) / (a1*b2 - a2*b1)

	return [2]float64{x, y}
}

func calcIntersection(pEndButton, mStartButton, mEndButton [2]int) [2]float64 {
	mStart := float64(mStartButton[1]) / float64(mStartButton[0])
	mEnd := float64(mEndButton[1]) / float64(mEndButton[0])
	// pStartX := float64(pStartButton[0])
	// pStartY := float64(pStartButton[1])
	pEndX := float64(pEndButton[0])
	pEndY := float64(pEndButton[1])

	x := (-mEnd*pEndX + pEndY) / (mStart - mEnd)

	y := mStart * (x)

	return [2]float64{x, y}
}

func sanityCheck(machine Machine, prizePermutations [][2]int) {
	xPrize := prizePermutations[0][0]*machine.buttonA[0] + prizePermutations[0][1]*machine.buttonB[0]
	yPrize := prizePermutations[0][0]*machine.buttonA[1] + prizePermutations[0][1]*machine.buttonB[1]

	if xPrize != machine.prize[0] || yPrize != machine.prize[1] {
		fmt.Println("something weird here", machine, prizePermutations)
	}
}

func FindPossiblePrizes(data []string) int {
	machinesDetails := getMachineDetails(data)
	totalTokens := 0

	for _, machine := range machinesDetails {
		// fmt.Println(machine)
		prizePermutations := findPrizePermutations(machine)
		if len(prizePermutations) == 0 {
			continue
		}
		fmt.Println(machine, prizePermutations)
		sanityCheck(machine, prizePermutations)
		tokens := prizePermutations[0][0]*3 + prizePermutations[0][1]
		// fmt.Println(tokens)
		for _, permutation := range prizePermutations {
			if tmpTokens := permutation[0]*3 + permutation[1]; tmpTokens < tokens {
				tokens = tmpTokens
			}
		}
		totalTokens += tokens
		// fmt.Println("machine done")

	}
	fmt.Println("")
	return totalTokens
}

func FindPrizeCost(data []string) int {
	cost := 0
	machines := getMachineDetails(data)
	for _, machine := range machines {
		cost += findIntersection(machine)
	}
	return cost
}

func main() {
	data := []string{}

	file, _ := os.Open("./data")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, line)
	}

	totalTokens := FindPossiblePrizes(data)
	// totalTokens := FindPrizeCost(data)
	fmt.Println(totalTokens)
}
