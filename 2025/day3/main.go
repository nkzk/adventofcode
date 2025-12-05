package main

import (
	"fmt"
	"log"
	"github.com/nkzk/adventofcode/utils"
)

func main() {
	var input []string
	err := utils.ReadFile("./2025/day3/input", &input)
	if err != nil {
		log.Panicf("failed to read file: %v", err)
	}

	part1 := Part1(input)
	fmt.Printf("part 1: %d\n", part1)

	part2 := Part2(input)
	fmt.Printf("part 2: %d\n", part2)
}
