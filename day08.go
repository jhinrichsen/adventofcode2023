package adventofcode2023

import (
	"fmt"
	"strings"
)

const (
	start  = "AAA"
	finish = "ZZZ"
)

type D08Puzzle struct {
	Instructions string
	Idx          int // index of next instruction
	Network      map[string]Tuple[string]
	Current      string
}

// complete returns nil if the network has all connections.
func (a D08Puzzle) complete() error {
	for k, v := range a.Network {
		if _, ok := a.Network[v.A]; !ok {
			return fmt.Errorf("node %q references left node %q which does not exist", k, v.A)
		}
		if _, ok := a.Network[v.B]; !ok {
			return fmt.Errorf("node %q references right node %q which does not exist", k, v.A)
		}
	}
	return nil
}

func (a *D08Puzzle) Next() bool {
	t := a.Network[a.Current]
	if a.Instructions[a.Idx] == 'L' {
		a.Current = t.A
	} else {
		a.Current = t.B
	}
	if a.Current == finish {
		return false
	}

	a.Idx++
	if a.Idx == len(a.Instructions) {
		a.Idx = 0
	}
	return true
}

// next2 moves to the next node without checking for finish (for Part 2)
func (a *D08Puzzle) next2() {
	t := a.Network[a.Current]
	if a.Instructions[a.Idx] == 'L' {
		a.Current = t.A
	} else {
		a.Current = t.B
	}

	a.Idx++
	if a.Idx == len(a.Instructions) {
		a.Idx = 0
	}
}

func Day08(d8 D08Puzzle, part1 bool) uint {
	if part1 {
		var steps uint
		for d8.Next() {
			steps++
		}
		return 1 + steps
	}

	starts := d8.startNodes()

	// Find cycle length for each starting node
	var cycleLengths []uint
	for _, start := range starts {
		d := d8
		d.Network = d8.Network
		d.Current = start

		var steps uint
		for !finishNode(d.Current) {
			d.next2()
			steps++
		}
		cycleLengths = append(cycleLengths, steps)
	}

	// Calculate LCM of all cycle lengths
	result := cycleLengths[0]
	for i := 1; i < len(cycleLengths); i++ {
		result = lcm(result, cycleLengths[i])
	}
	return result
}

// gcd calculates the greatest common divisor
func gcd(a, b uint) uint {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// lcm calculates the least common multiple
func lcm(a, b uint) uint {
	return a * b / gcd(a, b)
}

func (a D08Puzzle) startNodes() (nodes []string) {
	for k := range a.Network {
		if startNode(k) {
			nodes = append(nodes, k)
		}
	}
	return
}

func startNode(s string) bool {
	return s[len(s)-1] == 'A'
}

func finishNode(s string) bool {
	return s[len(s)-1] == 'Z'
}

func NewDay08(lines []string) (D08Puzzle, error) {
	const (
		minLines = 4 // instructions, empty, 2 nodes
	)

	var d8 D08Puzzle
	d8.Network = make(map[string]Tuple[string], len(lines)-2)

	if len(lines) <= minLines {
		return d8, fmt.Errorf("want at least %d lines but got %d", minLines, len(lines))
	}

	d8.Instructions = lines[0]

	for i, line := range lines[2:] {
		var node, left, right string // AAA = (BBB, CCC)

		before, after, ok := strings.Cut(line, "=")
		if !ok {
			return d8, fmt.Errorf("line %d: want %q but got %q", i+1, "a = b", line)
		}

		node = strings.TrimSpace(before)
		s := strings.TrimSpace(after)
		if !strings.HasPrefix(s, "(") || !strings.HasSuffix(s, ")") {
			return d8, fmt.Errorf("line %d: want %q but got %q", i+1, "a = (b)", s)
		}
		s = s[1 : len(s)-1]
		left, right, ok = strings.Cut(s, ",")
		if !ok {
			return d8, fmt.Errorf("line %d: want %q but got %q", i+1, "a = (b, c)", s)
		}
		left = strings.TrimSpace(left)
		right = strings.TrimSpace(right)
		d8.Network[node] = Tuple[string]{left, right}
	}

	// make sure we have a start, and an end
	d8.Current = start
	return d8, d8.complete()
}
