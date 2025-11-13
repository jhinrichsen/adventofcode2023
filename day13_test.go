package adventofcode2023

import (
	"testing"
)

func TestDay13Part1Example(t *testing.T) {
	const want = 405
	lines := linesFromFilename(t, exampleFilename(13))
	grids := NewDay13(lines)
	got := Day13(grids, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay13Part1(t *testing.T) {
	const want = 33356
	lines := linesFromFilename(t, filename(13))
	grids := NewDay13(lines)
	got := Day13(grids, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay13Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(13))
	b.ResetTimer()
	for b.Loop() {
		grids := NewDay13(lines)
		_ = Day13(grids, true)
	}
}

func TestDay13Part2(t *testing.T) {
	const want = 0
	lines := linesFromFilename(t, filename(13))
	grids := NewDay13(lines)
	got := Day13(grids, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
