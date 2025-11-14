package adventofcode2023

import (
	"strconv"
	"strings"
)

// Day04Part1V1 is the old version with many allocations (kept for comparison)
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

func Day04(buf []byte, part1 bool) (uint, error) {
	if part1 {
		var points uint
		lineStart := 0
		for i := range buf {
			if buf[i] == '\n' {
				wins := countWinningNumbers(buf[lineStart:i])
				points += A131577(wins)
				lineStart = i + 1
			}
		}
		return points, nil
	}

	// Part 2: Use fixed-size array to avoid allocation
	// AoC inputs have <256 cards, so this is safe
	var ns [256]uint
	var cardCount uint

	// Initialize first pass: count cards and set initial counts to 1
	lineStart := 0
	for i := range buf {
		if buf[i] == '\n' {
			ns[cardCount] = 1
			cardCount++
			lineStart = i + 1
		}
	}

	// Second pass: process winning numbers and update card counts
	lineStart = 0
	cardIdx := uint(0)
	for i := range buf {
		if buf[i] == '\n' {
			wins := countWinningNumbers(buf[lineStart:i])
			for j := range wins {
				idx := 1 + cardIdx + j
				ns[idx] += ns[cardIdx]
			}
			lineStart = i + 1
			cardIdx++
		}
	}

	// return total number of scratchcards
	var n uint
	for i := range cardCount {
		n += ns[i]
	}
	return n, nil
}

// countWinningNumbers parses a card line and returns the number of winning numbers
// Format: "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
func countWinningNumbers(line []byte) uint {
	var winning [100]bool // numbers are in range 1-99, use index directly
	var wins uint

	// Find the colon (skip "Card N:")
	i := 0
	for i < len(line) && line[i] != ':' {
		i++
	}
	i += 2 // skip ": "

	// Parse winning numbers (before |)
	var num uint
	hasNum := false
	for i < len(line) && line[i] != '|' {
		c := line[i]
		if c >= '0' && c <= '9' {
			num = num*10 + uint(c-'0')
			hasNum = true
		} else if hasNum {
			// End of number
			winning[num] = true
			num = 0
			hasNum = false
		}
		i++
	}
	i++ // skip '|'

	// Parse our numbers (after |) and count matches
	num = 0
	hasNum = false
	for i < len(line) {
		c := line[i]
		if c >= '0' && c <= '9' {
			num = num*10 + uint(c-'0')
			hasNum = true
		} else if hasNum {
			// End of number - check if it's winning
			if winning[num] {
				wins++
			}
			num = 0
			hasNum = false
		}
		i++
	}
	// Check last number
	if hasNum && winning[num] {
		wins++
	}

	return wins
}

// A131577 returns the OEIS sequence [A131577] 'zero followed by powers of 2'.
// A131577:[https://oeis.org/A131577]
func A131577(n uint) uint {
	if n == 0 {
		return 0
	}
	return 1 << (n - 1)
}
