package main

import (
	"fmt"
	"github.com/WulffTL/advent-of-code-2020/utils"
	"math"
)

type slope struct {
	rise int
	run  int
}

func main() {
	inputPath := "day5/puzzle-input.txt"
	input := utils.GetArrayOfStringsFromFile(inputPath)

	IDs := getIDs(input)
	answer := GetSeat(IDs)

	fmt.Print(answer)
}

func GetSeat(IDs map[int64]bool) int64 {
	foundFront := false
	for i := int64(1); i < 1023; i++ {
		_, foundHigher := IDs[i+1]
		if foundFront && !foundHigher {
			return i + 1
		}
		if foundHigher {
			foundFront = true
		}
	}

	return -1
}

func GetHighestID(input []string) int64 {
	maxID := int64(0)

	for i := 0; i < len(input); i++ {
		ID := GetIDForBoardingPass(input[i])
		if ID > maxID {
			maxID = ID
		}
	}

	return maxID
}

func GetIDForBoardingPass(input string) int64 {
	row := GetRow(input)
	col := GetCol(input)
	ID := getSeatID(row, col)

	return ID
}

func GetRow(input string) int64 {
	rows := input[0:7]
	min := int64(0)
	max := int64(128)
	for i := 0; i < len(rows); i++ {
		size := max - min
		if string(rows[i]) == "F" {
			max = max - int64(math.Max(float64(size/2), 1))
		} else if string(rows[i]) == "B" {
			min = min + int64(math.Max(float64(size/2), 1))
		}
	}

	return max - 1
}

func GetCol(input string) int64 {
	rows := input[7:10]
	min := int64(0)
	max := int64(8)
	for i := 0; i < len(rows); i++ {
		size := max - min
		if string(rows[i]) == "L" {
			max = max - size/2
		} else if string(rows[i]) == "R" {
			min = min + size/2
		}
	}
	return max - 1
}

func getSeatID(row, col int64) int64 {
	return 8*row + col
}

func getIDs(input []string) map[int64]bool {
	IDs := map[int64]bool{}

	for i := 0; i < len(input); i++ {
		ID := GetIDForBoardingPass(input[i])
		IDs[ID] = true
	}

	return IDs
}
