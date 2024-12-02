package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type NumList struct {
	leftList  []int
	rightList []int
}

func (n *NumList) popMin() int {
  leftMin := slices.Min(n.leftList)
  rightMin := slices.Min(n.rightList)

	leftMinIndex := getMinIdenx(leftMin, n.leftList)
	rightMinIndex := getMinIdenx(rightMin, n.rightList)

	n.leftList = append(n.leftList[:leftMinIndex], n.leftList[leftMinIndex+1:]...)
	n.rightList = append(n.rightList[:rightMinIndex], n.rightList[rightMinIndex+1:]...)

	return distance(leftMin, rightMin)
}

func (n *NumList) DistanceSum() int {
	distanceSum := 0

	for ok := true; ok; ok = (len(n.leftList) > 0) {
		distanceSum += n.popMin()
	}

	return distanceSum
}

func (n *NumList) SimilaritySum() int {
  similaritySum := 0

  for _, v := range n.leftList {
    similaritySum += getSimilarity(v, n.rightList)
  }
  return similaritySum
}

func getMinIdenx(minVal int, arr []int) int {
	for i, v := range arr {
		if v == minVal {
			return i
		}
	}
	return -1
}

func distance(leftMin, rightMin int) int {
	if leftMin < rightMin {
		return rightMin - leftMin
	}

	return leftMin - rightMin
}

func getSimilarity(leftVal int, rightList []int) int {
  multiplier := 0
  for _, v := range rightList {
    if v == leftVal {
      multiplier += 1
    }
  }
  return leftVal * multiplier
}

func main() {
	numList := NumList{}

	file, _ := os.Open("./data")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		leftVal, _ := strconv.Atoi(parts[0])
		rightVal, _ := strconv.Atoi(parts[1])

		numList.leftList = append(numList.leftList, leftVal)
		numList.rightList = append(numList.rightList, rightVal)
	}
  fmt.Printf("SimilaritySum: %d\n", numList.SimilaritySum())
  fmt.Printf("DistanceSum: %d\n", numList.DistanceSum())
}
