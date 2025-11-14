package adventofcode2023

import "strings"

type Day20Puzzle map[string]*struct {
	typ    byte // 'b' = broadcast, '%' = flip-flop, '&' = conjunction
	name   string
	dests  []string
	state  bool            // for flip-flops: on/off
	inputs map[string]bool // for conjunctions: last pulse from each input
}

func NewDay20(lines []string) (Day20Puzzle, error) {
	puzzle := make(Day20Puzzle)

	// First pass: create all modules
	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.Split(line, " -> ")
		src := parts[0]
		dests := strings.Split(parts[1], ", ")

		m := &struct {
			typ    byte
			name   string
			dests  []string
			state  bool
			inputs map[string]bool
		}{
			dests:  dests,
			inputs: make(map[string]bool),
		}

		if src == "broadcaster" {
			m.typ = 'b'
			m.name = "broadcaster"
		} else {
			m.typ = src[0]
			m.name = src[1:]
		}

		puzzle[m.name] = m
	}

	// Second pass: register inputs for conjunction modules
	for name, m := range puzzle {
		for _, dest := range m.dests {
			if destMod, ok := puzzle[dest]; ok && destMod.typ == '&' {
				destMod.inputs[name] = false // initially low
			}
		}
	}

	return puzzle, nil
}

func Day20(puzzle Day20Puzzle, part1 bool) uint {
	if !part1 {
		return 0
	}

	type pulse struct {
		from, to string
		high     bool
	}

	var lowCount, highCount uint

	for i := 0; i < 1000; i++ {
		queue := []pulse{{from: "button", to: "broadcaster", high: false}}

		for len(queue) > 0 {
			p := queue[0]
			queue = queue[1:]

			if p.high {
				highCount++
			} else {
				lowCount++
			}

			m, ok := puzzle[p.to]
			if !ok {
				continue
			}

			switch m.typ {
			case 'b': // broadcaster
				for _, dest := range m.dests {
					queue = append(queue, pulse{from: m.name, to: dest, high: p.high})
				}

			case '%': // flip-flop
				if !p.high { // only respond to low pulses
					m.state = !m.state
					for _, dest := range m.dests {
						queue = append(queue, pulse{from: m.name, to: dest, high: m.state})
					}
				}

			case '&': // conjunction
				m.inputs[p.from] = p.high
				allHigh := true
				for _, inputHigh := range m.inputs {
					if !inputHigh {
						allHigh = false
						break
					}
				}
				sendHigh := !allHigh
				for _, dest := range m.dests {
					queue = append(queue, pulse{from: m.name, to: dest, high: sendHigh})
				}
			}
		}
	}

	return lowCount * highCount
}
