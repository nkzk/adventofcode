package main

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

func Part2(input []string) int {
	var sum int

	orderMap := make(map[int]map[int]bool, 0)

	updates := make([]string, 0)

	for _, line := range input {
		if strings.Contains(line, "|") {
			parts := strings.Split(line, "|")
			part1, err := strconv.Atoi(parts[0])
			if err != nil {
				log.Panicf("failed to convert string to int")
			}

			part2, err := strconv.Atoi(parts[1])
			if err != nil {
				log.Panic("failed to convert string to int")
			}

			if _, exists := orderMap[part1]; !exists {
				orderMap[part1] = make(map[int]bool)
			}

			orderMap[part1][part2] = true
		} else if strings.Contains(line, ",") {
			updates = append(updates, line)
		}
	}

	fmt.Printf("updates: %+v\n", orderMap)

	updatesInt := make([][]int, len(updates))

	for i, update := range updates {
		parts := strings.Split(update, ",")
		for _, part := range parts {
			number, err := strconv.Atoi(part)
			if err != nil {
				log.Panic("failed to convert string to int")
			}

			updatesInt[i] = append(updatesInt[i], number)
		}
	}

	for _, update := range updatesInt {
		order := true
		for i := range update {
			for j := i + 1; j <= len(update)-1; j++ {
				if _, exists := orderMap[update[i]][update[j]]; !exists {
					order = false
				}
			}
		}

		if order == false {
			// sort -> add
			slices.SortFunc(update, func(a, b int) int {
				// if a < b
				for x := range orderMap[a] {
					if b == x {
						return -1
					}
				}
				return 1
			})

			// add
			length := len(update) - 1
			fmt.Printf("%+v\n", update)
			fmt.Printf("adding: %d\n", update[length/2])
			sum += update[length/2]
		}
	}

	return sum
}
