package main

import (
	"fmt"
	// "math"
	"strconv"
	"strings"

	aoc "aoc-24/lib"
)

type Instruction struct {
	opcode  uint8
	operand uint8
}

type Computer struct {
	registerA int
	registerB uint8
	registerC uint8

	program map[int8]Instruction
}

func translateOperand(computer Computer, operand uint8) uint8 {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return uint8(computer.registerA)
	case 5:
		return computer.registerB
	case 6:
		return computer.registerC
	case 7:
		return 99
	}
	return 99
}

func adv(computer *Computer, operand uint8) {
	combo := translateOperand(*computer, operand)
	numerator := computer.registerA
	// denominator := math.Pow(2, float64(combo))
	// denominator := 1 << combo

	// computer.registerA = numerator / int(denominator)
	computer.registerA = numerator >> combo
}

func bdv(computer *Computer, operand uint8) {
	combo := translateOperand(*computer, operand)
	numerator := uint8(computer.registerA)
	// denominator := math.Pow(2, float64(combo))
	// denominator := 1 << combo

	// computer.registerB = numerator / int(denominator)
	computer.registerB = numerator >> combo
}

func cdv(computer *Computer, operand uint8) {
	combo := translateOperand(*computer, operand)
	numerator := uint8(computer.registerA)
	// denominator := math.Pow(2, float64(combo))
	// denominator := 1 << combo

	// computer.registerC = numerator / int(denominator)
	computer.registerC = numerator >> combo
}

func bxl(computer *Computer, operand uint8) {
	computer.registerB = computer.registerB ^ operand
}

func bst(computer *Computer, operand uint8) {
	combo := translateOperand(*computer, operand)
	computer.registerB = combo & 0b111
}

func bxc(computer *Computer) {
	computer.registerB = computer.registerB ^ computer.registerC
}

func out(computer Computer, operand uint8) uint8 {
	combo := translateOperand(computer, operand)
	return combo & 0b111
}

func handleInstruction(computer *Computer, instruction Instruction, output *[]uint8) int8 {
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
			return int8(instruction.operand)
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

func RunInstructions(computer *Computer) []uint8 {
	output := []uint8{}
	progLen := int8(len(computer.program) * 2)
	var instrCtr int8

	for instrCtr < progLen {
		instruction := computer.program[instrCtr]

		if resp := handleInstruction(computer, instruction, &output); resp != -1 {
			instrCtr = resp
			continue
		}
		instrCtr += 2
	}
	return output
}

func runToNextOutputInstruction(computer *Computer, outputInstruction int8) bool {
	output := []uint8{}
	progLen := int8(len(computer.program) * 2)
	var instrCtr int8
	// fmt.Println("trying computer", computer)

	for instrCtr < progLen {
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

	// fmt.Println("trying output", output)

	if len(output) == 2 && output[0] == computer.program[outputInstruction].opcode &&
		output[1] == computer.program[outputInstruction].operand {
		return true
	}

	return false
}

func FindProgramCopy(computer *Computer) int {
	// regACtr := 90065166
	regACtr := 1
	regB := computer.registerB
	regC := computer.registerC
	allOutputInstructionsValid := false
	progLen := int8(len(computer.program) * 2)
	fmt.Println(progLen)
	for {
		computer.registerA = regACtr
		computer.registerB = regB
		computer.registerC = regC

		for i := int8(0); i < progLen; i += 2 {
			if validInstruction := runToNextOutputInstruction(computer, i); !validInstruction {
				allOutputInstructionsValid = false
				break
			}
			fmt.Println("passed instruction", i, computer, regACtr)
			allOutputInstructionsValid = true
		}
		if allOutputInstructionsValid {
			break
		}
		regACtr++
	}

	return regACtr
}

// func findNextValidInstruction(computer *Computer, outputInstruction int8, seed int) bool {
// 	if runToNextOutputInstruction(computer, outputInstruction) && computer.registerA == seed {
// 		return true
// 	}
// 	return false
// }

func stepThroughInstructions(
	computer *Computer,
	currentInstruction int8,
	seed, counter, ctrInc int,
) {
	// foundValidInstruction := false
	// noMoreValidInstructions := false
	// regA := computer.registerA
	// regB := computer.registerB
	// regC := computer.registerC

	if currentInstruction < 0 {
		return
	}
	// for i := currentInstruction; i >= 0; i -= 2 {
	for {
		computer.registerA = counter
		computer.registerB = 0
		computer.registerC = 0
		// computer.registerB = regB
		// computer.registerC = regC
		// fmt.Println("trying", computer)
		if runToNextOutputInstruction(computer, currentInstruction) {
			// fmt.Println(
			// 	"found possible instruction for",
			// 	currentInstruction,
			// 	"seed",
			// 	seed,
			// 	computer,
			// )
			if computer.registerA == seed {
				fmt.Println(
					"found valid instruction for",
					currentInstruction,
					"counter",
					counter,
					computer,
				)
				fmt.Printf("%b\n", int64(counter))
				if currentInstruction == 0 {
					fmt.Println("found valid regA value", counter)
					computer.registerA = counter
					return
				}
				stepThroughInstructions(
					computer,
					currentInstruction-2,
					counter,
					counter<<6,
					1,
				)
			} else if computer.registerA > seed {
				break
			}
		}
		counter += ctrInc
	}

	// }
}

func FindProgramCopyReverse(computer *Computer) int {
	progLen := int8(len(computer.program) * 2)

	// for regB := uint8(0); regB < 8; regB++ {
	// 	for regC := uint8(0); regC < 8; regC++ {
	//      computer.registerB = regB
	//      computer.registerC = regC
	stepThroughInstructions(computer, progLen-2, 0, 1, 1)
	// 	}
	// }
	// fmt.Println("")
	// fmt.Println(computer)

	// regACtr := 1
	//  tmpCtr := regACtr
	// regB := computer.registerB
	// regC := computer.registerC
	// // allOutputInstructionsValid := false
	// outputInstrValid := false
	// progLen := int8(len(computer.program) * 2)
	// for {
	// 	endVal := 0
	// 	tmpCtr = regACtr
	//    fmt.Println("starting regACtr at", regACtr)
	// 	for i := progLen - 2; i >= 0; i -= 2 {
	// 		outputInstrValid = false
	// 		for {
	// 			computer.registerA = tmpCtr
	// 			computer.registerB = regB
	// 			computer.registerC = regC
	// 			if runToNextOutputInstruction(computer, i) {
	// 				if computer.registerA == endVal {
	//            fmt.Println("passed instruction", i, computer, tmpCtr)
	// 					outputInstrValid = true
	// 					break
	// 				} else if computer.registerA > endVal {
	// 					fmt.Println("instruction", i, "broke because regA > endVal", computer.registerA, endVal)
	//            regACtr = endVal
	// 					break
	// 				}
	// 			}
	// 			tmpCtr += 1
	// 		}
	// 		if outputInstrValid {
	// 			endVal = tmpCtr
	// 			continue
	// 		}
	// 		break
	// 	}
	//
	// 	if outputInstrValid {
	// 		break
	// 	}
	// 	regACtr++
	//
	// }
	// return regACtr
	return computer.registerA
}

func initComputer(data []string) Computer {
	var regA, regB, regC int
	var prog string
	instrMap := map[int8]Instruction{}

	fmt.Sscanf(data[0], "Register A: %d", &regA)
	fmt.Sscanf(data[1], "Register B: %d", &regB)
	fmt.Sscanf(data[2], "Register C: %d", &regC)
	fmt.Sscanf(data[4], "Program: %s", &prog)

	instrStr := strings.Split(prog, ",")
	instrLen := int8(len(instrStr))

	for i := int8(0); i < instrLen; i += 2 {
		opcode, _ := strconv.Atoi(instrStr[i])
		operand, _ := strconv.Atoi(instrStr[i+1])
		instrMap[i] = Instruction{opcode: uint8(opcode), operand: uint8(operand)}
	}

	return Computer{
		registerA: regA,
		registerB: uint8(regB),
		registerC: uint8(regC),
		program:   instrMap,
	}
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
	fmt.Println("")
	fmt.Println(computer)

	computer.registerB = regB
	computer.registerC = regC
	copyComp := FindProgramCopyReverse(&computer)
	fmt.Println(copyComp)
}
