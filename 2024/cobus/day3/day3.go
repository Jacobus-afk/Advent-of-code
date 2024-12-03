package main

import (
	// "fmt"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findInstructionPattern(data string) []string {
	matchedPatternReg := []string{}
	// fmt.Println(data)
	for {
		startPos := strings.Index(data, "mul(")
		if startPos < 0 {
			break
		}
		data = data[startPos+4:]

		endPos := strings.Index(data, ")")
		if endPos < 0 {
			break
		}
		// fmt.Printf("startpos %d, endpos %d\n", startPos, endPos)
		matchedPattern := data[:endPos]
		// fmt.Println(matchedPattern)
		matchedPatternReg = append(matchedPatternReg, matchedPattern)
		data = data[endPos+1:]
	}

	return matchedPatternReg
}

func multiplyValidInstruction(dataEntry string) int {
	entries := strings.Split(dataEntry, ",")

	if len(entries) == 2 {
		val1, err1 := strconv.Atoi(entries[0])
		val2, err2 := strconv.Atoi(entries[1])

		if err1 == nil && err2 == nil {
			return val1 * val2
		}
	}
	return -1
}

func AddValidInstructions(data string) int {
	matchedPatternReg := findInstructionPattern(data)
	tally := 0

for _, match := range matchedPatternReg {
		val := multiplyValidInstruction(match)

		if val == -1 {
			// fmt.Printf("not valid instruction: %s\n", match)
			lastTryIndex := strings.LastIndex(match, "mul(") + 4
      if lastTryIndex < 4 {
        continue
      }
			// fmt.Printf("last try val: %s, index: %d\n", match[lastTryIndex:], lastTryIndex)
			val = multiplyValidInstruction(match[lastTryIndex:])
			if val == -1 {
				continue
			}
		}
		tally += val
	}
	return tally
}

func removeConditionalStatements(data string) string {
  validData := ""
	// fmt.Println(data)
	for {
		startPos := strings.Index(data, "don't()")
		if startPos < 0 {
      validData += data
			break
		}
    validData += data[:startPos]
		data = data[startPos+4:]

		endPos := strings.Index(data, "do()")
		if endPos < 0 {
			break
		}
		// fmt.Printf("startpos %d, endpos %d\n", startPos, endPos)
    data = data[endPos+4:]
	}

	return validData
}

func main() {
	lineReg := []string{}
  allInstr := ""
	// tally := 0
	file, _ := os.Open("./data")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineReg = append(lineReg, line)
	}
	for _, instr := range lineReg {
    allInstr += instr
		// tally += sum
	}
  sum := AddValidInstructions(allInstr)
	fmt.Println(sum)


  treatedInstr := removeConditionalStatements(allInstr)
  treatedSum := AddValidInstructions(treatedInstr)
  fmt.Println(treatedSum)


}
