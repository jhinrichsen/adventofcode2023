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

	minCoord := int64(200000000000000)
	maxCoord := int64(400000000000000)

	count := 0
	for i := 0; i < len(puzzle); i++ {
		for j := i + 1; j < len(puzzle); j++ {
			h1, h2 := puzzle[i], puzzle[j]

			// Check if paths intersect in XY plane
			// Line 1: (px1, py1) + t1 * (vx1, vy1)
			// Line 2: (px2, py2) + t2 * (vx2, vy2)
			// At intersection:
			// px1 + t1*vx1 = px2 + t2*vx2
			// py1 + t1*vy1 = py2 + t2*vy2
			// Rearranging:
			// t1*vx1 - t2*vx2 = px2 - px1
			// t1*vy1 - t2*vy2 = py2 - py1

			// Using determinants to solve:
			// det = vx1*(-vy2) - vy1*(-vx2) = vx1*vy2 - vy1*vx2
			det := h1.vx*h2.vy - h1.vy*h2.vx

			if det == 0 {
				// Parallel lines
				continue
			}

			// Solve for t1 and t2
			dx := h2.px - h1.px
			dy := h2.py - h1.py

			// t1 = (dx * vy2 - dy * vx2) / det
			// t2 = (dx * vy1 - dy * vx1) / det

			t1Num := dx*h2.vy - dy*h2.vx
			t2Num := dx*h1.vy - dy*h1.vx

			// Check if both are in the future (same sign as det for positive t)
			if (t1Num < 0) != (det < 0) || (t2Num < 0) != (det < 0) {
				// At least one is in the past
				continue
			}

			// Calculate intersection point
			// x = px1 + t1 * vx1 = px1 + (t1Num / det) * vx1
			// y = py1 + t1 * vy1 = py1 + (t1Num / det) * vy1

			// To avoid floating point, check bounds using scaled values
			// x * det = px1 * det + t1Num * vx1
			// We want: minCoord <= x <= maxCoord
			// Which means: minCoord * det <= x * det <= maxCoord * det (if det > 0)
			// Or: maxCoord * det <= x * det <= minCoord * det (if det < 0)

			xScaled := h1.px*det + t1Num*h1.vx
			yScaled := h1.py*det + t1Num*h1.vy

			var inBounds bool
			if det > 0 {
				inBounds = minCoord*det <= xScaled && xScaled <= maxCoord*det &&
					minCoord*det <= yScaled && yScaled <= maxCoord*det
			} else {
				inBounds = maxCoord*det <= xScaled && xScaled <= minCoord*det &&
					maxCoord*det <= yScaled && yScaled <= minCoord*det
			}

			if inBounds {
				count++
			}
		}
	}

	return uint(count)
}
