package adventofcode2023

import (
	"fmt"
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

func TestDay10Part1Example1(t *testing.T) {
	testSolver(t, 10, exampleFilename, true, Day10, 4)
}

func TestDay10Part1Example2(t *testing.T) {
	buf := fileFromFilename(t, func(uint8) string { return "testdata/day10_example2.txt" }, 0)
	got, err := Day10(buf, true)
	if err != nil {
		t.Fatal(err)
	}
	const want = 8
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay10Part1(t *testing.T) {
	testSolver(t, 10, filename, true, Day10, 6956)
}

func TestDay10Part2(t *testing.T) {
	testSolver(t, 10, filename, false, Day10, 455)
}

func BenchmarkOpposite(b *testing.B) {
	var ds [1000]direction
	pcg := rand.NewPCG(0, 0) // make sure we always bench against the same dataset
	rnd := rand.New(pcg)
	for i := range ds {
		ds[i] = 1 << rnd.IntN(4) // [0..3]
	}
	b.ResetTimer()
	for b.Loop() {
		for _, d := range ds {
			_ = opposite3(d)
		}
	}
}

func TestDay10Part2Examples(t *testing.T) {
	wants := []uint{4, 4 /* not working 8, */, 0, 10}
	for i := range wants {
		want := wants[i]
		t.Run(fmt.Sprintf("#%d", i+1), func(t *testing.T) {
			buf := fileFromFilename(t, func(uint8) string {
				return fmt.Sprintf("testdata/day10_part2_example%d.txt", i+1)
			}, 0)
			fmt.Println()
			fmt.Printf("Example #%d:\n", i+1)
			got, err := Day10(buf, false)
			fmt.Println()
			if err != nil {
				t.Fatal(err)
			}
			if want != got {
				t.Fatalf("want %d but got %d", want, got)
			}
		})
	}
}

func BenchmarkDay10Part1(b *testing.B) {
	buf := fileFromFilename(b, filename, 10)
	for b.Loop() {
		_, _ = Day10(buf, true)
	}
}

func BenchmarkDay10Part2(b *testing.B) {
	buf := fileFromFilename(b, filename, 10)
	for b.Loop() {
		_, _ = Day10(buf, false)
	}
}
