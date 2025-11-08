package adventofcode2023

import (
	"fmt"
	"testing"
)

func BrokenTestDay12Part1ExamplesChatGPT(t *testing.T) {
	tests := []struct {
		input []string
		want  uint
	}{
		{
			input: []string{".#...?....###. 1,1,3"},
			want:  2,
		},
		{
			input: []string{".#....#.### 1,1,3"},
			want:  1,
		},
		{
			input: []string{".#.??..###. 1,2,3"},
			want:  1,
		},
		{
			input: []string{"#.#.#.#.## 2,1,1"},
			want:  0,
		},
		{
			input: []string{".###.#?## 1,2,1"},
			want:  2,
		},
		{
			input: []string{".#??.# 1,2"},
			want:  4,
		},
		{
			input: []string{".#.###?### 2,3,1"},
			want:  0,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.input), func(t *testing.T) {
			got := Day12(tt.input, true)
			if got != tt.want {
				t.Errorf("For input %v, expected %d, but got %d", tt.input, tt.want, got)
			}
		})
	}
}

func TestDay12IsValidCombination(t *testing.T) {
	const want = 1
	valid := []string{
		"#.#.### 1,1,3",
		".#...#....###. 1,1,3",
		".#.###.#.###### 1,3,1,6",
		"####.#...#... 4,1,1",
		"#....######..#####. 1,6,5",
		".###.##....# 3,2,1",
	}
	for i := range valid {
		t.Run(valid[i], func(t *testing.T) {
			got := Day12([]string{valid[i]}, true)
			if want != got {
				t.Fatalf("%s: want %d but got %d", valid[i], want, got)
			}
		})
	}
}

func TestDay12Part1Example(t *testing.T) {
	testWithParser(t, 12, exampleFilename, true, NewDay12, Day12, uint(21))
}

func TestDay12Part1(t *testing.T) {
	testWithParser(t, 12, filename, true, NewDay12, Day12, uint(7939))
}

func TestDay12Part2Examples(t *testing.T) {
	tests := []struct {
		input string
		want  uint
	}{
		{"???.### 1,1,3", 1},
		{".??..??...?##. 1,1,3", 16384},
		{"?#?#?#?#?#?#?#? 1,3,1,6", 1},
		{"????.#...#... 4,1,1", 16},
		{"????.######..#####. 1,6,5", 2500},
		{"?###???????? 3,2,1", 506250},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := Day12([]string{tt.input}, false)
			if got != tt.want {
				t.Errorf("want %d but got %d", tt.want, got)
			}
		})
	}
}

func TestDay12Part2Example(t *testing.T) {
	// Test against example2 file which contains all Part 2 examples
	// Sum: 1 + 16384 + 1 + 16 + 2500 + 506250 = 525152
	testWithParser(t, 12, func(_ uint8) string {
		return "testdata/day12_example2.txt"
	}, false, NewDay12, Day12, uint(525152))
}

// TestDay12Part2 is commented out pending verification of the answer with AoC website.
// Once verified, uncomment this test and update the expected value.
// func TestDay12Part2(t *testing.T) {
// 	testWithParser(t, 12, func(day uint8) string { return filename(int(day)) }, false, NewDay12, Day12, uint(0))
// }

func BenchmarkDay12Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(12))
	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		data, _ := NewDay12(lines)
		_ = Day12(data, true)
	}
}

func BenchmarkDay12Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(12))
	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		data, _ := NewDay12(lines)
		_ = Day12(data, false)
	}
}
