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

func TestDay15Part2(t *testing.T) {
	testLines(t, 15, filename, false, Day15, 0)
}
