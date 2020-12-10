package main

import (
	"github.com/WulffTL/advent-of-code-2020/utils"
	"testing"
)

// TestGetAccumulatorBeforeInfiniteLoop ...
func TestGetAccumulatorBeforeInfiniteLoop(t *testing.T) {
	expected := 5

	path := "day8/test-cases/given-sample.txt"
	testCase := utils.GetArrayOfStringsFromFile(path)

	result := GetAccumulatorBeforeInfiniteLoop(testCase)

	if result != expected {
		t.Errorf("unexpected result, got %v but expected %v", result, expected)
	}
}

// TestGetAccumulatorAfterFix ...
func TestGetAccumulatorAfterFix(t *testing.T) {
	expected := 8

	path := "day8/test-cases/given-sample.txt"
	testCase := utils.GetArrayOfStringsFromFile(path)

	result := GetAccumulatorAfterFix(testCase)

	if result != expected {
		t.Errorf("unexpected result, got %v but expected %v", result, expected)
	}
}
