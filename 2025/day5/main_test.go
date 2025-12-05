package main

import (
	"log"
	"testing"

	"github.com/nkzk/adventofcode/utils"
)

func Test(t *testing.T) {
	var input []string
	err := utils.ReadFile("test-input", &input)
	if err != nil {
		log.Panicf("failed to read file: %v", err)
	}

	t.Run("Part 1 ", func(t *testing.T) {
		got := Part1_2(input)
		want := 3
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("Part 2 ", func(t *testing.T) {
		got := Part2(input)
		want := 14
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("Part 2_1", func(t *testing.T) {
		got := Part2([]string{
			"200-300",
			"100-101",
			"1-1",
			"2-2",
			"3-3",
			"1-3",
			"1-3",
			"2-2",
			"50-70",
			"10-10",
			"98-99",
			"99-99",
			"99-99",
			"99-100",
			"1-1",
			"2-1",
			"100-100",
			"100-100",
			"100-101",
			"200-300",
			"201-300",
			"202-300",
			"250-251",
			"98-99",
			"100-100",
			"100-101",
			"1-101",
		})
		want := 202
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

}
