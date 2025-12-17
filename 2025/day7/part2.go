package main

func Part2(input []string) int {
	var sum int
	grid := toGrid(input)
	y, x := findStart(grid)
	visited := make(map[string]bool)

	x += renderBeam(y+1, x, grid, visited)
	sum = (x ^ 2)
	return sum
}
