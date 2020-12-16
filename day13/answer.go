package main

import (
	"fmt"
	"github.com/WulffTL/advent-of-code-2020/utils"
	"math"
	"strconv"
	"strings"
)

func main() {
	inputPath := "day13/puzzle-input.txt"
	input := utils.GetArrayOfStringsFromFile(inputPath)

	answer := GetEarliestConcurrentSchedule(input)

	fmt.Print(answer)
}

// GetProduct ...
func GetProduct(input []string) int {
	minTime, busID := getEarliestBus(input)
	return minTime * busID
}

// GetEarliestConcurrentSchedule ...
func GetEarliestConcurrentSchedule(input []string) int {
	offsets := getBusOffsets(input)
	busIDs := formatBusIDs(input)
	firstBusValue := busIDs[0]
	add := int(firstBusValue)
	start := int(firstBusValue)

	for i := 0; i < len(busIDs); i++ {
		currentBusIDs := busIDs[:i+1]
		for !isAligned(int(start), currentBusIDs, offsets) {
			start += add
		}

		for i := 0; i < len(currentBusIDs); i++ {
			add = lcm(add, currentBusIDs[i])
		}
	}
	return start
}

func isAligned(start int, busIDs []int, offsets map[int]int) bool {
	for i := 0; i < len(busIDs); i++ {
		busID := busIDs[i]
		offset := offsets[busID]
		if (start+offset)%busID != 0 {
			return false
		}
	}
	return true
}

func getEarliestBus(input []string) (int, int) {
	waitTimes := getBusTimes(input)

	minTime := math.MaxInt64
	fastestBus := 0
	for busID, time := range waitTimes {
		if time < minTime {
			minTime = time
			fastestBus = busID
		}
	}
	return minTime, fastestBus
}

func getBusTimes(input []string) map[int]int {
	arrivalTime, _ := strconv.ParseInt(input[0], 10, 64)
	busIDs := formatBusIDs(input)
	waitTimes := map[int]int{}

	for i := 0; i < len(busIDs); i++ {
		busID := busIDs[i]
		a := (int(arrivalTime) % busID) - busID
		waitTime := math.Abs(float64(a))
		waitTimes[busID] = int(waitTime)
	}

	return waitTimes
}

func formatBusIDs(input []string) (busIDs []int) {
	currentBusIDs := strings.Split(input[1], ",")
	for i := 0; i < len(currentBusIDs); i++ {
		ID := currentBusIDs[i]
		busID, err := strconv.ParseInt(ID, 10, 64)
		if err == nil {
			busIDs = append(busIDs, int(busID))
		}
	}
	return busIDs
}

func getBusOffsets(input []string) map[int]int {
	offsets := map[int]int{}

	list := strings.Split(input[1], ",")
	for i := 0; i < len(list); i++ {
		ID := string(list[i])
		busID, err := strconv.ParseInt(ID, 10, 64)
		if err == nil {
			offsets[int(busID)] = i
		}
	}
	return offsets
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}
