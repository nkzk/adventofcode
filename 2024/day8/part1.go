package main

import (
	"fmt"
	"log/slog"

	"github.com/nkzk/adventofcode/utils/grid"
)

// returns slope (x,y) and error
func GetSlope(a, b string) (int, int, error) {
	ax, ay, err := grid.DecodeKey(a)
	if err != nil {
		return 0, 0, err
	}

	bx, by, err := grid.DecodeKey(b)
	if err != nil {
		return 0, 0, err
	}

	return bx - ax, by - ay, nil
}

func inBounds(x, y int, rowLength, columnLength int) bool {
	return x > rowLength && x >= 0 && y > columnLength && y > 0
}

func Part1(input []string) int {
	var sum int // total unique antinode locations

	// set up stuff
	cityGrid := grid.ArrayToGridMap(input)
	charPositions := make(map[string][]string)
	antiNodeMap := make(map[string]bool)

	// map characters and their positions
	for position, v := range cityGrid {
		if v != "." {
			if len(charPositions[v]) == 0 {
				charPositions[v] = []string{position}
			} else {
				charPositions[v] = append(charPositions[v], position)
			}
		}

	}

	// set bounds
	maxRow := len(input[0])
	maxColumn := len(input)

	fmt.Printf("maxRow: %d\nmaxColum: %d\n", maxRow, maxColumn)

	for key, positions := range charPositions {
		for i := 0; i < len(positions); i++ {
			fmt.Printf("\n\nmapping anti nodes for character '%s' on positions: %s\n", key, positions)
			MapAntiNodes(positions[i], positions, antiNodeMap, maxRow, maxColumn)
		}
	}

	// count unique occurences of antinodes in the antinode map
	for k := range antiNodeMap {
		if cityGrid[k] == "." {
			cityGrid[k] = "#"
		}
		sum += 1

	}

	// print the grid for fun
	grid.PrintGridMap(cityGrid, len(input[0]), len(input))

	// 265 too high
	// 259 too high
	return sum
}

func MapAntiNodes(startingPosition string, positions []string, antiNodeMap map[string]bool, maxRow, maxColumn int) {
	fmt.Printf("MapAntiNodes: startingposition: %s, positions, %+v\n", startingPosition, positions)

	x, y, err := grid.DecodeKey(startingPosition)
	if err != nil {
		slog.Error("failed to decode key", "err", err)
		return
	}

	for _, pos := range positions {
		dx, dy, err := GetSlope(startingPosition, pos)
		if err != nil {
			slog.Error("failed to get slope", "err", err)
			return
		}

		ax := x - dx
		ay := y - dy

		if ax >= 0 && ax < maxRow && ay >= 0 && ay < maxColumn && ax != x && ay != y {
			// fmt.Printf("New Antinode for %s: %d,%d\n", pos, ax, ay)
			fmt.Printf("for %s/%s, found: %d,%d\n", startingPosition, pos, ax, ay)
			antiNodeMap[grid.EncodeKey(ax, ay)] = true
		}
	}

	// MapAntiNodes(positions[0], positions[1:], antiNodeMap, maxRow, maxColumn)
}
