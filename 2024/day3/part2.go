package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func Part2(input []string) int {
	var sum int
	do := true
	for _, line := range input {
		re := regexp.MustCompile(`(mul\([0-9]{1,3},[0-9]{1,3}\))|(do\(\))|(don't\(\))`)
		fields := re.FindAllString(line, -1)
		fmt.Printf("found fields: %s\n", fields)

		for _, field := range fields {
			switch field {
			case "do()":
				do = true
			case "don't()":
				do = false
			default:
				if do {
					re := regexp.MustCompile(`[0-9]+`)
					numbers := re.FindAllString(field, -1)

					num1, err := strconv.Atoi(numbers[0])
					num2, err := strconv.Atoi(numbers[1])
					if err != nil {
						continue
					}

					sum += num1 * num2
				}
			}
		}
	}

	return sum
}
