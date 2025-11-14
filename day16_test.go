package adventofcode2023

import (
	"testing"
)

func TestDay16Part1Example(t *testing.T) {
	testLines(t, 16, exampleFilename, true, Day16, 46)
}

func TestDay16Part1(t *testing.T) {
	testLines(t, 16, filename, true, Day16, 7111)
}

func BenchmarkDay16Part1(b *testing.B) {
	benchLines(b, 16, true, Day16)
}

func TestDay16Part2(t *testing.T) {
	testLines(t, 16, filename, false, Day16, 0)
}
