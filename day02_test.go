package adventofcode2023

import (
	"testing"
)

func TestDay02Part1Example(t *testing.T) {
	testLines(t, 2, exampleFilename, true, Day02, 8)
}

func TestDay02Part1(t *testing.T) {
	testLines(t, 2, filename, true, Day02, 2207)
}

func TestDay02Part2Example(t *testing.T) {
	testLines(t, 2, exampleFilename, false, Day02, 2286)
}

func TestDay02Part2(t *testing.T) {
	testLines(t, 2, filename, false, Day02, 62241)
}

func BenchmarkDay02Part1(b *testing.B) {
	benchLines(b, 2, true, Day02)
}

func BenchmarkDay02Part2(b *testing.B) {
	benchLines(b, 2, false, Day02)
}
