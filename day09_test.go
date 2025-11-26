package adventofcode2023

import (
	"testing"
)

func TestDay09Part1Example(t *testing.T) {
	testSolver(t, 9, exampleFilename, true, Day09, 114)
}

func TestDay09Part1(t *testing.T) {
	testSolver(t, 9, filename, true, Day09, 2075724761)
}

func TestDay09V1Part1Example(t *testing.T) {
	testLines(t, 9, exampleFilename, true, Day09V1, 114)
}

func TestDay09V1Part1(t *testing.T) {
	testLines(t, 9, filename, true, Day09V1, 2075724761)
}

func TestDay09V2Part1Example(t *testing.T) {
	testSolver(t, 9, exampleFilename, true, Day09V2, 114)
}

func TestDay09V2Part1(t *testing.T) {
	testSolver(t, 9, filename, true, Day09V2, 2075724761)
}

func BenchmarkDay09Part1(b *testing.B) {
	buf := file(b, 9)
	b.ResetTimer()
	for b.Loop() {
		_, _ = Day09(buf, true)
	}
}

func BenchmarkDay09V1Part1(b *testing.B) {
	benchLines(b, 9, true, Day09V1)
}

func BenchmarkDay09V2Part1(b *testing.B) {
	buf := file(b, 9)
	b.ResetTimer()
	for b.Loop() {
		_, _ = Day09V2(buf, true)
	}
}

func TestDay09Part2Example(t *testing.T) {
	testSolver(t, 9, exampleFilename, false, Day09, 2)
}

func TestDay09Part2(t *testing.T) {
	testSolver(t, 9, filename, false, Day09, 1072)
}

func TestDay09V1Part2Example(t *testing.T) {
	testLines(t, 9, exampleFilename, false, Day09V1, 2)
}

func TestDay09V1Part2(t *testing.T) {
	testLines(t, 9, filename, false, Day09V1, 1072)
}

func TestDay09V2Part2Example(t *testing.T) {
	testSolver(t, 9, exampleFilename, false, Day09V2, 2)
}

func TestDay09V2Part2(t *testing.T) {
	testSolver(t, 9, filename, false, Day09V2, 1072)
}

func BenchmarkDay09Part2(b *testing.B) {
	buf := file(b, 9)
	b.ResetTimer()
	for b.Loop() {
		_, _ = Day09(buf, false)
	}
}

func BenchmarkDay09V1Part2(b *testing.B) {
	benchLines(b, 9, false, Day09V1)
}

func BenchmarkDay09V2Part2(b *testing.B) {
	buf := file(b, 9)
	b.ResetTimer()
	for b.Loop() {
		_, _ = Day09V2(buf, false)
	}
}
