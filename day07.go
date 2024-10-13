package adventofcode2023

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Card int
type HandType int

const (
	HighCard HandType = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func Day07(hands [][]Card, bids []int) (int, error) {

	// sort into ascending order

	ranks := copyDeck(hands)
	slices.SortFunc(ranks, func(a, b []Card) int {

		n1 := handType(a)
		n2 := handType(b)
		if n1 != n2 {
			return int(n1) - int(n2)
		}

		// equal hands: fallback to second ordering rule
		for i := range a {
			if a[i] < b[i] {
				return -1
			} else if a[i] > b[i] {
				return 1
			}
		}
		return 0
	})

	var total int
	for i, r := range ranks {
		rank := i + 1 // one-based
		// go back into original hands to find bid
		var bid int
		for j, h := range hands {
			if slices.Compare(r, h) == 0 {
				bid = bids[j]
				break
			}
		}
		total += rank * bid
	}
	return total, nil
}

func copyDeck(matrix [][]Card) [][]Card {
	duplicate := make([][]Card, len(matrix))
	for i := range matrix {
		duplicate[i] = make([]Card, len(matrix[i]))
		copy(duplicate[i], matrix[i])
	}
	return duplicate
}

func NewDay07(lines []string) ([][]Card, []int, error) {
	var (
		hands [][]Card
		bids  []int
	)
	for i, line := range lines {
		parts := strings.Fields(line)
		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("want 2 fields but got %d", len(parts))
		}

		// hands

		var hand []Card
		for j, b := range parts[0] {
			c, err := card(byte(b))
			if err != nil {
				return nil, nil, fmt.Errorf("line %d, card %d: %c is not a valid card", i, j, b)
			}
			hand = append(hand, c)
		}
		hands = append(hands, hand)

		// bids

		n, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, nil, fmt.Errorf("line %d: want numeric column 2 but got %s: %v", i, parts[1], err)
		}
		bids = append(bids, n)
	}
	return hands, bids, nil
}

func handType(hand []Card) HandType {
	const N = 13
	var cnt [N]int
	for _, b := range hand {
		cnt[b]++
	}
	slices.Sort(cnt[:])

	h1, h2 := cnt[N-1], cnt[N-2]
	if h1 == 5 {
		return FiveOfAKind
	}
	if h1 == 4 {
		return FourOfAKind
	}
	if h1 == 3 {
		if h2 == 2 {
			return FullHouse
		}
		return ThreeOfAKind
	}
	if h1 == 2 {
		if h2 == 2 {
			return TwoPair
		}
		return OnePair
	}
	return HighCard
}

func card(b byte) (Card, error) {
	var cards = []byte{'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'}
	for i, v := range cards {
		if b == v {
			return Card(i), nil
		}
	}
	return Card(0), fmt.Errorf("Unknown card: want one of %+v but got %c", cards, b)
}
