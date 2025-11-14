package adventofcode2023

import (
	"bytes"
)

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

func tiltWest(grid [][]byte) {
	rows := len(grid)
	if rows == 0 {
		return
	}
	cols := len(grid[0])

	for row := 0; row < rows; row++ {
		for col := 1; col < cols; col++ {
			if grid[row][col] == 'O' {
				newCol := col
				for newCol > 0 && grid[row][newCol-1] == '.' {
					newCol--
				}
				if newCol != col {
					grid[row][newCol] = 'O'
					grid[row][col] = '.'
				}
			}
		}
	}
}

func tiltSouth(grid [][]byte) {
	rows := len(grid)
	if rows == 0 {
		return
	}
	cols := len(grid[0])

	for col := 0; col < cols; col++ {
		for row := rows - 2; row >= 0; row-- {
			if grid[row][col] == 'O' {
				newRow := row
				for newRow < rows-1 && grid[newRow+1][col] == '.' {
					newRow++
				}
				if newRow != row {
					grid[newRow][col] = 'O'
					grid[row][col] = '.'
				}
			}
		}
	}
}

func tiltEast(grid [][]byte) {
	rows := len(grid)
	if rows == 0 {
		return
	}
	cols := len(grid[0])

	for row := 0; row < rows; row++ {
		for col := cols - 2; col >= 0; col-- {
			if grid[row][col] == 'O' {
				newCol := col
				for newCol < cols-1 && grid[row][newCol+1] == '.' {
					newCol++
				}
				if newCol != col {
					grid[row][newCol] = 'O'
					grid[row][col] = '.'
				}
			}
		}
	}
}

func runCycle(grid [][]byte) {
	tiltNorth(grid)
	tiltWest(grid)
	tiltSouth(grid)
	tiltEast(grid)
}

func gridToString(grid [][]byte) string {
	var buf bytes.Buffer
	for _, row := range grid {
		buf.Write(row)
		buf.WriteByte('\n')
	}
	return buf.String()
}

func copyGrid(src [][]byte) [][]byte {
	dst := make([][]byte, len(src))
	for i := range src {
		dst[i] = make([]byte, len(src[i]))
		copy(dst[i], src[i])
	}
	return dst
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
	grid := copyGrid(puzzle.grid)

	if part1 {
		tiltNorth(grid)
		return calculateLoad(grid)
	}

	seen := make(map[string]int)
	const totalCycles = 1000000000

	for i := 0; i < totalCycles; i++ {
		key := gridToString(grid)
		if prev, ok := seen[key]; ok {
			cycleLength := i - prev
			remaining := (totalCycles - i) % cycleLength
			for j := 0; j < remaining; j++ {
				runCycle(grid)
			}
			return calculateLoad(grid)
		}
		seen[key] = i
		runCycle(grid)
	}

	return calculateLoad(grid)
}
