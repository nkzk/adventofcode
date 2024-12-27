package main

import (
	"log/slog"
	"utils/grid"
)

func Part2(input []string) int {
	var sum int

	gridmap := grid.ArrayToGridMap(input)

	// key = coordinate, value = direction
	visited := make(map[string][2]int)

	var direction string
	var startingPosition string
	var visitFn func(row, col int)

	visitFn = func(row, col int) {
		x, y, _ := grid.DecodeKey(direction)

		if visited[grid.EncodeKey(row, col)] == [2]int{x, y} {
			slog.Info("loop detected")
			sum += 1
			return
		}

		slog.Info("visiting", "x", row, "y", col)

		if row < 0 || col < 0 || row >= len(input[0]) || col >= len(input) {
			return
		} else {
			x, y, _ := grid.DecodeKey(direction)
			visited[grid.EncodeKey(row, col)] = [2]int{x, y}

			// set next location based on direction
			dx := row + x
			dy := col + y

			// check if next is an obstruction
			if gridmap[grid.EncodeKey(dx, dy)] == "#" {
				slog.Info("next is obstruction")
				// change direction 90 degrees
				x, y, _ := grid.DecodeKey(direction)
				direction = grid.EncodeKey(y, -x)

				// overwrite next location based on new direction
				dx = row + y
				dy = col + -x

				slog.Info("changed direction and visiting.",
					"obstructionLocation", grid.EncodeKey(dx, dy),
					"direction", direction)
			}

			visitFn(dx, dy)
		}

	}

	for key, node := range gridmap {
		switch node {
		case ".":
			continue
		case "#":
			continue
		default:
			direction = directionMap[node]
			startingPosition = key
		}
	}

	slog.Info("part 1", "starting position", startingPosition)

	row, col, _ := grid.DecodeKey(startingPosition)

	visitFn(row, col)
	return sum
}
