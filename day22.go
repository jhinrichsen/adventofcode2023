package adventofcode2023

import (
	"sort"
)

type brick struct {
	x1, y1, z1 int
	x2, y2, z2 int
}

type Day22Puzzle []brick

func NewDay22(lines []string) (Day22Puzzle, error) {
	// Pre-allocate bricks with estimated capacity
	bricks := make(Day22Puzzle, 0, len(lines))

	for _, line := range lines {
		if line == "" {
			continue
		}

		var b brick
		// Parse format: x1,y1,z1~x2,y2,z2
		var i int
		for i < len(line) && line[i] != ',' {
			b.x1 = b.x1*10 + int(line[i]-'0')
			i++
		}
		i++ // skip comma
		for i < len(line) && line[i] != ',' {
			b.y1 = b.y1*10 + int(line[i]-'0')
			i++
		}
		i++ // skip comma
		for i < len(line) && line[i] != '~' {
			b.z1 = b.z1*10 + int(line[i]-'0')
			i++
		}
		i++ // skip tilde
		for i < len(line) && line[i] != ',' {
			b.x2 = b.x2*10 + int(line[i]-'0')
			i++
		}
		i++ // skip comma
		for i < len(line) && line[i] != ',' {
			b.y2 = b.y2*10 + int(line[i]-'0')
			i++
		}
		i++ // skip comma
		for i < len(line) {
			b.z2 = b.z2*10 + int(line[i]-'0')
			i++
		}

		bricks = append(bricks, b)
	}

	return bricks, nil
}

func Day22(puzzle Day22Puzzle, part1 bool) uint {
	// Sort bricks by lowest z coordinate
	bricks := make([]brick, len(puzzle))
	copy(bricks, puzzle)
	sort.Slice(bricks, func(i, j int) bool {
		return min(bricks[i].z1, bricks[i].z2) < min(bricks[j].z1, bricks[j].z2)
	})

	// Simulate falling
	for i := range bricks {
		maxZ := 1
		for j := 0; j < i; j++ {
			if overlapsXY(bricks[i], bricks[j]) {
				maxZ = max(maxZ, max(bricks[j].z1, bricks[j].z2)+1)
			}
		}
		// Move brick to rest at maxZ
		dz := min(bricks[i].z1, bricks[i].z2) - maxZ
		bricks[i].z1 -= dz
		bricks[i].z2 -= dz
	}

	// Build support relationships
	// Pre-allocate with estimated capacity (average ~3 connections per brick)
	supports := make([][]int, len(bricks))
	supportedBy := make([][]int, len(bricks))
	for i := range supports {
		supports[i] = make([]int, 0, 3)
		supportedBy[i] = make([]int, 0, 3)
	}

	for i := range bricks {
		for j := 0; j < i; j++ {
			if overlapsXY(bricks[i], bricks[j]) && max(bricks[j].z1, bricks[j].z2)+1 == min(bricks[i].z1, bricks[i].z2) {
				supports[j] = append(supports[j], i)
				supportedBy[i] = append(supportedBy[i], j)
			}
		}
	}

	if part1 {
		// Count bricks that can be safely disintegrated
		var count uint
		for i := range bricks {
			canRemove := true
			for _, above := range supports[i] {
				if len(supportedBy[above]) == 1 {
					// This brick is the sole supporter of 'above'
					canRemove = false
					break
				}
			}
			if canRemove {
				count++
			}
		}
		return count
	}

	// Part 2: Count total number of bricks that would fall
	var total uint
	// Reusable arrays to avoid allocations
	fallen := make([]bool, len(bricks))
	queue := make([]int, len(bricks))

	for i := range bricks {
		// Reset fallen array
		for j := range fallen {
			fallen[j] = false
		}
		fallen[i] = true

		// Use queue for BFS propagation with head/tail pointers
		head, tail := 0, 1
		queue[0] = i

		for head < tail {
			current := queue[head]
			head++

			// Check all bricks supported by current
			for _, above := range supports[current] {
				if fallen[above] {
					continue
				}

				// Check if all supporters of 'above' have fallen
				allFallen := true
				for _, supporter := range supportedBy[above] {
					if !fallen[supporter] {
						allFallen = false
						break
					}
				}

				if allFallen {
					fallen[above] = true
					queue[tail] = above
					tail++
				}
			}
		}

		// Count fallen bricks (excluding the initial brick i)
		var count uint
		for j := range fallen {
			if fallen[j] && j != i {
				count++
			}
		}
		total += count
	}

	return total
}

func overlapsXY(a, b brick) bool {
	return max(a.x1, a.x2) >= min(b.x1, b.x2) &&
		min(a.x1, a.x2) <= max(b.x1, b.x2) &&
		max(a.y1, a.y2) >= min(b.y1, b.y2) &&
		min(a.y1, a.y2) <= max(b.y1, b.y2)
}
