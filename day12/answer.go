package main

import (
	"fmt"
	"github.com/WulffTL/advent-of-code-2020/utils"
	"math"
	"strconv"
)

type Instruction struct {
	command string
	value   int
}

type Location struct {
	x int
	y int
}

func main() {
	inputPath := "day12/puzzle-input.txt"
	input := utils.GetArrayOfStringsFromFile(inputPath)

	answer := GetWaypointManhattanDistance(input)

	fmt.Print(answer)
}

// GetManhattanDistance ...
func GetManhattanDistance(input []string) int {
	instructions := getInstructionSet(input)

	xDistance := 0
	yDistance := 0
	currentDirection := "E"

	for i := 0; i < len(instructions); i++ {
		currentInstruction := instructions[i]

		command := currentInstruction.command

		if command == "F" {
			command = currentDirection
		}

		switch command {
		case "N":
			yDistance += currentInstruction.value
		case "S":
			yDistance -= currentInstruction.value
		case "E":
			xDistance += currentInstruction.value
		case "W":
			xDistance -= currentInstruction.value
		case "L":
			currentDirection = rotate(command, currentInstruction.value, currentDirection)
		case "R":
			currentDirection = rotate(command, currentInstruction.value, currentDirection)
		}
	}

	return int(math.Abs(float64(xDistance)) + math.Abs(float64(yDistance)))
}

// GetWaypointManhattanDistance ...
func GetWaypointManhattanDistance(input []string) int {
	instructions := getInstructionSet(input)

	xDistance := 0
	yDistance := 0
	waypointLocation := Location{
		x: 10,
		y: 1,
	}

	for i := 0; i < len(instructions); i++ {
		currentInstruction := instructions[i]

		command := currentInstruction.command

		switch command {
		case "N":
			waypointLocation.y += currentInstruction.value
		case "S":
			waypointLocation.y -= currentInstruction.value
		case "E":
			waypointLocation.x += currentInstruction.value
		case "W":
			waypointLocation.x -= currentInstruction.value
		case "L":
			waypointLocation = rotateWaypoint(command, currentInstruction.value, waypointLocation)
		case "R":
			waypointLocation = rotateWaypoint(command, currentInstruction.value, waypointLocation)
		case "F":
			xDistance += currentInstruction.value * waypointLocation.x
			yDistance += currentInstruction.value * waypointLocation.y
		}
	}

	fmt.Printf("x - %v, y - %v\n", xDistance, yDistance)
	return int(math.Abs(float64(xDistance)) + math.Abs(float64(yDistance)))
}

func rotateWaypoint(command string, value int, currentLocation Location) Location {
	if command == "L" {
		switch value {
		case 90:
			tempX := currentLocation.x
			currentLocation.x = -currentLocation.y
			currentLocation.y = tempX
		case 180:
			currentLocation.x *= -1
			currentLocation.y *= -1
		case 270:
			tempX := currentLocation.x
			currentLocation.x = currentLocation.y
			currentLocation.y = -tempX
		}
	}

	if command == "R" {
		switch value {
		case 90:
			tempX := currentLocation.x
			currentLocation.x = currentLocation.y
			currentLocation.y = -tempX
		case 180:
			currentLocation.x *= -1
			currentLocation.y *= -1
		case 270:
			tempX := currentLocation.x
			currentLocation.x = -currentLocation.y
			currentLocation.y = tempX
		}
	}

	return currentLocation
}

func rotate(command string, value int, currentDirection string) string {
	currentDegree := 0
	switch currentDirection {
	case "E":
		currentDegree = 0
	case "S":
		currentDegree = 90
	case "W":
		currentDegree = 180
	case "N":
		currentDegree = 270
	}

	if command == "L" {
		value *= -1
	}

	currentDegree += value
	currentDegree += 360
	currentDegree = currentDegree % 360

	finalDirection := "E"
	switch currentDegree {
	case 0:
		finalDirection = "E"
	case 90:
		finalDirection = "S"
	case 180:
		finalDirection = "W"
	case 270:
		finalDirection = "N"
	}

	return finalDirection
}

func getInstructionSet(input []string) []Instruction {
	instructions := make([]Instruction, len(input))
	for i := 0; i < len(input); i++ {
		instructions[i] = getInstruction(input[i])
	}
	return instructions
}

func getInstruction(raw string) Instruction {
	command := string(raw[0])
	value, _ := strconv.ParseInt(raw[1:], 10, 64)

	return Instruction{
		command: command,
		value:   int(value),
	}
}
