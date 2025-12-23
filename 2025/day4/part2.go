package main

import (
	"github.com/nkzk/adventofcode/utils/grid"
	"log"
)

func Part2(input []string) int {
	g := grid.ArrayToGridMap(input)
	return count2(g)
}

func count2(gridMap map[string]string) int {
	result := 0

	scanMap := true
	for scanMap {
		scanMap = false

		// scan whole gridMap
		for k, v := range gridMap {
			if v != "@" {
				continue
			}

			x, y, err := grid.DecodeKey(k)
			if err != nil {
				log.Fatalf("decode error: %v", err)
			}

			count := 0
			for _, d := range grid.AllDirections {
				nk := grid.EncodeKey(x+d.X, y+d.Y)
				if gridMap[nk] == "@" {
					count++
				}
			}

			if count < 4 {
				gridMap[k] = "x"
				result++
				// changed the grid so scan again
				scanMap = true
			}
		}
	}

	return result
}
