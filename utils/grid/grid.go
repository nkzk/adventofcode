package grid

import (
	"fmt"
)

// all directions, including diagonal
var AllDirections = []struct{ dx, dy int }{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	{-1, -1}, {-1, 1}, {1, -1}, {1, 1},
}

// top, bottom, left, right
var Directions = []struct{ dx, dy int }{
	{-1, 0}, {0, 1}, {1, 0}, {0, -1},
}

func EncodeKey(row, col int) string {
	return fmt.Sprintf("%d,%d", row, col)
}

func DecodeKey(key string) (int, int, error) {
	var row, col int
	_, err := fmt.Sscanf(key, "%d,%d", &row, &col)
	if err != nil {
		return 0, 0, err
	}
	return row, col, nil
}

func ArrayToGridMap(grid []string) map[string]string {
	hashmap := make(map[string]string)
	for row := range grid {
		for col, char := range grid[row] {
			key := EncodeKey(row, col)
			hashmap[key] = string(char)
		}
	}

	return hashmap
}

// func Bfs(grid map[string]string, startKey string) {
// 	startRow, startCol := DecodeKey(startKey)

// 	// visited grid
// 	visited := make(map[string]bool)

// 	type node struct {
// 		x, y int
// 	}

// 	queue := queue.LIFO[node]{}
// }
