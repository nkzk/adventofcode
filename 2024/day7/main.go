package main

import (
	"fmt"
	"github.com/nkzk/adventofcode/utils"
	"log"
)

func testFunc() {
	test := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	var recurseFn func(in []int)

	recurseFn = func(in []int) {
		if len(in) == 0 {
			return
		}

		fmt.Printf("%v\n", in)
		recurseFn(in[1:])
	}

	recurseFn(test)
}

func main() {
	var input []string
	err := utils.ReadFile("./2024/day7/input", &input)
	if err != nil {
		log.Panicf("failed to read file: %v", err)
	}

	part1 := Part1(input)
	fmt.Printf("part 1: %d\n", part1)

}
