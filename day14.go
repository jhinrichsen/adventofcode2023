package adventofcode2023

type Day14Puzzle struct {
	grid [][]byte
}

func NewDay14(lines []string) (Day14Puzzle, error) {
	grid := make([][]byte, len(lines))
	for i, line := range lines {
		grid[i] = []byte(line)
	}
	return Day14Puzzle{grid: grid}, nil
}

func tiltNorth(grid [][]byte) {
	rows := len(grid)
	if rows == 0 {
		return
	}
	cols := len(grid[0])

	for col := 0; col < cols; col++ {
		for row := 1; row < rows; row++ {
			if grid[row][col] == 'O' {
				newRow := row
				for newRow > 0 && grid[newRow-1][col] == '.' {
					newRow--
				}
				if newRow != row {
					grid[newRow][col] = 'O'
					grid[row][col] = '.'
				}
			}
		}
	}
}

func calculateLoad(grid [][]byte) uint {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])

	var load uint
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == 'O' {
				load += uint(rows - row)
			}
		}
	}
	return load
}

func Day14(puzzle Day14Puzzle, part1 bool) uint {
	if !part1 {
		return 0
	}

	grid := make([][]byte, len(puzzle.grid))
	for i := range puzzle.grid {
		grid[i] = make([]byte, len(puzzle.grid[i]))
		copy(grid[i], puzzle.grid[i])
	}

	tiltNorth(grid)
	return calculateLoad(grid)
}
