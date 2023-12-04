package main

import (
	"fmt"
	"strings"
	"unicode"
)

var transformations = map[string]string{
	"one":   "one1one",
	"two":   "two2two",
	"three": "three3three",
	"four":  "four4four",
	"five":  "five5five",
	"six":   "six6six",
	"seven": "seven7seven",
	"eight": "eight8eight",
	"nine":  "nine9nine",
}

func transform(input string) string {
	transformed := input
	for k, v := range transformations {
		transformed = strings.ReplaceAll(transformed, k, v)
	}
	return transformed
}

func parseDigits(line string) int {
	first, last := -1, -1
	for _, char := range line {
		if unicode.IsDigit(char) {
			if first == -1 {
				first = int(char - '0')
			}
			last = int(char - '0')
		}
	}
	return first*10 + last
}

func dayOne(input string) {
	split := strings.Split(input, "\n")
	sum := 0
	for _, line := range split {
		sum += parseDigits(transform(line))
	}
	fmt.Println(sum)
}
