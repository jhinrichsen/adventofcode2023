package adventofcode2023

import (
	"bufio"
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestLinesFromFilename(t *testing.T) {
	lines, err := linesFromFilename("testdata/helloworld.txt")
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 1 {
		t.Fatalf("want 1 line but got %d", len(lines))
	}
}

func TestLinesLinesAsNumbers(t *testing.T) {
	sample := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	ints, err := linesAsNumbers(sample)
	if err != nil {
		t.Fatal(err)
	}
	if len(ints) != len(sample) {
		t.Fatalf("want %d numbers but got %d", len(sample), len(ints))
	}
	for i := range sample {
		want := i + 1 // entries are 1-based
		got := ints[i]
		if want != got {
			t.Fatalf("line %d: want %d but got %d", i, want, got)
		}
	}
}

func BenchmarkBytesFromFilename(b *testing.B) {
	filenames, err := filepath.Glob("testdata/*.txt")
	if err != nil {
		b.Fatal(err)
	}
	readall := func() {
		for i := range filenames {
			_, _ = bytesFromFilename(filenames[i])
		}
	}
	// warm-up cache
	readall()
	b.ResetTimer()
	for range b.N {
		readall()
	}
}

func BenchmarkLinesFromFilename(b *testing.B) {
	filenames, err := filepath.Glob("testdata/*.txt")
	if err != nil {
		b.Fatal(err)
	}
	readall := func() {
		for i := range filenames {
			_, _ = linesFromFilename(filenames[i])
		}
	}
	// warm-up cache
	readall()
	b.ResetTimer()
	for range b.N {
		readall()
	}
}

const MAGIC_LONGEST_LINE = 307

func TestLongestLine(t *testing.T) {
	filenames, err := filepath.Glob("testdata/*.txt")
	if err != nil {
		t.Fatal(err)
	}

	var got uint
	for i := range filenames {
		buf, err := os.ReadFile(filenames[i])
		if err != nil {
			t.Fatal(err)
		}

		scanner := bufio.NewScanner(bytes.NewReader(buf))

		for scanner.Scan() {
			line := scanner.Text()
			lineLength := uint(len(line))
			if lineLength > got {
				got = lineLength
			}
		}

		if err := scanner.Err(); err != nil {
			t.Fatal(err)
		}
	}
	if MAGIC_LONGEST_LINE != got {
		t.Fatalf("want %d but got %d", MAGIC_LONGEST_LINE, got)
	}
}
