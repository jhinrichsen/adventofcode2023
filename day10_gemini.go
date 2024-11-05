package adventofcode2023

type Point struct {
	x, y int
}

func findFarthestTile(grid [][]byte) uint {
	// Find the starting point
	start := Point{-1, -1}
	for i, row := range grid {
		for j, cell := range row {
			if cell == 'S' {
				start = Point{i, j}
				break
			}
		}
		if start.x != -1 {
			break
		}
	}

	// Find the main loop
	visited := make(map[Point]bool)
	loop := []Point{start}
	stack := []Point{start}
	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for _, neighbor := range getNeighbors(curr, grid) {
			if !visited[neighbor] {
				visited[neighbor] = true
				loop = append(loop, neighbor)
				stack = append(stack, neighbor)
			}
		}
	}

	// Calculate distances from the starting point
	distances := make(map[Point]uint)
	queue := []Point{start}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		for _, neighbor := range getNeighbors(curr, grid) {
			if _, ok := distances[neighbor]; !ok {
				distances[neighbor] = distances[curr] + 1
				queue = append(queue, neighbor)
			}
		}
	}

	// Find the maximum distance in the loop
	maxDistance := uint(0)
	for _, p := range loop {
		if distances[p] > maxDistance {
			maxDistance = distances[p]
		}
	}

	return maxDistance
}

func getNeighbors(p Point, grid [][]byte) []Point {
	neighbors := []Point{}
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for _, dir := range directions {
		x, y := p.x+dir[0], p.y+dir[1]
		if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) && grid[x][y] != '.' {
			neighbors = append(neighbors, Point{x, y})
		}
	}
	return neighbors
}

func Day10Gemini(grid [][]byte) uint {
	farthestDistance := findFarthestTile(grid)
	return farthestDistance
}
