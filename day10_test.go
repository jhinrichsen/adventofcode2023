package adventofcode2023

import (
	"math/rand/v2"
	"testing"
)

func TestConnects(t *testing.T) {
	var wants = []struct {
		char    byte
		want1   direction
		want2   direction
		nowant1 direction
		nowant2 direction
	}{
		{'L', North, East, South, West},
		{'|', North, South, East, West},
		{'J', North, West, South, East},
		{'7', South, West, North, East},
		{'F', South, East, North, West},
		{'-', West, East, North, South},
	}
	for i := range wants {
		want := wants[i].want1 | wants[i].want2
		got := connects[wants[i].char]
		if want != got {
			t.Fatalf("want %c (%d) = %d but got %d", wants[i].char, wants[i].char, want, got)
		}
	}
	for i := range wants {
		want := true
		got := hasConnection(wants[i].char, wants[i].want1)
		if want != got {
			t.Fatalf("want %c (%d) = %t but got %t", wants[i].char, wants[i].char, want, got)
		}
		got = hasConnection(wants[i].char, wants[i].want2)
		if want != got {
			t.Fatalf("want %c (%d) = %t but got %t", wants[i].char, wants[i].char, want, got)
		}
		want = false
		got = hasConnection(wants[i].char, wants[i].nowant1)
		if want != got {
			t.Fatalf("want %c (%d) = %t but got %t", wants[i].char, wants[i].char, want, got)
		}
		got = hasConnection(wants[i].char, wants[i].nowant2)
		if want != got {
			t.Fatalf("want %c (%d) = %t but got %t", wants[i].char, wants[i].char, want, got)
		}
	}
}

func TestOpposite(t *testing.T) {
	var tt = []struct {
		in   direction
		want direction
	}{
		{North, South},
		{South, North},
		{West, East},
		{East, West},
	}
	for i := range tt {
		want := tt[i].want
		got := opposite2(tt[i].in)
		if want != got {
			t.Fatalf("want %b but got %b", want, got)
		}
	}
}

func TestOtherHor(t *testing.T) {
	const want = East
	got := other('-', West)
	if want != got {
		t.Fatalf("want %b but got %b", want, got)
	}
}

func TestOtherVer(t *testing.T) {
	const want = North
	got := other('|', South)
	if want != got {
		t.Fatalf("want %b but got %b", want, got)
	}
}

func TestDay10Part1Example(t *testing.T) {
	const want = 4
	lines, err := linesFromFilename(exampleFilename(10))
	if err != nil {
		t.Fatal(err)
	}
	got := Day10(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay10Part1ExampleChatGPT(t *testing.T) {
	const want = 4
	got, err := Day10ChatGPT()
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay10Part1Example2(t *testing.T) {
	const want = 8
	lines, err := linesFromFilename("testdata/day10_example2.txt")
	if err != nil {
		t.Fatal(err)
	}
	got := Day10(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay10Part1(t *testing.T) {
	const want = 6956
	lines, err := linesFromFilename(filename(10))
	if err != nil {
		t.Fatal(err)
	}
	got := Day10(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkOpposite(b *testing.B) {
	var ds [1000]direction
	pcg := rand.NewPCG(0, 0) // make sure we always bench against the same dataset
	rnd := rand.New(pcg)
	for i := range ds {
		ds[i] = 1 << rnd.IntN(4) // [0..3]
	}
	b.ResetTimer()
	for range b.N {
		for _, d := range ds {
			_ = opposite3(d)
		}
	}
}

func BenchmarkDay10Part1(b *testing.B) {
	lines, err := linesFromFilename(filename(10))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for range b.N {
		_ = Day10(lines)
	}
}

func TestDay10Part2(t *testing.T) {
	const want = 0
	lines, err := linesFromFilename(filename(10))
	if err != nil {
		t.Fatal(err)
	}
	got := Day10(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
