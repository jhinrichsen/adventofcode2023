package adventofcode2023

import (
	"testing"
)

func TestDay17Part1Example(t *testing.T) {
	testLines(t, 17, exampleFilename, true, Day17, 102)
}

func TestDay17Part1(t *testing.T) {
	testLines(t, 17, filename, true, Day17, 1065)
}

func BenchmarkDay17Part1(b *testing.B) {
	benchLines(b, 17, true, Day17)
}

func TestDay17Part2(t *testing.T) {
	testLines(t, 17, filename, false, Day17, 0)
}

func BenchmarkDay17Part2(b *testing.B) {
	benchLines(b, 17, false, Day17)
}
