package adventofcode2023

import (
	"testing"
)

func TestDay06Part1Example(t *testing.T) {
	const want = 288
	lines := linesFromFilename(t, exampleFilename(6))
	ts, ds, err := NewDay06(lines)
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day06(ts, ds)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay06Part1(t *testing.T) {
	const want = 440000
	lines := linesFromFilename(t, filename(6))
	ts, ds, err := NewDay06(lines)
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day06(ts, ds)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay06Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(6))
	ts, ds, err := NewDay06(lines)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for b.Loop() {
		_, _ = Day06(ts, ds)
	}
}

func TestDay06Part2Example(t *testing.T) {
	const want = 71503
	got, err := Day06([]int{71530}, []int{940200})
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay06Part2(t *testing.T) {
	const want = 26187338
	got, err := Day06([]int{42686985}, []int{284100511221341})
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay06Part2(b *testing.B) {
	for b.Loop() {
		_, _ = Day06([]int{42686985}, []int{284100511221341})
	}
}
