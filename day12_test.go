package adventofcode2023

import (
	"fmt"
	"testing"
)

func BrokenTestDay12Part1ExamplesChatGPT(t *testing.T) {
	tests := []struct {
		input []string
		want  int
	}{
		// Test cases with different inputs
		{
			input: []string{".#...?....###. 1,1,3"},
			want:  2, // Example with valid combinations
		},
		{
			input: []string{".#....#.### 1,1,3"},
			want:  1, // One valid combination
		},
		{
			input: []string{".#.??..###. 1,2,3"},
			want:  1, // Multiple '?' characters, one valid
		},
		{
			input: []string{"#.#.#.#.## 2,1,1"},
			want:  0, // No valid combinations
		},
		{
			input: []string{".###.#?## 1,2,1"},
			want:  2, // Valid combinations with '?'
		},
		{
			input: []string{".#??.# 1,2"},
			want:  4, // Valid combinations with multiple '?'
		},
		{
			input: []string{".#.###?### 2,3,1"},
			want:  0, // Invalid checksum after replacement
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.input), func(t *testing.T) {
			got := Day12(tt.input)
			if got != tt.want {
				t.Errorf("For input %v, expected %d, but got %d", tt.input, tt.want, got)
			}
		})
	}
}

func TestDay12Part1Example(t *testing.T) {
	const want = 21
	lines, err := linesFromFilename(exampleFilename(12))
	if err != nil {
		t.Fatal(err)
	}
	got := Day12(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay12Part1(t *testing.T) {
	const want = 7939
	lines, err := linesFromFilename(filename(12))
	if err != nil {
		t.Fatal(err)
	}
	got := Day12(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay12Part1(b *testing.B) {
	lines, err := linesFromFilename(filename(12))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for range b.N {
		_ = Day12(lines)
	}
}

func TestDay12Part2(t *testing.T) {
	const want = 0
	lines, err := linesFromFilename(filename(12))
	if err != nil {
		t.Fatal(err)
	}
	got := Day12(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
