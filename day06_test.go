package adventofcode2023

import (
	"testing"
)

func TestDay06Part1Example(t *testing.T) {
	testWithParser(t, 6, exampleFilename, true, NewDay06, Day06, 288)
}

func TestDay06Part1(t *testing.T) {
	testWithParser(t, 6, filename, true, NewDay06, Day06, 440000)
}

func TestDay06Part2Example(t *testing.T) {
	testWithParser(t, 6, exampleFilename, false, NewDay06, Day06, 71503)
}

func TestDay06Part2(t *testing.T) {
	testWithParser(t, 6, filename, false, NewDay06, Day06, 26187338)
}

func BenchmarkDay06Part1(b *testing.B) {
	benchWithParser(b, 6, true, NewDay06, Day06)
}

func BenchmarkDay06Part2(b *testing.B) {
	benchWithParser(b, 6, false, NewDay06, Day06)
}
