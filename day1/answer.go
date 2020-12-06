package main

import (
	"fmt"

	"github.com/WulffTL/advent-of-code-2020/utils"
)

func main() {
	inputPath := "day1/puzzle-input.txt"
	input := utils.GetArrayFromFile(inputPath)
	sumValue := int64(2020)

	a, b, c := GetTripletOfSums(input, sumValue)
	fmt.Print(a * b * c)
}

func GetPairOfSums(entries []int64, expectedSum int64) (int64, int64) {
	var possiblePairs = make(map[int64]int64)

	for i := 0; i < len(entries); i++ {
		currentValue := entries[i]

		pair, ok := possiblePairs[currentValue]
		if ok {
			return currentValue, pair
		}

		possiblePairs[expectedSum-currentValue] = currentValue
	}

	return -1, -1
}

func GetTripletOfSums(entries []int64, expectedSum int64) (int64, int64, int64) {
	for i := 0; i < len(entries); i++ {
		currentValue := entries[i]
		remainderNeeded := expectedSum - currentValue

		possiblePair1, possiblePair2 := GetPairOfSums(entries, remainderNeeded)

		if possiblePair1 != currentValue && possiblePair2 != currentValue && possiblePair1 != -1 {
			return currentValue, possiblePair1, possiblePair2
		}
	}

	return -1, -1, -1
}
