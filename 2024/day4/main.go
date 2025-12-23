package main

import (
	"fmt"
	"github.com/nkzk/adventofcode/utils"
	"log"
)

type Input [][]rune

func lineProcessor(line string, result *Input) error {
	if len(line) > 0 {
		(*result) = append((*result), []rune(line))
	}
	return nil
}

func main() {
	input := make(Input, 0)
	err := utils.ReadFile_old[Input]("./2024/day4/input", lineProcessor, &input)
	if err != nil {
		log.Panicf("failed to read file: %v", err)
	}

	sum := Part1(input)
	fmt.Printf("part 1: %d\n", sum)

	// got := Part2(input)
	// fmt.Printf("part 2: %d\n", got)
}
