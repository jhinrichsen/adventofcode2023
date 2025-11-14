package adventofcode2023

import (
	"strconv"
	"strings"
)

// unfold replicates the pattern and checksums 5 times for Part 2.
func unfold(pattern string, checksums []uint) (string, []uint) {
	patterns := make([]string, 5)
	for i := range 5 {
		patterns[i] = pattern
	}
	unfolded := make([]uint, 0, len(checksums)*5)
	for range 5 {
		unfolded = append(unfolded, checksums...)
	}
	return strings.Join(patterns, "?"), unfolded
}

// countArrangements uses dynamic programming with memoization to count valid arrangements.
// This is efficient for part 2 where brute force permutation would explode.
func countArrangements(pattern string, checksums []uint) uint {
	memo := make(map[[3]int]uint)

	var solve func(pos, groupIdx, groupLen int) uint
	solve = func(pos, groupIdx, groupLen int) uint {
		// Base case: reached end of pattern
		if pos == len(pattern) {
			// Valid if we've matched all groups and not in middle of a group
			if groupIdx == len(checksums) && groupLen == 0 {
				return 1
			}
			// Or if we're on the last group and it matches
			if groupIdx == len(checksums)-1 && groupLen == int(checksums[groupIdx]) {
				return 1
			}
			return 0
		}

		key := [3]int{pos, groupIdx, groupLen}
		if val, ok := memo[key]; ok {
			return val
		}

		var count uint
		c := pattern[pos]

		// Try placing '.'
		if c == '.' || c == '?' {
			if groupLen == 0 {
				// Not in a group, continue
				count += solve(pos+1, groupIdx, 0)
			} else if groupIdx < len(checksums) && groupLen == int(checksums[groupIdx]) {
				// Finish current group
				count += solve(pos+1, groupIdx+1, 0)
			}
		}

		// Try placing '#'
		if c == '#' || c == '?' {
			if groupIdx < len(checksums) && groupLen < int(checksums[groupIdx]) {
				// Continue or start a group
				count += solve(pos+1, groupIdx, groupLen+1)
			}
		}

		memo[key] = count
		return count
	}

	return solve(0, 0, 0)
}

// NewDay12 parses the input lines (no parsing needed, just validation).
func NewDay12(lines []string) ([]string, error) {
	return lines, nil
}

// Day12 solves both parts based on the part1 flag.
func Day12(lines []string, part1 bool) uint {
	var total uint
	for i := range lines {
		parts := strings.Fields(lines[i])
		if len(parts) != 2 {
			continue
		}

		pattern := parts[0]
		checksumStr := parts[1]

		// Parse checksums
		csParts := strings.Split(checksumStr, ",")
		checksums := make([]uint, len(csParts))
		for j := range csParts {
			val, err := strconv.Atoi(csParts[j])
			if err != nil {
				continue
			}
			checksums[j] = uint(val)
		}

		if part1 {
			// Part 1: Use DP approach (faster than permutation)
			total += countArrangements(pattern, checksums)
		} else {
			// Part 2: Use DP approach with unfolded input
			unfoldedPattern, unfoldedChecksums := unfold(pattern, checksums)
			total += countArrangements(unfoldedPattern, unfoldedChecksums)
		}
	}
	return total
}
