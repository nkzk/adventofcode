package utils

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func SumLineProcessor(line string, result *int) error {
	numbers := strings.Fields(line)
	for _, num := range numbers {
		n, err := strconv.Atoi(num)
		if err != nil {
			return fmt.Errorf("failed to convert string to int: %v", err)
		}

		*result += n
	}
	return nil
}

// Example usage and test
func TestReadFile(t *testing.T) {
	var total int

	err := ReadFile("./test-file.txt", SumLineProcessor, &total)
	if err != nil {
		t.Fatalf("ReadFile returned an error: %v", err)
	}

	expected := 21
	if total != expected {
		t.Errorf("expected sum %d, got %d", expected, total)
	}
}
