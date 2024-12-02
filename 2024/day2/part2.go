package main

import "fmt"

func remove(slice []int, i int) []int {
	newSlice := make([]int, len(slice)-1)
	copy(newSlice, slice[:i])
	copy(newSlice[i:], slice[i+1:])
	return newSlice

}

func IsReportSafe(input []int) (int, bool) {
	increasing := false
	decreasing := false
	safe := true
	failingIndex := 0

	for i := range input {
		if i != 0 {
			if input[i] > input[i-1] {
				increasing = true
				if decreasing {
					safe = false
					failingIndex = i
				}

				diff := input[i] - input[i-1]
				if diff < 1 || diff > 3 {
					safe = false
					failingIndex = i
				}
			} else {
				decreasing = true
				if increasing {
					safe = false
					failingIndex = i
				}

				diff := input[i-1] - input[i]
				if diff < 1 || diff > 3 {
					safe = false
					failingIndex = i
				}
			}
		} else {
			continue
		}

	}

	if safe {
		fmt.Printf("safe\n")
		return 0, true
	}

	return failingIndex, false
}
