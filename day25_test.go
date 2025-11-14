package adventofcode2023

import (
	"strings"
	"testing"
)

func TestDay25Part1Example(t *testing.T) {
	input := `jqt: rhn xhk nvd
rsh: frs pzl lsr
xhk: hfx
cmg: qnr nvd lhk bvb
rhn: xhk bvb hfx
bvb: xhk hfx
pzl: lsr hfx nvd
qnr: nvd
ntq: jqt hfx bvb xhk
nvd: lhk
lsr: lhk
rzs: qnr cmg lsr rsh
frs: qnr lhk lsr`

	lines := strings.Split(input, "\n")
	puzzle, err := NewDay25(lines)
	if err != nil {
		t.Fatal(err)
	}
	got := Day25(puzzle, true)
	const want = 54
	if got != want {
		t.Errorf("Example: got %d, want %d", got, want)
	}
}

func TestDay25Part1(t *testing.T) {
	testWithParser(t, 25, filename, true, NewDay25, Day25, 580800)
}

func BenchmarkDay25Part1(b *testing.B) {
	benchWithParser(b, 25, true, NewDay25, Day25)
}
