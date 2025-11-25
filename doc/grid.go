package adventofcode2023

import "image"

// NewGrid creates a 2D array of size [X, Y].
func NewGrid(x, y int) Grid {
	var g Grid
	for i := 0; i < y; i++ {

	}
	return g
}

type Grid struct {
	DimX, DimY int
	Points     [][]int
}

func (a Grid) Contains(x, y int) bool {
	return x >= 0 &&
		y >= 0 &&
		x < a.DimX &&
		y < a.DimY
}

func (a Grid) Get(x, y int) int {
	return a.Points[y][x]
}

func (a Grid) Set(x, y, n int) {
	a.Points[y][x] = n
}

func (a Grid) ContainsPoint(p image.Point) bool {
	return a.Contains(p.X, p.Y)
}

// Con4 returns a minimum of 2, and a maximum of 4 neighbors (north, east, south, west).
func (a Grid) Con4(x, y int) []image.Point {
	var ps []image.Point
	for _, p := range []image.Point{
		{x, y - 1},
		{x + 1, y},
		{x, y + 1},
		{x - 1, y},
	} {
		if a.ContainsPoint(p) {
			ps = append(ps, p)
		}
	}
	return ps
}

// Con8 returns a minimum of 3, and a maximum of 8 neighbors (north, northeast, east, southeast, south, southwest, west, northwest).
func (a Grid) Con8(x, y int) []image.Point {
	var ps []image.Point
	for _, p := range []image.Point{
		{x, y - 1},
		{x + 1, y - 1},
		{x + 1, y},
		{x + 1, y + 1},
		{x, y + 1},
		{x - 1, y + 1},
		{x - 1, y},
		{x - 1, y - 1},
	} {
		if a.ContainsPoint(p) {
			ps = append(ps, p)
		}
	}
	return ps
}
