package adventofcode2023

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func linesFromFilename(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return []string{}, err
	}
	return linesFromReader(f)
}

func linesFromReader(r io.Reader) ([]string, error) {
	var lines []string
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		line := sc.Text()
		lines = append(lines, line)
	}
	return lines, nil
}

func exampleFilename(day int) string {
	return fmt.Sprintf("testdata/day%02d_example.txt", day)
}

func filename(day int) string {
	return fmt.Sprintf("testdata/day%02d.txt", day)
}

func lineAsNumbers(line string) ([]int, error) {
	var (
		n   int
		ns  []int
		err error
	)
	for _, s := range strings.Fields(line) {
		n, err = strconv.Atoi(s)
		if err != nil {
			break
		}
		ns = append(ns, n)
	}
	return ns, err
}

// linesAsNumber converts strings into integer.
func linesAsNumbers(lines []string) ([]int, error) {
	var is []int
	for i := range lines {
		n, err := strconv.Atoi(lines[i])
		if err != nil {
			msg := "error in line %d: cannot convert %q to number"
			return is, fmt.Errorf(msg, i, lines[i])
		}
		is = append(is, n)
	}
	return is, nil
}

/*
func numbersFromFilename(filename string) ([]int, error) {
	ls, err := linesFromFilename(filename)
	if err != nil {
		return nil, err
	}
	return linesAsNumbers(ls)
}
*/

// bytesFromFilename reads newline separated lines from a file and returns them as [][]byte.
// casting string([]byte) has a runtime overhead because of internal memory allocation.
// len() is the number of lines, len([0]) is the length of the first line.
// Both indices start at top left and go bottom right.
//
// A..
// ...
// .Z.
// [0][0] == A
// [2][1] == Z

func bytesFromFilename(filename string) ([][]byte, error) {
	buf, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
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

	return result, nil
}
