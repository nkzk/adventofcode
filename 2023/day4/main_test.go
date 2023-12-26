package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

// Solves Part 1, returning points for each match/line. Sum is calculated in main/test function
func PartOne(line string) int {
	numbers := make(map[int]int)
	cardPart := strings.Split(line, ":")
	card := strings.Split(cardPart[1], "|")

	// numbers on scratcher
	for _, str := range strings.Fields(card[0]) {
		num, err := strconv.Atoi(str)
		if err != nil {
			return 0
		}
		numbers[num] = num
	}
	sumMatch := 0
	// check if numbers are a match.
	for _, str := range strings.Fields(card[1]) {
		num, err := strconv.Atoi(str)
		if err != nil {
			return 0
		}
		_, exists := numbers[num]
		if exists {
			if sumMatch == 0 {
				sumMatch = 1
			} else {
				sumMatch += sumMatch
			}
		}
	}
	return sumMatch
}

func TestPartOne(t *testing.T) {

	input := "./test-input"
	file, err := os.Open(input)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sum int
	wantSum := 13
	for scanner.Scan() {
		sum += PartOne(scanner.Text())
	}
	if sum != wantSum {
		t.Errorf("want %d got %d", wantSum, sum)
	}
}

// Returns returns card id, number of matches
func PartTwo(line string) (int, int) {
	numbers := make(map[int]int)
	cardPart := strings.Split(line, ":")
	card := strings.Split(cardPart[1], "|")
	idPart := strings.Fields(cardPart[0])
	id, _ := strconv.Atoi(idPart[1])
	// numbers on scratcher
	for _, str := range strings.Fields(card[0]) {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Printf("FAILED TO CONVERT STRING TO INT: %s", err)
			return id, 0
		}
		numbers[num] = num
	}
	sumMatch := 0
	fmt.Printf("Processing Card %d: ", id)
	// check if numbers are a match.
	for _, str := range strings.Fields(card[1]) {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Printf("FAILED TO CONVERT STRING TO INT: %s", err)
			return id, 0
		}
		_, exists := numbers[num]
		if exists {
			if sumMatch == 0 {
				sumMatch = 1
			} else {
				sumMatch++
			}
		}
	}
	fmt.Printf("| Matches: %d\n", sumMatch)
	return id, sumMatch
}

func TestPartTwo(t *testing.T) {
	// input := "./test-input"
	input := "./input"
	file, err := os.Open(input)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sum int
	wantSum := 30
	var cards []string
	var numCards = make(map[int]int)
	for scanner.Scan() {
		cards = append(cards, scanner.Text())
	}

	// process copies
	for _, card := range cards {
		id, points := PartTwo(card)
		numCards[id] += 1
		for j := 1; j <= points; j++ {
			next := id + j
			if next <= len(cards) {
				numCards[next] += numCards[id]
			}
		}
	}
	sum = 0
	for i := 1; i <= len(cards); i++ {
		sum += numCards[i]
		fmt.Printf("Card %d, %d (sum: %d)\n", i, numCards[i], sum)
	}
	if sum != wantSum {
		t.Errorf("want %d got %d", wantSum, sum)
	}

}
