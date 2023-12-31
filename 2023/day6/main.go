package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// d = hold-ms * remaining-time
// d = heldButtonMS * (maxTime - heldbButtonMS)
// d =  heldButtonMS * maxTime - heldButtonMS^2
// d = -heldButtonMS^2 + heldbuttonMS*MaxTime
// d = ax^2+bx
// max:
// 	-b / 2a (a = -1, b = totalTime)
/*
	-b +/-  sqr(b2-4ac)
	-------------------
	2a


		ax2 + bx + c
	a = 1
	b = -totalTime
	c = x
*/

// Returns max hold time based on totalTime
//
//	maxHoldTime := func(totalTime int) int {
//		return (-totalTime / 2 * (-1))
//	}
//
// Returns distance traveled based on hold-time and totalTime available
func distanceTraveled(holdTime int, totalTime int) int {
	// d=holdTime×(totalTime−holdTime)
	return holdTime * (totalTime - holdTime)
}

// Returns holdtime based on distance and total time
func holdTime(distance int, totalTime int) (int, int) {
	if float64(totalTime)*float64(totalTime)-float64(4)*(-float64(1)) < 0 {
		fmt.Printf("no solutions")
		return 0, 0
	}
	holdTime1 := (-float64(totalTime) + math.Sqrt(float64(totalTime)*float64(totalTime)-float64(4)*(-float64(1))*(-float64(distance)))) / 2 * (-1)
	holdTime2 := (-float64(totalTime) - math.Sqrt(float64(totalTime)*float64(totalTime)-float64(4)*(-float64(1))*(-float64(distance)))) / 2 * (-1)
	return int(math.Round(holdTime1)), int(math.Round(holdTime2))
}

// To solve Part 1: Returns number of ways to win based on already recorded record and total time
func waysToWin(totalTime int, recordDistance int) int {
	var count int
	// recordDistance holdtime
	holdTime1, holdTime2 := holdTime(recordDistance, totalTime)
	for i := min(holdTime1, holdTime2); i <= max(holdTime1, holdTime2); i++ {
		if distanceTraveled(i, totalTime) > recordDistance {
			count++
		}
	}
	return count
}

func PartOne() {
	input := "./input"
	file, err := os.Open(input)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	races := make(map[int]int)
	var times []string
	var distances []string
	for scanner.Scan() {
		// var time int
		// var distance int
		line := scanner.Text()
		if strings.Contains(line, "Time") {
			times = strings.Fields(strings.TrimPrefix(line, "Time:"))
		} else if strings.Contains(line, "Distance") {
			distances = strings.Fields(strings.TrimPrefix(line, "Distance:"))
		}
	}
	for i := range times {
		key, err := strconv.Atoi(times[i])
		if err != nil {
			os.Exit(2)
		}
		distance, err := strconv.Atoi(distances[i])
		if err != nil {
			os.Exit(3)
		}
		races[key] = distance
	}
	sum := 1
	for time, distance := range races {
		sum *= waysToWin(time, distance)
	}
	fmt.Printf("Sum: %d\n", sum)
}
func PartTwo() {
	fmt.Printf("Ways to beat the record in this one much longer race: %d\n", waysToWin(55826490, 246144110121111))
}

func main() {
	PartOne()
	PartTwo()
}
