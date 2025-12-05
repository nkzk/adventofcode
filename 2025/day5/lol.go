package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// Slow lol
func Part1(input []string) int {
	var sum int

	fmt.Printf("parsing input...\n")
	freshIdRanges := []string{}
	availableIds := []string{}

	for _, line := range input {
		if strings.Contains(line, "-") {
			freshIdRanges = append(freshIdRanges, line)
		} else {
			availableIds = append(availableIds, line)
		}
	}

	fmt.Printf("populating map...\n")
	freshIds := make(map[int]bool)

	for _, r := range freshIdRanges {
		x, y, ok := strings.Cut(r, "-")
		if !ok {
			log.Fatalf("failed to cut %s", r)
		}

		start, err := strconv.Atoi(x)
		if err != nil {
			log.Fatalf("%v", err)
		}

		end, err := strconv.Atoi(y)
		if err != nil {
			log.Fatalf("%v", err)
		}

		for i := start; i <= end; i++ {
			freshIds[i] = true
		}
	}

	fmt.Printf("counting available fresh ingredients...\n")
	for _, s := range availableIds {
		id, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("%v", err)
		}
		_, ok := freshIds[id]
		if !ok {
			continue
		}
		sum++
	}
	return sum
}
