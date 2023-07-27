package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func isNumeric(astr string) bool {
	return regexp.MustCompile(`\d`).MatchString(astr)
}

func numDifferentIntegers(word string) int {
	// Convert to a slice of chars
	chars := strings.Split(word, "")

	// Remove all duplicated
	seen := make(map[string]bool)
	result := []string{}

	// Remove duplicated
	for _, chr := range chars {
		if _, ok := seen[chr]; !ok {
			seen[chr] = true
			if isNumeric(chr) && chr != "0" {
				result = append(result, chr)
			}
		}
	}
	log.Println("All unique integers: ", result)

	return len(result)
}

func main() {
	testString := "a11111bcsd222222"
	count := numDifferentIntegers(testString)
	fmt.Printf("%v", count)
}
