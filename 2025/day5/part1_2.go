package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Part1_2(input []string) int {
	var sum int

	freshIdRanges := []string{}
	availableIds := []string{}

	for _, line := range input {
		if strings.Contains(line, "-") {
			freshIdRanges = append(freshIdRanges, line)
		} else {
			availableIds = append(availableIds, line)
		}
	}

	fmt.Printf("counting available fresh ingredients...\n")
	for _, s := range availableIds {
		id, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("%v", err)
		}

	loopThroughAllRanges:
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

			if id >= start && id <= end {
				sum++
				break loopThroughAllRanges
			}
		}

	}
	return sum
}
