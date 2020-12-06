package main

import (
	"testing"

	"github.com/WulffTL/advent-of-code-2020/utils"
)

func TestGetPairOfSums(t *testing.T) {
	testSum := int64(2020)
	expected1 := int64(1721)
	expected2 := int64(299)

	path := "day1/test-cases/given-sample.txt"
	testCase := utils.GetArrayFromFile(path)

	result1, result2 := GetPairOfSums(testCase, testSum)

	if !(result1 == expected1 && result2 == expected2) && !(result1 == expected2 && result2 == expected1) {
		t.Errorf("unexpected result, got %v and %v expected %v and %v", result1, result2, expected1, expected2)
	}
}

func TestGetTripletOfSums(t *testing.T) {
	testSum := int64(2020)
	expected1 := int64(979)
	expected2 := int64(366)
	expected3 := int64(675)

	path := "day1/test-cases/given-sample.txt"
	testCase := utils.GetArrayFromFile(path)

	result1, result2, result3 := GetTripletOfSums(testCase, testSum)

	if result1*result2*result3 != expected1*expected2*expected3 {
		t.Errorf("unexpected result, got %v, %v, and %v expected %v, %v, and %v", result1, result2, result3, expected1, expected2, expected3)
	}
}
