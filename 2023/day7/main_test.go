package main

import (
	"fmt"
	"testing"
)

func TestPartOne(t *testing.T) {
	// var rankedCards []string
	cards := make(map[string]int)
	ReadFile("test-input", cards)
	var cardArray []Card
	for key, value := range cards {
		cardArray = append(cardArray, Card{Hand: key, Bet: value})
	}
	fmt.Printf("Card array: %v\n", cardArray)

	sortedCardArray := QuickSort(cardArray, 0, len(cardArray)-1)

	fmt.Printf("Sorted array: %v\n", sortedCardArray)
	// card, rank := CompareHands("32T3K", "T55J5")
	// fmt.Printf("biggest hand: %s, rank: %d\n", card, rank)

	sortedArrayTest := []Card{
		{Hand: "32T3K", Bet: 765},
		{Hand: "KTJJT", Bet: 220},
		{Hand: "KK677", Bet: 28},
		{Hand: "T55J5", Bet: 684},
		{Hand: "QQQJA", Bet: 483},
	}
	t.Run("Sorted Card Array", func(t *testing.T) {
		for i := range sortedCardArray {
			if sortedCardArray[i].Hand != sortedArrayTest[i].Hand {
				t.Errorf("want %s, got %s", sortedArrayTest[i].Hand, sortedCardArray[i].Hand)
			}
		}
	})

	t.Run("Sum total winnings", func(t *testing.T) {
		want := 6440
		got := SumWinnings(sortedCardArray)

		if got != want {
			t.Errorf("want %d, got %d", want, got)
		}

	})
}
func TestPartTwo(t *testing.T) {
	cards := make(map[string]int)
	ReadFile("test-input", cards)
	var cardArray []Card
	for key, value := range cards {
		cardArray = append(cardArray, Card{Hand: key, Bet: value})
	}
	fmt.Printf("Card array p2 %v\n", cardArray)

	sortedCardArray := QuickSort(cardArray, 0, len(cardArray)-1)

	fmt.Printf("Sorted array p2: %v\n", sortedCardArray)

	sortedArrayTest := []Card{
		{Hand: "32T3K", Bet: 765},
		{Hand: "KK677", Bet: 28},
		{Hand: "T55J5", Bet: 684},
		{Hand: "QQQJA", Bet: 483},
		{Hand: "KTJJT", Bet: 220},
	}
	t.Run("Sorted Card Array", func(t *testing.T) {
		for i := range sortedCardArray {
			if sortedCardArray[i].Hand != sortedArrayTest[i].Hand {
				t.Errorf("want %s, got %s", sortedArrayTest[i].Hand, sortedCardArray[i].Hand)
			}
		}
	})

	t.Run("Sum total winnings", func(t *testing.T) {
		want := 5905
		got := SumWinnings(sortedArrayTest)

		if got != want {
			t.Errorf("want %d, got %d", want, got)
		}

	})
}

func TestPartTwoExtra(t *testing.T) {
	cards := make(map[string]int)
	ReadFile("test-input-2", cards)
	var cardArray []Card
	for key, value := range cards {
		cardArray = append(cardArray, Card{Hand: key, Bet: value})
	}

	sortedCardArray := QuickSort(cardArray, 0, len(cardArray)-1)

	t.Run("Sum total winnings 2", func(t *testing.T) {
		want := 1343
		got := SumWinnings(sortedCardArray)

		if got != want {
			t.Errorf("want %d, got %d", want, got)
		}

	})
}

func TestPartTwoRankCard(t *testing.T) {
	testCases := []struct {
		hand string
		bid  int
		rank int // Expected rank
	}{
		{"AAAAA", 2, 7},  // Five of a Kind
		{"22222", 3, 7},  // Five of a Kind
		{"AAAAK", 5, 6},  // Four of a Kind
		{"22223", 7, 6},  // Four of a Kind
		{"AAAKK", 11, 5}, // Full House
		{"22233", 13, 5}, // Full House
		{"AAAKQ", 17, 4}, // Three of a Kind
		{"22234", 19, 4}, // Three of a Kind
		{"AAKKQ", 23, 3}, // Two Pair
		{"22334", 29, 3}, // Two Pair
		{"AAKQJ", 31, 4}, // Three of a Kind with joker
		{"22345", 37, 2}, // One Pair
		{"AKQJT", 41, 2}, // One Pair
		{"23456", 43, 1}, // High Card
		{"A2345", 0, 1},  // High Card
		{"AJ234", 0, 2},  // One Pair with Joker
		{"AACC2", 0, 3},  // Two Pair
		{"AA234", 0, 2},  // One Pair
		{"A2345", 0, 1},  // High Card
		{"AJ234", 0, 2},  // One Pair with Joker
		{"AACC2", 0, 3},  // Two Pair
		{"JJJJJ", 0, 7},  // Five of a kind
		{"JJJJK", 0, 7},  // Five of a kind
		{"JKKKK", 0, 7},  // Five of a kind
		{"JJAA1", 0, 6},  // Four of a kind
	}

	for _, tc := range testCases {
		kindCount := make(map[string]int)
		for _, kind := range tc.hand {
			kindCount[string(kind)]++
		}

		gotRank := rankCard(kindCount)
		if gotRank != tc.rank {
			t.Errorf("Hand %s expected rank %d, got %d", tc.hand, tc.rank, gotRank)
		}
	}
}

func TestPartTwoCompareHands(t *testing.T) {
	testCases := []struct {
		hand1  string
		hand2  string
		winner string // expected winning hand
	}{
		{"AAAAA", "22222", "AAAAA"}, // Both Five of a Kind; A higher than 2
		{"AAAAK", "22223", "AAAAK"}, // Both Four of a Kind; A higher than 2
		{"AAAKK", "22233", "AAAKK"}, // Both Full House; A higher than 2
		{"AAAKQ", "22234", "AAAKQ"}, // Both Three of a Kind; A higher than 2
		{"AAKKQ", "22334", "AAKKQ"}, // Both Two Pair; A higher than 2
		{"AAKQJ", "22345", "AAKQJ"}, // Three of a Kind with Joker vs One Pair
		{"AKQJT", "23456", "AKQJT"}, // Both High Card; A higher than 2
		{"AAAKK", "KKKAA", "AAAKK"}, // Both Full House; Triples compared first
		{"AAKQJ", "KKQJ2", "AAKQJ"}, // Both One Pair; A Pair higher than K Pair
		{"AJ234", "KJ234", "AJ234"}, // Both One Pair with Joker; A higher than K
		{"AA234", "KK234", "AA234"}, // Both One Pair; A Pair higher than K Pair
		{"A2345", "K2345", "A2345"}, // Both High Card; A higher than K
		{"JJJJJ", "22222", "22222"}, // Jokers make Five of a Kind
		{"JJJJK", "KKKK2", "JJJJK"}, // Jokers make Five of a Kind
		{"JKKKK", "2222A", "JKKKK"}, // Jokers make Five of a Kind
		{"JJAAK", "AAAKK", "JJAAK"}, // Jokers make Four of a Kind
	}

	for _, tc := range testCases {
		winner, _ := CompareHands(tc.hand1, tc.hand2)
		if winner != tc.winner {
			t.Errorf("Expected winner between %s and %s is %s, got %s", tc.hand1, tc.hand2, tc.winner, winner)
		}
	}
}
func TestPartTwoSortHands(t *testing.T) {
	handsWithBets := []Card{
		{Hand: "AAAAA", Bet: 2},
		{Hand: "22222", Bet: 3},
		{Hand: "AAAAK", Bet: 5},
		{Hand: "22223", Bet: 7},
		{Hand: "AAAKK", Bet: 11},
		{Hand: "22233", Bet: 13},
		{Hand: "AAKQJ", Bet: 31},
		{Hand: "AAAKQ", Bet: 17},
		{Hand: "22234", Bet: 19},
		{Hand: "AAKKQ", Bet: 23},
		{Hand: "22334", Bet: 29},
		{Hand: "22345", Bet: 37},
		{Hand: "AKQJT", Bet: 41},
		{Hand: "23456", Bet: 43},
	}

	sortedHands := QuickSort(handsWithBets, 0, len(handsWithBets)-1)

	// Expected order
	expectedOrder := []string{"AAAAA", "22222", "AAAAK", "22223", "AAAKK", "22233", "AAKQJ", "AAAKQ", "22234", "AAKKQ", "22334", "22345", "AKQJT", "23456"}

	// Check if the sorted hands are in the expected order
	for i, expectedHand := range expectedOrder {
		if sortedHands[i].Hand != expectedHand {
			t.Errorf("Expected hand at position %d is %s, got %s", i, expectedHand, sortedHands[i].Hand)
		}
	}
}
