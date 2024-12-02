package main

import (
	"fmt"
	"strconv"
	"strings"
	"utils"
)

type Result [][]int

func lineProcessor(line string, result *Result) error {
	lineContent := strings.Fields(line)

	for i, num := range lineContent {
		n, err := strconv.Atoi(num)
		if err != nil {
			return fmt.Errorf("failed to convert string to int: %w", err)
		}

		(*result)[i] = append((*result)[i], n)
	}

	return nil
}

// assert that lineprocessorimplements the utils.LineProcessor signature
var _ utils.LineProcessor[Result] = lineProcessor

func main() {
	resultInput := Result{
		make([]int, 0),
		make([]int, 0),
	}

	err := utils.ReadFile[Result]("./2024/day1/input", lineProcessor, &resultInput)
	if err != nil {
		fmt.Println(err)
		return
	}

	sumPart1 := part1(resultInput)
	sumPart2 := part2(resultInput)

	fmt.Printf("Part 1: %d\n", sumPart1)
	fmt.Printf("Part 2: %d\n", sumPart2)
}
