package adventofcode2023

import (
	"os"
	"testing"
)

func TestDay01Part1ExampleV1(t *testing.T) {
	const want = 142
	lines, err := linesFromFilename(exampleFilename(1))
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day01V1(lines)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay01Part1ExampleV2(t *testing.T) {
	const want = 142
	buf, err := os.ReadFile(exampleFilename(1))
	if err != nil {
		t.Fatal(err)
	}
	got := Day01V2(buf)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay01Part1(t *testing.T) {
	const want = 55130
	buf, err := os.ReadFile(filename(1))
	if err != nil {
		t.Fatal(err)
	}
	got := Day01V2(buf)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay01V1(b *testing.B) {
	lines, err := linesFromFilename(filename(1))
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		_, _ = Day01V1(lines)
	}
}

func BenchmarkDay01V2(b *testing.B) {
	buf, err := os.ReadFile(filename(1))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Day01V2(buf)
	}
}

// BenchmarkDay01Large processes a 1 MB input file, the original input
// repeated until 1.000.000 bytes are filled, plus the rest of the line.
func BenchmarkDay01Large(b *testing.B) {
	buf, err := os.ReadFile("testdata/day01_large.txt")
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Day01V2(buf)
	}
}
