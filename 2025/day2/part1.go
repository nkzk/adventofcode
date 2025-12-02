package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func Part1(input []string) int {
	var sum int

	// get ranges, comma-separated
	// each range: start-end
	// check for invalid id's:
	// invalid: any ID which is made only of some sequence of digits repeated twice
	// 55 (5 twice)
	// 6464 (64 twice)
	// 123123 (123 twice)

	// find the invalid IDs, add them to sum

	// think i only need to check numbers with even
	//   number of digits

	// cut id in half
	// if 1st half and 2nd half is the same and  == id
	// invalid!

	allRanges := input[0]

	ranges := strings.Split(allRanges, ",")

	for _, r := range ranges {
		start, end, ok := strings.Cut(r, "-")
		if !ok {
			log.Printf("error: couldnt cut %s. no '-'", r)
			os.Exit(1)
		}
		startNumber, err := strconv.Atoi(start)
		endNumber, err := strconv.Atoi(end)
		if err != nil {
			log.Printf("failed: %v", err)
			os.Exit(2)
		}

		for i := startNumber; i <= endNumber; i++ {
			if len(strconv.Itoa(i))%2 != 0 {
				continue
			}

			number := strconv.Itoa(i)
			firstHalf := number[:len(number)/2]
			secondHalf := number[len(number)/2:]

			if firstHalf != secondHalf {
				continue
			}

			sum += i
		}
	}

	return sum
}
