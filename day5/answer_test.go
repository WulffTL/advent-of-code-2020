package main

import (
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

func TestGetIDForBoardingPass(t *testing.T) {
	testCases := []testInput{
		{
			input:    "BFFFBBFRRR",
			expected: 567,
		},
		{
			input:    "FFFBBBFRRR",
			expected: 119,
		},
		{
			input:    "BBFFBBFRLL",
			expected: 820,
		},
		{
			input:    "FFFFFFFLLL",
			expected: 0,
		},
		{
			input:    "BBBBBBBRRR",
			expected: 1023,
		},
	}

	for i := 0; i < len(testCases); i++ {
		result := GetIDForBoardingPass(testCases[i].input)

		if result != testCases[i].expected {
			t.Errorf("unexpected result, got %v but expected %v", result, testCases[i].expected)
		}
	}

}

func TestGetRowColForBoardingPass(t *testing.T) {
	testCases := []testRowColFromInput{
		{
			input:       "BFFFBBFRRR",
			expectedRow: 70,
			expectedCol: 7,
		},
		{
			input:       "BBBBBBFLLL",
			expectedRow: 126,
			expectedCol: 0,
		},
		{
			input:       "FFFBBBFRRR",
			expectedRow: 14,
			expectedCol: 7,
		},
		{
			input:       "BBFFBBFRLL",
			expectedRow: 102,
			expectedCol: 4,
		},
	}

	for i := 0; i < len(testCases); i++ {
		row := GetRow(testCases[i].input)
		col := GetCol(testCases[i].input)

		if row != testCases[i].expectedRow || col != testCases[i].expectedCol {
			t.Errorf("unexpected result, got row=%v and col=%v but expected row=%v and col=%v", row, col, testCases[i].expectedRow, testCases[i].expectedCol)
		}
	}

}
