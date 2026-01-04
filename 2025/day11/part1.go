package main

import (
	"log"
	"strings"
)

func Part1(input []string) int {
	var sum int

	adjacency := map[string][]string{}

	for i, line := range input {
		in, outputs, ok := strings.Cut(line, ":")
		if !ok {
			log.Fatalf("failed to parse input on line %d", i)
		}

		outputFields := strings.Fields(outputs)

		x := []string{}
		for _, val := range outputFields {
			x = append(x, val)
		}

		adjacency[in] = x
	}

	sum += part1("you", adjacency)

	return sum
}

func part1(val string, adjacency map[string][]string) int {
	if val == "out" {
		return 1
	}

	sum := 0

	for _, adjacent := range adjacency[val] {
		sum += part1(adjacent, adjacency)
	}

	return sum
}
