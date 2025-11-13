package adventofcode2023

func NewDay13(lines []string) [][]string {
	var grids [][]string
	var current []string

	for _, line := range lines {
		if line == "" {
			if len(current) > 0 {
				grids = append(grids, current)
				current = nil
			}
		} else {
			current = append(current, line)
		}
	}
	if len(current) > 0 {
		grids = append(grids, current)
	}

	return grids
}

func findVerticalReflection(grid []string) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	cols := len(grid[0])

	for col := 1; col < cols; col++ {
		if isVerticalReflection(grid, col) {
			return col
		}
	}
	return 0
}

func isVerticalReflection(grid []string, col int) bool {
	cols := len(grid[0])

	for i := 0; i < min(col, cols-col); i++ {
		left := col - 1 - i
		right := col + i

		for row := 0; row < len(grid); row++ {
			if grid[row][left] != grid[row][right] {
				return false
			}
		}
	}
	return true
}

func findHorizontalReflection(grid []string) int {
	if len(grid) == 0 {
		return 0
	}

	rows := len(grid)

	for row := 1; row < rows; row++ {
		if isHorizontalReflection(grid, row) {
			return row
		}
	}
	return 0
}

func isHorizontalReflection(grid []string, row int) bool {
	rows := len(grid)

	for i := 0; i < min(row, rows-row); i++ {
		above := row - 1 - i
		below := row + i

		if grid[above] != grid[below] {
			return false
		}
	}
	return true
}

func Day13(grids [][]string, part1 bool) uint {
	if !part1 {
		return 0
	}

	var total uint
	for _, grid := range grids {
		if vr := findVerticalReflection(grid); vr > 0 {
			total += uint(vr)
		} else if hr := findHorizontalReflection(grid); hr > 0 {
			total += uint(hr * 100)
		}
	}
	return total
}
