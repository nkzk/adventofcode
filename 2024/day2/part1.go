package main

import (
	"fmt"
)

// Part 1, find safe reports
// Safe: All levels are increasing or decreasing
// Diff is 3 >= 1

// Return count of safe reports
func CountSafeReports(input [][]int) (int, error) {
	var safeCount int

	for _, array := range input {
		increasing := false
		decreasing := false
		safe := true

		for i := range array {
			if i != 0 {
				if array[i] > array[i-1] {
					increasing = true

					diff := array[i] - array[i-1]
					if diff < 1 || diff > 3 {
						safe = false
					}
				} else {
					decreasing = true

					diff := array[i-1] - array[i]
					if diff < 1 || diff > 3 {
						safe = false
					}
				}
			} else {
				continue
			}

		}
		if increasing && decreasing {
			safe = false
		}

		if safe {
			safeCount += 1
			fmt.Printf("safe\n")
		}
	}

	return safeCount, nil
}
