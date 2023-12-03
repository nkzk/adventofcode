package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var words = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

var maxLenWords = func() int {
	maxLen := 0
	for word, _ := range words {
		if len(word) > maxLen {
			maxLen = len(word)
		}
	}
	return maxLen
}()

var minLenWords = func() int {
	minLen := maxLenWords
	for word, _ := range words {
		if len(word) < minLen {
			minLen = len(word)
		}
	}
	return minLen
}()

func WordToNumber(word string) (int, error) {
	if !(len(word) >= minLenWords) {
		return 0, errors.New("string too short to contain numbers")
	}

	for key, value := range words {
		if len(word) >= len(key) && key == word[:len(key)] {
			return value, nil
		}
	}
	return 0, errors.New("could not find number")
}

func ConcatTwoIntegers(a int, b int) int {
	// https://www.tutorialspoint.com/python-program-to-concatenate-two-integer-values-into-one
	var number int
	multiplier := 1
	for multiplier <= b {
		multiplier *= 10
	}
	number = a*multiplier + b
	return number
}

func ConvertToInt(char string) (int, error) {
	number, err := strconv.Atoi(char)
	if err != nil {
		return number, err
	}
	return number, nil
}

func GetCalibration(line string) int {
	var first, last int
	var foundFirst bool

	for i, char := range strings.Split(line, "") {
		num, err := ConvertToInt(char)
		if err == nil {
			if !foundFirst {
				foundFirst = true
				first = num
				last = num
			} else {
				last = num
			}
		} else {
			num, err := WordToNumber(line[i:])
			if err == nil {
				if !foundFirst {
					foundFirst = true
					first = num
					last = num
				} else {
					last = num
				}
			}
		}
	}
	return ConcatTwoIntegers(first, last)
}

func Sum(lines []int) int {
	var sum int
	for _, num := range lines {
		sum += num
	}
	return sum
}

func main() {
	input := "./input"
	file, err := os.Open(input)
	if err != nil {
		fmt.Printf("Failed to open file")
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var numbers []int
	for scanner.Scan() {
		numbers = append(numbers, GetCalibration(scanner.Text()))
	}

	sum := Sum(numbers)
	fmt.Println(sum)
}
