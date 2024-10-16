package adventofcode2023

import (
	"testing"
)

func TestDay08Part1Example(t *testing.T) {
	const want = 6
	lines, err := linesFromFilename(exampleFilename(8))
	diet(t, err)
	d8, err := NewDay08(lines)
	diet(t, err)
	got := Day08Part1(d8)
	diet(t, err)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay08Part1(t *testing.T) {
	const want = 18727
	lines, err := linesFromFilename(filename(8))
	diet(t, err)
	d8, err := NewDay08(lines)
	diet(t, err)
	got := Day08Part1(d8)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay08Part1(b *testing.B) {
	lines, err := linesFromFilename(filename(8))
	if err != nil {
		b.Fatal(err)
	}
	d8, err := NewDay08(lines)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Day08Part1(d8)
	}
}

func TestDay08Part2Example(t *testing.T) {
	const (
		filename = "testdata/day08_example_part2.txt"
		want     = 6
	)
	lines, err := linesFromFilename(filename)
	diet(t, err)
	d8, err := NewDay08(lines)
	diet(t, err)
	got := Day08Part2(d8)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay08Part2(t *testing.T) {
	if testing.Short() {
		t.Skip("brute force approach will run forever")
	}
	const want = 0
	lines, err := linesFromFilename(filename(8))
	diet(t, err)
	d8, err := NewDay08(lines)
	diet(t, err)
	got := Day08Part2(d8)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func diet(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}
