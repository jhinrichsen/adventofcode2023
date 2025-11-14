package adventofcode2023

func findVerticalReflection(lines []string, start, end int) int {
	if start >= end || len(lines[start]) == 0 {
		return 0
	}

	cols := len(lines[start])

	for col := 1; col < cols; col++ {
		if isVerticalReflection(lines, start, end, col) {
			return col
		}
	}
	return 0
}

func isVerticalReflection(lines []string, start, end, col int) bool {
	cols := len(lines[start])

	for i := 0; i < min(col, cols-col); i++ {
		left := col - 1 - i
		right := col + i

		for row := start; row < end; row++ {
			if lines[row][left] != lines[row][right] {
				return false
			}
		}
	}
	return true
}

func findHorizontalReflection(lines []string, start, end int) int {
	if start >= end {
		return 0
	}

	rows := end - start

	for row := 1; row < rows; row++ {
		if isHorizontalReflection(lines, start, end, row) {
			return row
		}
	}
	return 0
}

func isHorizontalReflection(lines []string, start, end, row int) bool {
	rows := end - start

	for i := 0; i < min(row, rows-row); i++ {
		above := start + row - 1 - i
		below := start + row + i

		if lines[above] != lines[below] {
			return false
		}
	}
	return true
}

func Day13(lines []string, part1 bool) uint {
	if !part1 {
		return 0
	}

	var total uint
	start := 0

	for i := 0; i <= len(lines); i++ {
		if i == len(lines) || lines[i] == "" {
			if i > start {
				if vr := findVerticalReflection(lines, start, i); vr > 0 {
					total += uint(vr)
				} else if hr := findHorizontalReflection(lines, start, i); hr > 0 {
					total += uint(hr * 100)
				}
			}
			start = i + 1
		}
	}

	return total
}
