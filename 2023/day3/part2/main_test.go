package main

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestSumPartTwo(t *testing.T) {

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

	// wantPart1 := 16345
	// wantPart2 := 46835

	wantSum := 467835
	gotSum := Sum(schematic)
	if gotSum != wantSum {
		t.Errorf("got %d, want %d", gotSum, wantSum)
	}

}
