package adventofcode2023

import (
	"slices"
	"strconv"
	"strings"
)

// Day09V1 returns the sum of the next values for an OASIS sequence.
// This implementation implements lessons learned from day 8 part 2 - visually check the input
// and only do as little as possible.
func Day09V1(lines []string) int {
	const DIM = 100

	var (
		ns [DIM][DIM]int
	)
	var total int
	for _, line := range lines {
		parts := strings.Fields(line)
		for j, s := range parts {
			n, _ := strconv.Atoi(s)
			ns[0][j] = n
		}

		// derive deltas
		y := 1
		for {
			finished := true
			for x := 1; x < 1+len(parts)-y; x++ {
				dx := ns[y-1][x] - ns[y-1][x-1]
				ns[y][x-1] = dx
				if dx != 0 {
					finished = false
				}
			}
			if finished {
				break
			}
			y++
		}

		// y is index into [0 ...] series
		x := len(parts) - y

		// work way up
		next := 0
		for y >= 0 {
			next += ns[y][x-1]
			ns[y][x] = next
			y--
			x++
		}
		total += next
	}
	return total
}

// Day09V2 returns the sum of the next values for an OASIS sequence.
// This implementation implements lessons learned from day 8 part 2 - visually check the input
// and only do as little as possible.
func Day09V2(buf []byte, part1 bool) int {
	const (
		LINES   = 200 // Number of input lines
		MAXNUMS = 30  // Max numbers per line (input has 21, use 30 for safety)
	)

	var (
		ns [LINES][MAXNUMS]int

		// state machine for parser
		negative bool
		n        int // current number building from digits as we go
		idx      int // index into current sequence
	)

	startN := func() {
		negative = false
		n = 0
	}

	digit := func(d int) {
		n = 10*n + d
	}

	endN := func() {
		if negative {
			n = -n
		}
		ns[0][idx] = n
		idx++
	}

	neg := func() {
		negative = true
	}

	newline := func() int {
		if !part1 {
			slices.Reverse(ns[0][:idx])
		}
		// derive deltas
		y := 1
		for {
			finished := true
			for x := 1; x < 1+idx-y; x++ {
				dx := ns[y-1][x] - ns[y-1][x-1]
				ns[y][x-1] = dx
				if dx != 0 {
					finished = false
				}
			}
			if finished {
				break
			}
			y++
		}

		// y is index into [0 ...] series
		x := idx - y

		// work way up
		next := 0
		for y >= 0 {
			next += ns[y][x-1]
			ns[y][x] = next
			y--
			x++
		}
		idx = 0
		return next
	}

	var total int
	for _, b := range buf {
		if b >= '0' && b <= '9' {
			digit(int(b - '0'))
		} else if b == ' ' {
			endN()
			startN()
		} else if b == '-' {
			neg()
		} else if b == '\n' {
			endN()
			total += newline()
			startN()
		}
	}
	return total
}
