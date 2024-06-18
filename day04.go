package adventofcode2023

import (
	"strconv"
	"strings"
)

func Day04(lines []string) (uint, error) {
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
