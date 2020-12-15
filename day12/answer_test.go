package main

import (
	"github.com/WulffTL/advent-of-code-2020/utils"
	"testing"
)

// TestGetManhattanDistance ...
func TestGetManhattanDistance(t *testing.T) {
	expected := 25

	path := "day12/test-cases/given-sample.txt"
	testCase := utils.GetArrayOfStringsFromFile(path)

	result := GetManhattanDistance(testCase)

	if result != expected {
		t.Errorf("unexpected result, got %v but expected %v", result, expected)
	}
}

func TestRotate(t *testing.T) {
	result := rotate("L", 90, "W")
	expected := "S"

	if result != expected {
		t.Errorf("unexpected result, got %v but expected %v", result, expected)
	}
}

// TestGetWaypointManhattanDistance ...
func TestGetWaypointManhattanDistance(t *testing.T) {
	expected := 286

	path := "day12/test-cases/given-sample.txt"
	testCase := utils.GetArrayOfStringsFromFile(path)

	result := GetWaypointManhattanDistance(testCase)

	if result != expected {
		t.Errorf("unexpected result, got %v but expected %v", result, expected)
	}
}
