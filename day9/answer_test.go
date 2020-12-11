package main

import (
	"github.com/WulffTL/advent-of-code-2020/utils"
	"testing"
)

// TestGetAccumulatorBeforeInfiniteLoop ...
func TestGetAccumulatorBeforeInfiniteLoop(t *testing.T) {
	expected := int64(127)

	path := "day9/test-cases/given-sample.txt"
	testCase := utils.GetArrayOfStringsFromFile(path)

	preambleLength := 5
	result := GetFirstViolatingNumber(testCase, preambleLength)

	if result != expected {
		t.Errorf("unexpected result, got %v but expected %v", result, expected)
	}
}

// TestGetEncryptionWeakness ...
func TestGetEncryptionWeakness(t *testing.T) {
	expected := int64(62)

	path := "day9/test-cases/given-sample.txt"
	testCase := utils.GetArrayOfStringsFromFile(path)

	previousAnswer := int64(127)
	result := GetEncryptionWeakness(testCase, previousAnswer)

	if result != expected {
		t.Errorf("unexpected result, got %v but expected %v", result, expected)
	}
}
