package main

import (
	"log"
	"testing"
	"utils"
)

func Test(t *testing.T) {
	var input []string
	err := utils.ReadFile("test-input", &input)
	if err != nil {
		log.Panicf("failed to read file: %v", err)
	}

	t.Run("Part 1 ", func(t *testing.T) {
		got := Part1(input)
		want := 41
		if got != want {
			t.Fatalf("got %d want %d", got, want)
		}
	})

	t.Run("Part 2 ", func(t *testing.T) {
		got := Part2(input)
		want := 0
		if got != want {
			t.Fatalf("got %d want %d", got, want)
		}
	})

}
