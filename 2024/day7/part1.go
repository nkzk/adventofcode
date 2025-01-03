package main

import (
	"strconv"
	"strings"
)

func ParseInput(input []string) [][]int {
	result := make([][]int, 0)

	for _, line := range input {
		parts := strings.Split(line, ":")
		left := strings.TrimSpace(parts[0])
		right := strings.TrimSpace(parts[1])

		sumNum, err := strconv.Atoi(left)
		if err != nil {
			continue
		}

		row := []int{sumNum}
		rightNumbers := strings.Fields(right)

		for _, num := range rightNumbers {
			n, _ := strconv.Atoi(num)
			row = append(row, n)
		}
		result = append(result, row)
	}

	return result
}

func calculate(a, b int, operator byte) int {
	sum := 0
	switch operator {
	case '+':
		sum = a + b
	case '*':
		sum = a * b
	}

	return sum
}

func calculateMatch(expectedSum, sum int, input []int) bool {
	if len(input) == 0 {
		return sum == expectedSum
	}

	if sum > expectedSum {
		return false
	}

	if calculateMatch(expectedSum, calculate(sum, input[0], '+'), input[1:]) {
		return true
	}

	return calculateMatch(expectedSum, calculate(sum, input[0], '*'), input[1:])
}

func Part1(input []string) int {
	var sum int

	rows := ParseInput(input)
	for _, row := range rows {
		if calculateMatch(row[0], 0, row) {
			sum += row[0]
		}
	}

	return sum
}
