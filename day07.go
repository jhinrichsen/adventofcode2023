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
	TwoPairs
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func Day07(hands [][]Card, bids []uint, useJoker bool) (uint, error) {

	// sort into ascending order

	ranked := copyDeck(hands)
	slices.SortFunc(ranked, func(a, b []Card) int {

		n1 := handType(a, useJoker)
		n2 := handType(b, useJoker)
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
			// equal, continue with next card
		}
		return 0
	})

	var total uint
	for i, r := range ranked {
		rank := 1 + uint(i) // one-based
		// go back into original hands to find bid
		var bid uint
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

func NewDay07(lines []string, part1 bool) ([][]Card, []uint, error) {
	var (
		hands [][]Card
		bids  []uint
	)
	for i, line := range lines {
		parts := strings.Fields(line)
		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("want 2 fields but got %d", len(parts))
		}

		// hands

		var hand []Card
		for j, b := range parts[0] {
			c, err := card(byte(b), part1)
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
		bids = append(bids, uint(n))
	}
	return hands, bids, nil
}

func handType(hand []Card, useJoker bool) HandType {

	// sort by count ascending

	const N = 13
	var cnt [N]int
	for _, b := range hand {
		if useJoker { // consider joker?
			if b == 0 { // joker do not count for hand type
				continue
			}
		}
		cnt[b]++
	}
	slices.Sort(cnt[:])

	h1, h2 := cnt[N-1], cnt[N-2]
	var t HandType
	if h1 == 5 {
		t = FiveOfAKind
	} else if h1 == 4 {
		t = FourOfAKind
	} else if h1 == 3 {
		if h2 == 2 {
			t = FullHouse
		} else {
			t = ThreeOfAKind
		}
	} else if h1 == 2 {
		if h2 == 2 {
			t = TwoPairs
		} else {
			t = OnePair
		}
	} else {
		t = HighCard
	}

	if !useJoker {
		return t
	}

	n := joker(hand)
	for i := 0; i < n; i++ {
		switch t {
		case FiveOfAKind:
			t = FiveOfAKind
		case FourOfAKind:
			t = FiveOfAKind
		case FullHouse:
			t = FourOfAKind
		case ThreeOfAKind:
			t = FourOfAKind
		case TwoPairs:
			t = FullHouse
		case OnePair:
			t = ThreeOfAKind
		case HighCard:
			t = OnePair
		default:
			panic("incomplete switch case")
		}
	}
	return t
}

func card(b byte, joker bool) (Card, error) {
	var cards []byte
	if joker {
		cards = []byte{'J', '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'Q', 'K', 'A'}
	} else {
		cards = []byte{'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'}
	}
	for i, v := range cards {
		if b == v {
			return Card(i), nil
		}
	}
	return Card(0), fmt.Errorf("Unknown card: want one of %+v but got %c", cards, b)
}

// joker returns number of Joker in hand.
func joker(hand []Card) (n int) {
	for _, h := range hand {
		if h == 0 {
			n++
		}
	}
	return
}
