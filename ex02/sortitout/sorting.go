package sortitout

import (
	"log"
	"regexp"
	"sort"
	"strconv"
)

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

func SortInt(data []string) []int {
	out := make([]int, len(data))

	if !validateInt(data) {
		log.Panic("Input data must be numeric for sorting")
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
		log.Panic("Input data must be numeric for sorting")
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
		log.Panic("Input data must be non-numeric for sorting")
		panic(data)
	}

	sort.Strings(data)

	return data
}
