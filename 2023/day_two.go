package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var knownColors = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func parseGame(input string) int {
	re := regexp.MustCompile(`Game (\d+)`)
	matches := re.FindStringSubmatch(input)
	num, _ := strconv.Atoi(matches[1])
	return num
}

func parseColors(input string) map[string]int {
	colors := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	re := regexp.MustCompile(`(\d+) (red|green|blue)`)
	matches := re.FindAllStringSubmatch(input, -1)

	for _, match := range matches {
		num, _ := strconv.Atoi(match[1])
		color := match[2]
		colors[color] = max(colors[color], num)
	}

	return colors
}

func dayTwo(input string) {
	split := strings.Split(input, "\n")
	sum, powerSum := 0, 0
	for _, line := range split {
		power, invalid := 1, 0
		game := parseGame(line)
		colors := parseColors(line)
		for color, num := range colors {
			power *= num
			if num > knownColors[color] {
				invalid++
			}
		}
		powerSum += power
		if invalid == 0 {
			sum += game
		}
	}

	fmt.Printf("sum of possible games: %d\n", sum)
	fmt.Printf("power sum of all games : %d\n", powerSum)
}
