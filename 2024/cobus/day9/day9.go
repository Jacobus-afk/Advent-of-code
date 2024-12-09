package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findSameValLength(pos, val, counter int, blocksSlice []int) int {
	tally := 0
	for {
		pos += counter
		tally++
		if blocksSlice[pos] != val {
			break
		}
	}
	return tally
}

func freeupSpacePerFile(blocksSlice []int) {
	backwards := len(blocksSlice) - 1
	fileID := blocksSlice[backwards]
	// forwards := 0

	for {
		if file := blocksSlice[backwards]; file == fileID {
			fileID--
			// fmt.Println(backwards)
			if file == 1 {
				break
			}
			fileSize := findSameValLength(backwards, file, -1, blocksSlice)
			fileEnd := backwards - fileSize + 1
			// fmt.Println(fileSize)

			for forwards := 0; forwards < fileEnd; forwards++ {
				if blocksSlice[forwards] != -1 {
					continue
				}

				freeSpace := findSameValLength(forwards, -1, 1, blocksSlice)
				if fileSize <= freeSpace {
					for i := 0; i < fileSize; i++ {
						blocksSlice[fileEnd+i] = -1
						blocksSlice[forwards+i] = file
					}
					break
				}

			}

		}

		backwards--
	}
	// fmt.Println(blocksSlice)
}

// func freeupSpacePerFile(blocksSlice []int) {
// 	back := len(blocksSlice) - 1
// 	front := 0
// 	// backFile := blocksSlice[back]
//
// 	for {
// 		for blocksSlice[front] != -1 {
// 			front++
// 		}
// 		if front >= back {
// 			break
// 		}
// 		freeSpace := findSameValLength(front, -1, 1, blocksSlice)
//
//     // fmt.Println(front, blocksSlice[front], freeSpace)
//
//     if file := blocksSlice[back]; file != -1 {
//       fileSpace := findSameValLength(back, file, -1, blocksSlice)
//       backLimit := back - fileSpace
//
//       fmt.Println(fileSpace, freeSpace, back, file, backLimit)
//
//       if fileSpace <= freeSpace {
//         frontLimit := front + fileSpace
//         for front < frontLimit {
//           // fmt.Println(block, front, frontLimit)
//
//           blocksSlice[front] = file
//           front++
//         }
//
//         for back > backLimit {
//           blocksSlice[back] = -1
//           back--
//         }
//       } else {
//         back = backLimit
//       }
//
//       // back -= backFileSpace
//     }
// 	}
// }

func freeUpSpaceFragmented(blocksSlice []int) {
	back := len(blocksSlice) - 1
	front := 0
	for {
		for blocksSlice[front] != -1 {
			front++
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
				tmpSlice[i] = index / 2
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
	// fmt.Println(blocksSlice)
	// freeUpSpaceFragmented(blocksSlice)
	freeupSpacePerFile(blocksSlice)

	// fmt.Println(blocksSlice)
	for pos, val := range blocksSlice {
		if val == -1 {
			continue
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
