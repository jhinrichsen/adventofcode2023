package adventofcode2023

import (
	"testing"
)

func TestDay08Part1Example(t *testing.T) {
	const want = 6
	lines := linesFromFilename(t, exampleFilename(8))
	d8, err := NewDay08(lines)
	diet(t, err)
	got := Day08(d8, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay08Part1(t *testing.T) {
	const want = 18727
	lines := linesFromFilename(t, filename(8))
	d8, err := NewDay08(lines)
	diet(t, err)
	got := Day08(d8, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay08Part2Example(t *testing.T) {
	const (
		filename = "testdata/day08_example_part2.txt"
		want     = 6
	)
	lines := linesFromFilename(t, filename)
	d8, err := NewDay08(lines)
	diet(t, err)
	got := Day08(d8, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

/* TODO remove
func TestDay08Part2(t *testing.T) {
	const want = 18024643846273
	lines := linesFromFilename(t, filename(8))
	d8, err := NewDay08(lines)
	diet(t, err)
	got := Day08Part2(d8)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
*/

func TestDay08Part2(t *testing.T) {
	testWithParser(t, 8, filename, false, NewDay08, Day08, 18024643846273)
}

func diet(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func BenchmarkDay08Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(8))
	for b.Loop() {
		d8, _ := NewDay08(lines)
		_ = Day08(d8, true)
	}
}

func BenchmarkDay08Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(8))
	for b.Loop() {
		d8, _ := NewDay08(lines)
		_ = Day08(d8, false)
	}
}
