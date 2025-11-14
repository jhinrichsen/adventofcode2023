package adventofcode2023

import (
	"testing"
)

func TestDay14Part1Example(t *testing.T) {
	testWithParser(t, 14, exampleFilename, true, NewDay14, Day14, 136)
}

func TestDay14Part1(t *testing.T) {
	testWithParser(t, 14, filename, true, NewDay14, Day14, 108792)
}

func BenchmarkDay14Part1(b *testing.B) {
	benchWithParser(b, 14, true, NewDay14, Day14)
}

func BenchmarkDay14Part2(b *testing.B) {
	benchWithParser(b, 14, false, NewDay14, Day14)
}

func TestDay14Part2Example(t *testing.T) {
	testWithParser(t, 14, exampleFilename, false, NewDay14, Day14, 64)
}

func TestDay14Part2(t *testing.T) {
	testWithParser(t, 14, filename, false, NewDay14, Day14, 99118)
}
