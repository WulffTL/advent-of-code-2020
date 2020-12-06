package main

import (
	"testing"

	"github.com/WulffTL/advent-of-code-2020/utils"
)

func TestGetValidPassportCount(t *testing.T) {
	expected := 2

	path := "day4/test-cases/given-sample.txt"
	testCase := utils.GetArrayOfStringsFromFile(path)

	result := GetValidPassportCount(testCase)

	if result != expected {
		t.Errorf("unexpected result, got %v but expected %v", result, expected)
	}
}

func TestGetValidPassportCountAllInvalid(t *testing.T) {
	expected := 0

	path := "day4/test-cases/given-sample-invalid.txt"
	testCase := utils.GetArrayOfStringsFromFile(path)

	result := GetValidPassportCount(testCase)

	if result != expected {
		t.Errorf("unexpected result, got %v but expected %v", result, expected)
	}
}

func TestGetValidPassportCountAllValid(t *testing.T) {
	expected := 4

	path := "day4/test-cases/given-sample-valid.txt"
	testCase := utils.GetArrayOfStringsFromFile(path)

	result := GetValidPassportCount(testCase)

	if result != expected {
		t.Errorf("unexpected result, got %v but expected %v", result, expected)
	}
}
