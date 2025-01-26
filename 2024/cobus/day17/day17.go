package main

import (
	"fmt"
	// "math"
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
	case 0, 1, 2, 3:
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
	// denominator := math.Pow(2, float64(combo))
	// denominator := 1 << combo

	// computer.registerA = numerator / int(denominator)
	computer.registerA = numerator >> combo
}

func bdv(computer *Computer, operand int) {
	combo := translateOperand(*computer, operand)
	numerator := computer.registerA
	// denominator := math.Pow(2, float64(combo))
	// denominator := 1 << combo

	// computer.registerB = numerator / int(denominator)
	computer.registerB = numerator >> combo
}

func cdv(computer *Computer, operand int) {
	combo := translateOperand(*computer, operand)
	numerator := computer.registerA
	// denominator := math.Pow(2, float64(combo))
	// denominator := 1 << combo

	// computer.registerC = numerator / int(denominator)
	computer.registerC = numerator >> combo
}

func bxl(computer *Computer, operand int) {
	computer.registerB = computer.registerB ^ operand
}

func bst(computer *Computer, operand int) {
	combo := translateOperand(*computer, operand)
	computer.registerB = combo & 7
}

func bxc(computer *Computer) {
	computer.registerB = computer.registerB ^ computer.registerC
}

func out(computer Computer, operand int) int {
	combo := translateOperand(computer, operand)
	return combo & 7
}

func handleInstruction(computer *Computer, instruction Instruction, output *[]int) int {
	switch instruction.opcode {
	case 0:
		adv(computer, instruction.operand)
	case 1:
		bxl(computer, instruction.operand)
	case 2:
		bst(computer, instruction.operand)
	case 3:
		if computer.registerA != 0 {
			// instrCtr = instruction.operand
			// continue
			return instruction.operand
		}
	case 4:
		bxc(computer)
	case 5:
		outval := out(*computer, instruction.operand)
		*output = append(*output, outval)
	case 6:
		bdv(computer, instruction.operand)
	case 7:
		cdv(computer, instruction.operand)
	}
	return -1
}

func RunInstructions(computer *Computer) []int {
	output := []int{}
	instrCtr := 0

	for instrCtr < len(computer.program)*2 {
		instruction := computer.program[instrCtr]

		if resp := handleInstruction(computer, instruction, &output); resp != -1 {
			instrCtr = resp
			continue
		}
		instrCtr += 2
	}
	return output
}

func runToNextOutputInstruction(computer *Computer, outputInstruction int) bool {
	output := []int{}
	instrCtr := 0

	for instrCtr < len(computer.program)*2 {
		instruction := computer.program[instrCtr]

		if resp := handleInstruction(computer, instruction, &output); resp != -1 {
			if len(output) >= 2 {
				break
			}
			instrCtr = resp
			continue
		}
		instrCtr += 2
	}

	if len(output) == 2 && output[0] == computer.program[outputInstruction].opcode &&
		output[1] == computer.program[outputInstruction].operand {
    // fmt.Println("passed instruction", outputInstruction)
		return true
	}

	return false
}

func FindProgramCopy(computer *Computer) int {
	regACtr := 1
	regB := computer.registerB
	regC := computer.registerC
	allOutputInstructionsValid := false
	for {
		computer.registerA = regACtr
		computer.registerB = regB
		computer.registerC = regC
    // fmt.Println("trying computer", computer)

		for i := 0; i < len(computer.program) * 2; i += 2 {
			if validInstruction := runToNextOutputInstruction(computer, i); !validInstruction {
				allOutputInstructionsValid = false
				break
			}
			allOutputInstructionsValid = true
		}
		if allOutputInstructionsValid {
			break
		}
		regACtr++
	}

  return regACtr
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
  regB := computer.registerB
  regC := computer.registerC
	output := RunInstructions(&computer)
	for _, entry := range output {
		fmt.Printf("%d,", entry)
	}

  computer.registerB = regB
  computer.registerC = regC
  copyComp := FindProgramCopy(&computer)
  fmt.Println(copyComp)
}
