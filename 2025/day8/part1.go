package main

import (
	"fmt"
	"log"
	"math"
	"slices"
	"strconv"
	"strings"
)

func Part1(input []string, pairs int) int {
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

	sum += shortestUnion(edges, dsu, pairs)

	return sum
}

func shortestUnion(edges []Edge, dsu *DSU, numberOfPairs int) int {
	for i := 0; i < numberOfPairs; i++ {
		dsu.Union(edges[i].x, edges[i].y)
	}

	sizes := []int{}
	seen := map[*Node]bool{}

	for _, x := range dsu.nodes {
		root := dsu.Find(x.Coordinates)
		if root == nil || seen[root] {
			continue
		}
		seen[root] = true
		sizes = append(sizes, root.Size)
	}

	slices.SortFunc(sizes, func(a, b int) int {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	})

	sum := 1
	for _, size := range sizes[len(sizes)-3:] {
		sum *= size
		fmt.Printf("adding %d\n", size)
	}

	return sum
}

type Edge struct {
	x, y     Coordinates
	distance float64
}

type Coordinates struct {
	X int
	Y int
	Z int
}

type Node struct {
	Size   int // just number of descendants
	Parent *Node
	Coordinates
}

type DSU struct {
	nodes map[Coordinates]*Node
}

func New() *DSU {
	return &DSU{
		nodes: make(map[Coordinates]*Node, 0),
	}
}

func (d *DSU) Add(c Coordinates) {
	_, ok := d.nodes[c]
	if !ok {
		d.nodes[c] = &Node{
			Parent:      nil,
			Coordinates: c,
			Size:        1,
		}
	}
}

func (d *DSU) Find(c Coordinates) *Node {
	node, ok := d.nodes[c]
	if !ok {
		return nil
	}

	root := node

	for root.Parent != nil {
		root = root.Parent
	}

	return root
}

// https://en.wikipedia.org/wiki/Disjoint-set_data_structure
// "union by size"
func (d *DSU) Union(a, b Coordinates) bool {
	x := d.Find(a)
	y := d.Find(b)

	// not found or already in same set
	if x == nil || y == nil || x == y {
		return false
	}

	// make x the new root
	y.Parent = x

	// update size
	x.Size = x.Size + y.Size

	return true
}

// https://en.wikipedia.org/wiki/Euclidean_distance
func Distance(p Coordinates, q Coordinates) float64 {
	return math.Sqrt(
		math.Pow(float64(q.X)-float64(p.X), 2) +
			math.Pow(float64(q.Y)-float64(p.Y), 2) +
			math.Pow(float64(q.Z)-float64(p.Z), 2))
}

func ParseInput(input []string) []Coordinates {
	result := []Coordinates{}
	for _, line := range input {
		fields := strings.FieldsFunc(line, func(r rune) bool {
			return string(r) == ","
		})

		x, err := strconv.Atoi(fields[0])
		if err != nil {
			log.Fatalf("failed to convert %s to x", fields[0])
		}
		y, err := strconv.Atoi(fields[1])
		if err != nil {
			log.Fatalf("failed to convert %s to y", fields[1])
		}
		z, err := strconv.Atoi(fields[2])
		if err != nil {
			log.Fatalf("failed to convert %s to z", fields[2])
		}
		result = append(result, Coordinates{
			X: x,
			Y: y,
			Z: z,
		})
	}
	return result
}
