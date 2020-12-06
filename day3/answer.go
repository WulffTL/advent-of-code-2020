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
	inputPath := "day3/puzzle-input.txt"
	input := utils.GetArrayOfStringsFromFile(inputPath)

	answer := GetTreesForMultipleAngles(input)
	fmt.Print(answer)
}

func GetTreesForAngle(trees []string, rise int, run int) int {
	x := 1
	y := 1
	treeCount := 0
	for {
		x += run
		y += rise
		if y > len(trees) {
			return treeCount
		}
		line := trees[y-1]
		if isTreeAtPosition(line, x) {
			treeCount++
		}
	}
}

func isTreeAtPosition(line string, position int) bool {
	newPosition := math.Mod(float64(position), float64(len(line)))
	if newPosition == 0 {
		newPosition = float64(len(line))
	}
	return string(line[int(newPosition-1)]) == "#"
}

func GetTreesForMultipleAngles(input []string) int {
	slopes := []slope{
		{
			rise: 1,
			run:  1,
		},
		{
			rise: 1,
			run:  3,
		},
		{
			rise: 1,
			run:  5,
		},
		{
			rise: 1,
			run:  7,
		},
		{
			rise: 2,
			run:  1,
		},
	}

	answer := 1
	for i := 0; i < len(slopes); i++ {
		rise := slopes[i].rise
		run := slopes[i].run
		treeCount := GetTreesForAngle(input, rise, run)
		answer *= treeCount
	}

	return answer
}
