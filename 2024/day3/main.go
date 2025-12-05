package main

import (
	"fmt"
	"log"
	"github.com/nkzk/adventofcode/utils"
)

type Input []string

func lineProcessor(line string, result *Input) error {
	(*result) = append((*result), line)
	return nil
}

func main() {
	input := make(Input, 0)
	err := utils.ReadFile_old[Input]("./2024/day3/input", lineProcessor, &input)
	if err != nil {
		log.Panicf("failed to read file: %v", err)
	}

	sum := Part1(input)
	fmt.Printf("part 1: %d\n", sum)

	got := Part2(input)
	fmt.Printf("part 2: %d\n", got)
}
