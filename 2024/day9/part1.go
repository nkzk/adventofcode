package main

import (
	"fmt"
)

func parseDiskMap(diskmap string) []int {
	var result []int

	blockIndex := 0
	for i, char := range diskmap {
		var numberOfBlocks int
		var freeSpace int

		if i%2 == 0 {
			numberOfBlocks = int(char - '0')
			for range numberOfBlocks {
				result = append(result, blockIndex)

			}
			blockIndex++
		} else {
			freeSpace = int(char - '0')
			for range freeSpace {
				result = append(result, -1)
			}
		}
	}

	return result
}

func moveBlocks(diskmap []int) []int {
	fmt.Printf("start:\n%v\n", diskmap)

	moveIndex := 0
	freeSpaceIndexes := make([]int, 0)

	for i, val := range diskmap {
		if val == -1 {
			freeSpaceIndexes = append(freeSpaceIndexes, i)
		}
	}
	fmt.Printf("len: %d\n", len(freeSpaceIndexes))
	fmt.Printf("freespace indexes:\n%v\n", freeSpaceIndexes)

	for i := len(diskmap) - 1; moveIndex < len(freeSpaceIndexes) && i >= len(diskmap)-len(freeSpaceIndexes); i-- {
		if diskmap[i] != -1 {
			move := freeSpaceIndexes[moveIndex]
			diskmap[move] = diskmap[i]
			diskmap[i] = -1
			moveIndex++
		}
	}

	return diskmap

}

func checksum(diskmap []int) int {
	var sum int

	for i, val := range diskmap {
		if val != -1 {
			sum += i * val
		}
	}

	return sum
}

func Part1(input []string) int {
	parsedDiskMap := parseDiskMap(input[0])
	fmt.Printf("Parsed diskmap\n")
	movedDiskMap := moveBlocks(parsedDiskMap)

	fmt.Printf("Moved blocks\n")
	return checksum(movedDiskMap)
}
