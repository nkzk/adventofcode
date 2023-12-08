package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Bag struct {
	red   int
	green int
	blue  int
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

	bag := &Bag{
		red:   12,
		green: 13,
		blue:  14,
	}

	var sum int
	var sumPartTwo int
	for scanner.Scan() {
		sum += ReadGames(bag, scanner.Text())
		sumPartTwo += ReadGamesPartTwo(scanner.Text())
	}
	fmt.Printf("Sum part 1: %d\n", sum)
	fmt.Printf("Sum part 2: %d\n", sumPartTwo)
}

func ReadGames(bag *Bag, line string) int {

	gameData := strings.Split(line, ":")
	id, _ := strconv.Atoi(strings.Split(gameData[0], " ")[1])
	gameRounds := strings.Split(gameData[1], ";")

	for _, round := range gameRounds {
		newGame := &Bag{}
		set := strings.Split(round, ",")
		for _, cube := range set {
			parts := strings.Split(cube, " ")
			amount, _ := strconv.Atoi(parts[1])
			color := parts[2]
			switch color {
			case "red":
				newGame.red += amount
			case "green":
				newGame.green += amount
			case "blue":
				newGame.blue += amount
			}
		}
		if newGame.red > bag.red || newGame.green > bag.green || newGame.blue > bag.blue {
			return 0
		}
	}
	return id
}

func ReadGamesPartTwo(line string) int {
	power := 0
	gameData := strings.Split(line, ":")
	gameRounds := strings.Split(gameData[1], ";")
	minBag := &Bag{}
	for _, round := range gameRounds {

		set := strings.Split(round, ",")
		for _, cube := range set {
			parts := strings.Split(cube, " ")
			amount, _ := strconv.Atoi(parts[1])
			color := parts[2]
			switch color {
			case "red":
				if amount >= minBag.red {
					minBag.red = amount
				}
			case "green":
				if amount >= minBag.green {
					minBag.green = amount
				}
			case "blue":
				if amount >= minBag.blue {
					minBag.blue = amount
				}
			}
		}
	}
	power = minBag.red * minBag.green * minBag.blue
	return power
}
