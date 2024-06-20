package adventofcode2023

import (
	"strconv"
	"testing"
)

func TestDay05Part1Example(t *testing.T) {
	const want = 35
	lines, err := linesFromFilename(exampleFilename(5))
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day05(lines)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay05Part1(t *testing.T) {
	const want = 340994526
	lines, err := linesFromFilename(filename(5))
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day05(lines)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay05Part1(b *testing.B) {
	lines, err := linesFromFilename(filename(5))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Day05(lines)
	}
}

func TestDay05Part2(t *testing.T) {
	const want = 340994526
	lines, err := linesFromFilename(filename(5))
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day05(lines)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestExampleMap(t *testing.T) {
	samples := []struct {
		from, into uint
	}{
		{48, 48},
		{49, 49},
		{50, 52},
		{51, 53},
		{96, 98},
		{97, 99},
		{98, 50},
		{99, 51},
	}

	rs := Ranges{
		NewRange(50, 98, 2),
		NewRange(52, 50, 48),
	}

	// test range
	want := samples[2].into
	got := uint(int(samples[2].from) + rs[1].Delta)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}

	// test ranges
	for i, sample := range samples {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			in := sample.from
			want := sample.into
			got := rs.Do(in)
			if want != got {
				t.Fatalf("%d: want %d but got %d\n", in, want, got)
			}
		})
	}
}
