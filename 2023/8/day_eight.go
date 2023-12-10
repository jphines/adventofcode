package main

import (
	"fmt"
	"regexp"
	"strings"
)

type Graph struct {
	Directions string
	Nodes      map[string]*Node
}

type Node struct {
	Name       string
	Neighbours []string
}

func (n *Node) String() string {
	return fmt.Sprintf("%v: %#v", n.Name, n.Neighbours)
}

func parseInput(input string) *Graph {
	lines := strings.Split(input, "\n")
	directions := lines[0]

	nodes := make(map[string]*Node)
	for i := 2; i < len(lines); i++ {
		line := lines[i]
		re := regexp.MustCompile(`(\w+) = \((\w+), (\w+)\)`)
		matches := re.FindAllStringSubmatch(line, -1)
		name := matches[0][1]
		neighbours := []string{matches[0][2], matches[0][3]}
		nodes[name] = &Node{name, neighbours}
	}
	return &Graph{directions, nodes}
}

func debug(is ...interface{}) {
	//for _, i := range is {
	//	fmt.Printf("%#v\n", i)
	//}
}

func walk(graph *Graph, node string, stop func(string) bool) int {
	count := 0
	for {
		for _, char := range graph.Directions {
			if char == 'L' {
				node = graph.Nodes[node].Neighbours[0]
			} else {
				node = graph.Nodes[node].Neighbours[1]
			}
			count++
			if stop(node) {
				return count
			}
		}
	}
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(integers ...int) int {
	if len(integers) == 1 {
		return integers[0]
	}

	a, b := integers[0], integers[1]
	result := a * b / gcd(a, b)

	for i := 2; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}

func partOne(input string) {
	graph := parseInput(input)
	count := walk(graph, "AAA", func(node string) bool {
		return node == "ZZZ"
	})
	fmt.Printf("ans: %d\n", count)
}

func partTwo(input string) {
	graph := parseInput(input)
	nodes := []string{}
	for name, _ := range graph.Nodes {
		if strings.HasSuffix(name, "A") {
			nodes = append(nodes, name)
		}
	}

	cycles := make([]int, len(nodes))
	for i, start := range nodes {
		count := walk(graph, start, func(node string) bool {
			return strings.HasSuffix(node, "Z")
		})
		cycles[i] = count
	}
	ans := lcm(cycles...)
	fmt.Printf("ans: %d\n", ans)
}
