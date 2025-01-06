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

func PrintGridMap(grid map[string]string, x, y int) {
	fmt.Printf("   ")
	for i := 0; i < x; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\n")
	for i := 0; i < x; i++ {
		if i < 10 {
			fmt.Printf("%d  ", i)
		} else {
			fmt.Printf("%d ", i)
		}
		for j := 0; j < y; j++ {
			if char, ok := grid[EncodeKey(i, j)]; ok {
				fmt.Printf("%s ", char)
			} else {
				fmt.Printf(". ")
			}
		}

		fmt.Printf("\n")

	}
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
