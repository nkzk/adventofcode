package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(path string, result *[]string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if len(scanner.Text()) > 0 {
			(*result) = append((*result), scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("failed to scan file: %w", err)
	}

	return nil
}

type LineProcessor[Result any] func(line string, result *Result) error

// ReadFile reads a line from path,
// and takes a ReadLineFunc to read each line and modify a Result.
func ReadFile_old[Result any](path string,
	processor LineProcessor[Result], result *Result) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		err := processor(scanner.Text(), result)
		if err != nil {
			return fmt.Errorf("failed to read line: %w", err)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("failed to scan file: %w", err)
	}

	return nil
}
