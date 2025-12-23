package main

import (
	"slices"
)

func Part2(input []string) int {
	var sum int
	dsu := New()
	coordinates := ParseInput(input)
	for _, c := range coordinates {
		dsu.Add(c)
	}

	// create all edges
	edges := []Edge{}
	for i := 0; i < len(coordinates)-1; i++ {
		for j := i + 1; j < len(coordinates); j++ {
			edges = append(edges, Edge{
				x:        coordinates[i],
				y:        coordinates[j],
				distance: Distance(coordinates[i], coordinates[j]),
			})
		}
	}

	// sort by distance
	slices.SortFunc(edges, func(a, b Edge) int {
		if a.distance < b.distance {
			return -1
		}
		if a.distance > b.distance {
			return 1
		}
		return 0
	})

	jumpboxes := len(coordinates)
	for _, e := range edges {
		if dsu.Union(e.x, e.y) {
			jumpboxes--
			if jumpboxes == 1 {
				return e.x.X * e.y.X
			}
		}
	}
	return sum
}
