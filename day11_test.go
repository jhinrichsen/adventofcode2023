package adventofcode2023

import (
	"strconv"
	"testing"
)

func TestDay11Part2Examples(t *testing.T) {
	t.Skip("testdata/day11_example.txt does not exist")
	tests := []struct {
		expansion uint
		want      uint
	}{
		{10, 1030},
		{100, 8410},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i+1), func(t *testing.T) {
			lines := linesFromFilename(t, exampleFilename(11))
			got := day11Solver(lines, tt.expansion)
			if tt.want != got {
				t.Fatalf("want %d but got %d", tt.want, got)
			}
		})
	}
}

func TestDay11Part1(t *testing.T) {
	testLines(t, 11, filename, true, Day11, 9563821)
}

func TestDay11Part2(t *testing.T) {
	testLines(t, 11, filename, false, Day11, 827_009_909_817)
}

func BenchmarkDay11Part1(b *testing.B) {
	benchLines(b, 11, true, Day11)
}

func BenchmarkDay11Part2(b *testing.B) {
	benchLines(b, 11, false, Day11)
}
