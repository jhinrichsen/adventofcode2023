package adventofcode2023

import (
	"math"
	"strings"
)

func Day05(lines []string, part1 bool) (uint, error) {
	seeds, err := lineAsNumbers(lines[0][len("seeds: "):])
	if err != nil {
		return 0, err
	}

	rrs, err := parseDay05(lines[2:])
	if err != nil {
		return 0, err
	}

	var n0 uint = math.MaxUint
	if part1 {
		// iterative brute force
		for _, seed := range seeds {
			n := uint(seed)
			for i := range rrs {
				n = rrs[i].Do(n)
			}
			n0 = min(n0, n)
		}
	} else {
		var total int
		// for part 2, the iterative approach is not going anywhere
		// estimated runtime is around 41 days
		//  => do not iterate single seeds, but consolidate consecutive ranges

		for i := 0; i < len(seeds); i += 2 {
			// fmt.Printf("seed range: [%d..%d] = %d\n", seeds[i], seeds[i]+seeds[i+1], seeds[i+1])
			total += seeds[i+1]
			// fmt.Printf("total: %d\n", total)
			/*
				for j := seeds[i]; j < seeds[i]+seeds[i+1]; j++ {
					n := all(uint(j))
					n0 = min(n0, n)
				}
			*/
		}
	}
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

func Merge(r1, r2 Range) Ranges {
	// case 1: equal
	if r1.Min == r2.Min && r1.Max == r2.Max {
		return Ranges{Range{r1.Min, r1.Max, r1.Delta + r2.Delta}}
	}

	// swap predicates
	p1 := r1.Min > r2.Min
	p2 := r1.Min == r2.Min && r1.Max < r2.Max // case 4
	if p1 || p2 {
		tmp := r1
		r1 = r2
		r2 = tmp
	}

	// case 2: non overlapping
	if r1.Max < r2.Min {
		return Ranges{
			r1, r2,
		}
	}

	// case 3: inside
	if r1.Min < r2.Min && r1.Max > r2.Max {
		return Ranges{Range{r1.Min, r2.Min - 1, r1.Delta},
			Range{r2.Min, r2.Max, r1.Delta + r2.Delta},
			Range{r2.Max + 1, r1.Max, r1.Delta}}
	}

	// case 4
	if r1.Min == r2.Min && r1.Max > r2.Max {
		return Ranges{
			Range{r1.Min, r2.Max, r1.Delta + r2.Delta},
			Range{r2.Max + 1, r1.Max, r1.Delta},
		}
	}

	// case 5
	if r1.Min < r2.Min && r1.Max < r2.Max {
		return Ranges{
			Range{r1.Min, r2.Min - 1, r1.Delta},
			Range{r2.Min, r1.Max, r1.Delta + r2.Delta},
			Range{r1.Max + 1, r2.Max, r2.Delta},
		}
	}

	// case 6
	if r1.Min < r2.Min && r1.Max == r2.Max {
		return Ranges{
			Range{r1.Min, r2.Min, r1.Delta},
			Range{r2.Min + 1, r1.Max, r1.Delta + r2.Delta},
		}
	}
	panic("unimplemented case")
}
