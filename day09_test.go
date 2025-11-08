package adventofcode2023

import (
	"os"
	"testing"
)

func TestDay09V1Part1Example(t *testing.T) {
	const want = 114
	lines := linesFromFilename(t, exampleFilename(9))
	got := Day09V1(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay09V2Part1Example(t *testing.T) {
	const (
		part1 = true
		want  = 114
	)
	buf, err := os.ReadFile(exampleFilename(9))
	if err != nil {
		t.Fatal(err)
	}
	got := Day09V2(buf, part1)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay09V1Part1(t *testing.T) {
	const want = 2075724761
	lines := linesFromFilename(t, filename(9))
	got := Day09V1(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay09V2Part1(t *testing.T) {
	const (
		part1 = true
		want  = 2075724761
	)
	buf, err := os.ReadFile(filename(9))
	if err != nil {
		t.Fatal(err)
	}
	got := Day09V2(buf, part1)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay09V1Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(9))
	b.ResetTimer()
	for b.Loop() {
		_ = Day09V1(lines)
	}
}

func BenchmarkDay09V2Part1(b *testing.B) {
	buf, _ := os.ReadFile(filename(9))
	b.ResetTimer()
	for b.Loop() {
		_ = Day09V2(buf, true)
	}
}

func BenchmarkDay09Part1(b *testing.B) {
	buf, _ := os.ReadFile(filename(9))
	b.ResetTimer()
	b.ReportAllocs()
	for b.Loop() {
		_ = Day09V2(buf, true)
	}
}

func TestDay09Part2Example(t *testing.T) {
	const (
		part1 = false
		want  = 2
	)
	buf, err := os.ReadFile(exampleFilename(9))
	if err != nil {
		t.Fatal(err)
	}
	got := Day09V2(buf, part1)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay09Part2(t *testing.T) {
	const (
		part1 = false
		want  = 1072
	)
	buf, err := os.ReadFile(filename(9))
	if err != nil {
		t.Fatal(err)
	}
	got := Day09V2(buf, part1)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay09Part2(b *testing.B) {
	buf, _ := os.ReadFile(filename(9))
	b.ResetTimer()
	for b.Loop() {
		_ = Day09V2(buf, false)
	}
}
