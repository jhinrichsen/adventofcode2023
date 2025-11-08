package adventofcode2023

import (
	"testing"
)

func TestDay00Part1Example(t *testing.T) {
	const want = 0
	lines := linesFromFilename(t, exampleFilename(0))
	got := Day00(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay00Part1(t *testing.T) {
	const want = 0
	lines := linesFromFilename(t, filename(0))
	got := Day00(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay00Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(0))
	b.ResetTimer()
	for b.Loop() {
		_ = Day00(lines)
	}
}

func TestDay00Part2(t *testing.T) {
	const want = 0
	lines := linesFromFilename(t, filename(0))
	got := Day00(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
