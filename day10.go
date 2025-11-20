package adventofcode2023

import (
	"fmt"
	"image"
	"math/bits"
)

type direction uint8

const (
	North direction = 1 << iota
	South
	West
	East
)

// opposite returns the opposite direction, or the same direction for illegal directions.
func opposite1(d direction) direction {
	switch d {
	case North:
		return South
	case South:
		return North
	case West:
		return East
	case East:
		return West
	}
	return d
}

func opposite2(d direction) direction {
	switch d {
	case North:
		d = South
	case South:
		d = North
	case West:
		d = East
	case East:
		d = West
	}
	return d
}

func opposite3(in direction) direction {
	var out direction
	bitPosition := bits.Len8(uint8(in))
	if bitPosition%2 == 0 {
		out = in >> 1
	} else {
		out = in << 1
	}
	return out
}

var connects = [128]direction{
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, // 0
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, // 16
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, West | East, 0, 0, // 32
	0, 0, 0, 0, 0, 0, 0, South | West, 0, 0, 0, 0, 0, 0, 0, 0, // 48
	0, 0, 0, 0, 0, 0, South | East, 0, 0, 0, North | West, 0, North | East, 0, 0, 0, // 64
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, // 80
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, // 96
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, North | South, 0, 0, 0, // 112
}

func hasConnection(idx byte, d direction) bool {
	return (connects[idx] & d) != 0
}

func other(idx byte, d direction) direction {
	return connects[idx] & ^d
}

func Day10(buf []byte, part1 bool) (uint, error) {
	// Parse buf into [][]byte
	var lines [][]byte
	lineStart := 0
	for i := 0; i < len(buf); i++ {
		if buf[i] == '\n' {
			line := make([]byte, i-lineStart)
			copy(line, buf[lineStart:i])
			lines = append(lines, line)
			lineStart = i + 1
		}
	}
	if lineStart < len(buf) {
		line := make([]byte, len(buf)-lineStart)
		copy(line, buf[lineStart:])
		lines = append(lines, line)
	}

	const (
		startChar = 'S'
	)
	part2 := !part1
	var notFound = image.Point{-1, -1}

	dim := image.Rectangle{image.Point{0, 0}, image.Point{len(lines[0]), len(lines)}}

	var vertices uint
	// Pre-allocate poly with maximum possible size (grid perimeter)
	// The polygon loop can't be larger than the perimeter of the grid
	maxPolySize := 2 * (len(lines[0]) + len(lines))
	var poly []image.Point
	if part2 {
		poly = make([]image.Point, 0, maxPolySize)
	}
	in := func(p image.Point) bool {
		for i := range poly {
			if p == poly[i] {
				return true
			}
		}
		return false
	}

	track := func(p image.Point) {
		vertices++ // keep a separate counter for part 1
		if part2 {
			poly = append(poly, p)
		}

		const TRACE_PATH = false
		if TRACE_PATH {
			fmt.Println()
			for y := range lines {
				for x := range lines[0] {
					p_ := image.Point{x, y}
					if p == p_ {
						fmt.Printf("*")
					} else if in(p_) {
						fmt.Printf("@")
					} else {
						fmt.Printf("%s", string(lines[y][x]))
					}
				}
				fmt.Println()
			}
			fmt.Println()
		}

		const TRACE_POINTS = false
		if TRACE_POINTS {
			fmt.Printf("%s\n", p)
		}
	}

	start := func() image.Point {
		for y := range lines {
			for x := range lines[0] {
				if lines[y][x] == startChar {
					return image.Point{x, y}
				}
			}
		}
		return notFound
	}()
	if start == notFound {
		return 0, fmt.Errorf("error: no starting point")
	}
	track(start)

	next, d := func(start image.Point) (image.Point, direction) {
		for _, e := range []struct {
			p image.Point
			d direction
		}{
			{image.Point{start.X, start.Y - 1}, South},
			{image.Point{start.X + 1, start.Y}, West},
			{image.Point{start.X, start.Y + 1}, North},
			{image.Point{start.X - 1, start.Y}, East},
		} {
			// precondition 1: must be on board
			// no puzzle has the starting point at the border, but why not just write some more generic code?
			if !e.p.In(dim) {
				continue
			}

			// precondition 2: exit must match neighbor's entrance
			cell := lines[e.p.Y][e.p.X]
			if hasConnection(cell, e.d) {
				return e.p, other(lines[e.p.Y][e.p.X], e.d)
			}
		}
		return notFound, 0
	}(start)
	if next == notFound {
		return 0, fmt.Errorf("error: no connecting tile from starting point %v", start)
	}
	track(next)

	for next != start {
		var delta image.Point
		switch d {
		case North:
			delta.Y--
		case South:
			delta.Y++
		case West:
			delta.X--
		case East:
			delta.X++
		}

		next = next.Add(delta)
		track(next) // in the last step, close polygon => poly[0] == poly[N]
		d = other(lines[next.Y][next.X], opposite1(d))
	}

	if part1 {
		// we counted a complete round trip, so the farthest point in our journey is
		return (vertices - 1) / 2, nil
	}

	const TRACE_GRID = false
	var inside uint
	for y := range lines {
		for x := range lines[y] {
			p := image.Point{x, y}
			// do not count the polygon itself, only embedded points
			if in(p) {
				if TRACE_GRID {
					fmt.Printf("-")
				}
			} else {
				n := wnPnPoly(p, poly)
				if n == 1 {
					inside++
					if TRACE_GRID {
						fmt.Printf("1")
					}
				} else {
					if TRACE_GRID {
						fmt.Printf("0")
					}
				}
			}
		}
		if TRACE_GRID {
			fmt.Println()
		}
	}
	return inside, nil
}
