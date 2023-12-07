package main

func main() {
	inputs := []string{
		"Time:      7  15   30\nDistance:  9  40  200",
		"Time:        42     89     91     89\nDistance:   308   1170   1291   1467",
	}
	for _, input := range inputs {
		daySix(input, parseOne)
		daySix(input, parseTwo)
	}
}
