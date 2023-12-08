package main

import "testing"

func TestReadGames(t *testing.T) {
	gameRule := &Bag{
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

func TestReadGamesPartTwo(t *testing.T) {
	endSum := 2286
	var sum int
	var gotSum int
	var tests = []struct {
		name    string
		input   string
		wantSum int
	}{
		{"Game 1 minimum is 4 red, 2 green, 6 blue. Power: 48", "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", 48},
		{"Game 2 minimum is 1 red, 3 green, 4 blue. Power: 12", "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", 12},
		{"Game 3 minimum is 20 red, 13 green, 6 blue. Power: 1560", "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", 1560},
		{"Game 4 minimum is 14 red, 3 green, 15 blue. Power: 630", "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", 630},
		{"Game 5 minimum is 6 red, 3 green, 2 blue. Power: 36", "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", 36},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotSum = ReadGamesPartTwo(tc.input)
			sum += gotSum
			if gotSum != tc.wantSum {
				t.Errorf("want %d, got %d", tc.wantSum, gotSum)
			}
		})
	}
	if endSum != sum {
		t.Errorf("end sum failed, want %d, got %d", endSum, sum)
	}

}
