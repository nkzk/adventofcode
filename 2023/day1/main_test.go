package main

import (
	"testing"
)

func TestConcatTwoIntegers(t *testing.T) {

	var tests = []struct {
		name   string
		inputA int
		inputB int
		want   int
	}{
		{"1 and 2 should be 12", 1, 2, 12},
		{"9 and 90 should be 990", 9, 90, 990},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := ConcatTwoIntegers(test.inputA, test.inputB)
			if got != test.want {
				t.Errorf("want %d got %d", test.want, got)
			}
		})
	}

}

func TestSum(t *testing.T) {

	wantedSum := 142
	lines := []int{12, 38, 15, 77}
	sum := Sum(lines)

	if wantedSum != sum {
		t.Errorf("Sum does not match wanted sum")
	}
}

func TestConvertToInt(t *testing.T) {
	var e error
	var tests = []struct {
		name    string
		input   string
		want    int
		wantErr error
	}{
		{"'12' should return 12", "12", 12, nil},
		{"'a' should return error", "a", 0, e},
		{"'a1' should return error", "a1", 0, e},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			number, err := ConvertToInt(test.input)
			if err != nil {
				e = err

			}
			if number != test.want {
				t.Errorf("got %d, want %d", number, test.want)
			}

		})
	}
}
func TestGetCalibration(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  int
	}{
		{"Line 1 should be 12", "1abc2", 12},
		{"Line 2 should be 38", "pqr3stu8vwx", 38},
		{"Line 3 should be 15", "a1b2c3d4e5f", 15},
		{"Line 4 should be 77", "treb7uchet", 77},
		{"Line 5 should be 0", "trebuchet", 0},
		{"Line 6 should be 29", "two1nine", 29},
		{"Line 7 should be 29", "7pqrstsixteen", 76},
		{"Line 8 should be 83", "eightwothree", 83},
		{"Line 9 should be 13", "abcone2threexyz", 13},
		{"Line 10 should be 24", "xtwone3four", 24},
		{"Line 11 should be 42", "4nineeightseven2", 42},
		{"Line 12 should be 14", "zoneight234", 14},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			number := GetCalibration(tc.input)
			if number != tc.want {
				t.Errorf("got %d, want %d", number, tc.want)
			}
		})
	}
}
