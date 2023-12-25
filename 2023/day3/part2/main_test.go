package main

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestSolvePartTwo(t *testing.T) {

	// read from file
	input := "./test-input"
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

	wantSum := 467835
	gotSum := SolvePartTwo(schematic)
	if gotSum != wantSum {
		t.Errorf("got %d, want %d", gotSum, wantSum)
	}
}
