package main

// for each X, look left, right, up, topleft, topright, down, bottomleft, bottomright
// 	each direction
// 	look for the letters X M A S in this order, if any letter mismatch, go next, if reaches full word +1

type Direction struct {
	x, y int
}

var directions = []Direction{
	{0, -1},  // left
	{0, 1},   // right
	{-1, 0},  // up
	{-1, -1}, // up left
	{-1, 1},  // up right
	{1, 0},   // down
	{1, -1},  // bottom left
	{1, 1},   // down right
}

func Part1(input Input) int {
	var count int
	word := "XMAS"
	checkWord := func(x, y int, dir Direction) bool {
		for i := 0; i < len(word); i++ {
			newX := x + i*dir.x
			newY := y + i*dir.y
			// Out of bounds or mismatch
			if newX < 0 || newX >= len(input) || newY < 0 || newY >= len(input[0]) || input[newX][newY] != rune(word[i]) {
				return false
			}
		}
		return true
	}

	for i, x := range input {
		for j := range x {
			if string(input[i][j]) == "X" {
				for _, dir := range directions {
					if checkWord(i, j, dir) {
						count++
					}
				}
			}
		}
	}
	return count
}
