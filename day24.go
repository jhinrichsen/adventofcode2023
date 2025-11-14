package adventofcode2023

type hailstone struct {
	px, py, pz int64
	vx, vy, vz int64
}

type Day24Puzzle []hailstone

func NewDay24(lines []string) (Day24Puzzle, error) {
	var puzzle Day24Puzzle

	for _, line := range lines {
		if line == "" {
			continue
		}

		var h hailstone
		// Parse format: px, py, pz @ vx, vy, vz
		var i int

		// Parse px
		negative := false
		for i < len(line) && (line[i] == ' ' || line[i] == '\t') {
			i++
		}
		if line[i] == '-' {
			negative = true
			i++
		}
		for i < len(line) && line[i] >= '0' && line[i] <= '9' {
			h.px = h.px*10 + int64(line[i]-'0')
			i++
		}
		if negative {
			h.px = -h.px
		}

		// Skip comma and spaces
		for i < len(line) && (line[i] == ',' || line[i] == ' ' || line[i] == '\t') {
			i++
		}

		// Parse py
		negative = false
		if line[i] == '-' {
			negative = true
			i++
		}
		for i < len(line) && line[i] >= '0' && line[i] <= '9' {
			h.py = h.py*10 + int64(line[i]-'0')
			i++
		}
		if negative {
			h.py = -h.py
		}

		// Skip comma and spaces
		for i < len(line) && (line[i] == ',' || line[i] == ' ' || line[i] == '\t') {
			i++
		}

		// Parse pz
		negative = false
		if line[i] == '-' {
			negative = true
			i++
		}
		for i < len(line) && line[i] >= '0' && line[i] <= '9' {
			h.pz = h.pz*10 + int64(line[i]-'0')
			i++
		}
		if negative {
			h.pz = -h.pz
		}

		// Skip @ and spaces
		for i < len(line) && (line[i] == '@' || line[i] == ' ' || line[i] == '\t') {
			i++
		}

		// Parse vx
		negative = false
		if line[i] == '-' {
			negative = true
			i++
		}
		for i < len(line) && line[i] >= '0' && line[i] <= '9' {
			h.vx = h.vx*10 + int64(line[i]-'0')
			i++
		}
		if negative {
			h.vx = -h.vx
		}

		// Skip comma and spaces
		for i < len(line) && (line[i] == ',' || line[i] == ' ' || line[i] == '\t') {
			i++
		}

		// Parse vy
		negative = false
		if line[i] == '-' {
			negative = true
			i++
		}
		for i < len(line) && line[i] >= '0' && line[i] <= '9' {
			h.vy = h.vy*10 + int64(line[i]-'0')
			i++
		}
		if negative {
			h.vy = -h.vy
		}

		// Skip comma and spaces
		for i < len(line) && (line[i] == ',' || line[i] == ' ' || line[i] == '\t') {
			i++
		}

		// Parse vz
		negative = false
		if line[i] == '-' {
			negative = true
			i++
		}
		for i < len(line) && line[i] >= '0' && line[i] <= '9' {
			h.vz = h.vz*10 + int64(line[i]-'0')
			i++
		}
		if negative {
			h.vz = -h.vz
		}

		puzzle = append(puzzle, h)
	}

	return puzzle, nil
}

func Day24(puzzle Day24Puzzle, part1 bool) uint {
	if !part1 {
		return 0
	}

	// Coordinates are ~2-5×10^14 (15 digits). Using int64 arithmetic causes overflow
	// when multiplying by determinants (~100-1000), exceeding int64 max (9×10^18).
	// float64 has ~15-17 significant digits of precision, sufficient for our 15-digit
	// coordinates and avoiding overflow. Go doesn't have float128, and big.Int would
	// be overkill for simple bounds checking.
	minCoord := 200000000000000.0
	maxCoord := 400000000000000.0

	count := 0
	for i := 0; i < len(puzzle); i++ {
		for j := i + 1; j < len(puzzle); j++ {
			h1, h2 := puzzle[i], puzzle[j]

			// Check if paths intersect in XY plane
			det := float64(h1.vx*h2.vy - h1.vy*h2.vx)

			if det == 0 {
				// Parallel lines
				continue
			}

			dx := float64(h2.px - h1.px)
			dy := float64(h2.py - h1.py)

			// t1 = (dx * vy2 - dy * vx2) / det
			// t2 = (dx * vy1 - dy * vx1) / det
			t1 := (dx*float64(h2.vy) - dy*float64(h2.vx)) / det
			t2 := (dx*float64(h1.vy) - dy*float64(h1.vx)) / det

			// Check if both are in the future
			if t1 < 0 || t2 < 0 {
				continue
			}

			// Calculate intersection point
			x := float64(h1.px) + t1*float64(h1.vx)
			y := float64(h1.py) + t1*float64(h1.vy)

			// Check if in bounds
			if x >= minCoord && x <= maxCoord && y >= minCoord && y <= maxCoord {
				count++
			}
		}
	}

	return uint(count)
}
