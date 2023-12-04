package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

const (
	runePeriod rune = '.'
	runeGear   rune = '*'
)

var directions = [][]int{
	{0, -1},  // left
	{1, -1},  // down left
	{-1, -1}, // up left
	{0, 1},   // right
	{1, 0},   // down
	{-1, 0},  // up
	{1, 1},   // down right
	{-1, 1},  // up right
}

type Grid struct {
	grid    [][]rune
	visited [][]bool
}

func NewGrid(input string) *Grid {
	rows := strings.Split(input, "\n")
	grid := make([][]rune, len(rows))
	visited := make([][]bool, len(rows))
	for i, row := range rows {
		grid[i] = []rune(row)
		visited[i] = make([]bool, len(grid[i]))
	}
	return &Grid{grid: grid, visited: visited}
}

func (g *Grid) buildDigit(row, start int) int {
	strNum := ""

	// build right
	for right := start; right < len(g.grid[row]); right++ {
		if !unicode.IsDigit(g.grid[row][right]) {
			break
		}
		strNum += string(g.grid[row][right])
		g.visited[row][right] = true
	}

	// build left
	for left := start - 1; left >= 0; left-- {
		if !unicode.IsDigit(g.grid[row][left]) {
			break
		}
		strNum = string(g.grid[row][left]) + strNum
		g.visited[row][left] = true
	}
	num, _ := strconv.Atoi(strNum)
	return num
}

func (g *Grid) validDigit(i, j int) bool {
	for col := j; col < len(g.grid[i]); col++ {
		if !unicode.IsDigit(g.grid[i][col]) {
			return false
		}
		for _, dir := range directions {
			xi, xj := i+dir[0], col+dir[1]
			if !g.canVisit(xi, xj) {
				continue
			}

			checkRune := g.grid[xi][xj]
			if checkRune == runePeriod || unicode.IsDigit(checkRune) {
				continue
			}
			return true
		}
	}
	return false
}

func (g *Grid) findAdjDigit(row, col int) int {
	for _, dir := range directions {
		xi, xj := row+dir[0], col+dir[1]
		if !g.canVisit(xi, xj) {
			continue
		}
		g.visited[xi][xj] = true

		if !unicode.IsDigit(g.grid[xi][xj]) {
			continue
		}

		digit := g.buildDigit(xi, xj)
		if digit > 0 {
			return digit
		}
	}
	return 0
}

func (g *Grid) canVisit(row, col int) bool {
	if row < 0 || row >= len(g.grid) || col < 0 || col >= len(g.grid[0]) || g.visited[row][col] {
		return false
	}
	return true
}

// findGearRatio the start of a number, looks for '*' rune
func (g *Grid) findGearRatio(row, col int) (int, int) {
	for x := col; x < len(g.grid[row]); x++ {
		// check if we're at the end of the digit
		if !unicode.IsDigit(g.grid[row][x]) {
			return -1, -1
		}
		for _, dir := range directions {
			xi, xj := row+dir[0], x+dir[1]
			if !g.canVisit(xi, xj) {
				continue
			}
			g.visited[xi][xj] = true

			if g.grid[xi][xj] == runeGear {
				return xi, xj
			}
		}
	}
	return -1, -1
}

func processGrid(input string, processDigit func(*Grid, int, int)) {
	grid := NewGrid(input)
	n, m := len(grid.grid), len(grid.grid[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if !grid.visited[i][j] && unicode.IsDigit(grid.grid[i][j]) {
				processDigit(grid, i, j)
			}
			grid.visited[i][j] = true
		}
	}
}

func dayThreePartOne(input string) {
	fmt.Println("part one")
	sum := 0
	processGrid(input, func(grid *Grid, i, j int) {
		digit := grid.buildDigit(i, j)
		if grid.validDigit(i, j) {
			sum += digit
		}
	})
	fmt.Println(sum)
}

func dayThreePartTwo(input string) {
	fmt.Println("part two")
	sum := 0
	processGrid(input, func(grid *Grid, i, j int) {
		digit := grid.buildDigit(i, j)
		gi, gj := grid.findGearRatio(i, j)
		if gi != -1 && gj != -1 {
			adjDigit := grid.findAdjDigit(gi, gj)
			if adjDigit > 0 {
				sum += digit * adjDigit
			}
		}
	})
	fmt.Println(sum)
}
