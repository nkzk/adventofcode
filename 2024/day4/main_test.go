package main

import (
	"log"
	"testing"
	"github.com/nkzk/adventofcode/utils"
)

func Test(t *testing.T) {
	input := make(Input, 0)
	err := utils.ReadFile_old[Input]("test-input", lineProcessor, &input)
	if err != nil {
		log.Panicf("failed to read file: %v", err)
	}

	want := 18

	t.Run("Part 1 ", func(t *testing.T) {
		got := Part1(input)
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
