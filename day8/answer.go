package main

import (
	"fmt"
	"github.com/WulffTL/advent-of-code-2020/utils"
	"strconv"
	"strings"
)

// Instruction ...
type Instruction struct {
	operation string
	value     int
}

func main() {
	inputPath := "day8/puzzle-input.txt"
	input := utils.GetArrayOfStringsFromFile(inputPath)

	answer := GetAccumulatorAfterFix(input)

	fmt.Print(answer)
}

// GetAccumulatorBeforeInfiniteLoop ...
func GetAccumulatorBeforeInfiniteLoop(input []string) int {
	instructions := map[int]bool{}
	currentIndex := 0
	accumulator := 0

	for {
		if currentIndex >= len(input) {
			return accumulator
		}
		_, ok := instructions[currentIndex]
		if ok {
			return accumulator
		} else {
			instructions[currentIndex] = true
		}

		currentInstruction := getInstruction(input[currentIndex])
		switch currentInstruction.operation {
		case "nop":
			currentIndex++
		case "acc":
			accumulator += currentInstruction.value
			currentIndex++
		case "jmp":
			currentIndex += currentInstruction.value
		}
	}
}

// GetAccumulatorAfterFix ...
func GetAccumulatorAfterFix(input []string) int {
	i := 0
	for {
		testInput, newIndex := swapNopJmp(input, i)
		if !isInfiniteLooped(testInput) {
			return GetAccumulatorBeforeInfiniteLoop(testInput)
		}
		i = newIndex
	}
}

func swapNopJmp(input []string, startIndex int) ([]string, int) {
	testInput := make([]string, len(input))
	copy(testInput, input)

	for i := startIndex; i < len(testInput); i++ {
		instruction := getInstruction(testInput[i])
		if instruction.operation == "jmp" {
			testInput[i] = strings.Replace(testInput[i], "jmp", "nop", 1)
			return testInput, i + 1
		}
		if instruction.operation == "nop" {
			testInput[i] = strings.Replace(testInput[i], "nop", "jmp", 1)
			return testInput, i + 1
		}
	}
	return testInput, 0
}

func isInfiniteLooped(input []string) bool {
	instructions := map[int]bool{}
	currentIndex := 0
	accumulator := 0

	for {
		if currentIndex >= len(input) {
			return false
		}
		_, ok := instructions[currentIndex]
		if ok {
			return true
		} else {
			instructions[currentIndex] = true
		}

		currentInstruction := getInstruction(input[currentIndex])
		switch currentInstruction.operation {
		case "nop":
			currentIndex++
		case "acc":
			accumulator += currentInstruction.value
			currentIndex++
		case "jmp":
			currentIndex += currentInstruction.value
		}
	}
}

func getInstruction(line string) Instruction {
	instruction := strings.Split(line, " ")
	value, _ := strconv.ParseInt(instruction[1], 10, 64)

	return Instruction{
		operation: instruction[0],
		value:     int(value),
	}
}
