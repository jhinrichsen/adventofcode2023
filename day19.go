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
	if !part1 {
		return 0
	}

	var total uint
	for _, p := range puzzle.parts {
		if processPart(p, puzzle.workflows) {
			total += uint(p.x + p.m + p.a + p.s)
		}
	}
	return total
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
