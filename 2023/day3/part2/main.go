package main

import (
	"bytes"
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

func SolvePartTwo(schematic [][]rune) int {
	var sum int
	var numbers []int
	visited := make([][]bool, len(schematic)) // to avoid counting adjacent digits (part of same number) twice
	for i := range visited {
		visited[i] = make([]bool, len(schematic[i]))
	}

	for y, line := range schematic {
		var count int
		rows := len(schematic)
		xLen := len(schematic[y])
		for x, _ := range line {
			if IsStar(schematic, x, y) {
				count = 0
				numbers = nil
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
					if !visited[dy][dx] {
						if y < 0 || y >= rows {
							continue
						}
						if x < 0 || x >= xLen {
							continue
						}
						if unicode.IsNumber(schematic[dy][dx]) {
							visited[dy][dx] = true
							number := make([]int, 0)
							d := int(schematic[dy][dx]) - '0'
							number = append(number, d)
							i := dx
							for i > 0 && unicode.IsNumber(schematic[dy][i-1]) {
								i--
								visited[dy][i] = true
								digit := int(schematic[dy][i] - '0')
								number = append([]int{digit}, number...)
							}
							j := dx
							for j < len(schematic[dy])-1 && unicode.IsNumber(schematic[dy][j+1]) {
								j++
								visited[dy][j] = true
								digit := int(schematic[dy][j]) - '0'
								number = append(number, digit)
							}
							var buf bytes.Buffer
							for _, digit := range number {
								buf.WriteString(fmt.Sprintf("%d", digit))
							}
							n, err := strconv.Atoi(buf.String())
							if err != nil {
								return 0
							}
							numbers = append(numbers, n)
							count++
						}
					}
				}
				if count == 2 {
					sum += numbers[0] * numbers[1]
				}
			}
		}
	}
	return sum
}
