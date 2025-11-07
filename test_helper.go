package adventofcode2023

import "testing"

// testDayPart is a generic test helper for standard day part tests.
func testDayPart[P any, R comparable](
	t *testing.T,
	day uint8,
	filenameFunc func(uint8) string,
	part1 bool,
	parser func([]string) (P, error),
	solver func(P, bool) R,
	want R,
) {
	t.Helper()
	lines, err := linesFromFilename(filenameFunc(day))
	if err != nil {
		t.Fatal(err)
	}
	puzzle, err := parser(lines)
	if err != nil {
		t.Fatal(err)
	}
	got := solver(puzzle, part1)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}
