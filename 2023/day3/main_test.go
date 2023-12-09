package main

import (
	"bufio"
	"os"
	"testing"
)

func TestReadSchematic(t *testing.T) {

	var schematic_test = [][]bool{}
	input := "./test-input"
	file, err := os.Open(input)
	if err != nil {
		t.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		ReadSchematic(schematic_test, scanner.Text())
	}

	// test schematic
}
