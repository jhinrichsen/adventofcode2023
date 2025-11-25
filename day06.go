package adventofcode2023

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Day06Puzzle struct {
	times     []int
	distances []int
}

func NewDay06(lines []string) (Day06Puzzle, error) {
	const (
		wantLines    = 2
		wantTime     = "Time:"
		wantDistance = "Distance:"
	)
	gotLines := len(lines)
	if wantLines != gotLines {
		return Day06Puzzle{}, fmt.Errorf("want %d lines but got %d", wantLines, gotLines)
	}

	// parse times
	parts := strings.Fields(lines[0])
	if parts[0] != wantTime {
		return Day06Puzzle{}, fmt.Errorf("want first line to start with %s but got %s", wantTime, parts[0])
	}
	var ts []int
	for i, p := range parts[1:] {
		n, err := strconv.Atoi(p)
		if err != nil {
			return Day06Puzzle{}, fmt.Errorf("line %d: cannot convert to numeric: %s: %v", i, p, err)
		}
		ts = append(ts, n)
	}

	// parse distances
	parts = strings.Fields(lines[1])
	if parts[0] != wantDistance {
		return Day06Puzzle{}, fmt.Errorf("want second line to start with %s but got %s", wantTime, parts[0])
	}
	var ds []int
	for i, p := range parts[1:] {
		n, err := strconv.Atoi(p)
		if err != nil {
			return Day06Puzzle{}, fmt.Errorf("line %d: cannot convert to numeric: %s: %v", i, p, err)
		}
		ds = append(ds, n)
	}

	return Day06Puzzle{times: ts, distances: ds}, nil
}

func Day06(puzzle Day06Puzzle, part1 bool) uint {
	if len(puzzle.times) != len(puzzle.distances) {
		return 0
	}

	if !part1 {
		// Part 2: Concatenate all digits into single race
		var timeStr, distStr string
		for i := range puzzle.times {
			timeStr += strconv.Itoa(puzzle.times[i])
			distStr += strconv.Itoa(puzzle.distances[i])
		}
		time, _ := strconv.Atoi(timeStr)
		dist, _ := strconv.Atoi(distStr)
		return uint(farther(time, dist))
	}

	// Part 1: Multiple races
	result := 1
	for i := range puzzle.times {
		n := farther(puzzle.times[i], puzzle.distances[i])
		result *= n
	}
	return uint(result)
}

// farther calculates the number of ways to beat the distance record.
// This is a quadratic equation: we want h*(t-h) > d
// Rearranging: -h² + h*t - d > 0, or h² - h*t + d < 0
// Using quadratic formula: h = (t ± sqrt(t² - 4d)) / 2
// The solutions give us the boundary points where distance equals d.
// We want integers strictly between these boundaries.
func farther(t, d int) int {
	// Convert to float64 for math operations
	tf := float64(t)
	df := float64(d)

	// Discriminant: t² - 4d
	discriminant := tf*tf - 4*df
	if discriminant < 0 {
		// No real solutions - can't beat the record
		return 0
	}

	sqrtDisc := math.Sqrt(discriminant)
	// Two roots of the quadratic equation
	h1 := (tf - sqrtDisc) / 2
	h2 := (tf + sqrtDisc) / 2

	// We need integers strictly greater than h1 and strictly less than h2
	// Use ceiling for lower bound and floor for upper bound
	minH := int(math.Floor(h1 + 1))
	maxH := int(math.Ceil(h2 - 1))

	// Handle edge case where roots are exact integers (equality case)
	// If h1 or h2 are integers, we need strict inequality
	if h1 == math.Floor(h1) {
		minH = int(h1) + 1
	}
	if h2 == math.Ceil(h2) {
		maxH = int(h2) - 1
	}

	// Count of integers in range [minH, maxH]
	if maxH < minH {
		return 0
	}
	return maxH - minH + 1
}
