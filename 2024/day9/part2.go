package main

import "fmt"

func moveFiles(diskmap []int) []int {
	fmt.Printf("moving files\n")
	moveIndex := 0
	chunkIndex := 0
	freeSpaceIndexes := make([]int, 0)
	freeSpaceChunks := []int{}

	// todo: look into free space chunks, [2] is wrong
	for i, val := range diskmap {
		if val == -1 {
			freeSpaceIndexes = append(freeSpaceIndexes, i)
			j := i + 1
			chunk := 1
			for {
				if diskmap[j] == -1 {
					chunk += 1
					j--
				} else {
					break
				}
			}

			freeSpaceChunks = append(freeSpaceChunks, chunk)
		}
	}

	fmt.Printf("free space chunks: %v\n", freeSpaceChunks)

	for i := len(diskmap) - 1; moveIndex < len(freeSpaceIndexes) && i >= len(diskmap)-len(freeSpaceIndexes); i-- {

		if diskmap[i] != -1 {
			j := i - 1
			chunk := []int{diskmap[i]}

			for {
				if diskmap[j] == diskmap[i] {
					chunk = append(chunk, diskmap[j])
					j--
				} else {
					break
				}
			}

			freeSpaceChunkSize := freeSpaceChunks[chunkIndex]
			fmt.Printf("index: %d, freespacechunksize: %d, chunksize: %d\n", freeSpaceIndexes[moveIndex], freeSpaceChunkSize, len(chunk))
			if freeSpaceChunkSize >= len(chunk) {
				move := freeSpaceIndexes[moveIndex]
				copy(diskmap[move:move+freeSpaceChunkSize], diskmap[j:i])
				// fmt.Printf("copied %d elements to starting index %d\n", copied, move)
				fmt.Printf("incrementing moveindex %d with %d\n", moveIndex, freeSpaceChunkSize)
				moveIndex += freeSpaceChunkSize
			} else {
				fmt.Printf("cant move %d to index %d\n", diskmap[i], freeSpaceIndexes[moveIndex])
			}

			chunkIndex++

		}
	}

	return diskmap
}

func Part2(input []string) int {
	fmt.Printf("part 2\n")
	parsedDiskMap := parseDiskMap(input[0])
	fmt.Printf("Parsed diskmap\n")
	movedDiskMap := moveFiles(parsedDiskMap)
	fmt.Printf("Moved blocks\n")

	return checksum(movedDiskMap)
}
