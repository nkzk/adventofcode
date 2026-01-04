package main

import (
	"fmt"
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

	for k, v := range adjacency {
		fmt.Printf("%s: %s\n", k, v)
	}

	return sum
}

// func part1(adjacency map[string][]string, visited map[string]bool, result int) int {

// }
