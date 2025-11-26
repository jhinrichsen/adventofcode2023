package adventofcode2023

import (
	"fmt"
	"testing"
)

func TestFloodFillWithHoles(t *testing.T) {
	// Input grid: 1 = boundary, 0 = empty
	grid := [][]int{
		{1, 1, 1, 1, 1, 1},
		{1, 0, 0, 1, 0, 1},
		{1, 0, 1, 1, 0, 1},
		{1, 1, 1, 0, 0, 1},
		{1, 1, 1, 1, 1, 1},
	}

	// Process the grid
	result := floodFillWithHoles(grid)

	// Print the result
	for _, row := range result {
		fmt.Println(row)
	}
}
