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
	const galaxy = '#'
	dimX, dimY := len(grid[0]), len(grid)

	points := func() []image.Point {
		var ps []image.Point
		for y := range dimY {
			for x := range dimX {
				if grid[y][x] == galaxy {
					ps = append(ps, image.Point{x, y})
				}
			}
		}
		return ps
	}()

	// reuse grid to highlight expanded rows and cols in [0]
	const (
		plain    = '0'
		expanded = plain + 1
	)
	for y := range dimY {
		grid[y][0] = expanded
	}
	for x := range dimX {
		grid[0][x] = expanded
	}
	for _, p := range points {
		grid[0][p.X] = plain
		grid[p.Y][0] = plain
	}

	// travel and count
	pairs := func() [][2]image.Point {
		var pairs [][2]image.Point

		// generate all unique pairs
		for i := 0; i < len(points); i++ {
			for j := i + 1; j < len(points); j++ {
				pairs = append(pairs, [2]image.Point{points[i], points[j]})
			}
		}

		return pairs
	}()
	var total uint
	var steps uint
	for i := range pairs {
		p1 := pairs[i][0]
		p2 := pairs[i][1]
		dx := abs(p2.X - p1.X)
		dy := abs(p2.Y - p1.Y)
		// manhattan distance
		steps = uint(dx + dy)

		// add expanded spaces
		// 1. horizontally
		x0 := uint(min(p1.X, p2.X))
		x1 := x0 + dx
		for x := x0; x < x1; x++ {
			steps += uint(grid[0][x] - plain)
		}
		// 2. vertically
		y0 := uint(min(p1.Y, p2.Y))
		y1 := y0 + dy
		for y := y0; y < y1; y++ {
			steps += uint(grid[y][0] - plain)
		}

		total += steps
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
