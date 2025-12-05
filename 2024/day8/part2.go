package main

import (
	"fmt"

	"github.com/nkzk/adventofcode/utils/grid"
)

func Part2(input []string) int {
	var sum int
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
			sum += 1
			fmt.Printf("\n\nmapping anti nodes for character '%s' on positions: %s\n", key, positions)
			MapAllAntiNodes(positions[i], positions, antiNodeMap, maxRow, maxColumn)
		}
	}

	for k := range antiNodeMap {
		if cityGrid[k] == "." {
			sum += 1
			cityGrid[k] = "#"
		}

	}

	grid.PrintGridMap(cityGrid, len(input[0]), len(input))
	return sum
}
func MapAllAntiNodes(startingPosition string, positions []string, antiNodeMap map[string]bool, maxRow, maxColumn int) {
	x, y, err := grid.DecodeKey(startingPosition)
	if err != nil {
		return
	}

	for _, pos := range positions {
		dx, dy, err := GetSlope(startingPosition, pos)
		if err != nil {
			return
		}

		ax := x - dx
		ay := y - dy
		for {
			if ax >= 0 && ax < maxRow && ay >= 0 && ay < maxColumn && ax != x && ay != y {
				fmt.Printf("for %s/%s, found: %d,%d\n", startingPosition, pos, ax, ay)

				antiNodeMap[grid.EncodeKey(ax, ay)] = true

				ax = ax - dx
				ay = ay - dy
			} else {
				break
			}

		}
	}
}
