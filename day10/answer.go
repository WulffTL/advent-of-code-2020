package main

import (
	"fmt"
	"github.com/WulffTL/advent-of-code-2020/utils"
	"sort"
	"strconv"
)

var cache = make(map[int]int)

func main() {
	inputPath := "day10/puzzle-input.txt"
	input := utils.GetArrayOfStringsFromFile(inputPath)

	answer := GetUniqueArrangementsCount(input)

	fmt.Print(answer)
}

// GetVoltageProduct ...
func GetVoltageProduct(input []string) int {
	voltages := sortVoltages(input)
	oneOff := 0
	threeOff := 0

	if voltages[0] == 1 {
		oneOff++
	} else if voltages[0] == 3 {
		threeOff++
	}

	for i := 0; i < len(voltages)-1; i++ {
		diff := voltages[i+1] - voltages[i]
		switch diff {
		case 1:
			oneOff++
		case 3:
			threeOff++
		}
	}
	// the last jump to your adapter is always 3
	threeOff++

	return oneOff * threeOff
}

func sortVoltages(input []string) []int {
	voltages := make([]int, len(input))
	for i := 0; i < len(input); i++ {
		voltage, _ := strconv.ParseInt(input[i], 10, 64)
		voltages[i] = int(voltage)
	}

	sort.Ints(voltages)

	return voltages
}

// GetUniqueArrangementsCount ...
func GetUniqueArrangementsCount(input []string) int {
	possibleAdapters := getPossibleAdapters(input)

	sortedVoltages := sortVoltages(input)
	var possibleFirstAdapters = [3]int{}
	for i := 0; i < 3; i++ {
		if sortedVoltages[i] < 4 {
			possibleFirstAdapters[i] = sortedVoltages[i]
		}
	}

	sum := 0
	for i := 0; i < len(possibleFirstAdapters); i++ {
		firstAdapter := possibleFirstAdapters[i]
		if firstAdapter == 0 {
			continue
		}
		sum += getAdapterCombinationsFromPoint(firstAdapter, possibleAdapters)
	}

	return sum
}

func getAdapterCombinationsFromPoint(point int, adapters map[int][]int) int {
	count, ok := cache[point]
	if ok {
		return count
	}
	adapterChildren := adapters[point]

	if len(adapterChildren) == 0 {
		return 1
	}

	for i := 0; i < len(adapterChildren); i++ {
		count += getAdapterCombinationsFromPoint(adapterChildren[i], adapters)
	}

	cache[point] = count

	return count
}

func getPossibleAdapters(input []string) map[int][]int {
	possibleAdapters := map[int][]int{}

	voltages := sortVoltages(input)

	for i := 0; i < len(voltages); i++ {
		lowVoltage := voltages[i]
		j := 1
		if i+j >= len(voltages) {
			continue
		}
		highVoltage := voltages[i+j]
		diff := highVoltage - lowVoltage
		for diff < 4 {
			newAdapters := append(possibleAdapters[lowVoltage], highVoltage)
			possibleAdapters[lowVoltage] = newAdapters
			j++
			if i+j < len(voltages) {
				lowVoltage = voltages[i]
				highVoltage = voltages[i+j]
				diff = highVoltage - lowVoltage
			} else {
				diff = 4
			}
		}
	}

	return possibleAdapters
}
