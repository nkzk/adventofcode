package main

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

func Part2(input []string) int {
	var sum int

	freshIdRanges := []string{}

	for _, line := range input {
		if strings.Contains(line, "-") {
			freshIdRanges = append(freshIdRanges, line)
		}
	}

	// sort by interval-start
	slices.SortFunc(freshIdRanges, func(a, b string) int {
		aS, _ := convert(a)
		bS, _ := convert(b)

		if aS < bS {
			return -1
		}

		if aS > bS {
			return 1
		}

		return 0
	})

	merged := mergeIntervals(freshIdRanges)
	for _, r := range merged {
		s, e := convert(r)
		sum += (e - s)
		if e != s {
			sum++
		}
	}

	return sum
}

func mergeIntervals(intervals []string) []string {
	changed := true
	for changed {
		changed = false
		for i, r := range intervals {
			if i+1 > len(intervals)-1 {
				break
			}

			nextInterval := ""

			nextInterval = intervals[i+1]
			nS, nE := convert(nextInterval)
			s, e := convert(r)

			if nS <= e+1 {
				copy(intervals[i:], intervals[i+1:])     // shift left
				intervals = intervals[:len(intervals)-1] // shrink array
				intervals[i] = fmt.Sprintf("%d-%d", s, max(e, nE))
				changed = true
			}
		}
	}
	return intervals
}

func convert(s string) (int, int) {
	x, y, ok := strings.Cut(s, "-")
	if !ok {
		log.Fatalf("failed to cut %s", s)
	}

	start, err := strconv.Atoi(x)
	if err != nil {
		log.Fatalf("%v", err)
	}

	end, err := strconv.Atoi(y)
	if err != nil {
		log.Fatalf("%v", err)
	}

	return start, end
}
