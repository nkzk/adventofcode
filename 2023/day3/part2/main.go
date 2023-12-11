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
			var startX int
			var endX int

			// check left
			i := dx
			for unicode.IsNumber(schematic[i][dy]) {
				startX = i
				i--
			}
			// check right
			j := dx
			for unicode.IsNumber(schematic[j][dy]) {
				endX = j
				j++
			}

			var number int
			for _, rune := range schematic[dy][startX:endX] {
				number += int(rune)
			}

			return number
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

	var ratioNum []int
	for y, runeLine := range schematic {
		newNumber = ""

		for x, rune := range runeLine {
			alreadyFound := false
			if unicode.IsNumber(rune) {
				if !alreadyFound {
					found, pos := IsAdjacentToStar(schematic, x, y)
					fmt.Printf("%s is adjacent to star\n", string(rune))
					if found {
						fmt.Printf("checking %d, %d\n", x, y)
						adjToStarNumber := IsAdjacentToNumber(schematic, pos.x, pos.y)
						if adjToStarNumber != 0 {
							alreadyFound = true
							ratioNum = append(ratioNum, adjToStarNumber)
						}
					}
					newNumber += string(rune)
				}
			} else {
				if newNumber != "" {
					num, _ = strconv.Atoi(newNumber)
				}
				ratioNum = append(ratioNum, num)
				newNumber = ""
				adjacent = false
				continue
			}
			if adjacent {
				fmt.Printf("helo")
			}
		}
	}
	fmt.Printf("rationum:%v\n", ratioNum)

	return sum
}
