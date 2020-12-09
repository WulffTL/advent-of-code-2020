package main

import (
	"fmt"
	"github.com/WulffTL/advent-of-code-2020/utils"
	"strings"
)

type slope struct {
	rise int
	run  int
}

func main() {
	inputPath := "day6/puzzle-input.txt"
	input := utils.GetArrayOfStringsFromFile(inputPath)

	answer := GetTotalMutualAnswerCount(input)

	fmt.Print(answer)
}

func GetTotalMutualAnswerCount(input []string) int {
	groups := getGroups(input)
	count := 0

	for i := 0; i < len(groups); i++ {
		count += getMutualAnswerCount(groups[i])
	}

	return count
}

func GetTotalAnswerCount(input []string) int {
	groups := getGroups(input)
	count := 0

	for i := 0; i < len(groups); i++ {
		count += getAnswerCount(groups[i])
	}

	return count
}

func getAnswerCount(input []string) int {
	answers := map[string]bool{}
	for i := 0; i < len(input); i++ {
		currentAnswers := strings.Split(input[i], "")
		for j := 0; j < len(currentAnswers); j++ {
			answer := string(currentAnswers[j])
			answers[answer] = true
		}
	}

	return len(answers)
}

func getMutualAnswerCount(input []string) int {
	answers := map[string]int{}
	groupSize := len(input)

	for i := 0; i < len(input); i++ {
		currentAnswers := strings.Split(input[i], "")
		for j := 0; j < len(currentAnswers); j++ {
			answer := string(currentAnswers[j])
			_, ok := answers[answer]
			if !ok {
				answers[answer] = 1
			} else {
				answers[answer]++
			}
		}
	}

	answerCount := 0
	for _, element := range answers {
		if element == groupSize {
			answerCount++
		}
	}

	return answerCount
}

func getGroups(input []string) [][]string {
	var groups [][]string
	var currentGroup []string
	currentGroupIndex := 0

	for i := 0; i < len(input); i++ {
		if input[i] == "" {
			groups = append(groups, currentGroup)
			currentGroup = nil
			currentGroupIndex++
			continue
		}

		currentGroup = append(currentGroup, input[i])

		if i == len(input)-1 {
			groups = append(groups, currentGroup)
		}

	}

	return groups
}
