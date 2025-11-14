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

	energize := func(startRow, startCol, startDr, startDc int) uint {
		visited := make(map[pos]bool)
		energized := make(map[[2]int]bool)
		queue := []pos{{row: startRow, col: startCol, dr: startDr, dc: startDc}}

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

			tile := lines[beam.row][beam.col]
			var nextBeams []pos

			switch tile {
			case '.':
				nextBeams = append(nextBeams, pos{beam.row + beam.dr, beam.col + beam.dc, beam.dr, beam.dc})
			case '/':
				dr, dc := -beam.dc, -beam.dr
				nextBeams = append(nextBeams, pos{beam.row + dr, beam.col + dc, dr, dc})
			case '\\':
				dr, dc := beam.dc, beam.dr
				nextBeams = append(nextBeams, pos{beam.row + dr, beam.col + dc, dr, dc})
			case '|':
				if beam.dc != 0 {
					nextBeams = append(nextBeams, pos{beam.row - 1, beam.col, -1, 0})
					nextBeams = append(nextBeams, pos{beam.row + 1, beam.col, 1, 0})
				} else {
					nextBeams = append(nextBeams, pos{beam.row + beam.dr, beam.col + beam.dc, beam.dr, beam.dc})
				}
			case '-':
				if beam.dr != 0 {
					nextBeams = append(nextBeams, pos{beam.row, beam.col - 1, 0, -1})
					nextBeams = append(nextBeams, pos{beam.row, beam.col + 1, 0, 1})
				} else {
					nextBeams = append(nextBeams, pos{beam.row + beam.dr, beam.col + beam.dc, beam.dr, beam.dc})
				}
			}

			queue = append(queue, nextBeams...)
		}

		return uint(len(energized))
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
