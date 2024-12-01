package main

import "fmt"

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

func part2(input Result) int {
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
