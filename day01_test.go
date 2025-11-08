package adventofcode2023

import (
	"os"
	"testing"
)

// NewDay01 parser converts lines to byte buffer for Day01
func NewDay01(lines []string) ([]byte, error) {
	var size int
	for i := range lines {
		size += len(lines[i]) + 1 // +1 for newline
	}
	buf := make([]byte, 0, size)
	for i := range lines {
		buf = append(buf, []byte(lines[i])...)
		buf = append(buf, '\n')
	}
	return buf, nil
}

func TestDay01Part1Example(t *testing.T) {
	testWithParser(t, 1, exampleFilename, true, NewDay01, Day01, 142)
}

func TestDay01Part1(t *testing.T) {
	testWithParser(t, 1, filename, true, NewDay01, Day01, 55130)
}

func BenchmarkDay01V1(b *testing.B) {
	lines := linesFromFilename(b, filename(1))
	b.ResetTimer()
	for b.Loop() {
		_, _ = Day01V1(lines, true)
	}
}

func BenchmarkDay01V2(b *testing.B) {
	buf := file(b, 1)
	for b.Loop() {
		_ = Day01(buf, true)
	}
}

func BenchmarkDay01Part1(b *testing.B) {
	buf := file(b, 1)
	for b.Loop() {
		_ = Day01(buf, true)
	}
}

func BenchmarkDay01Part2(b *testing.B) {
	buf := file(b, 1)
	for b.Loop() {
		_ = Day01(buf, false)
	}
}

// BenchmarkDay01Large processes a 1 MB input file, the original input
// repeated until 1.000.000 bytes are filled, plus the rest of the line.
func BenchmarkDay01Large(b *testing.B) {
	buf, _ := os.ReadFile("testdata/day01_large.txt")
	for b.Loop() {
		_ = Day01(buf, true)
	}
}
