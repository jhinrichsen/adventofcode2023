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
	// Parse lines into flat []byte array
	dimY := len(lines)
	dimX := len(lines[0])
	grid := make([]byte, dimX*dimY)

	for y, line := range lines {
		copy(grid[y*dimX:y*dimX+dimX], line)
	}

	const galaxy = '#'

	if expansion > 1 {
		expansion--
	}

	// Find all points with the galaxy symbol
	var points []image.Point
	for y := range dimY {
		for x := range dimX {
			if grid[y*dimX+x] == galaxy {
				points = append(points, image.Point{X: x, Y: y})
			}
		}
	}

	// mark all rows as expanded (column 0)
	for y := range dimY {
		grid[y*dimX] = 1
	}
	// mark all columns as expanded (row 0)
	for x := range dimX {
		grid[x] = 1
	}
	// unset cols and rows that contain galaxies
	for _, p := range points {
		grid[p.X] = 0        // column flag at row 0
		grid[p.Y*dimX] = 0   // row flag at column 0
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
				total += uint(grid[x]) * expansion
			}

			// Add expanded spaces (vertical)
			y0, y1 := min(p1.Y, p2.Y), max(p1.Y, p2.Y)
			for y := y0; y < y1; y++ {
				total += uint(grid[y*dimX]) * expansion
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
