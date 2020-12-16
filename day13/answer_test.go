package main

import (
	"github.com/WulffTL/advent-of-code-2020/utils"
	"testing"
)

// TestGetProduct ...
func TestGetProduct(t *testing.T) {
	expected := 295

	path := "day13/test-cases/given-sample.txt"
	testCase := utils.GetArrayOfStringsFromFile(path)

	result := GetProduct(testCase)

	if result != expected {
		t.Errorf("unexpected result, got %v but expected %v", result, expected)
	}
}

// TestGetEarliestConcurrentSchedule ...
func TestGetEarliestConcurrentSchedule(t *testing.T) {
	expected := 1068781
	path := "day13/test-cases/given-sample.txt"
	testCase := utils.GetArrayOfStringsFromFile(path)

	result := GetEarliestConcurrentSchedule(testCase)

	if result != expected {
		t.Errorf("unexpected result, got %v but expected %v", result, expected)
	}

	expected = 1202161486
	path = "day13/test-cases/given-sample-2.txt"
	testCase = utils.GetArrayOfStringsFromFile(path)

	result = GetEarliestConcurrentSchedule(testCase)

	if result != expected {
		t.Errorf("unexpected result, got %v but expected %v", result, expected)
	}

	expected = 1261476
	path = "day13/test-cases/given-sample-3.txt"
	testCase = utils.GetArrayOfStringsFromFile(path)

	result = GetEarliestConcurrentSchedule(testCase)

	if result != expected {
		t.Errorf("unexpected result, got %v but expected %v", result, expected)
	}

	expected = 779210
	path = "day13/test-cases/given-sample-4.txt"
	testCase = utils.GetArrayOfStringsFromFile(path)

	result = GetEarliestConcurrentSchedule(testCase)

	if result != expected {
		t.Errorf("unexpected result, got %v but expected %v", result, expected)
	}
}
