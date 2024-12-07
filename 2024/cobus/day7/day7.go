package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Operation func(uint64, uint64) uint64

// var OperatorMap = map[string]Operation{
// 	"add":      func(a, b int) int { return a + b },
// 	"multiply": func(a, b int) int { return a * b },
// }

var OperatorSlice = [3]Operation{
	func(a, b uint64) uint64 { return a + b },
	func(a, b uint64) uint64 { return a * b },
	func(a, b uint64) uint64 {
    total := strconv.FormatUint(a, 10) + strconv.FormatUint(b, 10)
    totalInt, _ := strconv.ParseUint(total, 10, 64)
    return totalInt
  },
}

func extractEquationInfo(line string) (uint64, []uint64) {
	numbers := []uint64{}
	answer := uint64(0)
	// for _, line := range equations {
	tmp := strings.Split(line, ":")
  answer, _ = strconv.ParseUint(tmp[0], 10, 64)
	// answer, _ = strconv.Atoi(tmp[0])
	tmp = strings.Fields(tmp[1])
	for _, entry := range tmp {
		// number, _ := strconv.Atoi(entry)
		number, _ := strconv.ParseUint(entry, 10, 64)
		numbers = append(numbers, number)
	}

	// }
	return answer, numbers
}

// func calculateResult(i int, numbers []int, operatorList map[string]Operation) {
//
//   if i >= len(numbers) {
//     return result
//   }
//
//     for _, operator := range operatorList {
//     tmpResult := operator(numbers[i])
//     result = calculateResult(i++, numbers, operatorList map[string]Operation)
//   }
//   // for i := 1; i < len(numbers); i++ {
//   //   result := numbers[i-1]
//   //   for _, operator := range operatorList {
//   //     result =
//   //   }
// }

func createEquationRecipes(operations uint64) [][]uint64 {
	base := float64(len(OperatorSlice))
	lineNums := math.Pow(base, float64(operations))
	// println(operations, int(base), int(lineNums))

	equationRecipes := make([][]uint64, uint64(lineNums))

	equationRecipes[0] = make([]uint64, operations)
	// fmt.Printf("first row: %d\n", equationRecipes[0])
	for row := 1; row < len(equationRecipes); row++ {
		equationLine := make([]uint64, operations)
		prevEquationLine := equationRecipes[row-1]
		copy(equationLine, prevEquationLine)
		// fmt.Printf("eq line: %d\n", equationLine)
		// test for carry
		// if prevEquationLine[0] + 1 >= len(OperatorSlice) {
		//   equationLine[0] = 0
		// }
		equationLine[0]++
		for col := 0; col < len(equationLine); col++ {
			// fmt.Printf("col %d val: %d\n", col, equationLine[col])
			if equationLine[col] >= uint64(len(OperatorSlice)) {
				equationLine[col] = 0
				if col+1 < len(equationLine) {
					equationLine[col+1]++
				}
			}
		}
		equationRecipes[row] = equationLine

	}

	return equationRecipes
}

func checkEquation(answer uint64, numbers []uint64) bool {
	operations := uint64(len(numbers) - 1)
	startNum := numbers[0]

	equationRecipes := createEquationRecipes(operations)
	// fmt.Println(equationRecipes)

	// for i := 1; i < len(numbers); i++ {
	for _, recipe := range equationRecipes {
		result := uint64(startNum)
		// fmt.Printf("\nstarting with %d\n", result)
		for i, operation := range recipe {
			result = OperatorSlice[operation](result, uint64(numbers[i+1]))
			// fmt.Printf("interim result after op type %d: %d\n", operation, result)
		}
		if result == answer {
      // suspicion is that value is too big, so let's try and look at long recipes
      if result > 101405378166 {
        fmt.Println(answer, numbers)
        fmt.Println(recipe)
      }
			return true
		}
		// }
	}

	return false
}

func TotalCalibrationResult(equations []string) uint64 {
	total := uint64(0)
	for _, line := range equations {
		answer, numbers := extractEquationInfo(line)

		if checkEquation(answer, numbers) {
      // fmt.Printf("got valid equation line for line %s\n", line)
      // fmt.Printf("answer: %d numbers: %d\n", answer, numbers)
			total += answer
		}
	}
	return total
}

func main() {
	fmt.Println("start")
	equations := []string{}

	file, _ := os.Open("./data")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		equations = append(equations, line)
	}

  total := TotalCalibrationResult(equations)
  fmt.Println(total)
}
