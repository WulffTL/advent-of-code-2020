package main

import (
	"testing"

	"github.com/WulffTL/advent-of-code-2020/utils"
)

func TestGetValidPasswords(t *testing.T) {
	expected := 2

	path := "day2/test-cases/given-sample.txt"
	testCase := utils.GetArrayOfStringsFromFile(path)

	result := GetValidPasswords(testCase)

	if result != expected {
		t.Errorf("unexpected result, got %v but expected %v", result, expected)
	}
}

func TestGetValidPasswords2(t *testing.T) {
	expected := 1

	path := "day2/test-cases/given-sample.txt"
	testCase := utils.GetArrayOfStringsFromFile(path)

	result := GetValidPasswords2(testCase)

	if result != expected {
		t.Errorf("unexpected result, got %v but expected %v", result, expected)
	}
}
