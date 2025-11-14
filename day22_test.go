package adventofcode2023

import (
	"testing"
)

func TestDay22Part1Example(t *testing.T) {
	lines := []string{
		"1,0,1~1,2,1",
		"0,0,2~2,0,2",
		"0,2,3~2,2,3",
		"0,0,4~0,2,4",
		"2,0,5~2,2,5",
		"0,1,6~2,1,6",
		"1,1,8~1,1,9",
	}
	puzzle, err := NewDay22(lines)
	if err != nil {
		t.Fatal(err)
	}
	got := Day22(puzzle, true)
	const want = 5
	if got != want {
		t.Errorf("Example: got %d, want %d", got, want)
	}
}

func TestDay22Part1(t *testing.T) {
	testWithParser(t, 22, filename, true, NewDay22, Day22, 499)
}

func BenchmarkDay22Part1(b *testing.B) {
	benchWithParser(b, 22, true, NewDay22, Day22)
}

func TestDay22Part2(t *testing.T) {
	testWithParser(t, 22, filename, false, NewDay22, Day22, 0)
}

func BenchmarkDay22Part2(b *testing.B) {
	benchWithParser(b, 22, false, NewDay22, Day22)
}
