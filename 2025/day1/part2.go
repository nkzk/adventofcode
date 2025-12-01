package main

import (
	"log"
	"os"
	"strconv"
)

func Part2(input []string) int {
	var sum int
	b := 50

	for _, line := range input {
		r := rotation(line)
		sum += numberOfZeroPasses(b, r)

		b = mod(b+r, 100)

	}

	return sum
}

func rotation(l string) int {
	direction := l[:1]
	if direction != "L" && direction != "R" {
		log.Printf("invalid direction, expected L or R but got: %s", direction)
		os.Exit(1)
	}

	n, err := strconv.Atoi(l[1:])
	if err != nil {
		log.Printf("fatal error, check input! failed to convert %s to int: %v", l[1:], err)
		os.Exit(2)
	}

	if direction == "L" {
		return n * -1
	}

	return n
}

func numberOfZeroPasses(start, rotation int) int {
	if rotation > 0 {
		return floorDiv(start+rotation, 100) - floorDiv(start, 100)
	}

	if rotation < 0 {
		return floorDiv(start-1, 100) - floorDiv(start+rotation-1, 100)
	}

	return 0
}

func floorDiv(a, b int) int {
	if a >= 0 || a%b == 0 {
		return a / b
	}

	return (a / b) - 1
}
