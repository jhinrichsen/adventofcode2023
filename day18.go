package adventofcode2023

import "strings"

type Day18Puzzle struct {
	instructions []struct {
		dir   byte
		dist  int
		color string
	}
}

func NewDay18(lines []string) (Day18Puzzle, error) {
	var puzzle Day18Puzzle
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Fields(line)
		dir := parts[0][0]
		dist := 0
		for i := 0; i < len(parts[1]); i++ {
			dist = dist*10 + int(parts[1][i]-'0')
		}
		color := parts[2][2 : len(parts[2])-1]
		puzzle.instructions = append(puzzle.instructions, struct {
			dir   byte
			dist  int
			color string
		}{dir, dist, color})
	}
	return puzzle, nil
}

func Day18(puzzle Day18Puzzle, part1 bool) uint {
	if !part1 {
		return 0
	}

	x, y := 0, 0
	vertices := [][2]int{{0, 0}}
	perimeter := 0

	for _, inst := range puzzle.instructions {
		perimeter += inst.dist
		switch inst.dir {
		case 'R':
			x += inst.dist
		case 'L':
			x -= inst.dist
		case 'D':
			y += inst.dist
		case 'U':
			y -= inst.dist
		}
		vertices = append(vertices, [2]int{x, y})
	}

	area := 0
	for i := 0; i < len(vertices)-1; i++ {
		area += vertices[i][0] * vertices[i+1][1]
		area -= vertices[i+1][0] * vertices[i][1]
	}
	if area < 0 {
		area = -area
	}
	area /= 2

	total := area + perimeter/2 + 1
	return uint(total)
}
