package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func freeUpSpace(blocksSlice []int) {
	back := len(blocksSlice) - 1
  front := 0
	for {
    for {
      if blocksSlice[front] == -1 {
        break
      }
      front ++
    }
		if front >= back {
			break
		}

		if block := blocksSlice[back]; block != -1 {
      blocksSlice[front] = block
      blocksSlice[back] = -1
		}
		back--
  }
}

// func freeUpSpace(blockMap string) string {
// 	back := len(blockMap) - 1
// 	for {
// 		front := strings.Index(blockMap, ".")
// 		if front < 0 || front >= back {
// 			break
// 		}
// 		// if blockMap[front] != '.' {
// 		//   front ++
// 		// }
//
// 		if block := blockMap[back]; block != '.' {
// 			padding := strings.Repeat(".", len(blockMap)-back)
// 			blockMap = blockMap[:front] + string(block) + blockMap[front+1:back] + padding
// 		}
//
// 		back--
//
// 	}
// 	return blockMap
// }

func translateDiskMap(diskMap string) (string, []int) {
	blocks := ""
  blocksSlice := []int{}

	for index, size := range diskMap {
		amount := int(size - '0')
    tmpSlice := make([]int, amount)
		if index%2 == 0 {
			blocks += strings.Repeat(strconv.Itoa(index/2), amount)
      for i := range tmpSlice {
        tmpSlice[i] = index/2
      }
		} else {
			blocks += strings.Repeat(".", amount)
      for i := range tmpSlice {
        tmpSlice[i] = -1
      }
		}
    blocksSlice = append(blocksSlice, tmpSlice...)
	}

	return blocks, blocksSlice
}

func CalcCheckSum(diskMap string) int {
	checkSum := 0

	_, blocksSlice := translateDiskMap(diskMap)
	freeUpSpace(blocksSlice)

  fmt.Println(blocksSlice)
  for pos, val := range blocksSlice{
    if val == -1 {
      break
    }
    tally := pos * val
    checkSum += tally
  }
	// for pos, char := range cleanedUpMap {
	// 	if char == '.' {
	// 		break
	// 	}
	// 	tally := pos * int(char-'0')
	// 	checkSum += tally
	// }

	return checkSum
}

func main() {
	diskMap := ""

	file, _ := os.Open("./data")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		diskMap = scanner.Text()
	}

	checkSum := CalcCheckSum(diskMap)

	fmt.Println(checkSum)
}
