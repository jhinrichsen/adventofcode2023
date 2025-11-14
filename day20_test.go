package adventofcode2023

import (
	"testing"
)

func TestDay20Part1Example1(t *testing.T) {
	lines := []string{
		"broadcaster -> a, b, c",
		"%a -> b",
		"%b -> c",
		"%c -> inv",
		"&inv -> a",
	}
	puzzle, err := NewDay20(lines)
	if err != nil {
		t.Fatal(err)
	}
	got := Day20(puzzle, true)
	const want = 32000000
	if got != want {
		t.Errorf("Example 1: got %d, want %d", got, want)
	}
}

func TestDay20Part1Example2(t *testing.T) {
	lines := []string{
		"broadcaster -> a",
		"%a -> inv, con",
		"&inv -> b",
		"%b -> con",
		"&con -> output",
	}
	puzzle, err := NewDay20(lines)
	if err != nil {
		t.Fatal(err)
	}
	got := Day20(puzzle, true)
	const want = 11687500
	if got != want {
		t.Errorf("Example 2: got %d, want %d", got, want)
	}
}

func TestDay20Part1(t *testing.T) {
	testWithParser(t, 20, filename, true, NewDay20, Day20, 879834312)
}

func BenchmarkDay20Part1(b *testing.B) {
	benchWithParser(b, 20, true, NewDay20, Day20)
}

func TestDay20Part2(t *testing.T) {
	testWithParser(t, 20, filename, false, NewDay20, Day20, 0)
}

func BenchmarkDay20Part2(b *testing.B) {
	benchWithParser(b, 20, false, NewDay20, Day20)
}
