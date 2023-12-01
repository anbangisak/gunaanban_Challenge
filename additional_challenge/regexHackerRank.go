// https://www.hackerrank.com/challenges/validating-credit-card-number/problem
// Validating Credit Card Numbers
// conditions
// It must start with a 4,5 or 6.
// It must contain exactly 16 digits.
// It must only consist of digits (0-9).
// It may have digits in groups of 4, separated by one hyphen "-".
// It must NOT use any other separator like ' ' , '_', etc.
// It must NOT have 4 or more consecutive repeated digits.

package main

import (
	"fmt"
	"regexp"
)

func main() {
	// golang doesn't support lookahead
	re, err := regexp.Compile(`((?!^[0-37-9])(?(?<=\d{4})-?|)(\d)(?!\2{2}-\2|\2-\2{2}|-\2{3})){16}`)
	if err != nil {
		fmt.Println(err)
	}
	inputList := []string{
		"4123456789123456", "5123-4567-8912-3456", "61234-567-8912-3456", "4123356789123456", "5133-3367-8912-3456",
		"5123 - 3567 - 8912 - 3456",
	}
	fmt.Println(re)
	for _, input := range inputList {
		// fmt.Println(input)
		fmt.Printf("%v : %v\n", input, re.MatchString(input))
	}
}
