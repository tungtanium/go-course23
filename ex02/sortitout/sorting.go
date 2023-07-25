package sortitout

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
)

func Testing() {
	fmt.Println("\nFrom `sortitout` package with love")
	// Sorting examples
	intSlc := []int{4, 60, 19, 92, 0, 1001}
	// fltSlc := []float64{0.01, 1.0, 0.8, 1.5, 2.3}
	// fltSlc2 := []float64{0.1, 1, 12, 5, 30}
	fltArr := [7]float64{200, 0.01, 0.1, 2.3, 12, 5, 30}
	strSlc := []string{"Foo", "Faz", "Apple", "Aubergine", "Zuccini", "Banana", "Strawberry"}

	sort.Ints(intSlc)
	fmt.Println(intSlc)

	sort.Float64s(fltArr[:])
	fmt.Println(fltArr)

	sort.Strings(strSlc)
	fmt.Println(strSlc)

}

func isNumeric(astr string) bool {
	return regexp.MustCompile(`\d`).MatchString(astr)
}

func validateInt(data []string) bool {
	for _, item := range data {
		_, err := strconv.ParseInt(item, 10, 64)
		if err != nil {
			return false
		}
	}
	return true
}

func validateFloat(data []string) bool {
	for _, item := range data {
		_, err := strconv.ParseFloat(item, 64)
		if err != nil {
			return false
		}
	}
	return true
}

func validateString(data []string) bool {
	for _, item := range data {
		if isNumeric(item) {
			return false
		}
	}
	return true
}

func ValidateDataType(data []string, dataType string) bool {
	switch dataType {
	case "int":
		return validateInt(data)
	case "float":
		return validateFloat(data)
	case "string":
		return validateString(data)
	default:
		message := fmt.Sprintf("%s type is unknown.", dataType)
		log.Panic(message)
		panic(message)
	}
}

func SortInt(data []string) []int {
	out := make([]int, len(data))

	if !validateInt(data) {
		panic(data)
	}

	for idx, item := range data {
		parsed, _ := strconv.Atoi(item)
		out[idx] = parsed
	}
	sort.Ints(out)
	return out
}

func SortFloat(data []string) []float64 {
	out := make([]float64, len(data))

	if !validateFloat(data) {
		panic(data)
	}

	for idx, item := range data {
		parsed, _ := strconv.ParseFloat(item, 64)
		out[idx] = parsed
	}
	sort.Float64s(out)
	return out
}

func SortString(data []string) []string {

	if !validateString(data) {
		panic(data)
	}

	sort.Strings(data)

	return data
}
