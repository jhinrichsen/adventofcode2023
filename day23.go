package adventofcode2023

type Day23Puzzle struct {
	grid  [][]byte
	start [2]int
	end   [2]int
}

func NewDay23(lines []string) (Day23Puzzle, error) {
	var puzzle Day23Puzzle
	puzzle.grid = make([][]byte, 0, len(lines))

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
	height := len(puzzle.grid)
	width := len(puzzle.grid[0])

	if part1 {
		// Part 1: simple DFS with slope enforcement
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
			moves := make([][2]int, 0, 4)
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

	// Part 2: compress graph into junctions and run DFS on compressed graph
	type pos struct{ x, y int }

	// Find all junctions (nodes with more than 2 neighbors, plus start and end)
	junctions := make(map[pos]bool)
	junctions[pos{puzzle.start[0], puzzle.start[1]}] = true
	junctions[pos{puzzle.end[0], puzzle.end[1]}] = true

	for y := range puzzle.grid {
		for x := range puzzle.grid[y] {
			if puzzle.grid[y][x] == '#' {
				continue
			}

			neighbors := 0
			for _, d := range [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
				nx, ny := x+d[0], y+d[1]
				if nx >= 0 && ny >= 0 && nx < width && ny < height && puzzle.grid[ny][nx] != '#' {
					neighbors++
				}
			}

			if neighbors > 2 {
				junctions[pos{x, y}] = true
			}
		}
	}

	// Build graph of junctions with edge weights
	type edge struct {
		to   pos
		dist int
	}
	graph := make(map[pos][]edge, len(junctions))
	for junction := range junctions {
		graph[junction] = make([]edge, 0, 4)
	}

	for junction := range junctions {
		// BFS from this junction to find reachable junctions
		visited := make(map[pos]bool)
		queue := []struct {
			p    pos
			dist int
		}{{junction, 0}}
		visited[junction] = true

		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]

			if curr.dist > 0 && junctions[curr.p] {
				// Reached another junction
				graph[junction] = append(graph[junction], edge{curr.p, curr.dist})
				continue
			}

			for _, d := range [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
				nx, ny := curr.p.x+d[0], curr.p.y+d[1]
				next := pos{nx, ny}

				if nx < 0 || ny < 0 || nx >= width || ny >= height {
					continue
				}
				if puzzle.grid[ny][nx] == '#' {
					continue
				}
				if visited[next] {
					continue
				}

				visited[next] = true
				queue = append(queue, struct {
					p    pos
					dist int
				}{next, curr.dist + 1})
			}
		}
	}

	// DFS on compressed graph
	start := pos{puzzle.start[0], puzzle.start[1]}
	end := pos{puzzle.end[0], puzzle.end[1]}
	visitedJunctions := make(map[pos]bool, len(junctions))

	var dfs func(p pos) int
	dfs = func(p pos) int {
		if p == end {
			return 0
		}

		visitedJunctions[p] = true
		maxDist := -1

		for _, e := range graph[p] {
			if visitedJunctions[e.to] {
				continue
			}

			dist := dfs(e.to)
			if dist >= 0 {
				maxDist = max(maxDist, dist+e.dist)
			}
		}

		visitedJunctions[p] = false
		return maxDist
	}

	result := dfs(start)
	return uint(result)
}
