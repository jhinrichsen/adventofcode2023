package adventofcode2023

import (
	"strconv"
	"strings"
)

func Day04Part1V1(lines []string) (uint, error) {
	var points uint
	m := make(map[uint]bool)
	for _, line := range lines {
		clear(m)
		nocards := strings.Split(line, ":")
		numbers := strings.Split(nocards[1], "|")

		// winning numbers
		for _, n := range strings.Fields(numbers[0]) {
			wn, err := strconv.Atoi(n)
			if err != nil {
				return 0, err
			}
			m[uint(wn)] = true
		}

		// numbers
		var factor uint = 1
		for _, n := range strings.Fields(numbers[1]) {
			nn, err := strconv.Atoi(n)
			if err != nil {
				return 0, err
			}
			u := uint(nn)
			if m[u] {
				factor *= 2
			}
		}
		points += factor / 2
	}
	return points, nil
}

func Day04(lines []string, part1 bool) (uint, error) {
	wins, err := NewDay04(lines)
	if err != nil {
		return 0, err
	}

	return day04(wins, part1), nil
}

func day04(wins []uint, part1 bool) uint {
	if part1 {
		var points uint
		for _, v := range wins {
			points += A131577(v)
		}
		return points
	}

	// in the beginning, there is one instance of every card
	var ns = make([]uint, len(wins))
	for i := range ns {
		ns[i] = 1
	}

	for i := range ns {
		for j := range wins[i] {
			idx := 1 + uint(i) + j
			ns[idx] += ns[i]
		}
	}

	// return total number of scratchcards
	var n uint
	for _, v := range ns {
		n += v
	}
	return n
}

// NewDay04 returns number of winning numbers for card at given index.
func NewDay04(lines []string) ([]uint, error) {
	const (
		base = 10
		bits = 8 // values in range [1..99]
	)
	var cards []uint
	m := make(map[uint]bool)
	for _, line := range lines {
		clear(m)
		nocards := strings.Split(line, ":")
		numbers := strings.Split(nocards[1], "|")

		// winning numbers
		for _, num := range strings.Fields(numbers[0]) {
			n, err := strconv.ParseUint(num, 10, bits)
			if err != nil {
				return nil, err
			}
			m[uint(n)] = true
		}

		// numbers
		var wins uint
		for _, num := range strings.Fields(numbers[1]) {
			n, err := strconv.ParseUint(num, base, bits)
			if err != nil {
				return nil, err
			}
			if m[uint(n)] {
				wins++
			}
		}
		cards = append(cards, wins)
	}
	return cards, nil
}

// A131577 returns the OEIS sequence [A131577] 'zero followed by powers of 2'.
// A131577:[https://oeis.org/A131577]
func A131577(n uint) uint {
	if n == 0 {
		return 0
	}
	return 1 << (n - 1)
}
