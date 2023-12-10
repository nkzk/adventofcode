package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func ReadSchematicLine(line string) []rune {
	var runes []rune

	for _, char := range line {
		runes = append(runes, char)
	}
	return runes
}

func IsStar(schematic [][]rune, x int, y int) bool {

	// check y bounds
	rows := len(schematic)
	if y < 0 || y >= rows {
		return false
	}
	// check x bounds
	xLen := len(schematic[y])
	if x < 0 || x >= xLen {
		return false
	}

	// check is Star
	if string(schematic[y][x]) == "*" {
		return true
	}
	return false
}

type pos struct {
	x int
	y int
}

func IsAdjacentToNumber(schematic [][]rune, x int, y int) int {
	rows := len(schematic)
	xLen := len(schematic[y])

	directions := []struct {
		x int
		y int
	}{
		{-1, 0}, {1, 0}, // Check X left/right
		{0, 1}, {0, -1}, // Check Y top/bot
		{-1, -1}, {1, -1}, // Top left/right
		{-1, 1}, {1, 1}, // Bottom left/right
	}
	for _, dir := range directions {
		dx, dy := x+dir.x, y+dir.y

		// check y bounds
		if y < 0 || y >= rows {
			continue
		}

		// check x bounds
		if x < 0 || x >= xLen {
			continue
		}

		if unicode.IsNumber(schematic[dx][dy]) {
			// check left

			// check right
			return true, &pos{x: dx, y: dy}
		}
	}

	return 0
}
func IsAdjacentToStar(schematic [][]rune, x int, y int) (bool, *pos) {

	rows := len(schematic)
	xLen := len(schematic[y])

	directions := []struct {
		x int
		y int
	}{
		{-1, 0}, {1, 0}, // Check X left/right
		{0, 1}, {0, -1}, // Check Y top/bot
		{-1, -1}, {1, -1}, // Top left/right
		{-1, 1}, {1, 1}, // Bottom left/right
	}

	for _, dir := range directions {
		dx, dy := x+dir.x, y+dir.y

		// check y bounds
		if y < 0 || y >= rows {
			continue
		}

		// check x bounds
		if x < 0 || x >= xLen {
			continue
		}

		if IsStar(schematic, dx, dy) {
			return true, &pos{x: dx, y: dy}
		}
	}
	return false, nil
}

func Sum(schematic [][]rune) int {

	var newNumber string
	var sum int
	var adjacent bool
	var num int

	for y, runeLine := range schematic {
		var startX int
		newNumber = ""
		for x, rune := range runeLine {
			if unicode.IsNumber(rune) {

				adj, pos := IsAdjacentToStar(schematic, x, y)

				if adj {
					adjacent = true
					if newNumber == "" {
						startX = x
					}
				}

				newNumber += string(rune)
			} else {

				if newNumber != "" {
					num, _ = strconv.Atoi(newNumber)
				}
				schematic[y][startX] = num
				newNumber = ""
				adjacent = false
				continue
			}
			if adjacent {
				fmt.Printf("helo")
			}
		}
	}
	return sum
}

func AdjacentToGear(schematic [][]rune, x int, y int) []int {
	var numbers []int

	rows := len(schematic)
	xLen := len(schematic[y])

	directions := []struct {
		x int
		y int
	}{
		{-1, 0}, {1, 0}, // Check X left/right
		{0, 1}, {0, -1}, // Check Y top/bot
		{-1, -1}, {1, -1}, // Top left/right
		{-1, 1}, {1, 1}, // Bottom left/right
	}
	for _, dir := range directions {
		x, y := x+dir.x, y+dir.y

		// check y bounds
		if y < 0 || y >= rows {
			continue
		}

		// check x bounds
		if x < 0 || x >= xLen {
			continue
		}

		if string(schematic[x][y]) == "*" {
			// for  _, dir2 := range directions {
			// 	dx, dy := dx + dir2.x, dy+dir.y
			//     // check y bounds
			// 	if dy < 0 || dy >= rows {
			// 		continue
			// 	}

			// 	// check x bounds
			// 	if dx < 0 || dx >= xLen {
			// 		continue
			// 	}
			// }
			numbers = append(numbers, int(schematic[x][y]))
		}
	}
	fmt.Printf("found *, returning adjanent: %v\n", numbers)
	return numbers
}
