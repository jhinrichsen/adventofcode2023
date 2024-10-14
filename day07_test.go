package adventofcode2023

import (
	"testing"
)

func TestDay07Part1Example(t *testing.T) {
	const (
		joker = false
		want  = 6440
	)
	lines, err := linesFromFilename(exampleFilename(07))
	if err != nil {
		t.Fatal(err)
	}
	hands, bids, err := NewDay07(lines, joker)
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day07(hands, bids, joker)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay07Part1(t *testing.T) {
	const (
		joker = false
		want  = 246424613
	)
	lines, err := linesFromFilename(filename(07))
	if err != nil {
		t.Fatal(err)
	}
	hands, bids, err := NewDay07(lines, joker)
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day07(hands, bids, joker)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay07Part1(b *testing.B) {
	const joker = false
	lines, err := linesFromFilename(filename(07))
	if err != nil {
		b.Fatal(err)
	}
	hands, bids, err := NewDay07(lines, joker)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Day07(hands, bids, joker)
	}
}

func TestCard(t *testing.T) {
	const want = 0
	var mu bool
	got, err := card('2', mu)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay07Part2Example(t *testing.T) {
	const (
		joker = true
		want  = 5905
	)
	lines, err := linesFromFilename(exampleFilename(07))
	if err != nil {
		t.Fatal(err)
	}
	hands, bids, err := NewDay07(lines, joker)
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day07(hands, bids, joker)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay07Part2Example2(t *testing.T) {
	const (
		joker = true
		want  = 6839
	)
	lines, err := linesFromFilename("testdata/day07_example2.txt")
	if err != nil {
		t.Fatal(err)
	}
	hands, bids, err := NewDay07(lines, joker)
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day07(hands, bids, joker)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay07Part2(t *testing.T) {
	const (
		joker = true
		// want  = 248822253 too high
		want = 248256639
	)
	lines, err := linesFromFilename(filename(07))
	if err != nil {
		t.Fatal(err)
	}
	hands, bids, err := NewDay07(lines, joker)
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day07(hands, bids, joker)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestQ2KJJ(t *testing.T) {
	const (
		want = ThreeOfAKind
		s    = "Q2KJJ"
	)
	hand, err := cards(s)
	if err != nil {
		t.Fatal(err)
	}
	got := handType(hand, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func cards(s string) ([]Card, error) {
	var hand = make([]Card, len(s))
	for i, b := range s {
		c, err := card(byte(b), true)
		if err != nil {
			return nil, err
		}
		hand[i] = c
	}
	return hand, nil
}

func BenchmarkDay07Part2(b *testing.B) {
	const joker = true
	lines, err := linesFromFilename(filename(07))
	if err != nil {
		b.Fatal(err)
	}
	hands, bids, err := NewDay07(lines, joker)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Day07(hands, bids, joker)
	}
}
