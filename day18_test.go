package adventofcode2023

import (
	"testing"
)

func TestDay18Part1Example(t *testing.T) {
	testWithParser(t, 18, exampleFilename, true, NewDay18, Day18, 62)
}

func TestDay18Part1(t *testing.T) {
	testWithParser(t, 18, filename, true, NewDay18, Day18, 45159)
}

func BenchmarkDay18Part1(b *testing.B) {
	benchWithParser(b, 18, true, NewDay18, Day18)
}

func TestDay18Part2(t *testing.T) {
	testWithParser(t, 18, filename, false, NewDay18, Day18, 0)
}

func BenchmarkDay18Part2(b *testing.B) {
	benchWithParser(b, 18, false, NewDay18, Day18)
}
