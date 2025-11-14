package adventofcode2023

import (
	"testing"
)

func TestDay13Part1Example(t *testing.T) {
	testLines(t, 13, exampleFilename, true, Day13, 405)
}

func TestDay13Part1(t *testing.T) {
	testLines(t, 13, filename, true, Day13, 33356)
}

func BenchmarkDay13Part1(b *testing.B) {
	benchLines(b, 13, true, Day13)
}

func TestDay13Part2(t *testing.T) {
	testLines(t, 13, filename, false, Day13, 0)
}
