package main

import (
	"log"
	"utils/grid"
)

func Part2(input []string) int {
	var sum int
	g := grid.ArrayToGridMap(input)
	// for i := 0; i <= 1000; i++ {
	sum += count2(g)
	// 	fmt.Printf("sum: %d\n", sum)
	// }

	return sum
}

func count2(gridMap map[string]string) int {
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
			gridMap[k] = "x"
			result += count2(gridMap)
		}
	}

	return result
}
