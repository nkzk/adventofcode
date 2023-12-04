package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
		sum += ReadGames(gameRule, scanner.Text())
	}
	fmt.Printf("Sum: %d", sum)
}

func ReadGames(gameRule *Game, line string) int {

	gameData := strings.Split(line, ":")
	id, err := strconv.Atoi(strings.Split(gameData[0], " ")[1])
	if err != nil {
		return 0
	}
	gameRound := strings.Split(gameData[1], ";")

	// number color
	type Cube struct {
		amount string
		color  string
	}

	numberColors := strings.Split(gameRound[0], ",")

	cubes := []Cube{}

	for _, numberColor := range numberColors {
		cubeStats := strings.Split(numberColor, " ")
		cubes = append(cubes, Cube{amount: cubeStats[0], color: cubeStats[1]})
	}
	fmt.Println(cubes)

	return id
}
