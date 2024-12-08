package main

func Part1(input []string) int {
	var sum int

	type position struct {
		x, y int
	}

	direction := map[string]position{
		"up":    {-1, 0},
		"down":  {1, 0},
		"left":  {0, -1},
		"right": {0, 1},
	}

	startingPosition := position{}

	for i, line := range input {
		for j, char := range line {
			switch string(char) {
			case ".":
				continue
			case "#":
				continue
			default:
				startingPosition = position{i, j}
			}
		}
	}

	// 	case "^":
	// case "<":
	// case ">":
	// case "v":

	return sum
}
