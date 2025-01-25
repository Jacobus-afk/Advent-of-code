package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

  aoc "aoc-24/lib"
)

type Instruction struct {
	opcode  int
	operand int
}

type Computer struct {
	registerA int
	registerB int
	registerC int

	program map[int]Instruction
}

func translateOperand(computer Computer, operand int) int {
  switch operand {
  case 0,1,2,3:
    return operand
  case 4:
    return computer.registerA
  case 5:
    return computer.registerB
  case 6:
    return computer.registerC
  case 7:
    return -7
  }
  return -1
}

func adv(computer *Computer, operand int) {
  combo := translateOperand(*computer, operand)
  numerator := computer.registerA
  denominator := math.Pow(2, float64(combo))

  computer.registerA = numerator / int(denominator)
}

func bdv(computer *Computer, operand int) {
  combo := translateOperand(*computer, operand)
  numerator := computer.registerA
  denominator := math.Pow(2, float64(combo))

  computer.registerB = numerator / int(denominator)
}

func cdv(computer *Computer, operand int) {
  combo := translateOperand(*computer, operand)
  numerator := computer.registerA
  denominator := math.Pow(2, float64(combo))

  computer.registerC = numerator / int(denominator)
}

func bxl(computer *Computer, operand int) {
  computer.registerB = computer.registerB ^ operand
}

func bst(computer *Computer, operand int) {
  combo := translateOperand(*computer, operand)
  computer.registerB = combo % 8
}

func bxc(computer *Computer) {
  computer.registerB = computer.registerB ^ computer.registerC
}

func out(computer Computer, operand int) int {
  combo := translateOperand(computer, operand)
  return combo % 8
}

func runInstructions(computer *Computer) []int {
  output := []int{}
  instrCtr := 0

  for instrCtr < len(computer.program) * 2 {
    instruction := computer.program[instrCtr]

    switch instruction.opcode {
    case 0:
      adv(computer, instruction.operand)
    case 1:
      bxl(computer, instruction.operand)
    case 2:
      bst(computer, instruction.operand)
    case 3:
      if computer.registerA != 0 {
        instrCtr = instruction.operand
        continue
      }
    case 4:
      bxc(computer)
    case 5:
      outval := out(*computer, instruction.operand)
      output = append(output, outval)
    case 6:
      bdv(computer, instruction.operand)
    case 7:
      cdv(computer, instruction.operand)
    }

    instrCtr += 2
  }
  return output
}

func initComputer(data []string) Computer {
	var regA, regB, regC int
	var prog string
	instrMap := map[int]Instruction{}

	fmt.Sscanf(data[0], "Register A: %d", &regA)
	fmt.Sscanf(data[1], "Register B: %d", &regB)
	fmt.Sscanf(data[2], "Register C: %d", &regC)
	fmt.Sscanf(data[4], "Program: %s", &prog)

	instrStr := strings.Split(prog, ",")

	for i := 0; i < len(instrStr); i += 2 {
		opcode, _ := strconv.Atoi(instrStr[i])
		operand, _ := strconv.Atoi(instrStr[i+1])
		instrMap[i] = Instruction{opcode: opcode, operand: operand}
	}

  return Computer{registerA: regA, registerB: regB, registerC: regC, program: instrMap}
}

func main() {
  data := aoc.ReadFileByLine("./data")

  computer := initComputer(data)
  output := runInstructions(&computer)
  for _, entry := range output {
    fmt.Printf("%d,", entry)
  }

}
