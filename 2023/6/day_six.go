package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func parseTwo(line string) []int {
	re := regexp.MustCompile(`(\d+)`)
	matches := re.FindAllStringSubmatch(line, -1)
	a := ""
	for _, match := range matches {
		a = a + match[0]
	}
	num, _ := strconv.Atoi(a)
	return []int{num}
}

func parseOne(line string) []int {
	re := regexp.MustCompile(`(\d+)`)
	matches := re.FindAllStringSubmatch(line, -1)

	nums := make([]int, len(matches))
	for i, match := range matches {
		num, _ := strconv.Atoi(match[0])
		nums[i] = num
	}
	return nums
}

func parseInput(input string, parser func(string) []int) [][]int {
	lines := strings.Split(input, "\n")
	output := make([][]int, len(lines))
	for i, line := range lines {
		nums := parser(line)
		output[i] = nums
	}
	return output
}

func calcBeats(time, best int) int {
	numBeats := 0
	for i := 0; i < time; i++ {
		speed, remaining := i, time-i
		distance := speed * remaining
		if distance > best {
			numBeats++
		}
	}
	return numBeats
}

func daySix(input string, parser func(string) []int) {
	output := parseInput(input, parser)
	waysToBeat := 1
	for j := 0; j < len(output[0]); j++ {
		time, best := output[0][j], output[1][j]
		ways := calcBeats(time, best)
		waysToBeat *= ways
	}
	fmt.Println(waysToBeat)
}
