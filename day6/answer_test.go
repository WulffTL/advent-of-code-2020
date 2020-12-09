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

func TestGetTotalAnswerCount(t *testing.T) {
	expected := 11

	path := "day6/test-cases/given-sample.txt"
	testCase := utils.GetArrayOfStringsFromFile(path)

	result := GetTotalAnswerCount(testCase)

	if result != expected {
		t.Errorf("unexpected result, got %v but expected %v", result, expected)
	}
}

func TestGetTotalMutualAnswerCount(t *testing.T) {
	expected := 6

	path := "day6/test-cases/given-sample.txt"
	testCase := utils.GetArrayOfStringsFromFile(path)

	result := GetTotalMutualAnswerCount(testCase)

	if result != expected {
		t.Errorf("unexpected result, got %v but expected %v", result, expected)
	}
}
