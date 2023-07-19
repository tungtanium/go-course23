package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func is_numeric(astr string) bool {
	return regexp.MustCompile(`\d`).MatchString(astr)
}

func validate_country_code(code string) bool {
	codeLength := len(code)
	if codeLength != 2 || is_numeric(code) {
		return false
	}

	return true
}

func reverse_name_order(name_to_reverse []string) string {
	name_length := len(name_to_reverse)
	first_name := name_to_reverse[0]
	last_name := name_to_reverse[name_length-1]

	if name_length == 2 {
		return fmt.Sprintf("%s %s", last_name, first_name)
	}
	middle_name := strings.Join(name_to_reverse[1:name_length-1], " ")
	return fmt.Sprintf("%s %s %s", last_name, middle_name, first_name)
}

func format_name_by_country(inputName []string, countryCode string) string {
	inputLength := len(inputName)

	if !validate_country_code(countryCode) {
		fmt.Println("Invalid country code format. Valid format example: VN")
	}

	switch inputLength {
	case 1:
		return string(inputName[0])
	default:
		countryCode = strings.ToUpper(countryCode)
		switch countryCode {
		// Countries using reversed name order. Country code in ISO3166 Alpha 2
		// https://gist.github.com/ssskip/5a94bfcd2835bf1dea52#file-iso3166-1-alpha2-json
		// https://obaninternational.com/blog/how-do-naming-conventions-vary-around-the-world
		case "CN", "JP", "KR", "KP", "VN", "TW", "HU", "BA", "HR", "RS":
			reversed := reverse_name_order(inputName)
			return reversed
		default:
			return strings.Join(inputName, " ")
		}
	}
}

func main() {
	args := os.Args[1:]
	countryCode := args[len(args)-1]
	inputName := args[:len(args)-1]

	if len(inputName) < 1 {
		fmt.Println("No name to order")
	}
	fmt.Println("Result: ", format_name_by_country(inputName, countryCode))

}
