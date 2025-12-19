package arraygrid

import (
	"os"
	"testing"
)

func TestGrid(t *testing.T) {
	input := []string{
		".......S.......",
		"...............",
		".......^.......",
		"...............",
		"......^.^......",
		"...............",
		".....^.^.^.....",
		"...............",
		"....^.^...^....",
		"...............",
		"...^.^...^.^...",
		"...............",
		"..^...^.....^..",
		"...............",
		".^.^.^.^.^...^.",
		"...............",
	}

	grid := New(input)
	grid.Print(os.Stdout)

	v := grid.Get(0, 7)
	if v != "S" {
		t.Fatalf("expected S, got %s", v)
	}

	grid.Set(2, 1, "X")
	if v := grid.Get(2, 1); v != "X" {
		t.Fatalf("failed to set X")
	}
}
