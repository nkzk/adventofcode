package main

import (
	"fmt"
	"reflect"
	"testing"
	"utils"
)

func TestRemoveSlice(t *testing.T) {
	slice := []int{3, 4, 7, 2, 1}
	want := []int{3, 7, 2, 1}

	got := remove(slice, 1)

	equal := reflect.DeepEqual(got, want)
	if !equal {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}

func TestMain(t *testing.T) {
	input := make(Input, 0)

	err := utils.ReadFile_old[Input]("test-input", lineProcessor, &input)
	if err != nil {
		t.Errorf("failed to read file: %v", err)
	}

	want := 2

	t.Run("Part 1", func(t *testing.T) {
		got, err := CountSafeReports(input)
		if err != nil {
			t.Errorf("got err: %v", err)
		}
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("Part 2", func(t *testing.T) {
		var sum int
		for _, report := range input {
			fmt.Printf("sending in report: %+v\n", report)
			_, safe := IsReportSafe(report)
			if safe {
				sum += 1
			} else {
				dampener := 0

				for i := range report {
					newReport := remove(report, i)
					fmt.Printf("testing new report: %+v\n", newReport)
					_, safe := IsReportSafe(newReport)
					if safe && dampener < 1 {
						dampener += 1
						sum += 1
					}
				}

			}
		}

		want := 4
		if sum != want {
			t.Errorf("got %d want %d, got, want", sum, want)
		}
	})

}
