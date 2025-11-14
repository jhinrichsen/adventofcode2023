package adventofcode2023

import "testing"

// testWithParser is a generic test helper for day part tests using a parser and solver.
func testWithParser[P any, R comparable](
	t *testing.T,
	day uint8,
	filenameFunc func(uint8) string,
	part1 bool,
	parser func([]string) (P, error),
	solver func(P, bool) R,
	want R,
) {
	t.Helper()
	lines := linesFromFilename(t, filenameFunc(day))
	puzzle, err := parser(lines)
	if err != nil {
		t.Fatal(err)
	}
	got := solver(puzzle, part1)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

// testSolver is a generic test helper for day part tests that work directly with []byte.
func testSolver[R comparable](
	t *testing.T,
	day uint8,
	filenameFunc func(uint8) string,
	part1 bool,
	solver func([]byte, bool) (R, error),
	want R,
) {
	t.Helper()
	buf := fileFromFilename(t, filenameFunc, day)
	got, err := solver(buf, part1)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

// testLines is a generic test helper for day part tests that work directly with []string lines.
func testLines[R comparable](
	t *testing.T,
	day uint8,
	filenameFunc func(uint8) string,
	part1 bool,
	solver func([]string, bool) R,
	want R,
) {
	t.Helper()
	lines := linesFromFilename(t, filenameFunc(day))
	got := solver(lines, part1)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}
