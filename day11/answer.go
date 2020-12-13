package main

import (
	"fmt"
	"github.com/WulffTL/advent-of-code-2020/utils"
)

func main() {
	inputPath := "day11/puzzle-input.txt"
	input := utils.GetArrayOfStringsFromFile(inputPath)

	answer := GetOccupiedSeatsBySight(input)

	fmt.Print(answer)
}

// GetOccupiedSeats ...
func GetOccupiedSeats(input []string) int {
	seatGrid := getSeatGrid(input)
	changedSeats, isChanged := getChangedSeats(seatGrid)

	for isChanged {
		changedSeats, isChanged = getChangedSeats(changedSeats)
	}

	return getCountOfOccupiedSeats(changedSeats)
}

// GetOccupiedSeatsBySight ...
func GetOccupiedSeatsBySight(input []string) int {
	seatGrid := getSeatGrid(input)
	changedSeats, isChanged := getChangedSeatsBySight(seatGrid)

	for isChanged {
		changedSeats, isChanged = getChangedSeatsBySight(changedSeats)
	}

	return getCountOfOccupiedSeats(changedSeats)
}

func getCountOfOccupiedSeats(seatGrid [][]string) int {
	count := 0
	for i := 0; i < len(seatGrid); i++ {
		for j := 0; j < len(seatGrid[i]); j++ {
			if seatGrid[i][j] == "#" {
				count++
			}
		}
	}
	return count
}

func getChangedSeats(seatGrid [][]string) ([][]string, bool) {
	newSeatGrid := make([][]string, len(seatGrid))
	for i := 0; i < len(seatGrid); i++ {
		newSeatGrid[i] = make([]string, len(seatGrid[i]))
	}
	isChanged := false

	for y := 0; y < len(seatGrid); y++ {
		for x := 0; x < len(seatGrid[y]); x++ {
			currentValue := seatGrid[y][x]
			newValue := getNewStateForSeat(x, y, seatGrid)
			newSeatGrid[y][x] = newValue
			if currentValue != newValue {
				isChanged = true
			}
		}
	}
	return newSeatGrid, isChanged
}

func getChangedSeatsBySight(seatGrid [][]string) ([][]string, bool) {
	newSeatGrid := make([][]string, len(seatGrid))
	for i := 0; i < len(seatGrid); i++ {
		newSeatGrid[i] = make([]string, len(seatGrid[i]))
	}
	isChanged := false

	for y := 0; y < len(seatGrid); y++ {
		for x := 0; x < len(seatGrid[y]); x++ {
			currentValue := seatGrid[y][x]
			newValue := getNewStateForSeatBySight(x, y, seatGrid)
			newSeatGrid[y][x] = newValue
			if currentValue != newValue {
				isChanged = true
			}
		}
	}
	return newSeatGrid, isChanged
}

func getNewStateForSeatBySight(seatX int, seatY int, seats [][]string) string {
	row := seats[seatY]
	seat := string(row[seatX])
	isFloor := seat == "."

	if isFloor {
		return "."
	}

	isOccupied := seat == "#"
	if isOccupied {
		return getNewStateForOccupiedSeatBySight(seatX, seatY, seats)
	}

	return getNewStateForUnoccupiedSeatBySight(seatX, seatY, seats)
}

func getNewStateForOccupiedSeatBySight(seatX int, seatY int, seats [][]string) string {
	occupiedCount := 0
	for y := -1; y <= 1; y++ {
		if seatY+y < 0 || seatY+y >= len(seats) {
			continue
		}
		for x := -1; x <= 1; x++ {
			if seatX+x < 0 || seatX+x >= len(seats[seatY+y]) {
				continue
			}
			if x == 0 && y == 0 {
				continue
			}
			newSeat := seats[seatY+y][seatX+x]

			newY := y
			newX := x
			for newSeat == "." {
				newY += y
				newX += x
				if seatY+newY < 0 || seatY+newY >= len(seats) {
					newSeat = "!"
					continue
				}
				if seatX+newX < 0 || seatX+newX >= len(seats[0]) {
					newSeat = "!"
					continue
				}
				newSeat = seats[seatY+newY][seatX+newX]
			}

			if newSeat == "#" {
				occupiedCount++
			}
			if occupiedCount == 5 {
				return "L"
			}
		}
	}

	return "#"
}

func getNewStateForUnoccupiedSeatBySight(seatX int, seatY int, seats [][]string) string {
	for y := -1; y <= 1; y++ {
		if seatY+y < 0 || seatY+y >= len(seats) {
			continue
		}
		for x := -1; x <= 1; x++ {
			if seatX+x < 0 || seatX+x >= len(seats[seatY+y]) {
				continue
			}
			newSeat := seats[seatY+y][seatX+x]
			newY := y
			newX := x
			for newSeat == "." {
				newY += y
				newX += x
				if seatY+newY < 0 || seatY+newY >= len(seats) {
					newSeat = "!"
					continue
				}
				if seatX+newX < 0 || seatX+newX >= len(seats[0]) {
					newSeat = "!"
					continue
				}
				newSeat = seats[seatY+newY][seatX+newX]
			}
			if newSeat == "#" {
				return "L"
			}
		}
	}

	return "#"
}

func getNewStateForSeat(seatX int, seatY int, seats [][]string) string {
	row := seats[seatY]
	seat := string(row[seatX])
	isFloor := seat == "."

	if isFloor {
		return "."
	}

	isOccupied := seat == "#"
	if isOccupied {
		return getNewStateForOccupiedSeat(seatX, seatY, seats)
	}

	return getNewStateForUnoccupiedSeat(seatX, seatY, seats)
}

func getNewStateForOccupiedSeat(seatX int, seatY int, seats [][]string) string {
	occupiedCount := 0
	for y := -1; y <= 1; y++ {
		if seatY+y < 0 || seatY+y >= len(seats) {
			continue
		}
		for x := -1; x <= 1; x++ {
			if seatX+x < 0 || seatX+x >= len(seats[seatY+y]) {
				continue
			}
			if x == 0 && y == 0 {
				continue
			}
			newSeat := seats[seatY+y][seatX+x]
			if newSeat == "#" {
				occupiedCount++
			}
			if occupiedCount == 4 {
				return "L"
			}
		}
	}

	return "#"
}

func getNewStateForUnoccupiedSeat(seatX int, seatY int, seats [][]string) string {
	for y := -1; y <= 1; y++ {
		if seatY+y < 0 || seatY+y >= len(seats) {
			continue
		}
		for x := -1; x <= 1; x++ {
			if seatX+x < 0 || seatX+x >= len(seats[seatY+y]) {
				continue
			}
			newSeat := seats[seatY+y][seatX+x]
			if newSeat == "#" {
				return "L"
			}
		}
	}

	return "#"
}

func getSeatGrid(input []string) [][]string {
	seatGrid := make([][]string, len(input))
	for i := 0; i < len(seatGrid); i++ {
		seatGrid[i] = make([]string, len(input[i]))
	}
	for y := 0; y < len(input); y++ {
		row := input[y]
		for x := 0; x < len(row); x++ {
			seatGrid[y][x] = string(row[x])
		}
	}

	return seatGrid
}
