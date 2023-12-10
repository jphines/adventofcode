package main

import (
	"strconv"
	"strings"
)

func parseInput(input string) [][]int {
	lines := strings.Split(input, "\n")
	grid := make([][]int, len(lines))
	for i, line := range lines {
		nums := []int{}
		for _, strNum := range strings.Split(line, " ") {
			num, err := strconv.Atoi(strNum)
			if err != nil {
				panic(err)
			}
			nums = append(nums, num)
		}
		grid[i] = nums
	}
	return grid
}

func isZero(arr []int) bool {
	for _, num := range arr {
		if num != 0 {
			return false
		}
	}
	return true
}

func prefix(arr []int) (int, int) {
	if isZero(arr) {
		return 0, 0
	}

	res := make([]int, len(arr)-1)
	for i := 0; i < len(arr)-1; i++ {
		res[i] = arr[i+1] - arr[i]
	}

	head, sum := prefix(res)
	newHead := arr[0] - head
	return newHead, sum + newHead
}

func suffix(arr []int) (int, int) {
	if isZero(arr) {
		return 0, 0
	}

	res := make([]int, len(arr)-1)
	for i := 0; i < len(arr)-1; i++ {
		res[i] = arr[i+1] - arr[i]
	}

	tail, sum := suffix(res)
	newTail := arr[len(arr)-1] + tail
	return newTail, sum + newTail
}

func partTwo(input string) {
	grid := parseInput(input)
	sum := 0
	for _, nums := range grid {
		t, _ := prefix(nums)
		sum += t
	}
	println(sum)
}

func partOne(input string) {
	grid := parseInput(input)
	sum := 0
	for _, nums := range grid {
		t, _ := suffix(nums)
		sum += t
	}
	println(sum)
}
