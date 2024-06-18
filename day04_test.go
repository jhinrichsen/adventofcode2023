package adventofcode2023

import (
	"testing"
)

func TestDay04Part1Example(t *testing.T) {
	const want = 13
	lines, err := linesFromFilename(exampleFilename(4))
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day04(lines)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay04Part1(t *testing.T) {
	const want = 26218
	lines, err := linesFromFilename(filename(4))
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day04(lines)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay04Part1(b *testing.B) {
	lines, err := linesFromFilename(filename(4))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Day04(lines)
	}
}

func TestDay04Part2(t *testing.T) {
	const want = 0
	lines, err := linesFromFilename(filename(4))
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day04(lines)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
