package adventofcode2023

import (
	"slices"
	"strings"
)

func Day05(lines []string) (uint, error) {
	seeds, err := lineAsNumbers(lines[0][len("seeds: "):])
	if err != nil {
		return 0, err
	}
	rrs, err := parseDay05(lines[2:])
	if err != nil {
		return 0, err
	}

	var results []uint
	for _, seed := range seeds {
		n := uint(seed)
		for i := range rrs {
			n = rrs[i].Do(n)
		}
		results = append(results, n)
	}
	n0 := min(slices.Min(results))
	return n0, nil
}

func parseDay05(lines []string) ([]Ranges, error) {
	var (
		rss []Ranges
		rs  Ranges
	)
	for _, line := range lines {
		if strings.HasSuffix(line, "map:") {
			rs = nil
		} else if len(line) == 0 {
			rss = append(rss, rs)
		} else {
			ns, err := lineAsNumbers(line)
			if err != nil {
				return rss, err
			}
			r := NewRange(uint(ns[0]), uint(ns[1]), uint(ns[2]))
			rs = append(rs, r)
		}
	}
	rss = append(rss, rs)
	return rss, nil
}

type Map struct {
	Destination, Source, Range uint
}

func NewRange(destination, source, length uint) Range {
	return Range{
		Min:   source,
		Max:   source + length - 1,
		Delta: int(destination - source),
	}
}

type Range struct {
	Min, Max uint // [min..max] included
	Delta    int
}

var IdentityRange = Range{0, 0, 0}

type Ranges []Range

func (a Ranges) Do(n uint) uint {
	r := a.Find(n)
	return uint(int(n) + r.Delta)
}

func (a Ranges) Find(n uint) Range {
	for _, r := range a {
		if r.Min <= n && n <= r.Max {
			return r
		}
	}
	return IdentityRange
}
