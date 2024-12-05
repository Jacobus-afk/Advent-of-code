package main

import (
	// "fmt"
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var PageOrderingRules = []string{
	"47|53",
	"97|13",
	"97|61",
	"97|47",
	"75|29",
	"61|13",
	"75|53",
	"29|13",
	"97|29",
	"53|29",
	"61|53",
	"97|53",
	"61|29",
	"47|13",
	"75|47",
	"97|75",
	"47|61",
	"75|61",
	"47|29",
	"75|13",
	"53|13",
}

func populateMapValues(currentMap map[int][]int, key, val int) {
	// fmt.Printf("key: %d, val: %d\n", key, val)
	mapVal, ok := currentMap[key]
	// fmt.Printf("mapVal: %d\n", mapVal)

	if ok {
		mapVal = append(mapVal, val)
		currentMap[key] = mapVal
	} else {
		currentMap[key] = []int{val}
	}
}

func buildOrderingRuleMap(rules []string) [2]map[int][]int {
	rightMap := make(map[int][]int)
	wrongMap := make(map[int][]int)

	for _, rule := range rules {
		ruleEntries := strings.Split(rule, "|")
		right, _ := strconv.Atoi(ruleEntries[0])
		wrong, _ := strconv.Atoi(ruleEntries[1])

		populateMapValues(rightMap, right, wrong)
		populateMapValues(wrongMap, wrong, right)
	}

	return [2]map[int][]int{rightMap, wrongMap}
}

func findPageInRule(ruleMap map[int][]int, rule, page int) bool {
	val, ok := ruleMap[rule]
	if ok {
		return slices.Contains(val, page)
	}
	return false
}

func checkValidOrdering(pages []int, orderingRuleMaps [2]map[int][]int) bool {
	rightMap := orderingRuleMaps[0]
	wrongMap := orderingRuleMaps[1]
	// fmt.Printf("rightmap: %d\nwrongmap: %d\n", rightMap, wrongMap)

	for i := 1; i < len(pages); i++ {
		beforePage := pages[i-1]
		afterPage := pages[i]

		// fmt.Printf("beforePage: %d, afterPage: %d\n", beforePage, afterPage)
		if findPageInRule(wrongMap, beforePage, afterPage) {
			return false
		}

		if findPageInRule(rightMap, beforePage, afterPage) {
			continue
		}

	}
	return true
}

func GetValidMiddlePages(pageUpdates [][]int, orderingRuleMaps [2]map[int][]int) int {
	sum := 0
	for _, pages := range pageUpdates {
		if checkValidOrdering(pages, orderingRuleMaps) {
			middleVal := pages[len(pages)/2]
			sum += middleVal
			// fmt.Printf("middleVal: %d\n", middleVal)
		}
	}
	return sum
}

func FixInvalidMiddlePages(pageUpdates [][]int, orderingRuleMaps [2]map[int][]int) int {
	return 0
}

func main() {
	lineReg := []string{}

	file, _ := os.Open("./orderingdata")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineReg = append(lineReg, line)
	}
	orderingRuleMaps := buildOrderingRuleMap(lineReg)
	// fmt.Println(orderingRuleMaps)

	pageUpdates := [][]int{}
	file, _ = os.Open("./data")
	defer file.Close()

	scanner = bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		// fmt.Println(parts)

		reportEntry := []int{}
		for _, part := range parts {
			// fmt.Println(part)
			partSlice := strings.Split(part, ",")
			for _, entry := range partSlice {
				intPart, _ := strconv.Atoi(entry)
				reportEntry = append(reportEntry, intPart)
			}
		}
		pageUpdates = append(pageUpdates, reportEntry)
	}
	// fmt.Println(pageUpdates)

	sum := GetValidMiddlePages(pageUpdates, orderingRuleMaps)
	fmt.Println(sum)
}
