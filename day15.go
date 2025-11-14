package adventofcode2023

import (
	"strings"
)

func hash(s string) uint {
	var current uint
	for i := 0; i < len(s); i++ {
		current += uint(s[i])
		current *= 17
		current %= 256
	}
	return current
}

func Day15(lines []string, part1 bool) uint {
	if !part1 {
		return 0
	}

	input := strings.Join(lines, "")
	steps := strings.Split(input, ",")

	var total uint
	for _, step := range steps {
		total += hash(step)
	}
	return total
}
