package adventofcode2023

import "strings"

type Day19Puzzle struct {
	workflows map[string]workflow
	parts     []part
}

type workflow struct {
	rules []rule
}

type rule struct {
	category byte
	op       byte
	value    int
	dest     string
}

type part struct {
	x, m, a, s int
}

type ratingRange struct {
	x, m, a, s [2]int // [min, max] inclusive
}

func NewDay19(lines []string) (Day19Puzzle, error) {
	var puzzle Day19Puzzle
	puzzle.workflows = make(map[string]workflow)

	inParts := false
	for _, line := range lines {
		if line == "" {
			inParts = true
			continue
		}

		if !inParts {
			// Parse workflow: "px{a<2006:qkq,m>2090:A,rfg}"
			idx := strings.IndexByte(line, '{')
			name := line[:idx]
			rulesStr := line[idx+1 : len(line)-1]

			var wf workflow
			for _, ruleStr := range strings.Split(rulesStr, ",") {
				if strings.Contains(ruleStr, ":") {
					// Conditional rule: "a<2006:qkq"
					parts := strings.Split(ruleStr, ":")
					cond := parts[0]
					dest := parts[1]

					category := cond[0]
					op := cond[1]
					value := 0
					for i := 2; i < len(cond); i++ {
						value = value*10 + int(cond[i]-'0')
					}

					wf.rules = append(wf.rules, rule{category, op, value, dest})
				} else {
					// Default rule: just destination
					wf.rules = append(wf.rules, rule{dest: ruleStr})
				}
			}
			puzzle.workflows[name] = wf
		} else {
			// Parse part: "{x=787,m=2655,a=1222,s=2876}"
			var p part
			line = line[1 : len(line)-1] // Remove { }
			for _, attr := range strings.Split(line, ",") {
				kv := strings.Split(attr, "=")
				val := 0
				for i := 0; i < len(kv[1]); i++ {
					val = val*10 + int(kv[1][i]-'0')
				}
				switch kv[0] {
				case "x":
					p.x = val
				case "m":
					p.m = val
				case "a":
					p.a = val
				case "s":
					p.s = val
				}
			}
			puzzle.parts = append(puzzle.parts, p)
		}
	}

	return puzzle, nil
}

func Day19(puzzle Day19Puzzle, part1 bool) uint {
	if part1 {
		var total uint
		for _, p := range puzzle.parts {
			if processPart(p, puzzle.workflows) {
				total += uint(p.x + p.m + p.a + p.s)
			}
		}
		return total
	}

	// Part 2: Count all possible rating combinations
	initial := ratingRange{
		x: [2]int{1, 4000},
		m: [2]int{1, 4000},
		a: [2]int{1, 4000},
		s: [2]int{1, 4000},
	}
	return countAccepted(initial, "in", puzzle.workflows)
}

func processPart(p part, workflows map[string]workflow) bool {
	current := "in"

	for current != "A" && current != "R" {
		wf := workflows[current]

		for _, r := range wf.rules {
			if r.category == 0 {
				// Default rule
				current = r.dest
				break
			}

			val := 0
			switch r.category {
			case 'x':
				val = p.x
			case 'm':
				val = p.m
			case 'a':
				val = p.a
			case 's':
				val = p.s
			}

			match := false
			if r.op == '<' {
				match = val < r.value
			} else if r.op == '>' {
				match = val > r.value
			}

			if match {
				current = r.dest
				break
			}
		}
	}

	return current == "A"
}

func countAccepted(rr ratingRange, workflowName string, workflows map[string]workflow) uint {
	if workflowName == "R" {
		return 0
	}
	if workflowName == "A" {
		// Calculate combinations: product of all range sizes
		return uint((rr.x[1] - rr.x[0] + 1) *
			(rr.m[1] - rr.m[0] + 1) *
			(rr.a[1] - rr.a[0] + 1) *
			(rr.s[1] - rr.s[0] + 1))
	}

	wf := workflows[workflowName]
	var total uint
	current := rr

	for _, r := range wf.rules {
		if r.category == 0 {
			// Default rule - send all remaining to destination
			total += countAccepted(current, r.dest, workflows)
			break
		}

		// Get the relevant range for this condition
		var rangePtr *[2]int
		switch r.category {
		case 'x':
			rangePtr = &current.x
		case 'm':
			rangePtr = &current.m
		case 'a':
			rangePtr = &current.a
		case 's':
			rangePtr = &current.s
		}

		matching := current

		if r.op == '<' {
			// Matching: [min, value-1]
			// Non-matching: [value, max]
			if rangePtr[0] < r.value {
				matchRange := *rangePtr
				matchRange[1] = min(matchRange[1], r.value-1)
				switch r.category {
				case 'x':
					matching.x = matchRange
				case 'm':
					matching.m = matchRange
				case 'a':
					matching.a = matchRange
				case 's':
					matching.s = matchRange
				}
				total += countAccepted(matching, r.dest, workflows)
			}

			if rangePtr[1] >= r.value {
				nonMatchRange := *rangePtr
				nonMatchRange[0] = max(nonMatchRange[0], r.value)
				switch r.category {
				case 'x':
					current.x = nonMatchRange
				case 'm':
					current.m = nonMatchRange
				case 'a':
					current.a = nonMatchRange
				case 's':
					current.s = nonMatchRange
				}
			} else {
				// No non-matching range
				break
			}
		} else { // r.op == '>'
			// Matching: [value+1, max]
			// Non-matching: [min, value]
			if rangePtr[1] > r.value {
				matchRange := *rangePtr
				matchRange[0] = max(matchRange[0], r.value+1)
				switch r.category {
				case 'x':
					matching.x = matchRange
				case 'm':
					matching.m = matchRange
				case 'a':
					matching.a = matchRange
				case 's':
					matching.s = matchRange
				}
				total += countAccepted(matching, r.dest, workflows)
			}

			if rangePtr[0] <= r.value {
				nonMatchRange := *rangePtr
				nonMatchRange[1] = min(nonMatchRange[1], r.value)
				switch r.category {
				case 'x':
					current.x = nonMatchRange
				case 'm':
					current.m = nonMatchRange
				case 'a':
					current.a = nonMatchRange
				case 's':
					current.s = nonMatchRange
				}
			} else {
				// No non-matching range
				break
			}
		}
	}

	return total
}
