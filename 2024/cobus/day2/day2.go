package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func decCheck(numbers []int) bool {
	diff := -99
	for i := 1; i < len(numbers); i++ {
		diff = numbers[i-1] - numbers[i]
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func incCheck(numbers []int) bool {
	diff := -99
	for i := 1; i < len(numbers); i++ {
		diff = numbers[i] - numbers[i-1]
		// fmt.Printf("diff between %d and %d: %d\n", numbers[i], numbers[i-1], diff)
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func incDecCheck(numbers []int) bool {
	if numbers[0] > numbers[1] {
		return decCheck(numbers)
	}
	return incCheck(numbers)
}

func GetSafeReports(reportList [][]int) int {
	safeReports := 0
	for _, report := range reportList {
		// fmt.Printf("trying %d\n", report)
		if incDecCheck(report) {
			// fmt.Println("Safe")
      // fmt.Printf("report: %d\n", report)
			safeReports++
		} else {
			for i := 0; i < len(report); i++ {
        poppedreport := make([]int, len(report))
        _ = copy(poppedreport, report)
				poppedreport = append(poppedreport[:i], poppedreport[i+1:]...)
        // fmt.Printf("original report: %d, trying %d\n", report, poppedreport)
				if incDecCheck(poppedreport) {
					// fmt.Println("Safe")
					safeReports++
					break
				}
			}
		}

	}
	return safeReports
}

func main() {
	reportList := [][]int{}
	file, _ := os.Open("./data")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		reportEntry := []int{}
		for _, part := range parts {
			intPart, _ := strconv.Atoi(part)
			reportEntry = append(reportEntry, intPart)
		}
		reportList = append(reportList, reportEntry)
	}

	safeReport := GetSafeReports(reportList)

	fmt.Println(safeReport)
}
