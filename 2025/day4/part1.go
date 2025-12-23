package main

import (
	"github.com/nkzk/adventofcode/utils/grid"
	"log"
)

func Part1(input []string) int {
	g := grid.ArrayToGridMap(input)
	return count(g)
}

func count(gridMap map[string]string) int {
	result := 0
	for k, v := range gridMap {
		if v != "@" {
			continue
		}

		count := 0
		x, y, err := grid.DecodeKey(k)

		for _, d := range grid.AllDirections {
			nx, ny := x+d.X, y+d.Y
			nk := grid.EncodeKey(nx, ny)

			if err != nil {
				log.Fatalf("failed to decode key %s: %v", k, err)
			}

			nv, ok := gridMap[nk]
			if !ok {
				continue
			}

			if nv == "@" {
				count++
			}
		}

		if count < 4 {
			result++
		}
	}

	return result
}
