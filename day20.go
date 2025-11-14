package adventofcode2023

type Day20Puzzle map[string]*struct {
	typ    byte // 'b' = broadcast, '%' = flip-flop, '&' = conjunction
	name   string
	dests  []string
	state  bool            // for flip-flops: on/off
	inputs map[string]bool // for conjunctions: last pulse from each input
}

func NewDay20(lines []string) (Day20Puzzle, error) {
	// Pre-allocate puzzle map with capacity
	puzzle := make(Day20Puzzle, len(lines))

	// First pass: create all modules
	for _, line := range lines {
		if line == "" {
			continue
		}

		// Parse inline without strings.Split
		// Format: "broadcaster -> a, b, c" or "%a -> b, c"
		arrowIdx := 0
		for i := 0; i < len(line)-3; i++ {
			if line[i] == ' ' && line[i+1] == '-' && line[i+2] == '>' && line[i+3] == ' ' {
				arrowIdx = i
				break
			}
		}

		src := line[:arrowIdx]
		destsStr := line[arrowIdx+4:]

		// Parse destinations
		dests := make([]string, 0, 4)
		start := 0
		for i := 0; i <= len(destsStr); i++ {
			if i == len(destsStr) || destsStr[i] == ',' {
				if i > start {
					dest := destsStr[start:i]
					// Skip comma and space
					if i < len(destsStr) && destsStr[i] == ',' {
						start = i + 2 // Skip ", "
					}
					dests = append(dests, dest)
				}
			}
		}

		m := &struct {
			typ    byte
			name   string
			dests  []string
			state  bool
			inputs map[string]bool
		}{
			dests:  dests,
			inputs: make(map[string]bool, 4),
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
	type pulse struct {
		from, to string
		high     bool
	}

	if part1 {
		var lowCount, highCount uint

		for range 1000 {
			queue := make([]pulse, 0, 100)
			queue = append(queue, pulse{from: "button", to: "broadcaster", high: false})

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

	// Part 2: Find when rx receives a low pulse
	// rx is typically fed by a conjunction module
	// Find that conjunction and track when each of its inputs sends high
	var rxFeeder string
	for name, m := range puzzle {
		for _, dest := range m.dests {
			if dest == "rx" {
				rxFeeder = name
				break
			}
		}
	}

	// Track cycle lengths for each input to rxFeeder
	cycles := make(map[string]uint)
	var buttonPresses uint

	for len(cycles) < len(puzzle[rxFeeder].inputs) {
		buttonPresses++
		queue := make([]pulse, 0, 100)
		queue = append(queue, pulse{from: "button", to: "broadcaster", high: false})

		for len(queue) > 0 {
			p := queue[0]
			queue = queue[1:]

			// Check if this is a high pulse to rxFeeder
			if p.to == rxFeeder && p.high {
				if _, seen := cycles[p.from]; !seen {
					cycles[p.from] = buttonPresses
				}
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
				if !p.high {
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

	// LCM of all cycle lengths
	result := uint(1)
	for _, cycle := range cycles {
		result = lcm(result, cycle)
	}
	return result
}
