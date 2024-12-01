package main

import (
	"reflect"
	"testing"
)

func TestDay1(t *testing.T) {
	result := Result{
		{
			3, 4, 2, 1, 3, 3,
		},
		{
			4, 3, 5, 3, 9, 3,
		},
	}

	wantSum := 11

	t.Run("Day1 Part 1", func(t *testing.T) {
		sum := part1(result)
		if sum != wantSum {
			t.Errorf("failed, want %d got %d", sum, wantSum)
		}
	})

	wantPart2 := 31

	t.Run("Day1 Part 2", func(t *testing.T) {
		sum := part2(result)
		if sum != wantPart2 {
			t.Errorf("failed, want %d, got %d", wantPart2, sum)
		}
	})

}

func TestSortList(t *testing.T) {
	input := []int{1, 5, 3}
	want := []int{1, 3, 5}
	t.Run("test sort list", func(t *testing.T) {
		result := sortList(input)
		if !reflect.DeepEqual(input, result) {
			t.Errorf("sort list failed, want %v, got %v", want, result)
		}
	})
}
