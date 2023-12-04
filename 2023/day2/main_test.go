package main

import "testing"

func NewGame(id, red, green, blue int) Game {
	return Game{id: id, red: red, green: green, blue: blue}
}
func TestReadGames(t *testing.T) {

	// var wantGames = map[int]map[int]Game{
	// 	0: {
	// 		0: NewGame(0, 12, 13, 14),
	// 	},
	// 	1: {
	// 		0: NewGame(1, 4, 0, 3),
	// 		1: NewGame(1, 1, 2, 6),
	// 		2: NewGame(1, 0, 2, 0),
	// 	},
	// 	2: {
	// 		0: NewGame(2, 0, 2, 1),
	// 		1: NewGame(2, 0, 0, 0),
	// 		2: NewGame(2, 0, 1, 1),
	// 	},
	// 	3: {
	// 		0: NewGame(3, 20, 8, 6),
	// 		1: NewGame(3, 4, 13, 5),
	// 		2: NewGame(3, 1, 5, 0),
	// 	},
	// 	4: {
	// 		0: NewGame(4, 3, 1, 6),
	// 		1: NewGame(4, 6, 3, 0),
	// 		2: NewGame(4, 14, 3, 15),
	// 	},
	// 	5: {
	// 		0: NewGame(5, 6, 3, 1),
	// 		1: NewGame(5, 1, 2, 2),
	// 	},
	// }

	gameRule := &Game{
		id:    0,
		red:   12,
		green: 13,
		blue:  14,
	}

	var tests = []struct {
		name  string
		input string
	}{
		{"Game 1 is possible", "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"},
		{"Game 2 is possible", "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue"},
		{"Game 3 is impossible", "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"},
		{"Game 4 is impossible", "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red"},
		{"Game 5 is possible", "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"},
	}

	var sum int

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			sum += ReadGames(gameRule, tc.input)
		})
	}
	if sum != 8 {
		t.Errorf("want 8, got %d", sum)
	}

}
