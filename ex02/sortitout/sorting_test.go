package sortitout

import (
	"fmt"
	"testing"
)

func TestValidateInt(t *testing.T) {
	testCases := []struct {
		input    []string
		expected bool
	}{
		{[]string{"Mot", "Hai", "Ba", "Bon"}, false},
		{[]string{"5", "6", "7", "8"}, true},
		{[]string{"10", "J", "Q", "K"}, false},
		{[]string{"7", "7", "8", "8", "9", "9"}, true},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %s", tc.input), func(t *testing.T) {
			ret := validateInt(tc.input)
			if ret != tc.expected {
				t.Errorf("Input: %s, expected %t, but got %t", tc.input, tc.expected, ret)
			}
		})
	}

}

func TestValidateFloat(t *testing.T) {
	testCases := []struct {
		input    []string
		expected bool
	}{
		{[]string{"10", "Boi", "Q", "Gia", "A"}, false},
		{[]string{"5", "6", "7", "8"}, true},
		{[]string{"0.001", "1.4", "11.5", "3.14"}, true},
		{[]string{"7", "7.1", "7.9", "8.0", "8.999", "9"}, true},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %s", tc.input), func(t *testing.T) {
			ret := validateFloat(tc.input)
			if ret != tc.expected {
				t.Errorf("Input: %s, expected %t, but got %t", tc.input, tc.expected, ret)
			}
		})
	}
}

func TestValidateString(t *testing.T) {
	testCases := []struct {
		input    []string
		expected bool
	}{
		{[]string{"Mot", "Hai", "Ba", "Bon"}, true},
		{[]string{"5", "6", "7", "8"}, false},
		{[]string{"10", "J", "Q", "K"}, false},
		{[]string{"J", "Q", "K", "A"}, true},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %s", tc.input), func(t *testing.T) {
			ret := validateString(tc.input)
			if ret != tc.expected {
				t.Errorf("Input: %s, expected %t, but got %t", tc.input, tc.expected, ret)
			}
		})
	}
}

// func TestSortString(t *testing.T) {
// 	testCases := []struct {
// 		input    []string
// 		expected []string
// 	}{
// 		{[]string{"Mot", "Hai", "Ba", "Bon"}, []string{"Ba", "Bon", "Hai", "Mot"}},
// 		{[]string{"Theta", "Omega", "Beta", "Gamma"}, []string{"Beta", "Gamma", "Omega", "Theta"}},
// 		{[]string{"Apple", "Absolute", "Adventure", "Anger"}, []string{"Absolute", "Adventure", "Anger", "Apple"}},
// 	}

// 	for _, tc := range testCases {
// 		t.Run(fmt.Sprintf("Input: %s", tc.input), func(t *testing.T) {
// 			ret := SortString(tc.input)
// 			// Need to convert ret & tc.expected to array to compare
// 			if bytes.Compare(ret, tc.expected) != 0 {
// 				t.Errorf("Input: %s, expected %s, but got %s", tc.input, tc.expected, ret)
// 			}
// 		})
// 	}
// }
