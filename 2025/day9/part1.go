package main

import (
	"log"
	"math"
	"strconv"
	"strings"
)

type Node struct {
	X int
	Y int
}

type Edge struct {
	x, y Node
	a, b int
}

func Part1(input []string) int {
	var sum int
	nodes := []Node{}
	for _, xy := range input {
		fields := strings.FieldsFunc(xy, func(r rune) bool {
			return string(r) == ","
		})
		x, err := strconv.Atoi(fields[0])
		if err != nil {
			log.Fatalf("failed to convert str to int: %s", fields[0])
		}
		y, err := strconv.Atoi(fields[1])
		if err != nil {
			log.Fatalf("failed to convert str to int: %s", fields[1])
		}

		nodes = append(nodes, Node{
			X: x,
			Y: y,
		})
	}

	edges := []Edge{}
	for _, x := range nodes {
		for _, y := range nodes {
			if x != y {
				edges = append(edges, Edge{
					x: x,
					y: y,
					a: int(math.Abs(float64(x.X)-float64(y.X)) + 1),
					b: int(math.Abs(float64(x.Y)-float64(y.Y)) + 1),
				})
			}
		}
	}

	for _, edge := range edges {
		area := int(edge.a * edge.b)
		if area > sum {
			sum = area
		}
	}
	return sum
}
