package adventofcode2023

import (
	"testing"
)

var day02ExampleTriple = Triple{12, 13, 14}

func TestDay02Part1Example(t *testing.T) {
	const want = 8
	lines, err := linesFromFilename(exampleFilename(2))
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day02(day02ExampleTriple, lines, true)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay02Part1(t *testing.T) {
	const want = 2207
	lines, err := linesFromFilename(filename(2))
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day02(day02ExampleTriple, lines, true)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay02Part1(b *testing.B) {
	lines, err := linesFromFilename(filename(2))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Day02(day02ExampleTriple, lines, true)
	}
}

func TestDay02Part2Example(t *testing.T) {
	const want = 2286
	lines, err := linesFromFilename(exampleFilename(2))
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day02(day02ExampleTriple, lines, false)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
func TestDay02Part2(t *testing.T) {
	const want = 62241
	lines, err := linesFromFilename(filename(2))
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day02(day02ExampleTriple, lines, false)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay02Part2(b *testing.B) {
	lines, err := linesFromFilename(filename(2))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = Day02(day02ExampleTriple, lines, false)
	}
}
