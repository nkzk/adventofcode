package main

import (
	"fmt"
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
		want := 3
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("Part 2 ", func(t *testing.T) {
		got := Part2(input)
		want := 6
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("Part 2,2", func(t *testing.T) {
		got := Part2([]string{"R150"})
		want := 2
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("Part 2,3", func(t *testing.T) {
		got := Part2([]string{"L150"})
		want := 2
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("Part 2,4", func(t *testing.T) {
		got := Part2([]string{"L50", "R1"})
		want := 1
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("Part 2,8", func(t *testing.T) {
		fmt.Printf("\nrunning part 2,5\n\n")
		got := Part2([]string{"L50", "L50"})
		want := 1
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
		fmt.Printf("\n\n")
	})

	t.Run("Part 2,5", func(t *testing.T) {
		got := Part2([]string{
			"L68",
			"L30",
			"R48",
			"L5",
			"R60",
			"L55",
			"L1",
			"L99",
			"R14",
			"L82",
			"R82", // 1
			"R1",  // 0
		})
		want := 7
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("Part 2,6", func(t *testing.T) {
		got := Part2([]string{"L650"})
		want := 7
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("Part 2,7", func(t *testing.T) {
		got := Part2([]string{"R50", "L50"})
		want := 1
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("Part 2,8", func(t *testing.T) {
		fmt.Printf("\nrunning 2,8\n\n")
		got := Part2([]string{"R50", "R50"})
		want := 1
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
		fmt.Printf("\n\n")
	})

}
