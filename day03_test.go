package adventofcode2023

import (
	"testing"
)

func TestDay03Part1Example(t *testing.T) {
	testWithParser(t, 3, exampleFilename, true, NewDay03, Day03, 4361)
}

func TestDay03Part1(t *testing.T) {
	testWithParser(t, 3, filename, true, NewDay03, Day03, 539713)
}

// TestColoredOutput is for manual visualization only - prints colored grid output.
// To run: go test -run TestColoredOutput -v
func TestColoredOutput(t *testing.T) {
	t.Skip("Skipping colored output test - produces large visual output for manual verification only")
	const want = 539713
	lines := linesFromFilename(t, filename(3))
	got := Day03ColoredLogging(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay03Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(3))
	b.ResetTimer()
	for b.Loop() {
		puzzle, _ := NewDay03(lines)
		_ = Day03(puzzle, true)
	}
}

func TestDay03Part2Example(t *testing.T) {
	testWithParser(t, 3, exampleFilename, false, NewDay03, Day03, 467835)
}

func TestDay03Part2_116_12(t *testing.T) {
	const want = 4032
	puzzle, _ := NewDay03([]string{
		".672.",
		".*...",
		".6...",
	})
	got := Day03(puzzle, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay03Part2(t *testing.T) {
	testWithParser(t, 3, filename, false, NewDay03, Day03, 84159075)
}

func BenchmarkDay03Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(3))
	b.ResetTimer()
	for b.Loop() {
		puzzle, _ := NewDay03(lines)
		_ = Day03(puzzle, false)
	}
}
