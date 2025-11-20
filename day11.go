package adventofcode2023

import (
	"image"
)

func Day11(lines []string, part1 bool) uint {
	expansion := uint(1)
	if !part1 {
		expansion = 1_000_000
	}
	return day11Solver(lines, expansion)
}

func day11Solver(lines []string, expansion uint) uint {
	// Parse lines into [][]byte
	grid := make([][]byte, len(lines))
	for i, line := range lines {
		grid[i] = []byte(line)
	}

	const galaxy = '#'
	dimX, dimY := len(grid[0]), len(grid)

	if expansion > 1 {
		expansion--
	}

	// Find all points with the galaxy symbol
	var points []image.Point
	for y := range dimY {
		for x := range dimX {
			if grid[y][x] == galaxy {
				points = append(points, image.Point{X: x, Y: y})
			}
		}
	}

	// mark all rows as expanded
	for y := range dimY {
		grid[y][0] = 1
	}
	// mark all columns as expanded
	for x := range dimX {
		grid[0][x] = 1
	}
	// unset cols and rows that contain galaxies
	for _, p := range points {
		grid[0][p.X] = 0
		grid[p.Y][0] = 0
	}

	var total uint
	for i := range points {
		p1 := points[i]
		for j := i + 1; j < len(points); j++ {
			p2 := points[j]

			// Manhattan distance
			dx := abs(p2.X - p1.X)
			dy := abs(p2.Y - p1.Y)
			total += uint(dx + dy)

			// Add expanded spaces (horizontal)
			x0, x1 := min(p1.X, p2.X), max(p1.X, p2.X)
			for x := x0; x < x1; x++ {
				total += uint(grid[0][x]) * expansion
			}

			// Add expanded spaces (vertical)
			y0, y1 := min(p1.Y, p2.Y), max(p1.Y, p2.Y)
			for y := y0; y < y1; y++ {
				total += uint(grid[y][0]) * expansion
			}
		}
	}
	return total
}

func abs(n int) uint {
	if n < 0 {
		return uint(-n)
	}
	return uint(n)
}
