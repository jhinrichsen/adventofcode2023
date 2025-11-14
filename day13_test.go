package adventofcode2023

import (
	"testing"
)

func TestDay13Part1Example(t *testing.T) {
	testWithParser(t, 13, exampleFilename, true, NewDay13, Day13, 405)
}

func TestDay13Part1(t *testing.T) {
	testWithParser(t, 13, filename, true, NewDay13, Day13, 33356)
}

func BenchmarkDay13Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(13))
	b.ResetTimer()
	for b.Loop() {
		puzzle, _ := NewDay13(lines)
		_ = Day13(puzzle, true)
	}
}

func TestDay13Part2(t *testing.T) {
	testWithParser(t, 13, filename, false, NewDay13, Day13, 0)
}
