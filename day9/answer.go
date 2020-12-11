package main

import (
	"fmt"
	"github.com/WulffTL/advent-of-code-2020/utils"
	"math"
	"strconv"
)

func main() {
	inputPath := "day9/puzzle-input.txt"
	input := utils.GetArrayOfStringsFromFile(inputPath)

	preambleLength := 25
	num := GetFirstViolatingNumber(input, preambleLength)
	answer := GetEncryptionWeakness(input, num)

	fmt.Print(answer)
}

func GetFirstViolatingNumber(input []string, preambleLength int) int64 {
	previousNumbers := make([]int64, preambleLength)

	for i := 0; i < preambleLength; i++ {
		previousNumbers[i], _ = strconv.ParseInt(input[i], 10, 64)
	}

	for i := preambleLength; i < len(input); i++ {
		num, _ := strconv.ParseInt(input[i], 10, 64)
		if isInvalidNumber(num, previousNumbers) {
			return num
		}

		index := i % preambleLength
		previousNumbers[index] = num
	}

	return -1
}

func GetEncryptionWeakness(input []string, num int64) int64 {
	start, end := getContiguousIndecies(input, num)
	min := int64(math.MaxInt64)
	max := int64(0)

	for i := start; i < end; i++ {
		num, _ := strconv.ParseInt(input[i], 10, 64)
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	return min + max
}

func getContiguousIndecies(input []string, num int64) (int, int) {
	for i := 0; i < len(input); i++ {
		sum := int64(0)
		length := 0
		for sum < num {
			currentNum, _ := strconv.ParseInt(input[i+length], 10, 64)
			sum += currentNum
			length++
		}

		if sum == num {
			return i, i + length
		}
	}
	return -1, -1
}

func isInvalidNumber(num int64, previousNumbers []int64) bool {
	possiblePairs := map[int64]bool{}

	for i := 0; i < len(previousNumbers); i++ {
		currentNum := previousNumbers[i]
		_, ok := possiblePairs[currentNum]
		if ok {
			return false
		}
		possiblePairs[num-currentNum] = true
	}

	return true
}
