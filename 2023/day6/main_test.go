package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestPartOne(t *testing.T) {
	input := "./input"
	file, err := os.Open(input)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	races := make(map[int]int)
	var times []string
	var distances []string
	for scanner.Scan() {
		// var time int
		// var distance int
		line := scanner.Text()
		if strings.Contains(line, "Time") {
			times = strings.Fields(strings.TrimPrefix(line, "Time:"))
		} else if strings.Contains(line, "Distance") {
			distances = strings.Fields(strings.TrimPrefix(line, "Distance:"))
		}
	}

	for i := range times {
		key, err := strconv.Atoi(times[i])
		if err != nil {
			os.Exit(2)
		}
		distance, err := strconv.Atoi(distances[i])
		if err != nil {
			os.Exit(3)
		}
		races[key] = distance
	}

	var tests = []struct {
		time           int
		recordDistance int
		waysToWin      int
	}{
		{7, 9, 4},
		{15, 40, 8},
		{30, 200, 9},
	}
	sum := 1
	for _, tc := range tests {
		t.Run(strconv.Itoa(tc.time), func(t *testing.T) {
			want := tc.waysToWin
			got := waysToWin(tc.time, tc.recordDistance)
			if want != got {
				t.Errorf("want %d, got %d\n", want, got)
			} else {
				sum *= got
			}
		})
	}

	t.Run("Check sum", func(t *testing.T) {
		want := 288
		got := sum
		if got != want {
			t.Errorf("want %d, got %d", want, got)
		}
	})

	// 7 ms, 8 mm
	// button, 1 sec =

	// parse input
	// find 1: least ms needed to beat record
	//      2: max ms possible

	// Q: Number of ways could beat the record in each race.

}

func TestPartTwo(t *testing.T) {
	want := 71503
	got := waysToWin(71530, 940200)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
