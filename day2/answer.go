package main

import (
	"fmt"
	"github.com/WulffTL/advent-of-code-2020/utils"
	"strconv"
	"strings"
)

func main() {
	inputPath := "day2/puzzle-input.txt"
	input := utils.GetArrayOfStringsFromFile(inputPath)

	answer := GetValidPasswords2(input)
	fmt.Print(answer)
}

// GetValidPasswords takes in an array of passwords and rules
// and returns how many of these passwords were valid
func GetValidPasswords(entries []string) int {
	validPasswordCount := 0
	for i := 0; i < len(entries); i++ {
		lowFreq, highFreq, letter, password := splitInput(entries[i])

		letterCount := int64(0)
		for j := 0; j < len(password); j++ {
			if string(password[j]) == letter {
				letterCount++
			}
		}

		if letterCount >= lowFreq && letterCount <= highFreq {
			validPasswordCount++
		}
	}

	return validPasswordCount
}

// GetValidPasswords2 takes in an array of passwords and rules
// and returns how many of these passwords were valid
func GetValidPasswords2(entries []string) int {
	validPasswordCount := 0
	for i := 0; i < len(entries); i++ {
		firstIndex, secondIndex, letter, password := splitInput(entries[i])

		letterCount := 0
		if string(password[firstIndex-1]) == letter {
			letterCount++
		}
		if string(password[secondIndex-1]) == letter {
			letterCount++
		}
		if letterCount == 1 {
			validPasswordCount++
		}
	}

	return validPasswordCount
}

func splitInput(input string) (int64, int64, string, string) {
	split := strings.Split(input, " ")

	freq := split[0]
	f := strings.Split(freq, "-")
	lowFreq, _ := strconv.ParseInt(f[0], 10, 64)
	highFreq, _ := strconv.ParseInt(f[1], 10, 64)

	letter := string(split[1][0])
	password := split[2]

	return lowFreq, highFreq, letter, password
}
