package adventofcode2023

import (
	"testing"
)

func TestDay07Part1Example(t *testing.T) {
	const want = 6440
	lines := linesFromFilename(t, exampleFilename(7))
	puzzle, err := NewDay07(lines, true)
	if err != nil {
		t.Fatal(err)
	}
	got := Day07(puzzle, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay07Part1(t *testing.T) {
	const want = 246424613
	lines := linesFromFilename(t, filename(7))
	puzzle, err := NewDay07(lines, true)
	if err != nil {
		t.Fatal(err)
	}
	got := Day07(puzzle, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay07Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(7))
	puzzle, err := NewDay07(lines, true)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for b.Loop() {
		_ = Day07(puzzle, true)
	}
}

func TestCard(t *testing.T) {
	const want = 0
	useJoker := false
	got, err := card('2', useJoker)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay07Part2Example(t *testing.T) {
	const want = 5905
	lines := linesFromFilename(t, exampleFilename(7))
	puzzle, err := NewDay07(lines, false)
	if err != nil {
		t.Fatal(err)
	}
	got := Day07(puzzle, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay07Part2(t *testing.T) {
	const want = 248256639
	lines := linesFromFilename(t, filename(7))
	puzzle, err := NewDay07(lines, false)
	if err != nil {
		t.Fatal(err)
	}
	got := Day07(puzzle, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestQ2KJJ(t *testing.T) {
	const (
		want = ThreeOfAKind
		s    = "Q2KJJ"
	)
	cs, err := cards(s)
	if err != nil {
		t.Fatal(err)
	}
	useJoker := true
	got := handType(Hand{cards: cs}, useJoker)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func cards(s string) ([]Card, error) {
	var hand = make([]Card, len(s))
	useJoker := true
	for i, b := range s {
		c, err := card(byte(b), useJoker)
		if err != nil {
			return nil, err
		}
		hand[i] = c
	}
	return hand, nil
}

func BenchmarkDay07Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(7))
	puzzle, err := NewDay07(lines, false)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for b.Loop() {
		_ = Day07(puzzle, false)
	}
}
