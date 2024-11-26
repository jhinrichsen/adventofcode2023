package adventofcode2023

import (
	"regexp"
	"strconv"
	"strings"
)

// Function to generate all combinations of '?' replacing with '.' or '#'
func generateCombinations(s string) []string {
	var combinations []string
	// Initialize the queue with the initial string
	queue := []string{s}

	// Process each string in the queue
	for len(queue) > 0 {
		// Take the next string from the queue
		current := queue[0]
		queue = queue[1:]

		// Find the index of the first '?' character
		idx := strings.Index(current, "?")
		if idx == -1 {
			// If there are no more '?', add this combination to the results
			combinations = append(combinations, current)
			continue
		}

		// Create new combinations by replacing '?' with '.' and '#'
		queue = append(queue, current[:idx]+"."+current[idx+1:])
		queue = append(queue, current[:idx]+"#"+current[idx+1:])
	}

	return combinations
}

// Function to validate each combination against the checksum
func isValidCombination(s string, checksum []int) bool {
	// Use regex to find consecutive '#' sequences
	re := regexp.MustCompile(`#+`)

	// Find all consecutive '#' blocks
	matches := re.FindAllString(s, -1)

	// Count the lengths of consecutive '#' blocks
	hashes := make([]int, len(matches))
	for i, match := range matches {
		hashes[i] = len(match)
	}

	// Compare the consecutive '#' sequences with the checksum
	if len(hashes) != len(checksum) {
		return false
	}

	for i := range hashes {
		if hashes[i] != checksum[i] {
			return false
		}
	}

	return true
}

// Day12 function that generates combinations and returns the count of valid ones
func Day12(inputs []string) int {
	validCount := 0

	// Iterate through each input line
	for _, input := range inputs {
		// Split the input into the pattern and checksum parts
		parts := strings.Split(input, " ")
		if len(parts) != 2 {
			continue
		}

		// The first part is the string of '#' and '.' (with possible '?')
		s := parts[0]

		// The second part is the checksum (comma-separated numbers)
		csStr := parts[1]
		csParts := strings.Split(csStr, ",")
		checksum := make([]int, len(csParts))
		for i, c := range csParts {
			val, err := strconv.Atoi(c)
			if err != nil {
				continue
			}
			checksum[i] = val
		}

		// Generate all possible combinations for the '?' characters
		combinations := generateCombinations(s)

		// Count valid combinations
		for _, comb := range combinations {
			if isValidCombination(comb, checksum) {
				validCount++
			}
		}
	}

	return validCount
}
