package adventofcode2023

import (
	"reflect"
	"slices"
	"strconv"
	"testing"
)

func TestDay05Part1Example(t *testing.T) {
	const want = 35
	lines := linesFromFilename(t, exampleFilename(5))
	got, err := Day05(lines, true)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

// The lowest location number can be obtained from seed number 82, which corresponds to soil 84, fertilizer 84,
// water 84, light 77, temperature 45, humidity 46, and location 46. So, the lowest location number is 46.
func TestDay05Part2Example(t *testing.T) {
	const want = 46
	lines := linesFromFilename(t, exampleFilename(5))
	got, err := Day05(lines, false)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay05Part1(t *testing.T) {
	const want = 340994526
	lines := linesFromFilename(t, filename(5))
	got, err := Day05(lines, true)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay05Part2(t *testing.T) {
	const want = 52210644
	lines := linesFromFilename(t, filename(5))
	got, err := Day05(lines, false)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay05Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(5))
	for b.Loop() {
		_, _ = Day05(lines, true)
	}
}

func BenchmarkDay05Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(5))
	for b.Loop() {
		_, _ = Day05(lines, false)
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

// Case 1
func TestDay05MergeEqual(t *testing.T) {
	r1 := Range{10, 19, 1}
	r2 := Range{10, 19, 2}
	want := Ranges{Range{10, 19, 3}}
	got := Merge(r1, r2)
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v but got %v", want, got)
	}
}

// Case 2
func TestDay05MergeNonOverlapping(t *testing.T) {
	r1 := Range{20, 29, 2}
	r2 := Range{10, 19, 1}
	want := Ranges{r2, r1}
	got := Merge(r1, r2)
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v but got %v", want, got)
	}
}

// Case 3
func TestDay05MergeInside(t *testing.T) {
	r1 := Range{10, 19, 1}
	r2 := Range{14, 15, 2}
	wantR := Ranges{Range{10, 13, 1}, Range{14, 15, 3}, Range{16, 19, 1}}
	gotR := Merge(r1, r2)
	if !reflect.DeepEqual(wantR, gotR) {
		t.Fatalf("want %v but got %v", wantR, gotR)
	}
	tts := []struct {
		in, out uint
	}{
		{13, 14},
		{14, 17},
		{15, 18},
		{16, 17},
	}
	for i, tt := range tts {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			want := tt.out
			got := gotR.Do(tt.in)
			if want != got {
				t.Fatalf("want %d but got %d", want, got)
			}
		})
	}
}

// Case 4
func TestDay05MergeShorter(t *testing.T) {
	r1 := Range{10, 19, 1}
	r2 := Range{10, 29, 2}
	wantR := Ranges{Range{10, 19, 3}, Range{20, 29, 2}}
	gotR := Merge(r1, r2)
	if !reflect.DeepEqual(wantR, gotR) {
		t.Fatalf("want %v but got %v", wantR, gotR)
	}
	tts := []struct {
		in, out uint
	}{
		{10, 13},
		{18, 21},
		{19, 22},
		{20, 22},
		{29, 31},
	}
	for i, tt := range tts {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			want := tt.out
			got := gotR.Do(tt.in)
			if want != got {
				t.Fatalf("want %d but got %d", want, got)
			}
		})
	}
}

// Case 5
func TestDay05MergeOverlapping(t *testing.T) {
	r1 := Range{10, 20, 1}
	r2 := Range{15, 25, 2}
	wantR := Ranges{Range{10, 14, 1}, Range{15, 20, 3}, Range{21, 25, 2}}
	gotR := Merge(r1, r2)
	if !reflect.DeepEqual(wantR, gotR) {
		t.Fatalf("want %v but got %v", wantR, gotR)
	}
	tts := []struct {
		in, out uint
	}{
		{10, 11},
		{14, 15},
		{15, 18},
		{16, 19},
		{19, 22},
		{20, 23},
		{21, 23},
		{25, 27},
	}
	for i, tt := range tts {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			want := tt.out
			got := gotR.Do(tt.in)
			if want != got {
				t.Fatalf("want %d but got %d", want, got)
			}
		})
	}
}

func TestDay05MergeRanges(t *testing.T) {
	lines := linesFromFilename(t, exampleFilename(5))
	rrs, err := parseDay05(lines[2:])
	if err != nil {
		t.Fatal(err)
	}
	// fmt.Printf("found %d maps\n", len(rrs))
	var total uint
	for _, rs := range rrs {
		// fmt.Printf("\t%+v\n", rs)
		total += uint(len(rs))
	}

	for _, rs := range rrs[1:] {
		condensed := rrs[0]
		for _, r1 := range condensed {
			for _, r2 := range rs {
				condensed = append(condensed, Merge(r1, r2)...)
			}
		}
		rrs[0] = condensed
	}

	// Sort ascending
	slices.SortFunc(rrs[0], func(a, b Range) int {
		// Sort on Min
		cmp := int(a.Min) - int(b.Min)
		if cmp != 0 {
			return cmp
		}
		// for the same Min, additionally sort on Max
		return int(a.Max) - int(b.Max)
	})
	/*
		for _, rs := range rrs[0] {
			fmt.Printf("%+v\n", rs)
		}
	*/

	// compress linear list of Ranges
	var cs Ranges
	base := rrs[0]
	for i := 1; i < len(base); i++ {
		cs = append(cs, Merge(base[0], base[i])...)
	}
	/*
		for _, rs := range cs {
			fmt.Printf("%+v\n", rs)
		}
	*/
	if len(cs) != 511511 {
		t.Fatalf("TODO broke %d", len(cs))
	}
}
