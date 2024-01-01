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

func rankCard(kindCount map[string]int) int {

	cardRank := map[string]int{
		"FiveOfAKind":  7,
		"FourOfAKind":  6,
		"FullHouse":    5,
		"ThreeOfAKind": 4,
		"TwoPair":      3,
		"OnePair":      2,
		"HighCard":     1,
	}

	// var three bool
	// var two bool
	maxCount := 0
	pairCounter := 0
	jokerCount := kindCount["J"]

	// count
	for key, count := range kindCount {
		if count > maxCount && key != "J" {
			maxCount = count
		}
		if count == 2 && key != "J" {
			pairCounter++
		}
	}
	maxCount += jokerCount
	switch maxCount {
	case 5:
		return cardRank["FiveOfAKind"]
	case 4:
		return cardRank["FourOfAKind"]
	case 3:
		if pairCounter == 1 && jokerCount == 0 {
			return cardRank["FullHouse"]
		}
		return cardRank["ThreeOfAKind"]
	case 2:
		if pairCounter == 2 {
			return cardRank["TwoPair"]
		}
		return cardRank["OnePair"]
	default:
		return cardRank["HighCard"]
	}
}

// Returns strongest hand
func CompareHands(hand1 string, hand2 string) (string, int) {
	kindValue := map[string]int{
		"J": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"T": 10,
		// "J": 11, // comment out for part two
		"Q": 12,
		"K": 13,
		"A": 14,
	}

	kindCount1 := make(map[string]int)
	kindCount2 := make(map[string]int)

	for _, kind := range hand1 {
		kindCount1[string(kind)] += 1
	}
	for _, kind := range hand2 {
		kindCount2[string(kind)] += 1
	}

	card1Rank := rankCard(kindCount1)
	// fmt.Printf(" %s got rank %d\n", hand1, card1Rank)
	card2Rank := rankCard(kindCount2)
	// fmt.Printf(" %s got rank %d\n", hand2, card2Rank)

	if card1Rank == card2Rank {
		for i := 0; i < len(hand1); i++ {
			if kindValue[string(hand1[i])] == kindValue[string(hand2[i])] {
				continue
			}
			if kindValue[string(hand1[i])] > kindValue[string(hand2[i])] {
				return hand1, card1Rank
			} else {
				return hand2, card2Rank
			}
		}
	}
	if card1Rank > card2Rank {
		return hand1, card1Rank
	} else {
		return hand2, card2Rank
	}
}

func QuickSort(arr []Card, low, high int) []Card {
	if low < high {
		var p int
		arr, p = Partition(arr, low, high)
		arr = QuickSort(arr, low, p-1)
		arr = QuickSort(arr, p+1, high)
	}
	return arr
}

func Partition(arr []Card, low int, high int) ([]Card, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		strongestHand, _ := CompareHands(arr[j].Hand, pivot.Hand)
		// fmt.Printf("Strongest hand: %s, %d\n", strongestHand, r)
		if arr[j].Hand != strongestHand {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

type Card struct {
	Hand string
	Bet  int
}

func SumWinnings(cardArray []Card) int {
	var sum int
	for i := range cardArray {
		add := cardArray[i].Bet * (i + 1)
		sum += add
	}
	return sum
}

func main() {
	cards := make(map[string]int)
	ReadFile("input", cards)

	var cardArray []Card
	for key, value := range cards {
		cardArray = append(cardArray, Card{Hand: key, Bet: value})
	}

	sortedCardArray := QuickSort(cardArray, 0, len(cardArray)-1)

	fmt.Println(SumWinnings(sortedCardArray)) // 249495061 too low
	// 249791369 too high

}
