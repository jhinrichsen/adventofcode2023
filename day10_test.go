package adventofcode2023

import (
	"testing"
)

func TestDay10Part1ExampleChatGPT(t *testing.T) {
	const want = 4
	got, err := Day10ChatGPT()
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay10Part1Example2(t *testing.T) {
	const want = 8
	lines, err := linesFromFilename("testdata/day10_example2.txt")
	if err != nil {
		t.Fatal(err)
	}
	got := Day10(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay10Part1(t *testing.T) {
	const want = 0
	lines, err := linesFromFilename(filename(10))
	if err != nil {
		t.Fatal(err)
	}
	got := Day10(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay10Part1(b *testing.B) {
	lines, err := linesFromFilename(filename(10))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Day10(lines)
	}
}

func TestDay10Part2(t *testing.T) {
	const want = 0
	lines, err := linesFromFilename(filename(10))
	if err != nil {
		t.Fatal(err)
	}
	got := Day10(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
