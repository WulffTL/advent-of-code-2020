package main

import (
	"github.com/WulffTL/advent-of-code-2020/utils"
	"testing"
)

// TestGetOccupiedSeats ...
func TestGetOccupiedSeats(t *testing.T) {
	expected := 37

	path := "day11/test-cases/given-sample.txt"
	testCase := utils.GetArrayOfStringsFromFile(path)

	result := GetOccupiedSeats(testCase)

	if result != expected {
		t.Errorf("unexpected result, got %v but expected %v", result, expected)
	}
}

// TestGetOccupiedSeatsBySight ...
func TestGetOccupiedSeatsBySight(t *testing.T) {
	expected := 26

	path := "day11/test-cases/given-sample.txt"
	testCase := utils.GetArrayOfStringsFromFile(path)

	result := GetOccupiedSeatsBySight(testCase)

	if result != expected {
		t.Errorf("unexpected result, got %v but expected %v", result, expected)
	}
}
