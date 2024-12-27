package grid

import "testing"

func TestDecodeKey(t *testing.T) {
	tests := []struct {
		input string
		wantX int
		wantY int
	}{
		{
			"1,4",
			1,
			4,
		},
		{
			"0,0",
			0,
			0,
		},
		{
			"-1,2",
			-1,
			2,
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			gotX, gotY, err := DecodeKey(test.input)
			if gotX != test.wantX || gotY != test.wantY {
				t.Errorf("want %d,%d, got %d,%d, err: %v", test.wantX, test.wantY, gotX, gotY, err)
			}
		})
	}
}
