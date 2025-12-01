package main

import (
	"log"
	"os"
	"strconv"
)

func Part1(input []string) int {
	var sum int
	b := 50

	for _, line := range input {
		b = rotate(b, line)
		if b == 0 {
			sum += 1
		}
	}

	return sum
}

func rotate(a int, rotation string) int {
	direction := rotation[:1]
	if direction != "L" && direction != "R" {
		log.Printf("invalid direction, expected L or R but got: %s", direction)
		os.Exit(1)
	}

	b, err := strconv.Atoi(rotation[1:])
	if err != nil {
		log.Printf("fatal error, check input! failed to convert %s to int: %v", rotation[1:], err)
		os.Exit(2)
	}

	if direction == "L" {
		return mod(a-b, 100)
	}

	return mod(a+b, 100)
}

func mod(a, b int) int {
	return ((a % b) + b) % b
}
