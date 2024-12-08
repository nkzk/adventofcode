package main

import (
	"log"
	"strconv"
	"strings"
)

func Part1(input []string) int {
	order := make(map[int]int)
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

			order[part1] = part2
		} else if strings.Contains(line, ",") {
			updates = append(updates, line)
		}
	}

	for _, update := range updates {
		parts := strings.Split(update, ",")
		for _, part := range parts {
			number, err := strconv.Atoi(part)
			if err != nil {
				log.Panic("failed to convert string to int")
			}
			
		}



	}

	return 0
}
