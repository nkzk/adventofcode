package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Part1(input []string) int {
	var sum int

	result := [][]string{}

	for _, line := range input {
		n := make([]string, 0)
		fields := strings.FieldsSeq(line)

		for field := range fields {
			n = append(n, field)
		}

		result = append(result, n)
	}

	for i := 0; i <= len(result[0])-1; i++ {
		cSum := 0
		for j := 0; j < len(result)-1; j++ {
			if j == 0 {
				cSum += toInt(result[j][i])
				continue
			}

			operation := result[len(input)-1][i]
			switch operation {
			case "*":
				cSum *= toInt(result[j][i])
			case "+":
				cSum += toInt(result[j][i])
			}
		}
		sum += cSum
		fmt.Printf("\n")
	}

	return sum
}

func toInt(s string) int {
	x, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("failed to convert %s to int: %v", s, err)
	}

	return x
}
