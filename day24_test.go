package adventofcode2023

import (
	"testing"
)

func TestDay24Part1Example(t *testing.T) {
	lines := []string{
		"19, 13, 30 @ -2,  1, -2",
		"18, 19, 22 @ -1, -1, -2",
		"20, 25, 34 @ -2, -2, -4",
		"12, 31, 28 @ -1, -2, -1",
		"20, 19, 15 @  1, -5, -3",
	}
	puzzle, err := NewDay24(lines)
	if err != nil {
		t.Fatal(err)
	}

	// For example, use test area [7, 27]
	minCoord := int64(7)
	maxCoord := int64(27)

	count := 0
	for i := 0; i < len(puzzle); i++ {
		for j := i + 1; j < len(puzzle); j++ {
			h1, h2 := puzzle[i], puzzle[j]

			det := h1.vx*h2.vy - h1.vy*h2.vx
			if det == 0 {
				continue
			}

			dx := h2.px - h1.px
			dy := h2.py - h1.py

			t1Num := dx*h2.vy - dy*h2.vx
			t2Num := dx*h1.vy - dy*h1.vx

			if (t1Num < 0) != (det < 0) || (t2Num < 0) != (det < 0) {
				continue
			}

			xScaled := h1.px*det + t1Num*h1.vx
			yScaled := h1.py*det + t1Num*h1.vy

			var inBounds bool
			if det > 0 {
				inBounds = minCoord*det <= xScaled && xScaled <= maxCoord*det &&
					minCoord*det <= yScaled && yScaled <= maxCoord*det
			} else {
				inBounds = maxCoord*det <= xScaled && xScaled <= minCoord*det &&
					maxCoord*det <= yScaled && yScaled <= minCoord*det
			}

			if inBounds {
				count++
			}
		}
	}

	const want = 2
	if count != want {
		t.Errorf("Example: got %d, want %d", count, want)
	}
}

func TestDay24Part1(t *testing.T) {
	testWithParser(t, 24, filename, true, NewDay24, Day24, 11098)
}

func BenchmarkDay24Part1(b *testing.B) {
	benchWithParser(b, 24, true, NewDay24, Day24)
}

func TestDay24Part2(t *testing.T) {
	testWithParser(t, 24, filename, false, NewDay24, Day24, 0)
}

func BenchmarkDay24Part2(b *testing.B) {
	benchWithParser(b, 24, false, NewDay24, Day24)
}
