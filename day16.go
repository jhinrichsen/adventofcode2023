package adventofcode2023

func Day16(lines []string, part1 bool) uint {
	if len(lines) == 0 {
		return 0
	}

	type pos struct {
		row, col int
		dr, dc   int
	}

	rows := len(lines)
	cols := len(lines[0])

	// Direction mapping: up=0, right=1, down=2, left=3
	dirIndex := func(dr, dc int) int {
		if dr == -1 {
			return 0
		} else if dc == 1 {
			return 1
		} else if dr == 1 {
			return 2
		} else {
			return 3
		}
	}

	// Pre-allocate reusable data structures for energize function
	const maxSize = 120 // Support grids up to 120×120
	var visited [maxSize][maxSize][4]bool
	var energized [maxSize][maxSize]bool

	energize := func(startRow, startCol, startDr, startDc int) uint {
		// Clear visited and energized arrays
		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				energized[i][j] = false
				for k := 0; k < 4; k++ {
					visited[i][j][k] = false
				}
			}
		}

		// Pre-allocate queue with reasonable capacity
		queue := make([]pos, 0, 1000)
		queue = append(queue, pos{row: startRow, col: startCol, dr: startDr, dc: startDc})

		for len(queue) > 0 {
			beam := queue[0]
			queue = queue[1:]

			if beam.row < 0 || beam.row >= rows || beam.col < 0 || beam.col >= cols {
				continue
			}

			dir := dirIndex(beam.dr, beam.dc)
			if visited[beam.row][beam.col][dir] {
				continue
			}
			visited[beam.row][beam.col][dir] = true
			energized[beam.row][beam.col] = true

			tile := lines[beam.row][beam.col]

			switch tile {
			case '.':
				queue = append(queue, pos{beam.row + beam.dr, beam.col + beam.dc, beam.dr, beam.dc})
			case '/':
				dr, dc := -beam.dc, -beam.dr
				queue = append(queue, pos{beam.row + dr, beam.col + dc, dr, dc})
			case '\\':
				dr, dc := beam.dc, beam.dr
				queue = append(queue, pos{beam.row + dr, beam.col + dc, dr, dc})
			case '|':
				if beam.dc != 0 {
					queue = append(queue, pos{beam.row - 1, beam.col, -1, 0})
					queue = append(queue, pos{beam.row + 1, beam.col, 1, 0})
				} else {
					queue = append(queue, pos{beam.row + beam.dr, beam.col + beam.dc, beam.dr, beam.dc})
				}
			case '-':
				if beam.dr != 0 {
					queue = append(queue, pos{beam.row, beam.col - 1, 0, -1})
					queue = append(queue, pos{beam.row, beam.col + 1, 0, 1})
				} else {
					queue = append(queue, pos{beam.row + beam.dr, beam.col + beam.dc, beam.dr, beam.dc})
				}
			}
		}

		// Count energized cells
		var count uint
		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				if energized[i][j] {
					count++
				}
			}
		}
		return count
	}

	if part1 {
		return energize(0, 0, 0, 1)
	}

	var maxEnergized uint
	for col := 0; col < cols; col++ {
		if e := energize(0, col, 1, 0); e > maxEnergized {
			maxEnergized = e
		}
		if e := energize(rows-1, col, -1, 0); e > maxEnergized {
			maxEnergized = e
		}
	}
	for row := 0; row < rows; row++ {
		if e := energize(row, 0, 0, 1); e > maxEnergized {
			maxEnergized = e
		}
		if e := energize(row, cols-1, 0, -1); e > maxEnergized {
			maxEnergized = e
		}
	}

	return maxEnergized
}
