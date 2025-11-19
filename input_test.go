package adventofcode2023

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"syscall"
	"testing"
	"unsafe"
)

func linesFromFilename(tb testing.TB, filename string) []string {
	tb.Helper()
	f, err := os.Open(filename)
	if err != nil {
		tb.Fatal(err)
	}
	lines := linesFromReader(tb, f)
	if b, ok := tb.(*testing.B); ok {
		b.ResetTimer()
	}
	return lines
}

func linesFromReader(tb testing.TB, r io.Reader) []string {
	tb.Helper()
	var lines []string
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		line := sc.Text()
		lines = append(lines, line)
	}
	if err := sc.Err(); err != nil {
		tb.Fatal(err)
	}
	return lines
}

func exampleFilename(day uint8) string {
	return fmt.Sprintf("testdata/day%02d_example.txt", int(day))
}

func filename(day uint8) string {
	return fmt.Sprintf("testdata/day%02d.txt", int(day))
}

// file reads the main input file bytes for day N (zero-padded).
func file(tb testing.TB, day uint8) []byte {
	tb.Helper()
	buf, err := os.ReadFile(filename(day))
	if err != nil {
		tb.Fatal(err)
	}
	if b, ok := tb.(*testing.B); ok {
		b.ResetTimer()
	}
	return buf
}

// fileFromFilename reads file bytes using a filename function (e.g., filename or exampleFilename).
func fileFromFilename(tb testing.TB, filenameFunc func(uint8) string, day uint8) []byte {
	tb.Helper()
	buf, err := os.ReadFile(filenameFunc(day))
	if err != nil {
		tb.Fatal(err)
	}
	if b, ok := tb.(*testing.B); ok {
		b.ResetTimer()
	}
	return buf
}

const (
	MagicMaxLines    = 140 // maximum number of lines for any puzzle input
	MagicLongestLine = 307 // longest line of any puzzle input
)

// bytesFromFilename reads newline separated lines from a file and returns them as [][]byte.
func bytesFromFilename(tb testing.TB, filename string) [][]byte {
	tb.Helper()
	f, err := os.Open(filename)
	if err != nil {
		tb.Fatal(err)
	}
	defer f.Close()

	buf, err := io.ReadAll(f)
	if err != nil {
		tb.Fatal(err)
	}

	var result [][]byte
	start := 0
	l := len(buf)

	for i := 0; i < l; i++ {
		if buf[i] == '\n' {
			result = append(result, append([]byte(nil), buf[start:i]...))
			start = i + 1
		}
	}

	// Check if there's any remaining text after the last newline
	if start < l {
		// Append the last line if it didn't end with a newline
		result = append(result, append([]byte(nil), buf[start:]...))
	}

	if b, ok := tb.(*testing.B); ok {
		b.ResetTimer()
	}
	return result
}

func DayAdapterV1(day func([][]byte, bool) (uint, error), filename string, part1 bool) (uint, error) {
	buf, err := os.ReadFile(filename)
	if err != nil {
		return 0, err
	}

	var result [][]byte
	start := 0
	l := len(buf)

	for i := 0; i < l; i++ {
		if buf[i] == '\n' {
			result = append(result, append([]byte(nil), buf[start:i]...))
			start = i + 1
		}
	}

	if start < l {
		result = append(result, append([]byte(nil), buf[start:]...))
	}

	return day(result, part1)
}

func DayAdapterV2(day func([][]byte, bool) (uint, error), filename string, part1 bool) (uint, error) {
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer func() { _ = f.Close() }()

	// Get the file size
	stat, err := f.Stat()
	if err != nil {
		return 0, err
	}

	size := int(stat.Size())
	if size == 0 {
		return 0, err
	}

	// Memory map the file
	data, err := syscall.Mmap(int(f.Fd()), 0, size, syscall.PROT_READ, syscall.MAP_PRIVATE)
	if err != nil {
		return 0, err
	}

	// Defer unmapping the memory
	defer func() { _ = syscall.Munmap(data) }()

	// Pre-allocate a fixed array for lines
	var lines [MagicMaxLines][]byte
	lineIndex := 0

	start := 0
	for i := 0; i < size; i++ {
		if data[i] == '\n' {
			lines[lineIndex] = unsafe.Slice(&data[start], i-start)
			lineIndex++
			start = i + 1
		}
	}

	// Handle the last line if it doesn't end with a newline
	if start < size {
		lines[lineIndex] = unsafe.Slice(&data[start], size-start)
		lineIndex++
	}

	return day(lines[:lineIndex], part1)
}
