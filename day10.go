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

func Day10(lines [][]byte, part1 bool) (uint, error) {
	const (
		markChar   = '1'
		unmarkChar = '0'
		startChar  = 'S'
	)

	r := image.Rectangle{image.Point{0, 0}, image.Point{len(lines), len(lines[0])}}

	start := func() image.Point {
		for y := range lines {
			for x := range lines[0] {
				if lines[y][x] == startChar {
					return image.Point{x, y}
				}
			}
		}
		panic(fmt.Sprintf("puzzle has no starting point %c", startChar))
	}()

	next, d := func(start image.Point) (image.Point, direction) {
		for _, e := range []struct {
			p image.Point
			d direction
		}{
			{image.Point{start.X, start.Y - 1}, South},
			{image.Point{start.X, start.Y + 1}, North},
			{image.Point{start.X - 1, start.Y}, East},
			{image.Point{start.X + 1, start.Y}, West},
		} {
			// precondition 1: must be on board
			// no puzzle has the starting point at the border, but why not just write some more generic code?
			if !e.p.In(r) {
				continue
			}

			// precondition 2: exit must match neighbor's entrance
			cell := lines[e.p.Y][e.p.X]
			if hasConnection(cell, e.d) {
				return e.p, other(lines[e.p.Y][e.p.X], e.d)
			}
		}
		panic(fmt.Sprintf("error: no connecting tile from starting point %v", start))
	}(start)

	var steps uint
	for steps = 1; next != start; steps++ {
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
		d = other(lines[next.Y][next.X], opposite1(d))

		if !part1 {
			lines[next.Y][next.X] = markChar
		}
	}

	if part1 {
		// we made a complete trip, so the farthest point in our journey is
		return steps / 2, nil
	}

	// reset all non-marked fields
	for y := range lines {
		for x := range lines[0] {
			if lines[y][x] != markChar {
				lines[y][x] = unmarkChar
			}
		}
	}
	for y := range lines {
		fmt.Println(lines[y])
	}
	return 42, nil
}
