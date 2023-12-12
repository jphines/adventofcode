package main

import (
	"fmt"
	"strings"
)

type Maze struct {
	Start   []int
	Board   [][]string
	Visited [][]bool
}

var pipes = map[string][][]int{
	"|": {{1, 0}, {-1, 0}},  // down, up / south, north
	"-": {{0, 1}, {0, -1}},  // right, left / east, west
	"L": {{0, 1}, {-1, 0}},  // right, down / north, west
	"J": {{0, -1}, {-1, 0}}, // left, down / north, east
	"7": {{0, -1}, {1, 0}},  // right, up / south, west
	"F": {{0, 1}, {1, 0}},   // left, up / south, east
}

var verticalPipes = map[string]int{
	"|": 0,
	"L": 1,
	"7": 1,
	"J": 2,
	"F": 2,
}

//-L|F7
//7S-7|
//L|7||
//-L-J|
//L|-JF

func (m *Maze) dfs(i, j, depth int) int {
	pipe := m.Board[i][j]
	for _, dir := range pipes[pipe] {
		xi, xj := i+dir[0], j+dir[1]
		if xi < 0 || xi >= len(m.Board) || xj < 0 || xj >= len(m.Board[0]) {
			continue
		}
		if m.Visited[xi][xj] {
			continue
		}
		m.Visited[xi][xj] = true
		depth = max(depth, m.dfs(xi, xj, depth+1))
	}
	return depth
}

func (m *Maze) Walk(i, j, depth int) int {
	m.Visited[i][j] = true
	for pipe, dirs := range pipes {
		for _, dir := range dirs {
			xi, xj := i+dir[0], j+dir[1]
			if xi < 0 || xi >= len(m.Board) || xj < 0 || xj >= len(m.Board[0]) {
				continue
			}
			if m.Visited[xi][xj] {
				continue
			}
			if m.Board[xi][xj] == pipe {
				m.Visited[xi][xj] = true
				mDepth := m.dfs(xi, xj, depth+1)
				depth = max(depth, mDepth)
			}
		}
	}

	return depth
}

func (m *Maze) Enclosed() int {
	count := 0
	for i, _ := range m.Board {
		in, previous := false, ""
		for j, _ := range m.Board[i] {
			// not a pipe and seen boundaries odd number of times
			if !m.Visited[i][j] && in {
				count++
			}
			// is a pipe
			if m.Visited[i][j] {
				pipe := m.Board[i][j]
				switch pipe {
				case "|":
					in = !in
					previous = ""
				case "L":
					previous = pipe
				case "7":
					if previous == "L" {
						in = !in
					}
					previous = ""
				case "F":
					previous = pipe
				case "J":
					if previous == "F" {
						in = !in
					}
					previous = ""
				}
			}
		}
	}
	return count
}

func NewMaze(input string) *Maze {
	split := strings.Split(input, "\n")
	maze := &Maze{
		Board:   make([][]string, len(split)),
		Visited: make([][]bool, len(split)),
	}
	for i, line := range split {
		maze.Board[i] = strings.Split(line, "")
		maze.Visited[i] = make([]bool, len(maze.Board[i]))
		for j, char := range maze.Board[i] {
			if char == "S" {
				maze.Start = []int{i, j}
			}
		}
	}
	return maze
}

func dayTen(input string) {
	maze := NewMaze(input)

	// part 1
	depth := maze.Walk(maze.Start[0], maze.Start[1], 0)
	fmt.Printf("Depth: %d\n", depth/2+1)

	// part 2
	e := maze.Enclosed()
	fmt.Printf("Enclosed: %d\n", e)
}
