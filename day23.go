package adventofcode2023

type Day23Puzzle struct {
	grid  [][]byte
	start [2]int
	end   [2]int
}

func NewDay23(lines []string) (Day23Puzzle, error) {
	var puzzle Day23Puzzle

	for y, line := range lines {
		if line == "" {
			continue
		}
		puzzle.grid = append(puzzle.grid, []byte(line))

		// Find start (top row)
		if y == 0 {
			for x, ch := range line {
				if ch == '.' {
					puzzle.start = [2]int{x, y}
					break
				}
			}
		}
	}

	// Find end (bottom row)
	lastRow := len(puzzle.grid) - 1
	for x, ch := range puzzle.grid[lastRow] {
		if ch == '.' {
			puzzle.end = [2]int{x, lastRow}
			break
		}
	}

	return puzzle, nil
}

func Day23(puzzle Day23Puzzle, part1 bool) uint {
	if !part1 {
		return 0
	}

	height := len(puzzle.grid)
	width := len(puzzle.grid[0])

	visited := make([][]bool, height)
	for i := range visited {
		visited[i] = make([]bool, width)
	}

	var dfs func(x, y int) int
	dfs = func(x, y int) int {
		if x == puzzle.end[0] && y == puzzle.end[1] {
			return 0
		}

		visited[y][x] = true
		maxDist := -1

		// Get possible moves
		moves := [][2]int{}
		cell := puzzle.grid[y][x]

		switch cell {
		case '>':
			moves = append(moves, [2]int{x + 1, y})
		case '<':
			moves = append(moves, [2]int{x - 1, y})
		case 'v':
			moves = append(moves, [2]int{x, y + 1})
		case '^':
			moves = append(moves, [2]int{x, y - 1})
		case '.':
			moves = append(moves, [2]int{x + 1, y}, [2]int{x - 1, y}, [2]int{x, y + 1}, [2]int{x, y - 1})
		}

		for _, move := range moves {
			nx, ny := move[0], move[1]

			if nx < 0 || ny < 0 || nx >= width || ny >= height {
				continue
			}
			if puzzle.grid[ny][nx] == '#' {
				continue
			}
			if visited[ny][nx] {
				continue
			}

			dist := dfs(nx, ny)
			if dist >= 0 {
				maxDist = max(maxDist, dist+1)
			}
		}

		visited[y][x] = false
		return maxDist
	}

	result := dfs(puzzle.start[0], puzzle.start[1])
	return uint(result)
}
