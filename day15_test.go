package adventofcode2023

import (
	"testing"
)

func TestDay15Part1Example(t *testing.T) {
	testLines(t, 15, exampleFilename, true, Day15, 1320)
}

func TestDay15Part1(t *testing.T) {
	testLines(t, 15, filename, true, Day15, 518107)
}

func BenchmarkDay15Part1(b *testing.B) {
	benchLines(b, 15, true, Day15)
}

func BenchmarkDay15Part2(b *testing.B) {
	benchLines(b, 15, false, Day15)
}

func TestDay15Part2Example(t *testing.T) {
	testLines(t, 15, exampleFilename, false, Day15, 145)
}

func TestDay15Part2(t *testing.T) {
	testLines(t, 15, filename, false, Day15, 303404)
}
