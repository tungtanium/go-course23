package main

import (
	"fmt"
	"testing"
)

func TestValidateCountryCode(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"US", true},
		{"PRC", false},
		{"VN", true},
		{"JPN", false},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %s", tc.input), func(t *testing.T) {
			result := validateCountryCode(tc.input)
			if result != tc.expected {
				t.Errorf("For input: %s, expected %t, but got %t", tc.input, tc.expected, result)
			}
		})
	}

}

func TestFormatNameByCountry(t *testing.T) {
	// inputName := []string{"Tung", "Pham", "Ly", "Son"}
	// countryCode := "VN"
	// expected := "Pham Ly Son Tung"
	testCases := []struct {
		inputName   []string // First, Last, Middle...
		countryCode string
		expected    string
	}{
		{[]string{"Peter", "Hernandez", "Gene"}, "US", "Peter Gene Hernandez"},
		{[]string{"Henry", "David", "Charles", "Albert"}, "UK", "Henry Charles Albert David"},
		{[]string{"Thao", "Nguyen", "Thi", "Thu"}, "VN", "Nguyen Thi Thu Thao"},
		{[]string{"Stefani", "Germanotta", "Joanne", "Angelina"}, "KR", "Germanotta Joanne Angelina Stefani"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %s %s", tc.inputName, tc.countryCode), func(t *testing.T) {
			result := formatNameByCountry(tc.inputName, tc.countryCode)
			if result != tc.expected {
				t.Errorf("For %s %s, expected %s, but got %s", tc.inputName, tc.countryCode, tc.expected, result)
			}
		})
	}
}
