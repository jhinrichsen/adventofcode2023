package adventofcode2023

import (
	"image"
)

// nPairs returns number of nPairs for N points.
func nPairs(points uint) uint {
	if points < 2 {
		return 0
	}
	return (points * (points - 1)) / 2
}

func Day11(grid [][]byte) uint {
	const (
		galaxy   = '#'
		plain    = 0
		expanded = plain + 1
	)
	dimX, dimY := len(grid[0]), len(grid)

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
		grid[y][0] = expanded
	}
	// mark all columns as expanded
	for x := range dimX {
		grid[0][x] = expanded
	}
	// unset cols and rows that contain galaxies
	for _, p := range points {
		grid[0][p.X] = plain
		grid[p.Y][0] = plain
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
				total += uint(grid[0][x])
			}

			// Add expanded spaces (vertical)
			y0, y1 := min(p1.Y, p2.Y), max(p1.Y, p2.Y)
			for y := y0; y < y1; y++ {
				total += uint(grid[y][0])
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

func sign[T int | int8 | int16 | int32 | int64](n T) int {
	if n < 0 {
		return -1
	}
	if n > 0 {
		return 1
	}
	return 0
}
