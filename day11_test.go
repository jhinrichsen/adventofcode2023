package adventofcode2023

import (
	"strconv"
	"testing"
)

func TestDay11Part1Example(t *testing.T) {
	const want = 374
	bytes, err := bytesFromFilename(exampleFilename(11))
	if err != nil {
		t.Fatal(err)
	}
	got := Day11(bytes, 1)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay11Part2Example(t *testing.T) {
	var tts = []struct {
		expansion uint
		want      uint
	}{
		{10, 1030},
		{100, 8410},
	}
	for i := range tts {
		t.Run(strconv.Itoa(i+1), func(t *testing.T) {
			want := tts[i].want
			bytes, err := bytesFromFilename(exampleFilename(11))
			if err != nil {
				t.Fatal(err)
			}
			got := Day11(bytes, tts[i].expansion)
			if want != got {
				t.Fatalf("want %d but got %d", want, got)
			}
		})
	}
}

func TestDay11Part1(t *testing.T) {
	const want = 9563821
	bytes, err := bytesFromFilename(filename(11))
	if err != nil {
		t.Fatal(err)
	}
	got := Day11(bytes, 1)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay11Part2(t *testing.T) {
	const want = 827_009_909_817
	bytes, err := bytesFromFilename(filename(11))
	if err != nil {
		t.Fatal(err)
	}
	got := Day11(bytes, 1_000_000)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay11Part1(b *testing.B) {
	bytes, err := bytesFromFilename(filename(11))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for b.Loop() {
		_ = Day11(bytes, 1)
	}
}

func BenchmarkDay11Part2(b *testing.B) {
	bytes, err := bytesFromFilename(filename(11))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for b.Loop() {
		_ = Day11(bytes, 1_000_000)
	}
}
