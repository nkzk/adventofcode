package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"utils"
)

// part 1
type Input [][]int

func lineProcessor(line string, input *Input) error {
	lineContent := strings.Fields(line)

	for i, num := range lineContent {
		n, err := strconv.Atoi(num)
		if err != nil {
			return fmt.Errorf("failed to convert string to int: %w", err)
		}

		(*input)[i] = append((*input)[i], n)
	}

	return nil
}

// assert that lineprocessorimplements the utils.LineProcessor signature
var _ utils.LineProcessor[Input] = lineProcessor

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

func day1(input Input) int {
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

// part 2

// return similarity score

// 1 3
// 2 3
// 3 3
// 3 4
// 3 5
// 4 9

// 1 - sort lists, nvm dont need
// 2 -  make counter
// 		[1]0
// 		[2]0
// 		[3]3
// 		[4]1
//
// 3 - go through right list
// 4 - get number, count untill next, add sum to hashmap [3]

// 5 - go through left list
// 6 - lookup hashmap index, add to sum

func day1_part2(input Input) int {
	// go through right list
	count := make(map[int]int)

	for _, number := range input[1] {
		count[number] += 1
	}

	// go through left list
	var sum int

	for _, number := range input[0] {
		_, ok := count[number]
		if !ok {
			sum += 0
		} else {
			fmt.Printf("adding %d * %d\n", number, count[number])
			sum += (number * count[number])
		}

	}

	return sum
}

func main() {
	input := Input{
		make([]int, 0),
		make([]int, 0),
	}

	err := utils.ReadFile[Input]("./2024/day1/input", lineProcessor, &input)
	if err != nil {
		fmt.Println(err)
		return
	}

	sumPart1 := day1(input)
	sumPart2 := day1_part2(input)

	fmt.Printf("Part 1: %d\n", sumPart1)
	fmt.Printf("Part 2: %d\n", sumPart2)
}
