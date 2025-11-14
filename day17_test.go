package adventofcode2023

import (
	"testing"
)

func TestDay17Part1Example(t *testing.T) {
	testLines(t, 17, exampleFilename, true, Day17, 102)
}

func TestDay17Part1(t *testing.T) {
	testLines(t, 17, filename, true, Day17, 1065)
}

func BenchmarkDay17Part1(b *testing.B) {
	benchLines(b, 17, true, Day17)
}

func TestDay17Part2Example(t *testing.T) {
	testLines(t, 17, exampleFilename, false, Day17, 94)
}

func TestDay17Part2Example2(t *testing.T) {
	lines := linesFromFilename(t, "testdata/day17_example2.txt")
	got := Day17(lines, false)
	const want = 71
	if got != want {
		t.Errorf("Part2 example2: got %d, want %d", got, want)
	}
}

func TestDay17Part2(t *testing.T) {
	testLines(t, 17, filename, false, Day17, 1249)
}

func BenchmarkDay17Part2(b *testing.B) {
	benchLines(b, 17, false, Day17)
}
