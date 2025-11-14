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
	var bricks Day22Puzzle

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
	if !part1 {
		return 0
	}

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
	supports := make([][]int, len(bricks))    // supports[i] = list of bricks that brick i supports
	supportedBy := make([][]int, len(bricks)) // supportedBy[i] = list of bricks that support brick i

	for i := range bricks {
		for j := 0; j < i; j++ {
			if overlapsXY(bricks[i], bricks[j]) && max(bricks[j].z1, bricks[j].z2)+1 == min(bricks[i].z1, bricks[i].z2) {
				supports[j] = append(supports[j], i)
				supportedBy[i] = append(supportedBy[i], j)
			}
		}
	}

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

func overlapsXY(a, b brick) bool {
	return max(a.x1, a.x2) >= min(b.x1, b.x2) &&
		min(a.x1, a.x2) <= max(b.x1, b.x2) &&
		max(a.y1, a.y2) >= min(b.y1, b.y2) &&
		min(a.y1, a.y2) <= max(b.y1, b.y2)
}
