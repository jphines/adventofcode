package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func sum(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func parseNumbers(numbers []string) map[int]bool {
	nums := map[int]bool{}
	for _, strNum := range numbers {
		if num, err := strconv.Atoi(strNum); err == nil {
			nums[num] = true
		}
	}
	return nums
}

func parseLine(line string) (map[int]bool, map[int]bool) {
	postHeading := strings.Split(line, ":")[1]
	split := strings.Split(postHeading, "|")
	left := parseNumbers(strings.Split(split[0], " "))
	right := parseNumbers(strings.Split(split[1], " "))
	return left, right
}

func dayFour(input string) {
	split := strings.Split(input, "\n")
	sumPoints := 0.0
	cards := make([]int, len(split))
	for i, _ := range cards {
		cards[i] = 1
	}
	for i, line := range split {
		left, right := parseLine(line)
		matches := 0.0
		for key, _ := range left {
			if right[key] {
				matches++
			}
		}
		if matches > 0.0 {
			sumPoints += math.Pow(2, matches-1)
		}
		for j := 1; j < int(matches)+1; j++ {
			cards[i+j] += cards[i]
		}
	}

	fmt.Println(sumPoints)
	fmt.Println(sum(cards))
}
