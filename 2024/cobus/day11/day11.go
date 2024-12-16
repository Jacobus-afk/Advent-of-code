package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

var mapMutex sync.RWMutex

func multiplyby2024(num uint64) uint64 {
  return (num << 11) - (num << 4) - (num << 3)
}

func blinkChecks(
	stone string,
	multMemos map[string]string,
	splitMemos map[string][2]string,
) (string, string) {
	if stone == "0" {
		return "1", ""
	}
	if stoneLen := uint64(len(stone)); stoneLen%2 == 0 {
    mapMutex.RLock()
		parts, ok := splitMemos[stone]
    mapMutex.RUnlock()
		if !ok {
			halfStone := stoneLen / 2
			part1 := stone[:halfStone]
			part2 := stone[halfStone:]
			number, _ := strconv.ParseUint(part2, 10, 64)
			part2 = strconv.FormatUint(number, 10)
			parts = [2]string{part1, part2}
      mapMutex.Lock()
			splitMemos[stone] = parts
      mapMutex.Unlock()
		}

		return parts[0], parts[1]
		// return part1, part2
	}

  mapMutex.RLock()
	answer, ok := multMemos[stone]
  mapMutex.RUnlock()
	if !ok {
		number, _ := strconv.ParseUint(stone, 10, 64)
		multiplied := multiplyby2024(number)
		answer = strconv.FormatUint(multiplied, 10)
    mapMutex.Lock()
		multMemos[stone] = answer
    mapMutex.Unlock()
	}
	return answer, ""
}

func Blinks(blinkCount int, stones []string) []string {
	multMemos := map[string]string{}
	splitMemos := map[string][2]string{}

	stoneReg := []string{}

	// resultsChannel := make(chan []string, len(stones))
  resultsSlice := make([][]string, len(stones))

	var wg sync.WaitGroup
	// wg.Add(len(stones))
	for pos := 0; pos < len(stones); pos++ {
		tempReg := []string{stones[pos]}
    wg.Add(1)
		go func(tmpReg []string, id int) {
      defer wg.Done()
			// fmt.Printf("tmpReg %s\n", tmpReg)
			for count := 0; count < blinkCount; count++ {
				fmt.Print(id)
				for l := 0; l < len(tmpReg); l++ {
					part1, part2 := blinkChecks(tmpReg[l], multMemos, splitMemos)
					// fmt.Println(part1, part2)
					tmpReg[l] = part1
					if part2 != "" {
						tmpReg = append(tmpReg[:l+1], tmpReg[l:]...)
						tmpReg[l+1] = part2
						l++
					}

				}
			}
			fmt.Printf("id %d done\n", id)
			// resultsChannel <- tmpReg
      resultsSlice[id] = tmpReg
		}(tempReg, pos)

		// fmt.Println(pos)
	}
	wg.Wait()
  // fmt.Println("got here")

	for i := 0; i < len(stones); i++ {
		// fmt.Printf("result %d\n", i)
		// r := <-resultsChannel
    r := resultsSlice[i]
		// fmt.Printf("result val: %s\n", r)
		stoneReg = append(stoneReg, r...)
	}
	return stoneReg
}

func main() {
	fmt.Println(runtime.NumCPU())
	stones := []string{}
	file, _ := os.Open("./data")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		stones = strings.Fields(line)
	}

	stoneReg := Blinks(75, stones)
	fmt.Println(len(stoneReg))
}
