package adventofcode2023

import (
	"testing"
)

func TestDay23Part1Example(t *testing.T) {
	lines := []string{
		"#.#####################",
		"#.......#########...###",
		"#######.#########.#.###",
		"###.....#.>.>.###.#.###",
		"###v#####.#v#.###.#.###",
		"###.>...#.#.#.....#...#",
		"###v###.#.#.#########.#",
		"###...#.#.#.......#...#",
		"#####.#.#.#######.#.###",
		"#.....#.#.#.......#...#",
		"#.#####.#.#.#########v#",
		"#.#...#...#...###...>.#",
		"#.#.#v#######v###.###v#",
		"#...#.>.#...>.>.#.###.#",
		"#####v#.#.###v#.#.###.#",
		"#.....#...#...#.#.#...#",
		"#.#########.###.#.#.###",
		"#...###...#...#...#.###",
		"###.###.#.###v#####v###",
		"#...#...#.#.>.>.#.>.###",
		"#.###.###.#.###.#.#v###",
		"#.....###...###...#...#",
		"#####################.#",
	}
	puzzle, err := NewDay23(lines)
	if err != nil {
		t.Fatal(err)
	}
	got := Day23(puzzle, true)
	const want = 94
	if got != want {
		t.Errorf("Example: got %d, want %d", got, want)
	}
}

func TestDay23Part1(t *testing.T) {
	testWithParser(t, 23, filename, true, NewDay23, Day23, 2130)
}

func BenchmarkDay23Part1(b *testing.B) {
	benchWithParser(b, 23, true, NewDay23, Day23)
}

func TestDay23Part2(t *testing.T) {
	testWithParser(t, 23, filename, false, NewDay23, Day23, 0)
}

func BenchmarkDay23Part2(b *testing.B) {
	benchWithParser(b, 23, false, NewDay23, Day23)
}
