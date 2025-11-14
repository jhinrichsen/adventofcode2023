package adventofcode2023

import (
	"strings"
)

type lens struct {
	label string
	focal uint
}

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
	input := strings.Join(lines, "")
	steps := strings.Split(input, ",")

	if part1 {
		var total uint
		for _, step := range steps {
			total += hash(step)
		}
		return total
	}

	boxes := make([][]lens, 256)

	for _, step := range steps {
		if strings.Contains(step, "=") {
			parts := strings.Split(step, "=")
			label := parts[0]
			focal := uint(parts[1][0] - '0')
			boxNum := hash(label)

			found := false
			for i := range boxes[boxNum] {
				if boxes[boxNum][i].label == label {
					boxes[boxNum][i].focal = focal
					found = true
					break
				}
			}
			if !found {
				boxes[boxNum] = append(boxes[boxNum], lens{label: label, focal: focal})
			}
		} else if strings.Contains(step, "-") {
			label := strings.TrimSuffix(step, "-")
			boxNum := hash(label)

			for i := range boxes[boxNum] {
				if boxes[boxNum][i].label == label {
					boxes[boxNum] = append(boxes[boxNum][:i], boxes[boxNum][i+1:]...)
					break
				}
			}
		}
	}

	var total uint
	for boxNum, box := range boxes {
		for slot, l := range box {
			power := uint(boxNum+1) * uint(slot+1) * l.focal
			total += power
		}
	}
	return total
}
