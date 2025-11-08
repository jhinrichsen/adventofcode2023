package adventofcode2023

import (
	"testing"
)

func TestDay04Part1V1Example(t *testing.T) {
	const want = 13
	lines := linesFromFilename(t, exampleFilename(4))
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
	lines := linesFromFilename(t, filename(4))
	got, err := Day04Part1V1(lines)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay04Part1V1(b *testing.B) {
	lines := linesFromFilename(b, filename(4))
	for b.Loop() {
		_, _ = Day04Part1V1(lines)
	}
}

func TestDay04Part1Example(t *testing.T) {
	testSolver(t, 4, exampleFilename, true, Day04, 13)
}

func TestDay04Part2Example(t *testing.T) {
	testSolver(t, 4, exampleFilename, false, Day04, 30)
}

func TestDay04Part1(t *testing.T) {
	testSolver(t, 4, filename, true, Day04, 26218)
}

func TestDay04Part2(t *testing.T) {
	testSolver(t, 4, filename, false, Day04, 9997537)
}

func BenchmarkDay04Part1(b *testing.B) {
	buf := file(b, 4)
	for b.Loop() {
		_, _ = Day04(buf, true)
	}
}

func BenchmarkDay04Part2(b *testing.B) {
	buf := file(b, 4)
	for b.Loop() {
		_, _ = Day04(buf, false)
	}
}
