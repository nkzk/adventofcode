package main

import (
	"fmt"
	"testing"
)

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

	var three bool
	var two bool
	two = false
	for key, count := range kindCount {
		fmt.Printf("In rankCard, kindCount: %s, %d\n", key, count)
		switch count {
		case 5:
			return cardRank["FiveOfAKind"]
		case 4:
			return cardRank["FourOfAKind"]
		case 3:
			three = true
		case 2:
			fmt.Printf("case 2\n")
			if !two {
				two = true
			} else {
				return cardRank["TwoPair"]
			}
		}
	}
	if three {
		return cardRank["ThreeOfAKind"]
	}
	if three && two {
		return cardRank["FullHouse"]
	}
	if two {
		return cardRank["OnePair"]
	}
	return cardRank["HighCard"]
}

// Returns strongest hand
func CompareHands(hand1 string, hand2 string) (string, int) {
	// 5 cards (runes)
	// A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, or 2
	kindValue := map[string]int{
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"T": 10,
		"J": 11,
		"Q": 12,
		"K": 13,
		"A": 14,
	}

	kindCount1 := make(map[string]int)
	kindCount2 := make(map[string]int)

	// rank the cards:
	// 1: count them
	for _, kind := range hand1 {
		kindCount1[string(kind)] += 1
	}
	for _, kind := range hand2 {
		kindCount2[string(kind)] += 1
	}
	fmt.Printf("Hand 1 (%s):\n", hand1)
	for k, v := range kindCount1 {
		fmt.Printf("%s: %d\n", k, v)
	}
	fmt.Printf("Hand 2 (%s):\n", hand2)
	for k, v := range kindCount2 {
		fmt.Printf("%s: %d\n", k, v)
	}
	// 2: rank them
	card1Rank := rankCard(kindCount1)
	card2Rank := rankCard(kindCount2)
	fmt.Printf("Card 1 rank: %d, Card 2 rank: %d\n", card1Rank, card2Rank)
	// Handle same rank based on type (first highest card wins)
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

func TestPartOne(t *testing.T) {
	// var rankedCards []string
	cards := make(map[string]int)
	ReadFile("input", cards)
	card, rank := CompareHands("32T3K", "T55J5")
	fmt.Printf("biggest hand: %s, rank: %d\n", card, rank)
}
