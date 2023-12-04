package main

import "testing"


func NewGame(id, red, green, blue int) Game {
	return Game{ID: id, red: red, green: green, blue: blue}
}
func TestReadGames(t *testing.t) {

	var wantGames = map[int]map[int]Game{	
		{ 0: map[int]Game{
			0: NewGame(0,12,13,14),
		},
		{ 1: map[int]Game{
			0: NewGame(1,4,0,3),
			1: NewGame(1,1,2,6),
			2: NewGame(1,0,2,0),
		}, 
		{ 2: map[int]Game{
			0: NewGame(2,0,2,1),
			1: NewGame(2,0,0,0),
			2: NewGame(2,0,0,0),
		},

	},
   }}

	testGame := &Game{}

	var tests = []struct {
		name  string
		input string
		want  Game
	}{
		{"Game 1 is possible", "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", },
		{"Game 2 is possible", "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", },
		{"Game 3 is impossible", "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", },
		{"Game 4 is impossible", "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", },
		{"Game 5 is possible", "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", },
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// 0 - rule, 1 - input-string
			games := ReadGames(test.input)
			if gammes.i
		})
	}
	
}