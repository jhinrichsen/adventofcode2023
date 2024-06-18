package adventofcode2023

import (
	"testing"
)

func TestDay04Part1V1Example(t *testing.T) {
	const want = 13
	lines, err := linesFromFilename(exampleFilename(4))
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day04Part1V1(lines)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay04Part1V1(t *testing.T) {
	const want = 26218
	lines, err := linesFromFilename(filename(4))
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day04Part1V1(lines)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay04Part1V1(b *testing.B) {
	lines, err := linesFromFilename(filename(4))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Day04Part1V1(lines)
	}
}

func testDay04Example(t *testing.T, part1 bool, want uint) {
	lines, err := linesFromFilename(exampleFilename(4))
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day04(lines, part1)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay04Part1Example(t *testing.T) {
	testDay04Example(t, true, 13)
}

func TestDay04Part2Example(t *testing.T) {
	testDay04Example(t, false, 30)
}

func TestDay04Part1(t *testing.T) {
	const want = 26218
	lines, err := linesFromFilename(filename(4))
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day04(lines, true)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay04Part2(t *testing.T) {
	const want = 9997537
	lines, err := linesFromFilename(filename(4))
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day04(lines, false)
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
		_, _ = Day04(lines, true)
	}
}

func BenchmarkDay04Part2(b *testing.B) {
	lines, err := linesFromFilename(filename(4))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Day04(lines, false)
	}
}
