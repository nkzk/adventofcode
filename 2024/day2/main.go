package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"utils"
)

type Input [][]int

func lineProcessor(line string, result *Input) error {
	var report []int

	for _, field := range strings.Fields(line) {
		num, err := strconv.Atoi(field)
		if err != nil {
			return fmt.Errorf("failed to convert string to int: %w", err)
		}

		report = append(report, num)
	}

	(*result) = append((*result), report)

	return nil
}

func main() {
	input := make(Input, 0)
	err := utils.ReadFile_old[Input]("./2024/day2/input.txt", lineProcessor, &input)
	if err != nil {
		log.Panicf("failed to read file: %v", err)
	}

	count, err := CountSafeReports(input)
	if err != nil {
		log.Panicf("failed to count reports: %v", err)
	}

	fmt.Printf("part 1: %d\n", count)

	var sum int
	for _, report := range input {
		fmt.Printf("sending in report: %+v\n", report)
		_, safe := IsReportSafe(report)
		if safe {
			sum += 1
		} else {
			dampener := 0
			for i := range report {
				newReport := remove(report, i)
				fmt.Printf("testing new report: %+v\n", newReport)
				_, safe := IsReportSafe(newReport)
				if safe && dampener < 1 {
					dampener += 1
					sum += 1
				}
			}
		}
	}
	fmt.Printf("part 2: %d", sum)

}
