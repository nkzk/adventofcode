package main

import (
	"log"
	"testing"
	"utils"
)

func Test(t *testing.T) {
	input := make(Input, 0)
	err := utils.ReadFile[Input]("test-input", lineProcessor, &input)
	if err != nil {
		log.Panicf("failed to read file: %v", err)
	}

	want := 161

	t.Run("Part 1 ", func(t *testing.T) {
		got := Part1(input)
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	input2 := make(Input, 0)
	err = utils.ReadFile[Input]("test-input-2", lineProcessor, &input2)
	if err != nil {
		log.Panicf("failed to read file: %v", err)
	}

	t.Run("Part 2 ", func(t *testing.T) {
		got := Part2(input2)
		want := 48
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
