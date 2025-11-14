package adventofcode2023

func findVerticalReflection(lines []string, start, end int) int {
	return findVerticalReflectionWithSmudges(lines, start, end, 0)
}

func findVerticalReflectionWithSmudges(lines []string, start, end, smudges int) int {
	if start >= end || len(lines[start]) == 0 {
		return 0
	}

	cols := len(lines[start])

	for col := 1; col < cols; col++ {
		if countVerticalMismatches(lines, start, end, col) == smudges {
			return col
		}
	}
	return 0
}

func countVerticalMismatches(lines []string, start, end, col int) int {
	cols := len(lines[start])
	mismatches := 0

	for i := 0; i < min(col, cols-col); i++ {
		left := col - 1 - i
		right := col + i

		for row := start; row < end; row++ {
			if lines[row][left] != lines[row][right] {
				mismatches++
			}
		}
	}
	return mismatches
}

func findHorizontalReflection(lines []string, start, end int) int {
	return findHorizontalReflectionWithSmudges(lines, start, end, 0)
}

func findHorizontalReflectionWithSmudges(lines []string, start, end, smudges int) int {
	if start >= end {
		return 0
	}

	rows := end - start

	for row := 1; row < rows; row++ {
		if countHorizontalMismatches(lines, start, end, row) == smudges {
			return row
		}
	}
	return 0
}

func countHorizontalMismatches(lines []string, start, end, row int) int {
	rows := end - start
	mismatches := 0

	for i := 0; i < min(row, rows-row); i++ {
		above := start + row - 1 - i
		below := start + row + i

		for col := 0; col < len(lines[above]); col++ {
			if lines[above][col] != lines[below][col] {
				mismatches++
			}
		}
	}
	return mismatches
}

func Day13(lines []string, part1 bool) uint {
	smudges := 0
	if !part1 {
		smudges = 1
	}

	var total uint
	start := 0

	for i := 0; i <= len(lines); i++ {
		if i == len(lines) || lines[i] == "" {
			if i > start {
				if vr := findVerticalReflectionWithSmudges(lines, start, i, smudges); vr > 0 {
					total += uint(vr)
				} else if hr := findHorizontalReflectionWithSmudges(lines, start, i, smudges); hr > 0 {
					total += uint(hr * 100)
				}
			}
			start = i + 1
		}
	}

	return total
}
