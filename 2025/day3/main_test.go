package main

import (
	"github.com/nkzk/adventofcode/utils"
	"log"
	"testing"
)

func Test(t *testing.T) {
	var input []string
	err := utils.ReadFile("test-input", &input)
	if err != nil {
		log.Panicf("failed to read file: %v", err)
	}

	t.Run("Part 1 ", func(t *testing.T) {
		got := Part1(input)
		want := 357
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("Part 2 ", func(t *testing.T) {
		got := Part2(input)
		want := 3121910778619
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

}

func TestCombine2(t *testing.T) {
	have := []int{1, 2, 3, 4}
	want := 1234
	result := combine2(have...)
	if result != want {
		t.Errorf("want: %d, got %d", want, result)
	}
}
