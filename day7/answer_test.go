package main

import (
	"github.com/WulffTL/advent-of-code-2020/utils"
	"testing"
)

type testInput struct {
	input    string
	expected int64
}

type testRowColFromInput struct {
	input       string
	expectedRow int64
	expectedCol int64
}

func TestGetTotalBagCountWhichCanContainGold(t *testing.T) {
	expected := 4

	path := "day7/test-cases/given-sample.txt"
	testCase := utils.GetArrayOfStringsFromFile(path)

	result := GetTotalBagCountWhichCanContainGold(testCase)

	if result != expected {
		t.Errorf("unexpected result, got %v but expected %v", result, expected)
	}
}

func TestGetTotalBagCountWithinGold(t *testing.T) {
	expected := 126

	path := "day7/test-cases/given-sample-2.txt"
	testCase := utils.GetArrayOfStringsFromFile(path)

	result := GetTotalBagCountWithinGold(testCase)

	if result != expected {
		t.Errorf("unexpected result, got %v but expected %v", result, expected)
	}
}
