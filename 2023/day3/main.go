package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func ReadSchematic(schematic [][]bool, line string) {
	var number []int
	for _, char := range line {
		var x, y int
		var newNumber int
		for unicode.IsNumber(char) {
			fmt.Printf("char: %s\n", string(char))
		}

		number = append(number, newNumber)

	}
	// for each line

	// for each character
	// if number, how many digits? whats the number
	// position in line, x, y, z

	// if line-1 og line+1 position x,y,z,x-1,z+1  != symbol = no number adjacent

}

func main() {
	var schematic = [][]bool{}
	input := "./input"
	file, err := os.Open(input)
	if err != nil {
		fmt.Printf("Failed to open file")
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		ReadSchematic(schematic, scanner.Text())
	}
}
