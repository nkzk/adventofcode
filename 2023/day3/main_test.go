package main

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestReadSchematic(t *testing.T) {

	input := "./test-input"
	file, err := os.Open(input)
	if err != nil {
		t.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var schematic [][]rune

	for scanner.Scan() {
		line := ReadSchematicLine(scanner.Text())
		schematic = append(schematic, line)
	}

	wantSum := 4361
	gotSum := Sum(schematic)

	if wantSum != gotSum {
		t.Errorf("want %d, got %d", wantSum, gotSum)
	}
}
