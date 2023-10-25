package adventofcode2023

import (
	"testing"
)

func TestDay00Part1Example(t *testing.T) {
	const want = 0
	lines, err := linesFromFilename(exampleFilename(0))
	if err != nil {
		t.Fatal(err)
	}
	got := Day00(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
