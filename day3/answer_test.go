package main

import (
	"testing"

	"github.com/WulffTL/advent-of-code-2020/utils"
)

func TestGetTreesForAngle(t *testing.T) {
	expected := 7
	rise := 1
	run := 3

	path := "day3/test-cases/given-sample.txt"
	testCase := utils.GetArrayOfStringsFromFile(path)

	result := GetTreesForAngle(testCase, rise, run)

	if result != expected {
		t.Errorf("unexpected result, got %v but expected %v", result, expected)
	}
}

func TestGetTreesForMultipleAngles(t *testing.T) {
	expected := 336

	path := "day3/test-cases/given-sample.txt"
	testCase := utils.GetArrayOfStringsFromFile(path)

	result := GetTreesForMultipleAngles(testCase)

	if result != expected {
		t.Errorf("unexpected result, got %v but expected %v", result, expected)
	}
}
