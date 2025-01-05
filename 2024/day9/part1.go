package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func parseDiskMap(diskmap string) string {
	var buffer bytes.Buffer
	var totalFreeSpace int

	blockIndex := 0
	for i, char := range diskmap {
		var numberOfBlocks int
		var freeSpace int

		if i%2 == 0 {
			numberOfBlocks = int(char - '0')
			for range numberOfBlocks {
				buffer.WriteString(strconv.Itoa(blockIndex))
			}
			blockIndex++
		} else {
			freeSpace = int(char - '0')
			totalFreeSpace += freeSpace
			for range freeSpace {
				buffer.WriteString(".")
			}
		}
	}

	return buffer.String()
}

// O(n^2)
func moveBlocks(diskmap string) string {
	if len(diskmap) <= 1 {
		return diskmap
	}

	diskMapRunes := []rune(diskmap)

	moveIndex := -1

	for i, char := range diskmap {
		if string(char) == "." {
			moveIndex = i
			break
		}
	}

	if moveIndex == -1 {
		return string(diskMapRunes)
	}
	lastChar := diskMapRunes[len(diskMapRunes)-1]

	if lastChar == '.' {
		return moveBlocks(string(diskMapRunes[:len(diskMapRunes)-1]))
	}

	diskMapRunes[moveIndex] = lastChar
	diskMapRunes[len(diskMapRunes)-1] = '.'

	fmt.Printf("%s\n", string(diskMapRunes))
	return moveBlocks(string(diskMapRunes))
}

func moveBlocks2(diskmap string) string {
	diskmapRunes := []rune(diskmap)
	moveIndex := len(diskmapRunes) - 1

	i := 0
	spaceUsed := 0
	for {
		if spaceUsed < freeSpace {
			if diskmapRunes[i] == '.' {
				for {
					if diskmapRunes[moveIndex] == '.' {
						moveIndex -= 1
					} else {
						break
					}
				}
				diskmapRunes[i] = diskmapRunes[moveIndex]
				spaceUsed++
			}
			i++
		} else {
			break
		}

	}

	fmt.Printf("%s\n", string(diskmapRunes))
	return string(diskmapRunes)
}

func checksum(diskmap string) int {
	var sum int

	for i, char := range diskmap {
		if string(char-'0') != "." {
			sum += i * int(char-'0')
		}
	}

	return sum
}

func Part1(input []string) int {
	parsedDiskMap := parseDiskMap(input[0])
	fmt.Printf("Parsed diskmap\n")
	movedDiskMap := moveBlocks2(parsedDiskMap)

	fmt.Printf("Moved blocks\n")
	return checksum(movedDiskMap)
}
