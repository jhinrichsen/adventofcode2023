package adventofcode2023

type Day21Puzzle struct {
	grid  [][]byte
	start [2]int
}

func NewDay21(lines []string) (Day21Puzzle, error) {
	var puzzle Day21Puzzle
	// Pre-allocate grid with estimated capacity
	puzzle.grid = make([][]byte, 0, len(lines))

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
	if part1 {
		return countReachable(puzzle, 64)
	}

	// Part 2: Infinite grid
	// The grid is 131x131, start is at (65,65)
	// Target steps: 26501365 = 202300 * 131 + 65
	// Use quadratic extrapolation
	gridSize := len(puzzle.grid)
	steps := 26501365
	n := steps / gridSize // 202300

	// Calculate values at 65, 65+131, 65+131*2
	y0 := countReachableInfinite(puzzle, gridSize/2)
	y1 := countReachableInfinite(puzzle, gridSize/2+gridSize)
	y2 := countReachableInfinite(puzzle, gridSize/2+gridSize*2)

	// Fit quadratic: f(n) = a*n^2 + b*n + c
	// f(0) = c = y0
	// f(1) = a + b + c = y1
	// f(2) = 4a + 2b + c = y2
	// Solving: a = (y2 - 2*y1 + y0) / 2
	//          b = y1 - y0 - a
	//          c = y0
	a := (y2 - 2*y1 + y0) / 2
	b := y1 - y0 - a
	c := y0

	return uint(a*n*n + b*n + c)
}

func countReachable(puzzle Day21Puzzle, maxSteps int) uint {
	type pos struct{ x, y int }

	height := len(puzzle.grid)
	width := len(puzzle.grid[0])

	// BFS to find minimum steps to each position
	// Pre-allocate with estimated capacity
	visited := make(map[pos]int, height*width/2)
	queue := make([]struct {
		p     pos
		steps int
	}, 0, 1000)
	queue = append(queue, struct {
		p     pos
		steps int
	}{p: pos{puzzle.start[0], puzzle.start[1]}, steps: 0})

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

func countReachableInfinite(puzzle Day21Puzzle, maxSteps int) int {
	type pos struct{ x, y int }

	height := len(puzzle.grid)
	width := len(puzzle.grid[0])

	// BFS on infinite grid
	// Pre-allocate with larger capacity for infinite grid
	visited := make(map[pos]int, 50000)
	queue := make([]struct {
		p     pos
		steps int
	}, 0, 5000)
	queue = append(queue, struct {
		p     pos
		steps int
	}{p: pos{puzzle.start[0], puzzle.start[1]}, steps: 0})

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

			// Get grid position (wrapping for infinite grid)
			gx := ((nx % width) + width) % width
			gy := ((ny % height) + height) % height

			if puzzle.grid[gy][gx] == '#' {
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
	count := 0
	targetParity := maxSteps % 2

	for _, steps := range visited {
		if steps <= maxSteps && steps%2 == targetParity {
			count++
		}
	}

	return count
}
