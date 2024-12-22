package main

import (
	"bufio"
	"fmt"
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
			machine.prize = [2]int{xcount, ycount}
			machinesDetails = append(machinesDetails, machine)
		}

	}

	return machinesDetails
}

func findPrizePermutations(machine Machine) [][2]int {
	prizePermutations := [][2]int{}
	aCount := 1
	aXVal := machine.buttonA[0] * aCount
	aYVal := machine.buttonA[1] * aCount
	bXVal := machine.buttonB[0]
	bYVal := machine.buttonB[1]
	prizeRestX := machine.prize[0] - aXVal
	prizeRestY := machine.prize[1] - aYVal

	for aXVal < machine.prize[0] && aYVal < machine.prize[1] {
		// fmt.Println(aCount, aXVal, aYVal, bXVal, bYVal, prizeRestX, prizeRestY)
		bModX := prizeRestX % bXVal
		bModY := prizeRestY % bYVal
		if bModX == 0 && bModY == 0 {
			// fmt.Println("got here")
			bCountX := prizeRestX / bXVal
			bCountY := prizeRestY / bYVal
			if bCountX == bCountY {
        prizePermutations = append(prizePermutations, [2]int{aCount, bCountX})
			}
		}

		aCount++
		aXVal = machine.buttonA[0] * aCount
		aYVal = machine.buttonA[1] * aCount
		prizeRestX = machine.prize[0] - aXVal
		prizeRestY = machine.prize[1] - aYVal
	}
	return prizePermutations
}

func FindPossiblePrizes(data []string) int {
	machinesDetails := getMachineDetails(data)
	totalTokens := 0

	for _, machine := range machinesDetails {
    // fmt.Println(machine)
		prizePermutations := findPrizePermutations(machine)
		fmt.Println(prizePermutations)
		if len(prizePermutations) == 0 {
			continue
		}
		tokens := prizePermutations[0][0]*3 + prizePermutations[0][1]
		// fmt.Println(tokens)
		for _, permutation := range prizePermutations {
			if tmpTokens := permutation[0]*3 + permutation[1]; tmpTokens < tokens {
				tokens = tmpTokens
			}
		}
		totalTokens += tokens
    fmt.Println("machine done")

	}
	fmt.Println("")
	return totalTokens
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
  fmt.Println(totalTokens)
}
