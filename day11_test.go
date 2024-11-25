package adventofcode2023

import (
	"testing"
)

func TestDay11Part1Example(t *testing.T) {
	const want = 374
	bytes, err := bytesFromFilename(exampleFilename(11))
	if err != nil {
		t.Fatal(err)
	}
	got := Day11(bytes)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay11Part1(t *testing.T) {
	const want = 9563821
	bytes, err := bytesFromFilename(filename(11))
	if err != nil {
		t.Fatal(err)
	}
	got := Day11(bytes)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay11Part1(b *testing.B) {
	bytes, err := bytesFromFilename(filename(11))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Day11(bytes)
	}
}

func TestDay11Part2(t *testing.T) {
	const want = 0
	bytes, err := bytesFromFilename(filename(11))
	if err != nil {
		t.Fatal(err)
	}
	got := Day11(bytes)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
