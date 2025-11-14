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
	if part1 {
		return solvePart1(puzzle)
	}
	return solvePart2(puzzle)
}

func solvePart1(puzzle Day24Puzzle) uint {

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

func solvePart2(puzzle Day24Puzzle) uint {
	// Strategy: In a reference frame moving with the rock's velocity (vx, vy, vz),
	// the rock is stationary and all hailstones must pass through the same point.
	// We search for velocities where this constraint holds.

	// Try different rock velocities in XY plane
	for vx := int64(-500); vx <= 500; vx++ {
		for vy := int64(-500); vy <= 500; vy++ {
			h0, h1 := puzzle[0], puzzle[1]

			// Adjusted velocities in reference frame moving with (vx, vy)
			dvx0, dvy0 := h0.vx-vx, h0.vy-vy
			dvx1, dvy1 := h1.vx-vx, h1.vy-vy

			// Find XY intersection of first two hailstones
			det := dvx0*dvy1 - dvy0*dvx1
			if det == 0 {
				continue // Parallel in XY plane
			}

			dx := h1.px - h0.px
			dy := h1.py - h0.py
			t0 := (dx*dvy1 - dy*dvx1) / det
			t1 := (dx*dvy0 - dy*dvx0) / det

			if t0 < 0 || t1 < 0 {
				continue // Intersection in the past
			}

			// XY coordinates of intersection point (rock's position)
			px := h0.px + t0*dvx0
			py := h0.py + t0*dvy0

			// Verify third hailstone also passes through this XY point
			h2 := puzzle[2]
			dvx2, dvy2 := h2.vx-vx, h2.vy-vy

			var t2 int64
			if dvx2 != 0 {
				if (px-h2.px)%dvx2 != 0 {
					continue
				}
				t2 = (px - h2.px) / dvx2
			} else if dvy2 != 0 {
				if (py-h2.py)%dvy2 != 0 {
					continue
				}
				t2 = (py - h2.py) / dvy2
			} else {
				continue // Stationary in XY
			}

			if t2 < 0 {
				continue
			}

			if h2.px+t2*dvx2 != px || h2.py+t2*dvy2 != py {
				continue // Third hailstone doesn't pass through same XY point
			}

			// Calculate vz analytically from the constraint that h0 and h1
			// reach the same Z coordinate at their respective intersection times
			// h0.pz + t0*(h0.vz - vz) = h1.pz + t1*(h1.vz - vz)
			// => vz = (h1.pz + t1*h1.vz - h0.pz - t0*h0.vz) / (t1 - t0)
			if t1 == t0 {
				continue
			}

			vzNum := h1.pz + t1*h1.vz - h0.pz - t0*h0.vz
			if vzNum%(t1-t0) != 0 {
				continue // vz must be an integer
			}
			vz := vzNum / (t1 - t0)

			pz := h0.pz + t0*(h0.vz-vz)

			// Verify this solution works for all hailstones
			valid := true
			for _, h := range puzzle {
				dvx, dvy, dvz := h.vx-vx, h.vy-vy, h.vz-vz

				// Find collision time
				var t int64
				found := false
				if dvx != 0 {
					if (px-h.px)%dvx == 0 {
						t = (px - h.px) / dvx
						found = true
					}
				} else if dvy != 0 {
					if (py-h.py)%dvy == 0 {
						t = (py - h.py) / dvy
						found = true
					}
				} else if dvz != 0 {
					if (pz-h.pz)%dvz == 0 {
						t = (pz - h.pz) / dvz
						found = true
					}
				}

				if !found || t < 0 {
					valid = false
					break
				}

				// Verify all three coordinates match at time t
				if h.px+t*dvx != px || h.py+t*dvy != py || h.pz+t*dvz != pz {
					valid = false
					break
				}
			}

			if valid {
				return uint(px + py + pz)
			}
		}
	}

	return 0
}
