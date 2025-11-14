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

type Hand struct {
	cards []Card
	bid   uint
}

func Day07(hands []Hand, useJoker bool) (uint, error) {

	// sort into ascending order

	slices.SortFunc(hands, func(a, b Hand) int {

		n1 := handType(a, useJoker)
		n2 := handType(b, useJoker)
		if n1 != n2 {
			return int(n1) - int(n2)
		}

		// equal hands: fallback to second ordering rule
		for i := range a.cards {
			if a.cards[i] < b.cards[i] {
				return -1
			} else if a.cards[i] > b.cards[i] {
				return 1
			}
			// equal, continue with next card
		}
		return 0
	})

	var total uint
	for i, h := range hands {
		rank := 1 + uint(i) // rank is one-based
		total += rank * h.bid
	}
	return total, nil
}

func NewDay07(lines []string, part1 bool) ([]Hand, error) {
	hands := make([]Hand, len(lines))
	for i, line := range lines {
		parts := strings.Fields(line)
		if len(parts) != 2 {
			return nil, fmt.Errorf("want 2 fields but got %d", len(parts))
		}

		// cards

		for j, b := range parts[0] {
			c, err := card(byte(b), part1)
			if err != nil {
				return nil, fmt.Errorf("line %d, card %d: %c is not a valid card", i, j, b)
			}
			hands[i].cards = append(hands[i].cards, c)
		}

		// bids

		n, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, fmt.Errorf("line %d: want numeric column 2 but got %s: %v", i, parts[1], err)
		}
		hands[i].bid = uint(n)
	}
	return hands, nil
}

func handType(h Hand, useJoker bool) HandType {

	// sort by count ascending

	const N = 13
	var cnt [N]int
	for _, b := range h.cards {
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
	switch h1 {
	case 5:
		t = FiveOfAKind
	case 4:
		t = FourOfAKind
	case 3:
		if h2 == 2 {
			t = FullHouse
		} else {
			t = ThreeOfAKind
		}
	case 2:
		if h2 == 2 {
			t = TwoPairs
		} else {
			t = OnePair
		}
	default:
		t = HighCard
	}

	if !useJoker {
		return t
	}

	n := joker(h)
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
	return Card(0), fmt.Errorf("unknown card: want one of %+v but got %c", cards, b)
}

// joker returns number of Joker in hand.
func joker(h Hand) (n int) {

	// neither slices nor iter packge can Count so far

	const j = 0 // joker is lowest card
	for _, c := range h.cards {
		if c == j {
			n++
		}
	}
	return
}
