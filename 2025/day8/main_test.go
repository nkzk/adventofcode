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
		got := Part1(input, 10)
		want := 40
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

    	t.Run("Part 2 ", func(t *testing.T) {
		got := Part2(input)
		want := 0
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

}
