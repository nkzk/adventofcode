package main

import (
	"sort"
)

func sortList(input []int) []int {
	sort.Ints(input)
	return input
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func part1(input Result) int {
	var sum int
	var resultArray []int

	for _, array := range input {
		array = sortList(array)
	}

	for i := range input[0] {
		resultArray = append(resultArray, abs(input[0][i]-input[1][i]))
	}

	for _, x := range resultArray {
		sum += x
	}

	return sum
}
