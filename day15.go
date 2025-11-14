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
	// Parse input directly without allocating intermediate strings
	// Input is typically a single line, but handle multiple lines
	input := lines[0]
	if len(lines) > 1 {
		input = strings.Join(lines, "")
	}

	if part1 {
		var total uint
		var current uint
		for i := 0; i < len(input); i++ {
			if input[i] == ',' {
				// End of step
				current = 0
			} else {
				current += uint(input[i])
				current *= 17
				current %= 256
				// Check if next char is comma or end
				if i+1 >= len(input) || input[i+1] == ',' {
					total += current
					current = 0
				}
			}
		}
		return total
	}

	// Part 2: Pre-allocate boxes with reasonable capacity per box
	boxes := make([][]lens, 256)
	for i := range boxes {
		boxes[i] = make([]lens, 0, 8) // Most boxes have few lenses
	}

	// Parse steps inline without string allocations
	start := 0
	for i := 0; i <= len(input); i++ {
		if i == len(input) || input[i] == ',' {
			if i > start {
				step := input[start:i]

				// Find operation: = or -
				opIdx := -1
				isAdd := false
				for j := 0; j < len(step); j++ {
					if step[j] == '=' {
						opIdx = j
						isAdd = true
						break
					} else if step[j] == '-' {
						opIdx = j
						break
					}
				}

				if opIdx >= 0 {
					label := step[:opIdx]
					boxNum := hash(label)

					if isAdd {
						focal := uint(step[opIdx+1] - '0')
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
					} else {
						// Remove operation
						for i := range boxes[boxNum] {
							if boxes[boxNum][i].label == label {
								boxes[boxNum] = append(boxes[boxNum][:i], boxes[boxNum][i+1:]...)
								break
							}
						}
					}
				}
			}
			start = i + 1
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
