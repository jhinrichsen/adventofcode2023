package adventofcode2023

type pattern struct {
	grid []string
}

func NewDay13(lines []string) []pattern {
	var patterns []pattern
	var current []string

	for _, line := range lines {
		if line == "" {
			if len(current) > 0 {
				patterns = append(patterns, pattern{grid: current})
				current = nil
			}
		} else {
			current = append(current, line)
		}
	}
	if len(current) > 0 {
		patterns = append(patterns, pattern{grid: current})
	}

	return patterns
}

func findVerticalReflection(p pattern) int {
	if len(p.grid) == 0 || len(p.grid[0]) == 0 {
		return 0
	}

	cols := len(p.grid[0])

	for col := 1; col < cols; col++ {
		if isVerticalReflection(p, col) {
			return col
		}
	}
	return 0
}

func isVerticalReflection(p pattern, col int) bool {
	cols := len(p.grid[0])

	for i := 0; i < min(col, cols-col); i++ {
		left := col - 1 - i
		right := col + i

		for row := 0; row < len(p.grid); row++ {
			if p.grid[row][left] != p.grid[row][right] {
				return false
			}
		}
	}
	return true
}

func findHorizontalReflection(p pattern) int {
	if len(p.grid) == 0 {
		return 0
	}

	rows := len(p.grid)

	for row := 1; row < rows; row++ {
		if isHorizontalReflection(p, row) {
			return row
		}
	}
	return 0
}

func isHorizontalReflection(p pattern, row int) bool {
	rows := len(p.grid)

	for i := 0; i < min(row, rows-row); i++ {
		above := row - 1 - i
		below := row + i

		if p.grid[above] != p.grid[below] {
			return false
		}
	}
	return true
}

func Day13(patterns []pattern, part1 bool) uint {
	if !part1 {
		return 0
	}

	var total uint
	for _, p := range patterns {
		if vr := findVerticalReflection(p); vr > 0 {
			total += uint(vr)
		} else if hr := findHorizontalReflection(p); hr > 0 {
			total += uint(hr * 100)
		}
	}
	return total
}
