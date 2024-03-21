package adventofcode2023

import (
	"testing"
)

func TestDay03Part1Example(t *testing.T) {
	const want = 4361
	lines, err := linesFromFilename(exampleFilename(0))
	if err != nil {
		t.Fatal(err)
	}
	got := Day03(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay03Part1(t *testing.T) {
	const want = 0
	lines, err := linesFromFilename(filename(0))
	if err != nil {
		t.Fatal(err)
	}
	got := Day03(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay03Part1(b *testing.B) {
	lines, err := linesFromFilename(filename(0))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Day03(lines)
	}
}

func TestDay03Part2(t *testing.T) {
	const want = 0
	lines, err := linesFromFilename(filename(0))
	if err != nil {
		t.Fatal(err)
	}
	got := Day03(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
