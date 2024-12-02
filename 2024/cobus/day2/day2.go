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
		if incDecCheck(report) {
			safeReports++
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
