package adventofcode2023

import (
	"testing"
)

func TestDay03Part1Example(t *testing.T) {
	const want = 4361
	lines, err := linesFromFilename(exampleFilename(3))
	if err != nil {
		t.Fatal(err)
	}
	got := Day03Part1(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay03Part1(t *testing.T) {
	// too low const want = 538121
	// too low const want = 538257
	const want = 539713
	lines, err := linesFromFilename(filename(3))
	if err != nil {
		t.Fatal(err)
	}
	got := Day03Part1(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestColoredOutput(t *testing.T) {
	const want = 539713
	lines, err := linesFromFilename(filename(3))
	if err != nil {
		t.Fatal(err)
	}
	got := Day03ColoredLogging(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay03Part1(b *testing.B) {
	lines, err := linesFromFilename(filename(3))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Day03Part1(lines)
	}
}

func TestDay03Part2Example(t *testing.T) {
	const want = 467835
	lines, err := linesFromFilename(exampleFilename(3))
	if err != nil {
		t.Fatal(err)
	}
	got := Day03Part2(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay03Part2_116_12(t *testing.T) {
	const want = 4032
	got := Day03Part2([]string{
		".672.",
		".*...",
		".6...",
	})
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay03Part2(t *testing.T) {
	// const want = 84138748 // too low
	const want = 84159075
	lines, err := linesFromFilename(filename(3))
	if err != nil {
		t.Fatal(err)
	}
	got := Day03Part2(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay03Part2(b *testing.B) {
	lines, err := linesFromFilename(filename(3))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for range b.N {
		_ = Day03Part2(lines)
	}
}
