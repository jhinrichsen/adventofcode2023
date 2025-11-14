package adventofcode2023

type pos16 struct {
	row, col int
	dr, dc   int
}

func Day16(lines []string, part1 bool) uint {
	if !part1 {
		return 0
	}

	return energize(lines, 0, 0, 0, 1)
}

func energize(grid []string, startRow, startCol, startDr, startDc int) uint {
	if len(grid) == 0 {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])

	visited := make(map[pos16]bool)
	energized := make(map[[2]int]bool)
	queue := []pos16{{row: startRow, col: startCol, dr: startDr, dc: startDc}}

	for len(queue) > 0 {
		beam := queue[0]
		queue = queue[1:]

		if beam.row < 0 || beam.row >= rows || beam.col < 0 || beam.col >= cols {
			continue
		}

		if visited[beam] {
			continue
		}
		visited[beam] = true
		energized[[2]int{beam.row, beam.col}] = true

		tile := grid[beam.row][beam.col]
		var nextBeams []pos16

		switch tile {
		case '.':
			nextBeams = append(nextBeams, pos16{beam.row + beam.dr, beam.col + beam.dc, beam.dr, beam.dc})
		case '/':
			dr, dc := -beam.dc, -beam.dr
			nextBeams = append(nextBeams, pos16{beam.row + dr, beam.col + dc, dr, dc})
		case '\\':
			dr, dc := beam.dc, beam.dr
			nextBeams = append(nextBeams, pos16{beam.row + dr, beam.col + dc, dr, dc})
		case '|':
			if beam.dc != 0 {
				nextBeams = append(nextBeams, pos16{beam.row - 1, beam.col, -1, 0})
				nextBeams = append(nextBeams, pos16{beam.row + 1, beam.col, 1, 0})
			} else {
				nextBeams = append(nextBeams, pos16{beam.row + beam.dr, beam.col + beam.dc, beam.dr, beam.dc})
			}
		case '-':
			if beam.dr != 0 {
				nextBeams = append(nextBeams, pos16{beam.row, beam.col - 1, 0, -1})
				nextBeams = append(nextBeams, pos16{beam.row, beam.col + 1, 0, 1})
			} else {
				nextBeams = append(nextBeams, pos16{beam.row + beam.dr, beam.col + beam.dc, beam.dr, beam.dc})
			}
		}

		queue = append(queue, nextBeams...)
	}

	return uint(len(energized))
}
