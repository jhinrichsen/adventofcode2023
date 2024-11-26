package adventofcode2023

import (
	"strconv"
	"strings"
)

const (
	low      = '.'
	high     = '#'
	wildcard = '?'
)

// permute generates all combinations of '?' replacing with '.' or '#'.
func permute(s string) []string {
	var (
		n         uint = 1 // number of wildcards
		val       byte = low
		p, period uint = 0, 1 // period counter and period length
	)

	// determine number of combinations
	for i := range s {
		if s[i] == wildcard {
			n *= 2
		}
	}
	if n == 1 { // no wildcards => no permutations
		return []string{s}
	}

	combinations := make([][]byte, n)
	for i := range combinations {
		combinations[i] = []byte(s)
	}

	// iterate sequentially, and let wildcard replacements fall down as in The Matrix.
	for i := range s {
		if s[i] != wildcard {
			continue
		}
		for y := range combinations {
			combinations[y][i] = val
			p++
			if p == period { // toggle?
				p = 0
				if val == low {
					val = high
				} else {
					val = low
				}
			}
		}
		period *= 2
	}
	ss := make([]string, len(combinations))
	for i := range combinations {
		ss[i] = string(combinations[i])
	}
	return ss
}

func isValid(s string, checksums []uint) bool {
	var (
		cons uint // counter for number of consecutive #
		idx  int  // index into checksums

		l        = len(checksums)
		previous = low // start on low
		valid    = true
	)

	rising := func() { // .#
		cons = 1
	}
	cont := func() { // ##
		cons++
	}
	check := func() {
		if idx >= l {
			valid = false
			return
		}
		valid = cons == checksums[idx]
	}
	falling := func() { // #.
		check()
		idx++
	}

	for i := range s {
		switch s[i] {
		case low:
			if previous == high {
				falling()
			}
		case high:
			if previous == low {
				rising()
			} else {
				cont()
			}
		}
		previous = rune(s[i])
		if !valid {
			return false
		}
	}
	// always end low
	if previous == high {
		falling()
	}

	// we must have processed all our checksums
	if idx != l {
		return false
	}
	return valid
}

// Day12 generates combinations and returns the count of valid ones.
func Day12(lines []string) uint {
	var n uint
	for i := range lines {
		// Split the input into the pattern and checksum parts
		parts := strings.Fields(lines[i])
		if len(parts) != 2 {
			continue
		}

		// The first part is the string of '#' and '.' (with possible '?')
		combinations := permute(parts[0])

		// The second part is the checksum (comma-separated numbers)
		csParts := strings.Split(parts[1], ",")
		checksums := make([]uint, len(csParts))
		for j := range csParts {
			val, err := strconv.Atoi(csParts[j])
			if err != nil {
				continue
			}
			checksums[j] = uint(val)
		}

		// Count valid combinations
		for _, comb := range combinations {
			if isValid(comb, checksums) {
				n++
			}
		}
	}
	return n
}
