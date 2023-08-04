package main

import (
	"fmt"
	"testing"
)

func TestNumDifferentIntegers(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{"a11111bcsd222222", 2},
		{"1a2b3c4d", 4},
		{"888aokeru8", 1},
		{"abcdefomcvfdgd", 0},
		{"0123456789", 9},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %s", tc.input), func(t *testing.T) {
			ret := numDifferentIntegers(tc.input)
			if ret != tc.expected {
				t.Errorf("Input: %s, expected %d, but got %d", tc.input, tc.expected, ret)
			}
		})
	}
}
