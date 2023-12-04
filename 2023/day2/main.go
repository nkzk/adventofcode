package main

import (
	"bufio"
	"fmt"
	"os"
)

type Game struct {
	id       int
	red      int
	green    int
	blue     int
	possible bool
}

func main() {
	input := "./input"
	file, err := os.Open(input)
	if err != nil {
		fmt.Printf("Failed to open file")
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	gameRule := &Game{
		id:    0,
		red:   12,
		green: 13,
		blue:  14,
	}

	var sum int

	for scanner.Scan() {
		sum += ReadGames(scanner.Text())
	}
	fmt.Printf("Sum: %d", sum)
}

func ReadGames(line string) int {

	return 0
}
