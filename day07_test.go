package adventofcode2023

import (
	"testing"
)

func TestDay07Part1Example(t *testing.T) {
	const want = 6440
	lines, err := linesFromFilename(exampleFilename(07))
	if err != nil {
		t.Fatal(err)
	}
	hands, bids, err := NewDay07(lines)
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day07(hands, bids)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay07Part1(t *testing.T) {
	const want = 246424613
	lines, err := linesFromFilename(filename(07))
	if err != nil {
		t.Fatal(err)
	}
	hands, bids, err := NewDay07(lines)
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day07(hands, bids)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay07Part1(b *testing.B) {
	lines, err := linesFromFilename(filename(07))
	if err != nil {
		b.Fatal(err)
	}
	hands, bids, err := NewDay07(lines)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Day07(hands, bids)
	}
}

/*
func TestDay07Part2(t *testing.T) {
	const want = 0
	lines, err := linesFromFilename(filename(07))
	if err != nil {
		t.Fatal(err)
	}
	got := Day07(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
*/

func TestCard(t *testing.T) {
	const want = 0
	got, err := card('2')
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
