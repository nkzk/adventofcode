package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func ReadSchematicLine(line string) []rune {
	var runes []rune

	for _, char := range line {
		runes = append(runes, char)
	}
	return runes
}
func IsSymbol(schematic [][]rune, x int, y int) bool {

	// check y bounds
	rows := len(schematic)
	if y < 0 || y >= rows {
		return false
	}
	// check x bounds
	xLen := len(schematic[y])
	if x < 0 || x >= xLen {
		return false
	}

	// check is symbol
	if !unicode.IsNumber(schematic[y][x]) && string(schematic[y][x]) != "." {
		return true
	}
	return false
}

func IsAdjacentToSymbol(schematic [][]rune, x int, y int) bool {

	// check x
	if IsSymbol(schematic, x-1, y) || IsSymbol(schematic, x+1, y) {
		return true
	}

	// check y
	if IsSymbol(schematic, x, y-1) || IsSymbol(schematic, x, y+1) {
		return true
	}

	// check diagonals
	if IsSymbol(schematic, x-1, y-1) || IsSymbol(schematic, x+1, y-1) || IsSymbol(schematic, x-1, y+1) || IsSymbol(schematic, x+1, y+1) {
		return true
	}
	return false
}

func Sum(schematic [][]rune) int {
	var newNumber string
	var sum int
	var adjacent bool
	var addNumber int
	for y, runeLine := range schematic {
		for x, rune := range runeLine {
			if unicode.IsNumber(rune) {
				if IsAdjacentToSymbol(schematic, x, y) {
					adjacent = true
				}
				newNumber += string(rune)
			} else {
				if len(newNumber) > 0 {
					addNumber, _ = strconv.Atoi(newNumber)
					if adjacent {
						sum += addNumber
					}
				}
				newNumber = ""
				adjacent = false
				continue
			}
		}
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

	var schematic [][]rune

	for scanner.Scan() {
		line := ReadSchematicLine(scanner.Text())
		schematic = append(schematic, line)
	}

	fmt.Println(Sum(schematic))
}
