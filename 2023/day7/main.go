package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadFile(fileName string, cardMap map[string]int) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("error reading file: %v", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		bet, err := strconv.Atoi(line[1])
		if err != nil {
			fmt.Printf("error parsing bet: %v", err)
			os.Exit(1)
		}
		_, exists := cardMap[line[0]]
		if !exists {
			cardMap[line[0]] = bet
		} else {
			fmt.Printf("found duplicate value") // found no duplicate values in real input, so do nothing.
		}
	}
}
