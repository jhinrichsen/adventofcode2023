package adventofcode2023

import (
	"testing"
)

func TestDay21Part1Example(t *testing.T) {
	lines := []string{
		"...........",
		".....###.#.",
		".###.##..#.",
		"..#.#...#..",
		"....#.#....",
		".##..S####.",
		".##..#...#.",
		".......##..",
		".##.#.####.",
		".##..##.##.",
		"...........",
	}
	puzzle, err := NewDay21(lines)
	if err != nil {
		t.Fatal(err)
	}
	got := countReachable(puzzle, 6)
	const want = 16
	if got != want {
		t.Errorf("Example: got %d, want %d", got, want)
	}
}

func TestDay21Part1(t *testing.T) {
	testWithParser(t, 21, filename, true, NewDay21, Day21, 3751)
}

func BenchmarkDay21Part1(b *testing.B) {
	benchWithParser(b, 21, true, NewDay21, Day21)
}

func TestDay21Part2(t *testing.T) {
	testWithParser(t, 21, filename, false, NewDay21, Day21, 0)
}

func BenchmarkDay21Part2(b *testing.B) {
	benchWithParser(b, 21, false, NewDay21, Day21)
}
