package adventofcode2023

import (
	"testing"
)

func TestDay02Part1Example(t *testing.T) {
	testWithParser(t, 2, exampleFilename, true, NewDay02, Day02, 8)
}

func TestDay02Part1(t *testing.T) {
	testWithParser(t, 2, filename, true, NewDay02, Day02, 2207)
}

func TestDay02Part2Example(t *testing.T) {
	testWithParser(t, 2, exampleFilename, false, NewDay02, Day02, 2286)
}

func TestDay02Part2(t *testing.T) {
	testWithParser(t, 2, filename, false, NewDay02, Day02, 62241)
}

func BenchmarkDay02Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(2))
	b.ResetTimer()
	for b.Loop() {
		puzzle, _ := NewDay02(lines)
		_ = Day02(puzzle, true)
	}
}

func BenchmarkDay02Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(2))
	b.ResetTimer()
	for b.Loop() {
		puzzle, _ := NewDay02(lines)
		_ = Day02(puzzle, false)
	}
}
