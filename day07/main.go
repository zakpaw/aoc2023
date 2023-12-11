package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var cardHierarchy = map[rune]int{
	'A': 0,
	'K': 1,
	'Q': 2,
	'J': 3,
	'T': 4,
	'9': 5,
	'8': 6,
	'7': 7,
	'6': 8,
	'5': 9,
	'4': 10,
	'3': 11,
	'2': 12,
}

type Hand struct {
	Cards []rune
	Bid   int
	Value int
}

func (h *Hand) setValue() {
	// Value e.g.
	// Five of a kind: 5*5 = 25
	// Full house: 3*3 + 2*2 = 13
	cardCounts := make(map[rune]int)
	for _, card := range h.Cards {
		cardCounts[card]++
	}
	for _, count := range cardCounts {
		h.Value += count * count
	}
}

func (h Hand) BetterThan(other Hand) bool {
	if h.Value != other.Value {
		return h.Value > other.Value
	}
	for i := range h.Cards {
		if cardHierarchy[h.Cards[i]] != cardHierarchy[other.Cards[i]] {
			return cardHierarchy[h.Cards[i]] < cardHierarchy[other.Cards[i]]
		}
	}
	return false
}

func parseInput(scanner *bufio.Scanner) []Hand {
	var hands []Hand
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		bid, _ := strconv.Atoi(parts[1])
		hand := Hand{
			Cards: []rune(parts[0]),
			Bid:   bid,
		}
		hand.setValue()
		hands = append(hands, hand)
	}
	return hands
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	hands := parseInput(scanner)

	sort.SliceStable(hands, func(i, j int) bool {
		return hands[j].BetterThan(hands[i])
	})

	resultP1 := 0
	for i, hand := range hands {
		resultP1 += hand.Bid * (i + 1)
	}

	fmt.Println("Total winnings:", resultP1)
}
