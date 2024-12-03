package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

// xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))

// regex: mul\([0-9],[0-9]\)
//  fields: mul(2,4) mul(5,5) mul(32,64) mul(11,8) mul(8,5)
// for each field, strconv.atoi, multiply, add up

func Part1(input []string) int {
	var sum int

	for _, line := range input {
		re := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)
		fields := re.FindAllString(line, -1)

		fmt.Printf("found fields: %s\n", fields)

		for _, field := range fields {
			re := regexp.MustCompile(`[0-9]+`)
			numbers := re.FindAllString(field, -1)

			sumNumbers := 1
			for _, number := range numbers {
				num, err := strconv.Atoi(number)
				if err != nil {
					log.Fatalf("failed converting string to int: %v", err)
				}

				sumNumbers *= num
			}

			sum += sumNumbers
		}

	}

	return sum
}
