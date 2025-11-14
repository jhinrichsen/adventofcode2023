package adventofcode2023

type Day18Puzzle []struct {
	dir   byte
	dist  int
	color string
}

func NewDay18(lines []string) (Day18Puzzle, error) {
	// Pre-allocate puzzle with estimated capacity
	puzzle := make(Day18Puzzle, 0, len(lines))

	for _, line := range lines {
		if line == "" {
			continue
		}

		// Parse inline without strings.Fields
		// Format: "R 6 (#70c710)"
		i := 0

		// Parse direction
		dir := line[0]
		i = 2 // Skip dir and space

		// Parse distance
		dist := 0
		for i < len(line) && line[i] >= '0' && line[i] <= '9' {
			dist = dist*10 + int(line[i]-'0')
			i++
		}

		// Skip " (#"
		i += 3

		// Parse color (6 chars)
		color := line[i : i+6]

		puzzle = append(puzzle, struct {
			dir   byte
			dist  int
			color string
		}{dir, dist, color})
	}
	return puzzle, nil
}

func Day18(puzzle Day18Puzzle, part1 bool) uint {
	x, y := 0, 0
	// Pre-allocate vertices with exact capacity
	vertices := make([][2]int, 0, len(puzzle)+1)
	vertices = append(vertices, [2]int{0, 0})
	perimeter := 0

	for _, inst := range puzzle {
		var dir byte
		var dist int

		if part1 {
			dir = inst.dir
			dist = inst.dist
		} else {
			// Decode color: first 5 hex digits = distance, last digit = direction
			dist = 0
			for i := 0; i < 5; i++ {
				dist = dist*16 + hexToInt(inst.color[i])
			}
			// Direction: 0=R, 1=D, 2=L, 3=U
			switch inst.color[5] {
			case '0':
				dir = 'R'
			case '1':
				dir = 'D'
			case '2':
				dir = 'L'
			case '3':
				dir = 'U'
			}
		}

		perimeter += dist
		switch dir {
		case 'R':
			x += dist
		case 'L':
			x -= dist
		case 'D':
			y += dist
		case 'U':
			y -= dist
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

func hexToInt(c byte) int {
	if c >= '0' && c <= '9' {
		return int(c - '0')
	}
	if c >= 'a' && c <= 'f' {
		return int(c - 'a' + 10)
	}
	if c >= 'A' && c <= 'F' {
		return int(c - 'A' + 10)
	}
	return 0
}
