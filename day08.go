package adventofcode2023

import (
	"fmt"
	"strings"
)

const (
	start  = "AAA"
	finish = "ZZZ"
)

type Tuple[E any] struct {
	A, B E
}

func (t Tuple[E]) Len() int {
	return 2
}

func (t Tuple[E]) String() string {
	return fmt.Sprintf("(%v/%v)", t.A, t.B)
}

type D8 struct {
	Instructions string
	Idx          int // index of next instruction
	Network      map[string]Tuple[string]
	Current      string
}

// complete returns nil if the network has all connections.
func (a D8) complete() error {
	for k, v := range a.Network {
		if _, ok := a.Network[v.A]; !ok {
			return fmt.Errorf("Node %q references left node %q which does not exist", k, v.A)
		}
		if _, ok := a.Network[v.B]; !ok {
			return fmt.Errorf("Node %q references right node %q which does not exist", k, v.A)
		}
	}
	return nil
}

func (a *D8) Next() bool {
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

func Day08(d8 D8) uint {
	var steps uint
	for d8.Next() {
		steps++
	}
	return 1 + steps
}

func NewDay08(lines []string) (D8, error) {
	const (
		minLines = 4 // instructions, empty, 2 nodes
	)

	var d8 D8
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
	if _, ok := d8.Network[start]; !ok {
		return d8, fmt.Errorf("error: missing start node %q", start)
	}
	if _, ok := d8.Network[finish]; !ok {
		return d8, fmt.Errorf("error: missing finish node %q", finish)
	}
	return d8, d8.complete()
}
