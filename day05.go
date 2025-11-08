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
		// Part 2: process seed ranges instead of individual seeds
		// Convert seed pairs to ranges
		var seedRanges []Range
		for i := 0; i < len(seeds); i += 2 {
			start := uint(seeds[i])
			length := uint(seeds[i+1])
			seedRanges = append(seedRanges, Range{
				Min:   start,
				Max:   start + length - 1,
				Delta: 0,
			})
		}

		// Process each seed range through all map layers
		currentRanges := seedRanges
		for _, mapLayer := range rrs {
			currentRanges = processRangesThroughMap(currentRanges, mapLayer)
		}

		// Find minimum location from all final ranges
		// The ranges are already in location space after processing through all layers
		for _, r := range currentRanges {
			n0 = min(n0, r.Min)
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

// processRangesThroughMap takes a set of ranges and processes them through a map layer
// It transforms the ranges from source space to destination space
func processRangesThroughMap(inputRanges []Range, mapLayer Ranges) []Range {
	var result []Range

	for _, inputRange := range inputRanges {
		// Split the input range based on overlaps with map ranges
		unmappedRanges := []Range{inputRange}
		var mappedRanges []Range

		for _, mapRange := range mapLayer {
			var newUnmapped []Range
			for _, ur := range unmappedRanges {
				// Check for overlap between unmapped range and map range
				if ur.Max < mapRange.Min || ur.Min > mapRange.Max {
					// No overlap, keep in unmapped
					newUnmapped = append(newUnmapped, ur)
					continue
				}

				// There is overlap, split the range
				// Before overlap - stays in source space (unmapped)
				if ur.Min < mapRange.Min {
					newUnmapped = append(newUnmapped, Range{
						Min:   ur.Min,
						Max:   mapRange.Min - 1,
						Delta: 0,
					})
				}

				// Overlapping part - transform to destination space
				overlapMin := max(ur.Min, mapRange.Min)
				overlapMax := min(ur.Max, mapRange.Max)
				mappedRanges = append(mappedRanges, Range{
					Min:   uint(int(overlapMin) + mapRange.Delta),
					Max:   uint(int(overlapMax) + mapRange.Delta),
					Delta: 0,
				})

				// After overlap - stays in source space (unmapped)
				if ur.Max > mapRange.Max {
					newUnmapped = append(newUnmapped, Range{
						Min:   mapRange.Max + 1,
						Max:   ur.Max,
						Delta: 0,
					})
				}
			}
			unmappedRanges = newUnmapped
		}

		// Add both mapped and unmapped ranges to result
		result = append(result, mappedRanges...)
		result = append(result, unmappedRanges...)
	}

	return result
}
