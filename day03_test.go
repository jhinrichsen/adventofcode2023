package adventofcode2023

import (
	"testing"
)

func TestDay03Part1Example(t *testing.T) {
	testLines(t, 3, exampleFilename, true, Day03, 4361)
}

func TestDay03Part1(t *testing.T) {
	testLines(t, 3, filename, true, Day03, 539713)
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
	benchLines(b, 3, true, Day03)
}

func TestDay03Part2Example(t *testing.T) {
	testLines(t, 3, exampleFilename, false, Day03, 467835)
}

func TestDay03Part2_116_12(t *testing.T) {
	const want = 4032
	lines := []string{
		".672.",
		".*...",
		".6...",
	}
	got := Day03(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay03Part2(t *testing.T) {
	testLines(t, 3, filename, false, Day03, 84159075)
}

func BenchmarkDay03Part2(b *testing.B) {
	benchLines(b, 3, false, Day03)
}
