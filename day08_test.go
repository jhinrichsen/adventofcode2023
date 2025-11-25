package adventofcode2023

import (
	"testing"
)

func TestDay08Part1Example(t *testing.T) {
	testWithParser(t, 8, exampleFilename, true, NewDay08, Day08, 6)
}

func TestDay08Part1(t *testing.T) {
	testWithParser(t, 8, filename, true, NewDay08, Day08, 18727)
}

func TestDay08Part2Example(t *testing.T) {
	testWithParser(t, 8, func(uint8) string { return "testdata/day08_example_part2.txt" }, false, NewDay08, Day08, 6)
}

func TestDay08Part2(t *testing.T) {
	testWithParser(t, 8, filename, false, NewDay08, Day08, 18024643846273)
}

func BenchmarkDay08Part1(b *testing.B) {
	benchWithParser(b, 8, true, NewDay08, Day08)
}

func BenchmarkDay08Part2(b *testing.B) {
	benchWithParser(b, 8, false, NewDay08, Day08)
}
