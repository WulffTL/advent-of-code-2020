package main

import (
	"github.com/WulffTL/advent-of-code-2020/utils"
	"testing"
)

// TestGetVoltageProduct ...
func TestGetVoltageProduct(t *testing.T) {
	expected := 220

	path := "day10/test-cases/given-sample.txt"
	testCase := utils.GetArrayOfStringsFromFile(path)

	result := GetVoltageProduct(testCase)

	if result != expected {
		t.Errorf("unexpected result, got %v but expected %v", result, expected)
	}

	expected = 35

	path = "day10/test-cases/given-sample-2.txt"
	testCase = utils.GetArrayOfStringsFromFile(path)

	result = GetVoltageProduct(testCase)

	if result != expected {
		t.Errorf("unexpected result, got %v but expected %v", result, expected)
	}
}

// TestGetUniqueArrangementsCount
func TestGetUniqueArrangementsCount(t *testing.T) {
	expected := 19208

	path := "day10/test-cases/given-sample.txt"
	testCase := utils.GetArrayOfStringsFromFile(path)

	result := GetUniqueArrangementsCount(testCase)

	if result != expected {
		t.Errorf("unexpected result, got %v but expected %v", result, expected)
	}
	expected = 8

	path = "day10/test-cases/given-sample-2.txt"
	testCase = utils.GetArrayOfStringsFromFile(path)

	result = GetUniqueArrangementsCount(testCase)

	if result != expected {
		t.Errorf("unexpected result, got %v but expected %v", result, expected)
	}

}
