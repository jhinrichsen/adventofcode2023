package adventofcode2023

// floodFill performs a flood fill on a grid starting from (x, y) and replaces target with replacement.
func floodFill(grid [][]int, x, y, target, replacement int) {
	rows := len(grid)
	cols := len(grid[0])

	if x < 0 || y < 0 || x >= rows || y >= cols || grid[x][y] != target {
		return
	}

	grid[x][y] = replacement

	// Directions for 4-connected grid (up, down, left, right)
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for _, dir := range directions {
		newX, newY := x+dir[0], y+dir[1]
		floodFill(grid, newX, newY, target, replacement)
	}
}

// floodFillWithHoles identifies inside, outside, and boundary regions in a grid.
func floodFillWithHoles(grid [][]int) [][]int {
	rows := len(grid)
	cols := len(grid[0])

	// Step 1: Mark the outer area as "outside" (use value 2)
	for i := 0; i < rows; i++ {
		if grid[i][0] == 0 { // Left edge
			floodFill(grid, i, 0, 0, 2)
		}
		if grid[i][cols-1] == 0 { // Right edge
			floodFill(grid, i, cols-1, 0, 2)
		}
	}
	for j := 0; j < cols; j++ {
		if grid[0][j] == 0 { // Top edge
			floodFill(grid, 0, j, 0, 2)
		}
		if grid[rows-1][j] == 0 { // Bottom edge
			floodFill(grid, rows-1, j, 0, 2)
		}
	}

	// Step 2: Label inside cells
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 0 { // Unvisited cell
				grid[i][j] = 3 // Mark as inside
			}
		}
	}

	return grid
}
