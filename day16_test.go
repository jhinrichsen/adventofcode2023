package adventofcode2023

import (
	"testing"
)

func TestDay16Part1Example(t *testing.T) {
	testWithParser(t, 16, exampleFilename, true, NewDay16, Day16, 46)
}

func TestDay16Part1(t *testing.T) {
	testWithParser(t, 16, filename, true, NewDay16, Day16, 7111)
}

func BenchmarkDay16Part1(b *testing.B) {
	benchWithParser(b, 16, true, NewDay16, Day16)
}

func TestDay16Part2(t *testing.T) {
	testWithParser(t, 16, filename, false, NewDay16, Day16, 0)
}
