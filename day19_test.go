package adventofcode2023

import (
	"testing"
)

func TestDay19Part1Example(t *testing.T) {
	testWithParser(t, 19, exampleFilename, true, NewDay19, Day19, 19114)
}

func TestDay19Part1(t *testing.T) {
	testWithParser(t, 19, filename, true, NewDay19, Day19, 425811)
}

func BenchmarkDay19Part1(b *testing.B) {
	benchWithParser(b, 19, true, NewDay19, Day19)
}

func TestDay19Part2(t *testing.T) {
	testWithParser(t, 19, filename, false, NewDay19, Day19, 0)
}

func BenchmarkDay19Part2(b *testing.B) {
	benchWithParser(b, 19, false, NewDay19, Day19)
}
