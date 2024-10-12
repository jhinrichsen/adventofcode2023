package adventofcode2023

import (
	"fmt"
	"strconv"
	"strings"
)

func Day06(ts, ds []int) (int, error) {
	if len(ts) != len(ds) {
		return 0, fmt.Errorf("need two arrays of same length but got %d and %d", len(ts), len(ds))
	}
	rc := 1
	for i := range ts {
		n := Farther(ts[i], ds[i])
		rc *= n
	}
	return rc, nil
}

func NewDay06(lines []string) ([]int, []int, error) {
	const (
		wantLines    = 2
		wantTime     = "Time:"
		wantDistance = "Distance:"
	)
	gotLines := len(lines)
	if wantLines != gotLines {
		return nil, nil, fmt.Errorf("want %d lines but got %d", wantLines, gotLines)
	}
	var ts, ds []int

	// parse times

	parts := strings.Fields(lines[0])
	if parts[0] != wantTime {
		return nil, nil, fmt.Errorf("want first line to start with %s but got %s", wantTime, parts[0])
	}
	for i, p := range parts[1:] {
		n, err := strconv.Atoi(p)
		if err != nil {
			return nil, nil, fmt.Errorf("line %d: cannot convert to numeric: %s: %v", i, p, err)
		}
		ts = append(ts, n)
	}

	// parse Distances

	parts = strings.Fields(lines[1])
	if parts[0] != wantDistance {
		return nil, nil, fmt.Errorf("want second line to start with %s but got %s", wantTime, parts[0])
	}
	for i, p := range parts[1:] {
		n, err := strconv.Atoi(p)
		if err != nil {
			return nil, nil, fmt.Errorf("line %d: cannot convert to numeric: %s: %v", i, p, err)
		}
		ds = append(ds, n)
	}

	return ts, ds, nil
}

// Farther calculates the distances for each possibility and returns the number of possibilities where distance is larger than d.
func Farther(t, d int) int {
	var n int
	for i := 1; i < t-1; i++ {
		speed := i
		travel := speed * (t - i)
		if travel > d {
			n++
		}
	}
	return n
}
