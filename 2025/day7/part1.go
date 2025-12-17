package main

import (
	"fmt"
	"log"
	"slices"
)

func Part1(input []string) int {
	var sum int
	grid := toGrid(input)
	print(grid)
	fmt.Println()

	y, x := findStart(grid)

	visited := make(map[string]bool)

	sum += renderBeam(y+1, x, grid, visited)

	return sum
}

func toGrid(input []string) [][]string {
	grid := make([][]string, 0, len(input)-1)
	for _, line := range input {
		x := make([]string, 0, len(input[0])-1)
		for _, char := range line {
			x = append(x, string(char))
		}
		grid = append(grid, x)
	}

	return grid
}

func print(grid [][]string) {
	for i, line := range grid {
		fmt.Printf("%-3d  %s\n", i, line)
	}
}

func findStart(grid [][]string) (y, x int) {
	i := slices.Index(grid[0], "S")
	if i == -1 {
		log.Fatalf("failed to find start")
	}
	return 0, i
}

func renderBeam(y, x int, input [][]string, visited map[string]bool) int {
	result := 0
	if x >= len(input[0])-1 || y >= len(input)-1 || y < 0 {
		return 0
	}

	if !visited[fmt.Sprintf("%d,%d", x, y)] {
		visited[fmt.Sprintf("%d,%d", x, y)] = true
	} else {
		return result
	}

	if input[y][x] == "." {
		input[y][x] = "|"
	}

	if input[y][x] == "^" {
		result++
		result += renderBeam(y+1, x+1, input, visited)
		result += renderBeam(y+1, x-1, input, visited)
		return result

	}

	result += renderBeam(y+1, x, input, visited)

	// fmt.Println()
	// print(input)

	return result
}
