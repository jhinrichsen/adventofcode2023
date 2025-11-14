package adventofcode2023

type Day21Puzzle struct {
	grid  [][]byte
	start [2]int
}

func NewDay21(lines []string) (Day21Puzzle, error) {
	var puzzle Day21Puzzle

	for y, line := range lines {
		if line == "" {
			continue
		}
		row := []byte(line)
		puzzle.grid = append(puzzle.grid, row)

		for x, ch := range row {
			if ch == 'S' {
				puzzle.start = [2]int{x, y}
				row[x] = '.' // Treat start as open space
			}
		}
	}

	return puzzle, nil
}

func Day21(puzzle Day21Puzzle, part1 bool) uint {
	if !part1 {
		return 0
	}

	return countReachable(puzzle, 64)
}

func countReachable(puzzle Day21Puzzle, maxSteps int) uint {
	type pos struct{ x, y int }

	height := len(puzzle.grid)
	width := len(puzzle.grid[0])

	// BFS to find minimum steps to each position
	visited := make(map[pos]int)
	queue := []struct {
		p     pos
		steps int
	}{{p: pos{puzzle.start[0], puzzle.start[1]}, steps: 0}}

	visited[pos{puzzle.start[0], puzzle.start[1]}] = 0

	dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.steps >= maxSteps {
			continue
		}

		for _, d := range dirs {
			nx, ny := curr.p.x+d[0], curr.p.y+d[1]

			if nx < 0 || ny < 0 || nx >= width || ny >= height {
				continue
			}
			if puzzle.grid[ny][nx] == '#' {
				continue
			}

			next := pos{nx, ny}
			if _, seen := visited[next]; seen {
				continue
			}

			visited[next] = curr.steps + 1
			queue = append(queue, struct {
				p     pos
				steps int
			}{next, curr.steps + 1})
		}
	}

	// Count positions reachable in exactly maxSteps
	// A position is reachable in exactly N steps if it's reachable in K steps
	// where K <= N and K has the same parity as N
	var count uint
	targetParity := maxSteps % 2

	for _, steps := range visited {
		if steps <= maxSteps && steps%2 == targetParity {
			count++
		}
	}

	return count
}
