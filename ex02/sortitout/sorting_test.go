package sortitout

import (
	"fmt"
	"reflect"
	"testing"
)

// Stack Oh.flow
func SliceToArray(in interface{}) interface{} {
	s := reflect.ValueOf(in)
	if s.Kind() != reflect.Slice {
		panic("not a slice")
	}
	t := reflect.ArrayOf(s.Len(), s.Type().Elem())
	a := reflect.New(t).Elem()
	reflect.Copy(a, s)
	return a.Interface()
}

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

func TestSortInt(t *testing.T) {
	testCases := []struct {
		input    []string
		expected []int
	}{
		{[]string{"4", "2", "8", "9"}, []int{2, 4, 8, 9}},
		{[]string{"1099", "1111", "1001", "1201"}, []int{1001, 1099, 1111, 1201}},
		{[]string{"13", "34", "47", "76"}, []int{13, 34, 47, 76}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %s", tc.input), func(t *testing.T) {
			ret := SortInt(tc.input)
			ret_arr := SliceToArray(ret)
			expected_arr := SliceToArray(tc.expected)

			if ret_arr != expected_arr {
				t.Errorf("Input: %s, expected %v, but got %v", tc.input, tc.expected, ret)
			}
		})
	}
}

func TestSortFloat(t *testing.T) {
	testCases := []struct {
		input    []string
		expected []float64
	}{
		{[]string{"0.01", "3.14", "7.8", "9.0"}, []float64{0.01, 3.14, 7.8, 9.0}},
		{[]string{"4.2", "2.0", "8", "0.0003"}, []float64{0.0003, 2.0, 4.2, 8}},
		{[]string{"20.4", "40.78", "33.3", "0"}, []float64{0, 20.4, 33.3, 40.78}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %s", tc.input), func(t *testing.T) {
			ret := SortFloat(tc.input)
			ret_arr := SliceToArray(ret)
			expected_arr := SliceToArray(tc.expected)

			if ret_arr != expected_arr {
				t.Errorf("Input: %s, expected %v, but got %v", tc.input, tc.expected, ret)
			}
		})
	}
}

func TestSortString(t *testing.T) {
	testCases := []struct {
		input    []string
		expected []string
	}{
		{[]string{"Mot", "Hai", "Ba", "Bon"}, []string{"Ba", "Bon", "Hai", "Mot"}},
		{[]string{"Theta", "Omega", "Beta", "Gamma"}, []string{"Beta", "Gamma", "Omega", "Theta"}},
		{[]string{"Apple", "Absolute", "Adventure", "Anger"}, []string{"Absolute", "Adventure", "Anger", "Apple"}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %s", tc.input), func(t *testing.T) {
			ret := SortString(tc.input)
			retArr := SliceToArray(ret)
			expectedArr := SliceToArray(tc.expected)
			if retArr != expectedArr {
				t.Errorf("Input: %s, expected %s, but got %s", tc.input, tc.expected, ret)
			}
		})
	}
}
